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
body_hash: 888901202fa31a17
---
# Verdi Pitch Engine

**Verdi Pitch Engine** is a containerized batch-conversion utility designed to run directly on enterprise NAS hardware (QNAP/Synology/Unraid). It processes lossless audio libraries, mathematically shifting them from standard 440 Hz tuning down to 432 Hz (-31.76 cents).

### Why 432 Hz? (Verdi's A)
The engine is named in honor of the Italian composer **Giuseppe Verdi**, who strongly advocated for the 432 Hz tuning standard (historically known as "Verdi's A"). According to widespread acoustic theories—and some popular conspiracy theories regarding the standardization of 440 Hz in the 20th century—440 Hz is considered unnaturally dissonant to human biology and consciousness. In contrast, 432 Hz is believed to be mathematically consistent with the patterns of the universe, resonating organically with human biophysics, the golden ratio, and the natural world. This engine seeks to restore digital music libraries to this natural, harmonious state.

### Architectural Motivation
In high-end residential audio environments (Roon, BluOS, Plex), applying real-time 432 Hz DSP introduces network latency, consumes continuous Roon Core CPU cycles, and often requires expensive proprietary hardware to mitigate electrical noise (e.g., 432 EVO). 

*Verdi Pitch Engine* solves this through **Data Gravity**—moving the compute to the storage. By utilizing a persistent Docker container running native `ffmpeg` filters (`rubberband` Time-Scale Modification) directly on the NAS, we pre-compute the acoustic math and save it to disk. The container acts as a batch processing service, preserving zero-trust computing rules while maximizing NAS resources.

### Core Features
* **High-Fidelity TSM (Time-Scale Modification):** Replaces legacy resampling techniques with `librubberband`, executing duration-preserving pitch shifting. Tracks are lowered to 432 Hz without altering their original tempo or acoustic timing.
* **Dynamic Bit-Depth & Sample Rate Preservation:** The engine mathematically detects the original bit depth and sample rate. High-resolution studio masters (e.g., 24-bit/96kHz, 192kHz) are natively preserved without downsampling, ensuring Audiophile-grade precision isn't crushed to standard 16-bit/44.1kHz during processing.
* **MQA Purification:** Master Quality Authenticated (MQA) files rely on proprietary, fragile high-frequency data folding embedded in the 24-bit noise floor. Time-scale modification fundamentally recalculates the waveform, naturally destroying the proprietary MQA layer. The engine outputs a pure, standard 24-bit 432 Hz FLAC, effectively freeing your music from MQA lock-in and hardware decoding requirements.
* **Zero Real-Time Overhead:** FLAC files are pre-processed, allowing network endpoints (Buchardt, BluOS, Denon) to stream them bit-perfectly without real-time DSP jitter.
* **Native Roon Integration:** Injects the official `VERSION=432 Hz` Vorbis metadata tag natively into both the new and original files. Roon will recognize the pitch-shifted tracks as a distinct release edition and explicitly badge the album in its UI!
* **Automated Sidecar Migration:** Seamlessly detects and duplicates all non-FLAC sidecar assets (like `.lrc` synced lyrics, `.pdf` digital booklets, and `.jpg` album covers) from the 440 Hz backup into the new 432 Hz output directory, guaranteeing a flawless presentation in your digital library.
* **Topology Aware (Recursive):** Recursively scans the source directory. You can point the engine at a single Album folder, an Artist root folder, or your entire Music library; it will perfectly replicate your nested hierarchy in the destination directory.
* **Real-Time CLI Progress Visualization:** Employs the `mpb/v8` multi-progress bar library to give you rich, terminal-based feedback. The CLI explicitly displays global batch completion percentage alongside real-time microsecond-level progression for each individual active worker thread.
* **Contextual Execution Safety:** To protect your enterprise NAS from catastrophic hang-states during large directory conversions, each lossless track conversion is wrapped in a strict 15-minute `context.WithTimeout` termination bound, ensuring silent `ffmpeg` execution failures or edge-case corrupted inputs never indefinitely deadlock the main execution pool.

### Hardware Requirements
- QNAP, Synology, or Unraid NAS capable of running Docker containers.
- Multi-core CPU recommended (Intel x86_64 architecture preferred). Default processing utilizes 4 workers.

### Execution & Deployment

Deploy the engine via the Antigravity Mage pipeline to your NAS infrastructure. Once deployed as a persistent container, you interact with it directly through your NAS Container Station terminal via the interactive `verdi-process` CLI dashboard.

#### 1. Deployment
Deploy the container to your remote QNAP/NAS host:
```bash
mage deploy
```

#### 2. Processing via Container Station
Open the **Execute** terminal for the `verdi-pitch-engine` container inside your NAS UI and use the interactive dashboard.

You can provide full QNAP absolute paths:
```bash
verdi-process "/share/Multimedia/Audio/Music/Artist/Nightmares on Wax/Smokers Delight"
```

Or utilize the ultra-fast relative pathing (automatically maps to your music root):
```bash
verdi-process "Artist/Nightmares on Wax/Smokers Delight"
```

The orchestration wrapper will:
1. Backup your master files to an appended `[440 Hz]` directory.
2. Enforce the `440 Hz` version tag on the master tracks.
3. Perform duration-preserving 432 Hz pitch shifting via the Go engine.
4. Auto-migrate your cover art, PDFs, and lyrics files to the newly minted `[432 Hz]` directory.

#### 3. Headless Batch Execution
If you want to process massive chunks of your library in an automated queue, use the new `verdi-batch` orchestration script.

```bash
verdi-batch "Artist" 10
```

The batch engine will natively:
* Scan the `Artist` directory for unprocessed FLAC albums.
* Filter out and strictly ignore albums that have already generated a `[432 Hz]` target or are already tagged as `[440 Hz]`.
* Sort the pending list alphabetically.
* Process the top 10 pending albums sequentially.

If you omit the numeric limit, it will default to processing all pending albums sequentially:
```bash
verdi-batch "Artist"
```
