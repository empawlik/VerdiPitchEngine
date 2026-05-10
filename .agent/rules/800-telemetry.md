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
updated: 2026-05-10
body_hash: 33c3237f051a9de0
tags: [telemetry, oracles, data-ingestion, normalization, state, go]
---

# 800-TELEMETRY: The Observation Layer (Sensors)

## I. The Data Mandate
The Telemetry Layer provides the ground truth. Its goal is to convert noisy external data into a structured `SystemState`.

## II. Mandatory Observation Principles
1.  **Normalization:** All external data must be mapped to internal Protobuf-generated Go structs immediately.
2.  **Temporal Integrity:** Every telemetry packet must include a `SourceTimestamp` and an `IngressTimestamp` (using `google.protobuf.Timestamp`).
3.  **Oracle Resilience:** Aggregate data from multiple independent sources where possible.
4.  **Unit Standardization:** All numerical data must be converted to internal standard units (e.g., Microns).

## III. State Reconstruction

| Component | Responsibility | Failure Mode |
| :--- | :--- | :--- |
| **Ingress** | Protocol-specific listeners (e.g., gRPC streams, message queue consumers). | Disconnect |
| **Normalizer** | Type conversion and unit scaling. | Mapping Error |
| **State Cache** | Maintaining the "Latest Known Good" state. | Staleness |

## IV. Technical Requirements
*   **Determinism Support:** Telemetry must be loggable to allow for replay.
*   **Backpressure:** Telemetry streams must implement backpressure (e.g., via buffered channels or rate limiting).
*   **Health Checks:** Every source must provide a `Heartbeat` signal.

## V. Prohibited Patterns
*   **Direct Passthrough:** Never pass raw JSON or other unstructured data from an external API to the Strategy logic.
*   **Implicit State:** The Intelligence layer should never "fetch" data; data must be "pushed".

## VI. SRE & Reliability Heuristics
When designing or reviewing Telemetry mechanisms, you must proactively enforce reliability standards:
1.  **SLO/SLA Alignment:** Prioritize operational metrics that directly align with service SLOs and error-budget allocations.
2.  **Alert Quality:** Require alert quality validation ensuring a high signal-to-noise ratio and operational actionability.
3.  **Capacity Monitoring:** Mandate capacity and saturation measurements tied to user-visible performance indicators.
4.  **Resilience Planning:** Define explicit rollback and degradation strategies for critical paths to mitigate cascading failures.
