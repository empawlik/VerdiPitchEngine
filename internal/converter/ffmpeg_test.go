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

func TestBuildFFmpegArgs(t *testing.T) {
	inPath := "in.flac"
	outPath := "out.flac"

	// Test 16-bit
	args16 := buildFFmpegArgs(inPath, outPath, 16)
	found16 := false
	for _, arg := range args16 {
		if arg == "s16" {
			found16 = true
			break
		}
	}
	if !found16 {
		t.Errorf("Expected 16-bit args to contain 's16', got: %v", args16)
	}

	// Test 24-bit
	args24 := buildFFmpegArgs(inPath, outPath, 24)
	found24 := false
	for _, arg := range args24 {
		if arg == "s32" {
			found24 = true
			break
		}
	}
	if !found24 {
		t.Errorf("Expected 24-bit args to contain 's32', got: %v", args24)
	}
}
