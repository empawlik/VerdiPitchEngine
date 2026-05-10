package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/empawlik/verdi-pitch-engine/internal/converter"
	"github.com/empawlik/verdi-pitch-engine/internal/fs"
)

var (
	Version = "dev"
	Build   = "unknown"
)

func main() {
	var workers int
	var inDir string
	var outDir string

	flag.IntVar(&workers, "workers", 4, "Number of concurrent workers")
	flag.StringVar(&inDir, "in", "/music_in", "Input directory containing FLAC files")
	flag.StringVar(&outDir, "out", "/music_out", "Output directory for pitch-shifted files")
	flag.Parse()

	// Check if specific flags were explicitly provided on the CLI
	isFlagPassed := func(name string) bool {
		found := false
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				found = true
			}
		})
		return found
	}

	// Optionally override via environment variables if set and flag was not explicitly provided
	if envWorkers := os.Getenv("VERDI_WORKERS"); envWorkers != "" && !isFlagPassed("workers") {
		if _, err := fmt.Sscanf(envWorkers, "%d", &workers); err != nil {
			log.Printf("Warning: failed to parse VERDI_WORKERS: %v", err)
		}
	}
	if envIn := os.Getenv("VERDI_IN"); envIn != "" && !isFlagPassed("in") {
		inDir = envIn
	}
	if envOut := os.Getenv("VERDI_OUT"); envOut != "" && !isFlagPassed("out") {
		outDir = envOut
	}

	if workers < 1 {
		workers = 1
	}

	fmt.Printf("\n🎶 \033[1;36mVERDI PITCH ENGINE\033[0m 🎶\n")
	fmt.Printf("\033[1;30m================================================\033[0m\n")
	fmt.Printf("🏷️  \033[1;34mVersion:\033[0m    %s (Build: %s)\n", Version, Build)
	fmt.Printf("🎯 \033[1;34mPurpose:\033[0m    High-fidelity lossless batch pitch-shifting (440 Hz -> 432 Hz)\n")
	fmt.Printf("📂 \033[1;34mInput Dir:\033[0m  %s\n", inDir)
	fmt.Printf("📂 \033[1;34mOutput Dir:\033[0m %s\n", outDir)
	fmt.Printf("⚙️  \033[1;34mWorkers:\033[0m    %d\n", workers)
	fmt.Printf("\033[1;30m================================================\033[0m\n\n")

	log.Printf("Starting initialization sequence...")

	tasks, err := fs.WalkAndCollect(inDir, outDir)
	if err != nil {
		log.Fatalf("Failed to collect tasks: %v", err)
	}

	if len(tasks) == 0 {
		log.Printf("No FLAC files found in %s", inDir)
		return
	}

	log.Printf("Found %d files to process.", len(tasks))

	converter.RunPool(tasks, workers)
}
