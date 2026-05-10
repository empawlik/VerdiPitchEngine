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
body_hash: 73cf9b4caf17a9a1
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

The application is deployed via Docker, commonly on QNAP or Synology NAS environments. The engine now utilizes a `docker-compose.yml` file to initialize an interactive Container Station environment.

```bash
docker build -t empawlik/verdi-pitch-engine:latest .
docker-compose up -d
```
### Dynamic Target Deployment

You can specify a custom NAS directory to process instead of the default `/share/DataVol1/Music` by setting the `VERDI_TARGET_DIR` environment variable before running `mage deploy`.

If you only specify the target directory (In-Place Mode), the system will automatically rename your target directory to `[440 Hz]` and output the newly pitch-shifted files into a freshly created folder with the `[432 Hz]` suffix!

```bash
# In-Place Mode: Will rename the album folder to "The North Borders [440 Hz]" and output to "The North Borders [432 Hz]"
VERDI_TARGET_DIR="/share/DataVol1/Music/The North Borders" mage deploy
```

You can optionally override both the input and output directories:

```bash
# Explicit Output Mode
VERDI_TARGET_DIR="/share/DataVol1/Music_New" VERDI_OUT_DIR="/share/DataVol1/Music_New_432" mage deploy
```

### Orchestration Scripts

Once the container is deployed, you execute the conversion logic directly from within the container environment (via Container Station or SSH).

#### Interactive Single-Album Processing
To process a single album immediately, use `verdi-process`. This maps relative paths automatically and handles version tag generation.

```bash
verdi-process "Artist/Nightmares on Wax/Smokers Delight"
```

#### Headless Batch Processing
To queue up an entire root folder of albums for unattended processing, use `verdi-batch`. This script intelligently filters out already processed albums and executes them alphabetically.

```bash
# Process a maximum of 10 pending albums from the Artist directory
verdi-batch "Artist" 10

# Process ALL pending albums in the directory (limit omitted or set to 'all')
verdi-batch "Artist"
```
