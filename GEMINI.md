---
project_name: "VerdiPitchEngine"
version: 0.3.0
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
updated: 2026-05-10
body_hash: bca0b33e0fe10c7c
tags: [dev-asset, audio, dsp]
---
# VerdiPitchEngine Agent Orchestration (GEMINI.md)

This file defines the base rules, orchestration paths, and interaction models for AI agents (e.g., Antigravity / Gemini) operating within the **VerdiPitchEngine** repository.

## 🏗️ Architecture & Context 

VerdiPitchEngine is a containerized Go application designed to perform high-fidelity batch Time-Scale Modification (TSM) on lossless audio files (pitch-shifting from 440 Hz to 432 Hz). It operates heavily on enterprise NAS hardware leveraging `ffmpeg` and `librubberband`.

Agents modifying this repository must adhere strictly to the following constraint pillars:

1. **Go Ecosystem Mandate:** VerdiPitchEngine is strictly a Go 1.26+ binary. The use of strict static typing is mandatory. Goroutines must be orchestrated safely using idiomatic WaitGroups, avoiding race conditions during batch concurrency.
2. **Zero-Trust Audio Processing:** The application reads from a strictly read-only (`:ro`) source volume to guarantee the absolute safety of the user's master FLAC library. Writes are directed to a separate volume.
3. **Data Gravity:** Compute is co-located with storage on NAS hardware (QNAP/Synology/Unraid) to prevent network latency during high-bandwidth media processing.

## 🧠 Knowledge Architecture

To govern AI behavior predictably, the repository relies on a structured knowledge path paradigm:

- **Hot Path (`.agent/rules/`):** Contains "Proactive" rules that are mandatory, high-priority constraints. You must check these directives before modifying any underlying structure.
- **Indexed Path (`.agent/docs/`):** General developmental documentation, Gap Analyses, and architectural plans.

## 🛠️ Execution & Automation Rules

1. **Commit Standardization:** Commits synthesized by the agent must strictly adhere to the `020-GIT` mathematical commit structures employed universally within the Antigravity organization.
2. **Mage Build System:** All lifecycle operations, linting, tests, and compliance checks MUST be routed through the `magefile.go` targets.
3. **1000-KEYS Compliance:** All Markdown documentation must retain and update valid `1000-KEYS` YAML frontmatter and `body_hash` validation signatures. Use `mage fix-markdown` to enforce this.

## 🤖 Operational Goal

- **Current Focus:** Finalize the Go-native infrastructure migration and verify the `ffmpeg` pipeline behavior for absolute bit-perfect audio conversion before v0.3.0 release tagging.
