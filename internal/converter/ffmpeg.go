package converter

import (
	"bytes"
	"fmt"
	"os/exec"
)

// ProcessFile invokes FFmpeg to pitch-shift the input FLAC file to the output FLAC file.
func ProcessFile(inPath, outPath string) error {
	// ffmpeg -v warning -i <input_file> -af "rubberband=pitch=432/440" -map_metadata 0 -c:a flac <output_file>
	cmd := exec.Command("ffmpeg",
		"-v", "warning",
		"-i", inPath,
		"-af", "rubberband=pitch=432/440",
		"-map_metadata", "0",
		"-c:a", "flac",
		outPath,
	)

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg failed: %w, stderr: %s", err, stderr.String())
	}

	return nil
}
