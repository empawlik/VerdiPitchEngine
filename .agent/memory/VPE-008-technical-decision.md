---
project_name: VerdiPitchEngine
version: 0.1.0
status: Active
priority: Medium
dev_stage: Production
agent_role: Core-Context
agent_weight: 4
asset_scope: Global
platform: CLI
tech_stack: [markdown]
dependencies: []
created: "2026-05-11"
updated: 2026-05-11
tags: [dev-asset, memory, vpe-008, technical]
body_hash: bafa59f9af38919f
---

# Technical Decision: VPE-008 Roon Compliance & QNAP Parsing

**Date:** 2026-05-11

## 1. Roon "Recent Activity" Mitigation
**Problem:** Applying new metadata tags to pitch-shifted media reset the file's modification timestamp to the present day, causing the Roon scanner to incorrectly flag archived historical albums as "Recently Added".
**Decision:** We adopted a strict requirement to use `metaflac --preserve-modtime` for all metadata enforcement steps. 
**Rationale:** By preserving the original 2012 timestamps directly onto the new 432 Hz output FLAC files, Roon interprets the media as old imports, keeping the recent activity feed perfectly clean. To update Roon's database with the new 432 Hz tag, users simply trigger a manual "Re-scan Album".

## 2. Zero-Trust QNAP Directory Pruning
**Problem:** During batch traversal on QNAP NAS hardware, the OS dynamically generates hidden directories (`.@__thumb`, `.@__desc`, `.AppleDouble`) that contain non-audio thumbnails falsely identified by standard recursive scanners. This resulted in fatal FFmpeg failures.
**Decision:** Hardcoded exclusion paths within both the Go binary and bash orchestration logic.
**Implementation:** 
- `walker.go` actively checks `strings.HasPrefix(info.Name(), ".@__")` and triggers `filepath.SkipDir`.
- Bash `find` arrays were appended with POSIX compliance: `-type d \( -name ".@__*" -o -name ".AppleDouble" \) -prune`.
