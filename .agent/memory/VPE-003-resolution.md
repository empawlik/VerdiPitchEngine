---
project_name: "VerdiPitchEngine"
version: 0.3.0
status: "active"
priority: "medium"
dev_stage: "development"
agent_role: "Core-Context"
agent_weight: 4.0
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell", "ffmpeg"]
dependencies: []
created: "2026-05-21"
updated: 2026-05-21
body_hash: efd3117a919c500b
tags: [dev-asset, memory, technical-resolution]
---

# Technical Resolution: VPE-003 Telemetry AI Integration

## Problem Space
The observation/telemetry layer in VerdiPitchEngine requires integration with the Vertex AI system to support high-assurance telemetry checkups. In addition, our static analysis flagged potential resource leaks and missing error propagation in the core FFmpeg processing and file traversal logic.

## Architectural Decisions & Resolutions

### 1. Mock-Tested Vertex AI Client Design
- **Decision:** Abstracted the Google Vertex AI SDK's client and generative model behaviors behind narrow interfaces (`client` and `generativeModel`) in `pkg/ai/client.go`.
- **Rationale:** The Vertex AI SDK utilizes gRPC transport under the hood and typically makes actual network calls. Direct unit testing would be fragile and require valid live service account credentials. The wrapper interfaces allow 100% test coverage using custom mock implementations (`mockClient` and `mockGenerativeModel`) without network calls.

### 2. Service Account Locking
- **Decision:** Required explicit configurations for `ProjectID` and `CredentialsPath` inside `NewEnterpriseClient` and enforced `option.WithAuthCredentialsFile(option.ServiceAccount, cfg.CredentialsPath)`.
- **Rationale:** Restricting client construction to authenticated files prevents the application from fallback searching default credentials, which could lead to utilizing a developer's personal `gcloud` local settings, violating Zero-Trust environment mandates.

### 3. Ffmpeg Scanner Error Checking
- **Decision:** Added explicit checking for `scanner.Err()` immediately after the `scanner.Scan()` loop finishes processing the FFmpeg output.
- **Rationale:** In standard Go, `bufio.Scanner` returns `nil` from `Scan()` on EOF or when an read error occurs. Failing to call `Err()` hides silent I/O truncation or descriptor closures, potentially resulting in partial or corrupted conversions.

### 4. Close & Remove Defer Verification
- **Decision:** Wrapped all `defer` close and remove operations in anonymous functions or explicitly discarded errors, ensuring no untracked return values.
- **Rationale:** Ensures clean file descriptor lifecycles across the worker pools and prevents descriptor leaks during heavy concurrent batch jobs.
