---
project_name: VerdiPitchEngine
version: 1.1.0
status: active
priority: high
dev_stage: Production
agent_role: Core-Context
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-09
body_hash: 9aeff7b6f69caa58
tags: [audit, observability, forensics, immutability, replayability, go]
---

# 600-AUDIT: The Forensic Layer (Black Box)

## I. The Audit Mandate
The system must maintain an immutable, append-only record of every decision and action. This record serves as the "Source of Truth" for system recovery and compliance.

## II. Mandatory Logging Principles
1.  **Immutability:** Once an audit log entry is written, it must be impossible to alter.
2.  **Atomicity (Log-Before-Act):** The **Execution Kernel** must successfully commit the **Intent** and **ValidationResult** to the audit log *before* dispatching the command to the **Action Layer**.
3.  **Contextual Linking:** Every log entry must be linked via a unique `IntentID`.
4.  **No PII/Secrets:** Audit logs must never contain plaintext private keys or sensitive PII.

## III. The Audit Lifecycle

| Event Type | Log Content | Trigger Layer |
| :--- | :--- | :--- |
| **Intent Created** | Proposed parameters, strategy ID, timestamp. | Intelligence |
| **Policy Evaluated**| Allow/Reject status, rule IDs triggered. | Validation |
| **Action Dispatched**| Execution payload, destination venue. | Action |
| **Action Receipt** | External TX hash, final status, latency. | Action |

## IV. Replayability Standards
*   **Deterministic Replay:** Logs must allow a developer to re-run the **Intelligence Layer** and produce the identical result.
*   **Structured Logs:** All logs must be emitted in a machine-readable format (e.g., JSON using the standard library `slog` package).

## V. Prohibited Patterns
*   **Selective Logging:** "Successful operations only" logging is forbidden.
*   **Local Buffering:** Never buffer audit logs in volatile memory without a flush guarantee.
