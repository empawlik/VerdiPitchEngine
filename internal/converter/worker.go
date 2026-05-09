package converter

import (
	"log"
	"os"
	"sync"
	"sync/atomic"

	"github.com/empawlik/verdi-pitch-engine/internal/fs"
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

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			for task := range taskCh {
				// Idempotency check: if output exists, skip
				if _, err := os.Stat(task.OutputPath); err == nil {
					atomic.AddInt32(&skipped, 1)
					log.Printf("[Worker %d] Skipped (already exists): %s", workerID, task.OutputPath)
					continue
				}

				log.Printf("[Worker %d] Processing: %s", workerID, task.InputPath)
				if err := ProcessFile(task.InputPath, task.OutputPath); err != nil {
					atomic.AddInt32(&errors, 1)
					log.Printf("[Worker %d] Error processing %s: %v", workerID, task.InputPath, err)
					// In case of error, we might have created a partial file. We could remove it.
					os.Remove(task.OutputPath)
				} else {
					atomic.AddInt32(&processed, 1)
					log.Printf("[Worker %d] Finished: %s", workerID, task.OutputPath)
				}
			}
		}(i)
	}

	wg.Wait()
	log.Printf("Worker pool completed. Processed: %d, Skipped: %d, Errors: %d",
		atomic.LoadInt32(&processed),
		atomic.LoadInt32(&skipped),
		atomic.LoadInt32(&errors))
}
