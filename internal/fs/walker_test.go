package fs

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestWalkAndCollect(t *testing.T) {
	// Setup
	inDir := t.TempDir()
	outDir := t.TempDir()

	// Create test files
	files := []string{
		"song1.flac",
		"sub/song2.FLAC",
		"sub/cover.jpg",
		"ignore.txt",
	}

	for _, f := range files {
		path := filepath.Join(inDir, f)
		err := os.MkdirAll(filepath.Dir(path), 0755)
		if err != nil {
			t.Fatal(err)
		}
		err = os.WriteFile(path, []byte("dummy data"), 0644)
		if err != nil {
			t.Fatal(err)
		}
	}

	// Test
	tasks, err := WalkAndCollect(inDir, outDir)
	if err != nil {
		t.Fatalf("WalkAndCollect returned an error: %v", err)
	}

	// Assert
	if len(tasks) != 2 {
		t.Fatalf("Expected 2 tasks, got %d", len(tasks))
	}

	expectedTasks := []Task{
		{
			InputPath:  filepath.Join(inDir, "song1.flac"),
			OutputPath: filepath.Join(outDir, "song1.flac"),
		},
		{
			InputPath:  filepath.Join(inDir, "sub/song2.FLAC"),
			OutputPath: filepath.Join(outDir, "sub/song2.FLAC"),
		},
	}

	// Verify the elements exist
	for _, expected := range expectedTasks {
		found := false
		for _, actual := range tasks {
			if reflect.DeepEqual(actual, expected) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected task %v not found in results: %v", expected, tasks)
		}
	}

	// Test error case (invalid input directory)
	_, err = WalkAndCollect("/path/does/not/exist/surely", outDir)
	if err == nil {
		t.Error("Expected error for non-existent input directory")
	}
}
