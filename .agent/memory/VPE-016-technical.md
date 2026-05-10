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
tech_stack: ["Go", "Shell", "ffmpeg", "Docker"]
dependencies: []
created: "2026-05-10"
updated: 2026-05-10
tags: [dev-asset, memory, technical, architecture, docker]
body_hash: 1f67bae32f4314fe
---

# Technical Memory: VPE-016 (Docker PID Namespace Isolation)

## Context
During the hardening phase of the Verdi Pitch Engine, the security guardrail that actively denies operation while an indexing media server (Roon, Plex, MinimServer) is running was found to be brittle. It was originally implemented at the `agent-proxy.sh` deployment level. This meant manual container restarts via the QNAP App Center completely bypassed the guardrail.

## Resolution
To establish a truly airtight Zero-Trust execution boundary, the process-tree inspection (`ps | grep`) logic was migrated directly into the interactive execution scripts (`verdi-process` and `verdi-batch`) running *inside* the container.

### Architectural Change
1. **Host PID Mapping**: Modified `docker-compose.yml` to run the container with `pid: "host"`.
   - **Why**: By default, Docker isolates container processes. `ps` running inside an isolated container will only return processes instantiated within that container. To view the underlying QNAP host process tree, the container must break out of the PID namespace isolation via `pid: "host"`.
2. **Internal Script Denial**: The `verdi-batch` and `verdi-process` scripts now natively analyze the host environment and perform hard-aborts if an illegal process is running. This effectively couples the safety check to the application execution rather than the deployment lifecycle, sealing the execution vulnerability.
