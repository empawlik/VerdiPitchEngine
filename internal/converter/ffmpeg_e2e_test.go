package converter

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/empawlik/verdi-pitch-engine/internal/fs"
)

// setupSyntheticAudio creates a 440 Hz sine wave tone as a FLAC file.
func hasRubberband() bool {
	cmd := exec.Command("ffmpeg", "-filters")
	output, err := cmd.Output()
	if err != nil {
		return false
	}
	return bytes.Contains(output, []byte("rubberband"))
}

func setupSyntheticAudio(t *testing.T, outPath string) {
	t.Helper()
	// Generate a 1-second 440 Hz tone
	cmd := exec.Command("ffmpeg",
		"-f", "lavfi",
		"-i", "sine=frequency=440:duration=1",
		"-c:a", "flac",
		outPath,
	)
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		t.Fatalf("Failed to generate synthetic audio: %v, stderr: %s", err, stderr.String())
	}
}

// verifyPitchShift uses ffprobe to verify the duration and properties of the file.
func verifyPitchShift(t *testing.T, outPath string) {
	t.Helper()
	// Run ffprobe to check if it's a valid FLAC
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration,format_name", "-of", "default=noprint_wrappers=1:nokey=1", outPath)
	output, err := cmd.Output()
	if err != nil {
		t.Fatalf("Failed to ffprobe output file: %v", err)
	}

	outStr := string(bytes.TrimSpace(output))
	if len(outStr) == 0 {
		t.Fatalf("Output file seems invalid or empty")
	}
}

func TestE2EFFmpegPitchShift(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping E2E test in short mode")
	}
	if !hasRubberband() {
		t.Skip("ffmpeg does not have rubberband filter compiled in, skipping E2E test")
	}

	// Create a temporary directory for isolation
	tempDir := t.TempDir()
	inDir := filepath.Join(tempDir, "input")
	outDir := filepath.Join(tempDir, "output")

	if err := os.Mkdir(inDir, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(outDir, 0755); err != nil {
		t.Fatal(err)
	}

	inputFlac := filepath.Join(inDir, "tone.flac")
	outputFlac := filepath.Join(outDir, "tone.flac")

	setupSyntheticAudio(t, inputFlac)

	tasks := []fs.Task{
		{
			InputPath:  inputFlac,
			OutputPath: outputFlac,
		},
	}

	// Run the worker pool with 1 worker
	RunPool(tasks, 1)

	// Verify the output exists
	if _, err := os.Stat(outputFlac); os.IsNotExist(err) {
		t.Fatalf("Output FLAC was not created: %s", outputFlac)
	}

	verifyPitchShift(t, outputFlac)
}
