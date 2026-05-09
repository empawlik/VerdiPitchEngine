---
project_name: VerdiPitchEngine
version: 1.1.0
status: active
priority: critical
dev_stage: "Production"
agent_role: Feature-Spec
agent_weight: 3
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell"]
dependencies: []
created: "2026-04-26"
updated: 2026-05-09
body_hash: e4ef92d260554ab6
tags: [safety, risk, guardrails, fail-closed, policy, go]
---

# 300-VALID: The Validation Layer (Guardrails)

## I. The Policy Authority
The Validation Layer (`/pkg/guardrail`) is the final arbiter of system safety. It evaluates an **Intent** against hard constraints, compliance rules, and risk parameters.

## II. Mandatory Safety Principles
1.  **Fail-Closed Logic:** If a rule evaluation is inconclusive, times out, or encounters an error, the result must be an absolute **REJECT**.
2.  **Immutability:** Validation rules must not modify the **Intent**. They provide a binary `ALLOW` or `REJECT` response.
3.  **Independence:** Guardrails must be independent of the Strategy. They do not share logic with the Intelligence Layer.
4.  **Zero Trust:** Never assume the Intelligence Layer has performed risk checks. Re-verify everything.

## III. Rule Categorization
*   **Global Invariants:** System-wide "Never" events.
*   **Velocity Limits:** Rate-limiting of actions over a sliding time window.
*   **Sanity Checks:** Logic bounds (e.g., schema validation).

## IV. Technical Requirements
*   **Latency:** Policy evaluation must be optimized for sub-millisecond response times.
*   **Auditability:** Every validation decision (Allow/Reject) must include a detailed trace.
*   **Side-Effect Free:** Validation logic may read from state caches but must never write to external environments.

## V. Prohibited Patterns
*   **"Soft" Warnings:** There are no warnings. An Intent is either compliant or it is rejected.
*   **External Dependencies:** Avoid blocking I/O inside validation logic.

## VI. Threat Model & Security Heuristics
When writing or reviewing code governed by the Validation layer, you must proactively evaluate the following:
1.  **Authentication & Privilege:** Verify authentication/authorization boundaries and identify privilege-escalation risks.
2.  **Input Resistance:** Require strict input validation and injection resistance mapping on all externally reachable paths.
3.  **Zero-Trust Secrets:** Mandate strict secret handling rules across code, config, runtime, and logging surfaces.
4.  **Cryptography:** Enforce cryptographic usage correctness and flag insecure default detection scenarios.
