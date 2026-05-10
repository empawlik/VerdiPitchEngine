package converter

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"sync/atomic"
	"time"

	"github.com/empawlik/verdi-pitch-engine/internal/fs"
	"github.com/vbauerster/mpb/v8"
	"github.com/vbauerster/mpb/v8/decor"
)

// RunPool starts a worker pool to process the given tasks.
func RunPool(tasks []fs.Task, numWorkers int) {
	taskCh := make(chan fs.Task, len(tasks))
	for _, t := range tasks {
		taskCh <- t
	}
	close(taskCh)

	var wg sync.WaitGroup
	var processed int32
	var skipped int32
	var errors int32

	p := mpb.New(mpb.WithWidth(60))
	totalBar := p.AddBar(int64(len(tasks)),
		mpb.PrependDecorators(
			decor.Name("Total Progress", decor.WC{W: 20, C: decor.DSyncSpaceR}),
			decor.CountersNoUnit("%d / %d", decor.WCSyncWidth),
		),
		mpb.AppendDecorators(
			decor.Percentage(decor.WCSyncWidth),
		),
	)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskCh {
				// Idempotency check: if output exists, skip
				if _, err := os.Stat(task.OutputPath); err == nil {
					atomic.AddInt32(&skipped, 1)
					totalBar.Increment()
					continue
				}

				name := filepath.Base(task.InputPath)
				if len(name) > 20 {
					name = name[:17] + "..."
				}

				log.Printf("➡️  [Worker %d] Started processing: %s", workerID, name)

				bar := p.AddBar(100, // placeholder, will be updated to total microseconds by ProcessFile
					mpb.BarRemoveOnComplete(),
					mpb.PrependDecorators(
						decor.Name(fmt.Sprintf("[W%d] %s", workerID, name), decor.WC{W: 25, C: decor.DSyncSpaceR}),
						decor.Percentage(decor.WCSyncSpace),
					),
					mpb.AppendDecorators(
						decor.OnComplete(
							decor.Elapsed(decor.ET_STYLE_GO, decor.WCSyncSpace), "Done!",
						),
					),
				)

				// Prevent file processing from hanging
				ctx, cancel := context.WithTimeout(context.Background(), 15*time.Minute)

				if err := ProcessFile(ctx, task.InputPath, task.OutputPath, bar); err != nil {
					atomic.AddInt32(&errors, 1)
					log.Printf("❌ [Worker %d] Error processing %s: %v", workerID, task.InputPath, err)
					os.Remove(task.OutputPath)
					bar.Abort(true)
				} else {
					atomic.AddInt32(&processed, 1)
					bar.SetTotal(bar.Current(), true)
					log.Printf("✅ [Worker %d] Finished processing: %s", workerID, name)
				}
				cancel()
				totalBar.Increment()
			}
		}(i)
	}

	wg.Wait()
	p.Wait()
	log.Printf("\nWorker pool completed. Processed: %d, Skipped: %d, Errors: %d",
		atomic.LoadInt32(&processed),
		atomic.LoadInt32(&skipped),
		atomic.LoadInt32(&errors))
}
