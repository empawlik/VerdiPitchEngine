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
body_hash: c8a2fc88f372c6d5
tags: [logic, strategy, determinism, intent-generation, go]
---

# 200-INTEL: The Intelligence Layer (Strategy)

## I. Functional Core
The Intelligence Layer (`/pkg/strategy`) is responsible for processing environmental telemetry and system state to produce a proposed **Intent**. It is a pure function of its inputs.

## II. Engineering Constraints
1.  **Strict Determinism:** Given the same inputs, the strategy must always produce the identical **Intent** hash.
2.  **Statelessness:** No local state persistence. Context must be passed in or retrieved from strict Protobuf-generated structs.
3.  **No Side Effects:** This layer is physically prohibited from:
    * Initiating network calls (I/O).
    * Accessing filesystem secrets.
    * Writing to databases.
4.  **Intent-Only Output:** All logic must terminate in the generation of an **Unsigned Intent** Protobuf message.

## III. Logic Boundaries
*   **Strategy Rules:** Reside here. These are "Goal-Oriented" (e.g., "If X occurs, we should do Y").
*   **Data Models:** All ML models or heuristic algorithms must be encapsulated here.
*   **Simulation Ready:** Logic must be able to run in `SIMULATED` mode against historical telemetry without modification.

## IV. Prohibited Patterns
*   **Leakage:** Never import from `/pkg/executor` or any credential-handling packages.
*   **Hardcoding:** Parameters must be configurable via the Kernel; never hardcode risk or execution limits (those belong in the Validation layer).

## V. Event Loop Routing & QoS (CTX-1295)
*   **Multi-Tiered Rate Limiting:** The Core Event Loop must implement dual-tier token buckets. High-alpha intents (e.g., liquidation defenses) must strictly route through an emergency reserve to bypass standard congestion limits.
*   **Non-Blocking Flush:** The saturation of the standard intent queue must never block the evaluation and flushing of the emergency queue.
