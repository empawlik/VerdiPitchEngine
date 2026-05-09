package fs

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Task represents a conversion job
type Task struct {
	InputPath  string
	OutputPath string
}

// WalkAndCollect finds all .flac files in inDir and maps them to outDir.
// It creates the necessary directory structure in outDir.
func WalkAndCollect(inDir, outDir string) ([]Task, error) {
	var tasks []Task

	err := filepath.Walk(inDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := strings.ToLower(filepath.Ext(path))
		relPath, err := filepath.Rel(inDir, path)
		if err != nil {
			return fmt.Errorf("failed to get relative path for %s: %w", path, err)
		}

		outPath := filepath.Join(outDir, relPath)
		outDirPath := filepath.Dir(outPath)

		// Create the target directory structure
		if err := os.MkdirAll(outDirPath, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", outDirPath, err)
		}

		if ext == ".flac" {
			tasks = append(tasks, Task{
				InputPath:  path,
				OutputPath: outPath,
			})
		} else {
			// Copy non-FLAC files (e.g., folder.jpg, cover.png, .pdf) natively
			inF, err := os.Open(path)
			if err == nil {
				defer inF.Close()
				outF, err := os.Create(outPath)
				if err == nil {
					defer outF.Close()
					// Use io.Copy to handle potentially large PDFs or images efficiently
					_, _ = io.Copy(outF, inF)
				}
			}
		}
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking input directory: %w", err)
	}

	return tasks, nil
}
