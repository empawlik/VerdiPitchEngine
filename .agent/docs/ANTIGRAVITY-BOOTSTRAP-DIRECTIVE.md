---
project_name: VerdiPitchEngine
version: 0.1.0
status: active
priority: high
dev_stage: development
agent_role: lead_engineer
agent_weight: 1.0
asset_scope: backend
platform: qnap
tech_stack: [go, ffmpeg]
dependencies: []
created: 2026-05-09
updated: 2026-05-10
tags: [audio, dsp]
body_hash: 6f2ef5515680e0f9
---
# ANTIGRAVITY BOOTSTRAP DIRECTIVE
# Project: Verdi Pitch Engine
# Target Environment: QNAP TS-453A (Intel x86_64) via Container Station (Docker)
# License Constraint: MIT

## I. Context & Architectural Mandate
You are acting as the lead software engineer developing a high-performance, containerized audio DSP utility. The application must be written in Go (Golang) and will orchestrate `ffmpeg` (with `librubberband`) to perform batch Time-Scale Modification (TSM). 
The goal is to mathematically pitch-shift standard 440 Hz lossless FLAC files down to 432 Hz (-31.76 cents) while preserving bit-perfect metadata and identical directory structures. 

This will run on an enterprise NAS. Resource efficiency, concurrency, and absolute data safety (zero-trust storage) are non-negotiable.

## II. Project Initialization phase
1. Initialize the Go module: `go mod init github.com/[YOUR-USERNAME]/verdi-pitch-engine`
2. Create the standard Go project layout:
   - `/cmd/verdi/main.go` (CLI entrypoint)
   - `/internal/converter/worker.go` (Goroutine worker pool logic)
   - `/internal/converter/ffmpeg.go` (Subprocess execution logic)
   - `/internal/fs/walker.go` (Directory traversal and replication logic)
3. Generate the `LICENSE` file (MIT).

## III. Core Functional Requirements
Implement the Go application with the following specifications:
* **Concurrency:** Implement a parameterized Goroutine worker pool. The number of concurrent workers should default to `4` (matching the QNAP CPU cores) but be configurable via CLI flags or environment variables.
* **Directory Walker:** Recursively scan the `/music_in` directory for `.flac` files. Replicate the exact relative folder tree in the `/music_out` directory.
* **Idempotency (Resumable):** Before processing, check if the output file already exists. If it does, skip it and log it. This allows the process to be interrupted and safely resumed.
* **FFmpeg Subprocess:** Use `os/exec` to invoke FFmpeg for each file. 
  - Required FFmpeg arguments: `-v warning -i <input_file> -af "rubberband=pitch=432/440" -map_metadata 0 -c:a flac <output_file>`
* **Logging & Progress:** Implement clean `stdout` logging detailing files processed, errors encountered, and overall worker pool progress.

## IV. Dockerization Strategy (Multi-Stage Build)
Generate a `Dockerfile` optimized for QTS Container Station. 
* **Stage 1 (Builder):** Use `golang:1.21-alpine` (or newer) to statically compile the Go binary.
* **Stage 2 (Runtime):** Use `ubuntu:22.04` (or similar Debian base) to ensure high-quality `ffmpeg` and `librubberband` packages are available.
  - Run `apt-get update && apt-get install -y ffmpeg`
  - Copy the compiled Go binary from the builder stage.
  - Set the entrypoint to the Go binary.

## V. Documentation
Generate a `README.md` that includes:
1. The architectural motivation (Data Gravity vs. Real-time hardware DSP).
2. Hardware requirements.
3. The exact `docker run` command required to execute the container, explicitly enforcing Read-Only (`:ro`) mounts for the source directory to prevent accidental library corruption. 
   - Example: `-v /share/DataVol1/Music:/music_in:ro -v /share/DataVol1/Music_432:/music_out:rw`

## VI. Execution
Please acknowledge these constraints and begin executing **Phase II** and **Phase III**. Generate the Go codebase step-by-step, ensuring strict error handling (no silent failures).