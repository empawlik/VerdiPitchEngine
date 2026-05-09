package converter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// getBitDepth uses ffprobe to extract the bits per sample of the input audio file.
func getBitDepth(filePath string) (int, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "stream=bits_per_sample",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe bit_depth failed: %w", err)
	}

	depthStr := strings.TrimSpace(out.String())
	depthStr = strings.Split(depthStr, "\n")[0]
	if depthStr == "" || depthStr == "N/A" {
		return 16, nil // Default fallback
	}
	depth, err := strconv.Atoi(depthStr)
	if err != nil {
		return 16, nil // Default fallback on parse error
	}
	return depth, nil
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func ProcessFile(inPath, outPath string) error {
	depth, _ := getBitDepth(inPath)

	// High-fidelity Time-Scale Modification (TSM) pitch-shift using librubberband.
	// This preserves the exact track duration (tempo) while shifting pitch from 440 Hz to 432 Hz.
	pitchRatio := float64(432) / float64(440)
	filter := fmt.Sprintf("rubberband=pitch=%f", pitchRatio)

	args := []string{
		"-v", "warning",
		"-i", inPath,
		"-af", filter,
		"-map", "0:a?",
		"-map", "0:v?",
		"-map_metadata", "0",
		"-metadata", "ENCODED_BY=VerdiPitchEngine",
		"-metadata", "PITCH_SHIFT=432Hz",
		"-metadata", "VERSION=432 Hz",
		"-c:a", "flac",
		"-c:v", "copy",
	}

	if depth >= 24 {
		// Preserve 24-bit audio (FLAC uses s32 internally for 24-bit)
		args = append(args, "-sample_fmt", "s32")
	} else {
		// Default to 16-bit
		args = append(args, "-sample_fmt", "s16", "-dither_method", "triangular_hp")
	}

	args = append(args, "-compression_level", "12", outPath)

	cmd := exec.Command("ffmpeg", args...)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderr.String())
	}

	return nil
}
