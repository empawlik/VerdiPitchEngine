package converter

import (
	"testing"
)

func TestProcessFile_FailsWithoutFFmpegOrInvalidFile(t *testing.T) {
	inPath := "dummy.flac"
	outPath := "out.flac"

	// This should fail because dummy.flac doesn't exist and/or ffmpeg might not be installed in the test environment
	err := ProcessFile(inPath, outPath)
	if err == nil {
		t.Error("Expected error from ProcessFile, got nil")
	}
}
