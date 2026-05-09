package converter

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"github.com/vbauerster/mpb/v8"
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

// getDuration uses ffprobe to extract the duration in seconds of the input audio file.
func getDuration(filePath string) (float64, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "format=duration",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe duration failed: %w", err)
	}

	durationStr := strings.TrimSpace(out.String())
	durationStr = strings.Split(durationStr, "\n")[0]
	if durationStr == "" || durationStr == "N/A" {
		return 0, nil
	}
	duration, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		return 0, nil // Default fallback
	}
	return duration, nil
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func buildFFmpegArgs(inPath, outPath string, depth int) []string {
	pitchRatio := float64(432) / float64(440)
	filter := fmt.Sprintf("rubberband=pitch=%f", pitchRatio)

	args := []string{
		"-v", "warning",
		"-progress", "pipe:1",
		"-nostats",
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
	return args
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func ProcessFile(ctx context.Context, inPath, outPath string, bar *mpb.Bar) error {
	depth, _ := getBitDepth(inPath)
	durationSec, _ := getDuration(inPath)
	totalMicrosec := int64(durationSec * 1000000)

	if bar != nil && totalMicrosec > 0 {
		bar.SetTotal(totalMicrosec, false)
	}

	// High-fidelity Time-Scale Modification (TSM) pitch-shift using librubberband.
	args := buildFFmpegArgs(inPath, outPath, depth)

	cmd := exec.CommandContext(ctx, "ffmpeg", args...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("ffmpeg failed to start: %w, stderr: %s", err, stderr.String())
	}

	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "out_time_us=") {
			usStr := strings.TrimPrefix(line, "out_time_us=")
			us, err := strconv.ParseInt(usStr, 10, 64)
			if err == nil && bar != nil && us > 0 {
				bar.SetCurrent(us)
			}
		}
	}

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderr.String())
	}

	if bar != nil && totalMicrosec > 0 {
		bar.SetCurrent(totalMicrosec)
	}

	return nil
}
