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
updated: 2026-05-09
tags: [audio, dsp]
body_hash: 4f2adedd0181e4b5
---
# Verdi Pitch Engine

**Verdi Pitch Engine** is a containerized batch-conversion utility designed to run directly on enterprise NAS hardware (QNAP/Synology/Unraid). It processes lossless audio libraries, mathematically shifting them from standard 440 Hz tuning down to 432 Hz (-31.76 cents).

### Architectural Motivation
In high-end residential audio environments (Roon, BluOS, Plex), applying real-time 432 Hz DSP introduces network latency, consumes continuous Roon Core CPU cycles, and often requires expensive proprietary hardware to mitigate electrical noise (e.g., 432 EVO). 

*Verdi Pitch Engine* solves this through **Data Gravity**—moving the compute to the storage. By utilizing a transient Docker container running `ffmpeg` and the `librubberband` audio filter directly on the NAS, we pre-compute the acoustic math and save it to disk.

### Core Features
* **Zero Real-Time Overhead:** FLAC files are pre-processed, allowing network endpoints (Buchardt, BluOS, Denon) to stream them bit-perfectly without real-time DSP jitter.
* **Bit-Perfect Metadata:** Ensures high-resolution album art, Roon tags, and ReplayGain data are flawlessly cloned to the output files.
* **Non-Destructive:** Mounts the source directory as Read-Only (`:ro`), making it structurally impossible to corrupt your original 440 Hz library.
* **Topology Aware:** Recursively scans the source directory and perfectly replicates your Artist/Album folder tree in the destination directory.

### Hardware Requirements
- QNAP, Synology, or Unraid NAS capable of running Docker containers.
- Multi-core CPU recommended (Intel x86_64 architecture preferred). Default processing utilizes 4 workers.

### Execution
Run the following `docker run` command, explicitly mapping your source and destination directories.

**CRITICAL:** Ensure the input directory is mounted as Read-Only (`:ro`) to prevent any accidental library corruption.

```bash
docker run -d \
  --name verdi-pitch-engine \
  -v /share/DataVol1/Music:/music_in:ro \
  -v /share/DataVol1/Music_432:/music_out:rw \
  -e VERDI_WORKERS=4 \
  empawlik/verdi-pitch-engine:latest
```
