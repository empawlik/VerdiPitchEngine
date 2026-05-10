package converter

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/vbauerster/mpb/v8"
)

var execCommand = exec.Command
var execCommandContext = exec.CommandContext

// getBitDepth uses ffprobe to extract the bits per sample or sample format of the input audio file.
func getBitDepth(filePath string) (int, error) {
	cmd := execCommand("ffprobe",
		"-v", "error",
		"-show_entries", "stream=bits_per_sample,bits_per_raw_sample,sample_fmt",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe bit_depth failed: %w", err)
	}

	output := strings.TrimSpace(out.String())
	lines := strings.Split(output, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "s32" || line == "s32p" {
			return 24, nil // FLAC uses s32 for 24-bit
		}
		if line == "s16" || line == "s16p" {
			return 16, nil
		}
		if val, err := strconv.Atoi(line); err == nil && val > 0 {
			if val == 24 || val == 32 {
				return 24, nil
			}
			return val, nil
		}
	}

	return 16, nil // Default fallback
}

// getDuration uses ffprobe to extract the duration in seconds of the input audio file.
func getDuration(filePath string) (float64, error) {
	cmd := execCommand("ffprobe",
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

// getSampleRate uses ffprobe to extract the sample rate of the input audio file.
func getSampleRate(filePath string) (string, error) {
	cmd := execCommand("ffprobe",
		"-v", "error",
		"-show_entries", "stream=sample_rate",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("ffprobe sample_rate failed: %w", err)
	}

	srStr := strings.TrimSpace(out.String())
	srStr = strings.Split(srStr, "\n")[0]
	if srStr == "" || srStr == "N/A" {
		return "", nil
	}
	return srStr, nil
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func buildFFmpegArgs(inPath, outPath string, depth int, sampleRate string) []string {
	pitchRatio := float64(432) / float64(440)
	filter := fmt.Sprintf("rubberband=pitch=%f", pitchRatio)

	args := []string{
		"-y",
		"-v", "warning",
		"-progress", "pipe:1",
		"-nostats",
		"-i", inPath,
		"-af", filter,
		"-map", "0:a?",
		"-map", "0:v?",
		"-map_metadata", "0",
		"-c:a", "flac",
		"-c:v", "copy",
	}

	if sampleRate != "" {
		args = append(args, "-ar", sampleRate)
	}

	if depth >= 24 {
		// Preserve 24-bit audio (FLAC uses s32 internally for 24-bit)
		args = append(args, "-sample_fmt", "s32")
	} else {
		// Default to 16-bit
		args = append(args, "-sample_fmt", "s16", "-dither_method", "triangular_hp")
	}

	args = append(args, "-compression_level", "5", "-f", "flac", outPath)
	return args
}

// inheritMetadata uses metaflac to strictly copy Vorbis comments and embedded pictures from source to destination.
func inheritMetadata(ctx context.Context, inPath, outPath string) error {
	// 1. Copy Vorbis comments (tags)
	var tagData bytes.Buffer
	exportTags := execCommandContext(ctx, "metaflac", "--export-tags-to=-", inPath)
	exportTags.Stdout = &tagData
	if err := exportTags.Run(); err != nil {
		return fmt.Errorf("failed to export tags: %w", err)
	}

	importTags := execCommandContext(ctx, "metaflac", "--remove-all-tags", "--import-tags-from=-", outPath)
	importTags.Stdin = &tagData
	if err := importTags.Run(); err != nil {
		return fmt.Errorf("failed to import tags: %w", err)
	}

	// 2. Set Verdi Pitch Engine custom tags
	setTags := execCommandContext(ctx, "metaflac",
		"--set-tag=ENCODED_BY=VerdiPitchEngine",
		"--set-tag=PITCH_SHIFT=432Hz",
		"--set-tag=VERSION=432 Hz",
		outPath)
	if err := setTags.Run(); err != nil {
		return fmt.Errorf("failed to set custom tags: %w", err)
	}

	// 3. Strip ffmpeg-generated pictures and copy original picture block
	removePic := execCommandContext(ctx, "metaflac", "--remove", "--block-type=PICTURE", "--dont-use-padding", outPath)
	_ = removePic.Run() // Ignore errors if no picture exists

	var picData bytes.Buffer
	exportPic := execCommandContext(ctx, "metaflac", "--export-picture-to=-", inPath)
	exportPic.Stdout = &picData
	if err := exportPic.Run(); err == nil && picData.Len() > 0 {
		importPic := execCommandContext(ctx, "metaflac", "--import-picture-from=-", outPath)
		importPic.Stdin = &picData
		_ = importPic.Run() // Ignore error if importing picture fails
	}

	return nil
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func ProcessFile(ctx context.Context, inPath, outPath string, bar *mpb.Bar) error {
	depth, _ := getBitDepth(inPath)
	durationSec, _ := getDuration(inPath)
	sampleRate, _ := getSampleRate(inPath)
	totalMicrosec := int64(durationSec * 1000000)

	if bar != nil && totalMicrosec > 0 {
		bar.SetTotal(totalMicrosec, false)
	}

	// Use a temporary file to avoid media scanners (e.g. Roon) locking the file while it's being written.
	tmpOutPath := outPath + ".tmp"

	// High-fidelity Time-Scale Modification (TSM) pitch-shift using librubberband.
	args := buildFFmpegArgs(inPath, tmpOutPath, depth, sampleRate)

	cmd := execCommandContext(ctx, "ffmpeg", args...)

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
		os.Remove(tmpOutPath) // cleanup partial file
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderr.String())
	}

	// Inject 1:1 metadata parity using metaflac
	if err := inheritMetadata(ctx, inPath, tmpOutPath); err != nil {
		os.Remove(tmpOutPath)
		return fmt.Errorf("metadata inheritance failed: %w", err)
	}

	// Preserve original timestamps on the tmp file BEFORE renaming it.
	// If we Chtimes after Rename, Roon's inotify watcher will scan the file the instant it is renamed
	// and will read the new mtime before we can restore it, falsely flagging it as a "Newly Added" album.
	if info, err := os.Stat(inPath); err == nil {
		_ = os.Chtimes(tmpOutPath, info.ModTime(), info.ModTime())
	}

	// Atomically rename the finished tmp file to the final output name
	if err := os.Rename(tmpOutPath, outPath); err != nil {
		return fmt.Errorf("failed to rename tmp file: %w", err)
	}

	if bar != nil && totalMicrosec > 0 {
		bar.SetCurrent(totalMicrosec)
	}

	return nil
}
