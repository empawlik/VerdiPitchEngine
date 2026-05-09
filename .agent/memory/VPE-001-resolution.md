---
project_name: "VerdiPitchEngine"
version: 0.1.0
status: "active"
priority: "high"
dev_stage: "development"
agent_role: "Core-Context"
agent_weight: 4.0
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell", "ffmpeg"]
dependencies: []
created: "2026-05-09"
updated: 2026-05-09
body_hash: 86063ffbe9b79f33
tags: [dev-asset, memory, technical-resolution]
---

# Technical Resolution: VPE-001 Bootstrap

## Problem Space
The Verdi Pitch Engine requires a high-assurance, fault-tolerant build architecture capable of scaling across QNAP/Unraid NAS hardware while rigidly complying with Antigravity 1000-KEYS documentation standards. The legacy scripts and documentation needed immediate replacement with type-safe Go automation and deterministic workflow wrappers.

## Architectural Decisions & Resolutions

### 1. Concurrency Model (Worker Pools)
- **Decision:** Implemented a fixed-size goroutine worker pool (`RunPool`) instead of unbounded concurrency.
- **Rationale:** Processing lossless FLAC files with FFmpeg is highly CPU/I/O bound. Unbounded goroutines would saturate NAS resources. The pool size is governed by `runtime.NumCPU()` but can be forcefully overridden via the `VERDI_WORKERS` environment variable for zero-trust environments.

### 2. Mage Build Standardization
- **Decision:** All project lifecycle tasks (linting, tests, static analysis, vulnerability checks) are orchestrated through `magefile.go`.
- **Rationale:** Eliminates bash portability issues (macOS vs Alpine). Hard-codes test coverage floors (80%) preventing pipeline progression if code lacks rigor. Incorporates `lint-markdown` dynamically into the CI loop.

### 3. Commit Wrappers
- **Decision:** Created `.agent/tmp/` structure to satisfy zero-trust sandbox rules during agentic commits, routing `gen-commit` operations through `scripts/gm-commit` to guarantee GPG signature compliance.
