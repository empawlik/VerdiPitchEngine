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
updated: 2026-05-10
body_hash: e0845081fb0e86a2
tags: [architecture, domain-map, logic, definitions]
---

# 010-MAP: The Antigravity Logical Domains

As the **Antigravity Principal Systems Architect**, I am defining the complete **Logical Domain Map**. In our framework, these domains represent the rigid lifecycle of data as it moves from observation to execution.

We use a **sparse 100-point indexing system**. This ensures that as the system grows, new sub-rules can be slotted into their respective functional families without disrupting the core hierarchy.

---

## 000 – Workspace Standards & Meta-Rules
- **Purpose:** Project governance and rule indexing.
- **Key Concept:** **The Rule of Law**.
- **Responsibility:** Defining how the workspace is organized, how rules are applied, and where documentation lives.

## 050 – Documentation
- **Purpose:** Live documentation and reference specifications.
- **Key Concept:** **Active Knowledge**.
- **Responsibility:** Maintaining the semantic source of truth for the project.

## 100 – Core Architecture & Orchestration
- **Purpose:** Defines the "Laws of Physics" for the workspace.
- **Key Concept:** The **Decoupled Execution Pattern** and **Kernel** logic.
- **Responsibility:** Governance of layer isolation and the high-level handoff between Intelligence, Validation, and Action.

## 200 – Intelligence (The Strategy Layer)
- **Purpose:** Deterministic decision-making logic.
- **Key Concept:** The **Intent** (Unsigned).
- **Responsibility:** Processing environmental telemetry and system state to decide "What should be done." No execution authority exists here.

## 300 – Validation (The Guardrail Layer)
- **Purpose:** Authoritative safety and risk policing.
- **Key Concept:** **Fail-Closed** logic.
- **Responsibility:** Evaluating **Intents** against global invariants, velocity limits, and compliance. This is the non-bypassable "Veto" layer.

## 400 – Action (The Execution Layer)
- **Purpose:** Protocol realization and side effects.
- **Key Concept:** **Idempotency**.
- **Responsibility:** Signing transactions, interacting with APIs/Gateways, and managing secrets. It converts a **Validated Intent** into a physical reality.

## 500 – Schema & IDL (The Contract Layer)
- **Purpose:** Strict type definition and consistency.
- **Key Concept:** **Single Source of Truth**.
- **Responsibility:** Managing **Protobuf** definitions and their generated Go structs. It ensures all layers share a common language for messages (Intents, Receipts, State) via gRPC.

## 600 – Audit & Forensics (The Ledger Layer)
- **Purpose:** Observability and system replayability.
- **Key Concept:** **Log-Before-Act**.
- **Responsibility:** Maintaining the immutable event store. It captures every state transition for post-mortem analysis and forensic recovery.

## 700 – Verification & Environments (The Target Layer)
- **Purpose:** Testing parity and promotion lifecycle.
- **Key Concept:** **SIMULATED vs. PRODUCTION**.
- **Responsibility:** Defining how logic is promoted from virtual playback (oracles) to live execution environments.

## 800 – Telemetry & State (The Observation Layer)
- **Purpose:** Data ingestion and normalization.
- **Key Concept:** **Oracle Integrity**.
- **Responsibility:** Formatting external data into a structured `SystemState` that the **Intelligence Layer** can consume deterministically.

## 900 – Language Standards (Go)
- **Purpose:** Implementation and style guidelines.
- **Key Concept:** **Idiomatic Precision**.
- **Responsibility:** Ensuring code quality, consistency, and adherence to language-specific best practices.



## 1000 – Security & Keys
- **Purpose:** Secrets management and key rotation policies.
- **Key Concept:** **Zero-Trust Access**.
- **Responsibility:** Handling of cryptographic material and access control policies (Antigravity Keys).

---

## Domain Relationship Summary

|**Index**|**Domain**|**Output Type**|**Authority Level**|
|---|---|---|---|
|**100**|**Core**|Control Flow|System Governor|
|**200**|**Intelligence**|`UnsignedIntent`|Advisory (Proposes)|
|**300**|**Validation**|`ValidationStatus`|Dictatorial (Permits)|
|**400**|**Action**|`ExecutionReceipt`|Executive (Performs)|
|**500**|**Schema**|Type Definitions|Contractual|
|**600**|**Audit**|Immutable Logs|Forensic|
|**700**|**Verification**|Test Results|Qualitatitve|
|**800**|**Telemetry**|`SystemState`|Informational|

---

### Architect's Enforcement Note

When adding a new file to `.agent/rules/`, you must first identify which 100-point bucket it falls into. For example, a rule about **Hardware Security Modules (HSMs)** would be assigned `410-hsm-signing.md`, as it falls under the **400 – Action** domain.
