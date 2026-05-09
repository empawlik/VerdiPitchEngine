package converter

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/empawlik/verdi-pitch-engine/internal/fs"
)

func TestRunPool(t *testing.T) {
	outDir := t.TempDir()

	// Task 1: Output already exists (should be skipped)
	existingOutPath := filepath.Join(outDir, "existing.flac")
	if err := os.WriteFile(existingOutPath, []byte("dummy"), 0644); err != nil {
		t.Fatalf("Failed to write dummy existing file: %v", err)
	}

	// Task 2: Output does not exist, should attempt processing and fail
	nonExistingOutPath := filepath.Join(outDir, "missing.flac")

	tasks := []fs.Task{
		{InputPath: "in1.flac", OutputPath: existingOutPath},
		{InputPath: "in2.flac", OutputPath: nonExistingOutPath},
	}

	// Running RunPool should process both tasks
	// 1 skipped, 1 error
	RunPool(tasks, 2)

	// We verify that existing file is untouched
	if _, err := os.Stat(existingOutPath); os.IsNotExist(err) {
		t.Error("Existing output file was deleted or missing")
	}

	// We verify that the non-existing output file might have been created and deleted by ProcessFile error handling
	if _, err := os.Stat(nonExistingOutPath); !os.IsNotExist(err) {
		t.Error("Expected non-existing output file to remain missing/deleted due to ProcessFile error")
	}
}
