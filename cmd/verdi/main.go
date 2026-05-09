package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/empawlik/verdi-pitch-engine/internal/converter"
	"github.com/empawlik/verdi-pitch-engine/internal/fs"
)

func main() {
	var workers int
	var inDir string
	var outDir string

	flag.IntVar(&workers, "workers", 4, "Number of concurrent workers")
	flag.StringVar(&inDir, "in", "/music_in", "Input directory containing FLAC files")
	flag.StringVar(&outDir, "out", "/music_out", "Output directory for pitch-shifted files")
	flag.Parse()

	// Optionally override via environment variables if set
	if envWorkers := os.Getenv("VERDI_WORKERS"); envWorkers != "" {
		if _, err := fmt.Sscanf(envWorkers, "%d", &workers); err != nil {
			log.Printf("Warning: failed to parse VERDI_WORKERS: %v", err)
		}
	}
	if envIn := os.Getenv("VERDI_IN"); envIn != "" {
		inDir = envIn
	}
	if envOut := os.Getenv("VERDI_OUT"); envOut != "" {
		outDir = envOut
	}

	if workers < 1 {
		workers = 1
	}

	log.Printf("Starting Verdi Pitch Engine")
	log.Printf("Input directory: %s", inDir)
	log.Printf("Output directory: %s", outDir)
	log.Printf("Workers: %d", workers)

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
