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
body_hash: cea26119fe66e4c6
---
# Verdi Pitch Engine Runbook

This runbook outlines the operational commands for developing, testing, and deploying the Verdi Pitch Engine.

## Build and Run

To build and run the go binary locally:
```bash
go build -o bin/verdi-pitch-engine ./cmd/verdi
./bin/verdi-pitch-engine -in /path/to/music_in -out /path/to/music_out -workers 4
```

## Mage Build System

This project uses [Mage](https://magefile.org/) as a build tool to automate compliance and linting tasks. 

Ensure you have Mage installed:
```bash
go install github.com/magefile/mage@latest
```

Available targets:
- `mage fmt`: Formats code via `go fmt`.
- `mage vet`: Analyzes code via `go vet`.
- `mage vulncheck`: Scans dependencies for known vulnerabilities via `govulncheck`.
- `mage lint`: Lints Go code via `golangci-lint`.
- `mage test`: Runs tests and generates a coverage report (`coverage.out` and `coverage.html`).
- `mage ci-test`: Runs tests optimized for CI environments.
- `mage check-coverage`: Enforces test coverage minimums (80%) using the custom verification script.
- `mage lint-markdown`: Lints Markdown files for `1000-KEYS` compliance using the central script.
- `mage fix-markdown`: Automatically adds or updates the `1000-KEYS` YAML frontmatter and `body_hash` on all `.md` files.
- `mage check`: Runs the full pipeline: formatting, vetting, linting (Go & Markdown), vulnerability checks, and coverage verification.

## Deployment

The application is deployed via Docker, commonly on QNAP or Synology NAS environments.

```bash
docker build -t empawlik/verdi-pitch-engine:latest .

docker run -d \
  --name verdi-pitch-engine \
  -v /share/DataVol1/Music:/music_in:ro \
  -v /share/DataVol1/Music_432:/music_out:rw \
  -e VERDI_WORKERS=4 \
  empawlik/verdi-pitch-engine:latest
```
