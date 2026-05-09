---
project_name: VerdiPitchEngine
version: 1.1.0
status: active
priority: high
dev_stage: "Production"
agent_role: Feature-Spec
agent_weight: 3
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell"]
dependencies: []
created: "2026-04-26"
updated: 2026-05-09
body_hash: 0ee21082a28497d0
tags: [execution, idempotency, secrets, protocol-adapters, side-effects, go]
---

# 400-ACTION: The Action Layer (Execution)

## I. The Execution Protocol
The Action Layer (`/pkg/executor`) is the only component with authority to create side effects. Its role is to realize a **Validated Intent** (e.g., via gRPC clients or REST API callers).

## II. Mandatory Execution Principles
1.  **Strict Idempotency:** Every operation must be indexed by a unique `IntentID`. Re-submitting the same ID must never result in duplicate side effects.
2.  **Secret Isolation:** This is the *only* layer permitted to interface with KMS or environment secrets.
3.  **One-Fail-All-Fail:** Multi-step transactions must be atomic.
4.  **Protocol Agnosticism (Upward):** The executor speaks the internal domain language to the Kernel, but speaks the protocol-specific language (e.g. gRPC, REST) to the external venue.

## III. Structural Requirements
*   **Adapters:** Use the Adapter Pattern to wrap external SDKs (e.g., the standard `http.Client`, gRPC clients).
*   **Retries:** Implement exponential backoff for transient network errors.
*   **Receipts:** Every execution must produce an **Execution Receipt**.

## IV. Technical Standards
*   **Timeout Management:** Every external call must have a `context.Context` with a deadline/timeout.
*   **Concurrency:** Use standard Go concurrency patterns (goroutines, channels) with care to ensure ordering where needed.

## V. Prohibited Patterns
*   **No Decision Logic:** The Action Layer does not decide "if" or "how much."
*   **Leaking Secrets:** Credentials must never leave the scope of the executor.
*   **State Assumption:** Never assume the external state has remained the same; verify locally if possible.
