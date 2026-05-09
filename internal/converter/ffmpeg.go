package converter

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

// getSampleRate uses ffprobe to extract the sample rate of the input audio file.
func getSampleRate(filePath string) (int, error) {
	cmd := exec.Command("ffprobe",
		"-v", "error",
		"-show_entries", "stream=sample_rate",
		"-of", "default=noprint_wrappers=1:nokey=1",
		filePath,
	)
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		return 0, fmt.Errorf("ffprobe failed: %w", err)
	}

	rateStr := strings.TrimSpace(out.String())
	// Some formats might have multiple streams returning multiline rates, grab the first one
	rateStr = strings.Split(rateStr, "\n")[0]
	rate, err := strconv.Atoi(rateStr)
	if err != nil {
		return 0, fmt.Errorf("failed to parse sample rate '%s': %w", rateStr, err)
	}
	return rate, nil
}

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func ProcessFile(inPath, outPath string) error {
	rate, err := getSampleRate(inPath)
	if err != nil {
		return err
	}

	// Pure math pitch-shift (slows down tempo by 1.8%)
	newRate := rate * 432 / 440
	filter := fmt.Sprintf("asetrate=%d,aresample=%d", newRate, rate)

	cmd := exec.Command("ffmpeg",
		"-v", "warning",
		"-i", inPath,
		"-af", filter,
		"-map_metadata", "0",
		"-c:a", "flac",
		"-sample_fmt", "s16",
		"-dither_method", "triangular_hp",
		"-compression_level", "12",
		outPath,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderr.String())
	}

	return nil
}
