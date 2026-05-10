---
project_name: VerdiPitchEngine
version: 1.1.0
status: active
priority: critical
dev_stage: Production
agent_role: Core-Context
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-10
body_hash: a40e7832490585d5
tags: [protobuf, grpc, types, contract, consistency, compatibility, go]
---

# 500-SCHEMA: The Data Contract (Protobuf)

## I. The Single Source of Truth
The Protobuf (`.proto`) definition files located in `/api/proto/` are the absolute authority for the system's data structures and service contracts. All data interchange between layers and services **must** be defined here.

## II. Mandatory Schema Principles
1.  **Strict Typing:** Use of `Any`, `Value`, or generic `bytes` fields for structured data is strictly prohibited for core primitives. All fields must be explicitly typed.
2.  **Backward Compatibility:** Breaking changes to message formats or service definitions must be managed carefully, following Protobuf best practices (e.g., using reserved fields, never re-using field numbers).
3.  **Semantic Precision:** Message and service names must reflect their role (e.g., `UnsignedIntent`, `ValidationResponse`, `ExecutionReceipt`).
4.  **Enforced Documentation:** Every message, field, and service must have comments explaining its purpose and units.
5.  **Code Generation**: Go structs are generated from these `.proto` files using `protoc`. These generated structs are the only valid data carriers for cross-layer communication.

## III. The Intent Contract
The **Intent** message is the most critical schema. It must contain:
*   **Header:** Metadata including `intent_id`, `timestamp`, and `origin_service`.
*   **Body:** The specific parameters of the desired action.
*   **Constraints:** Optional execution bounds.

## IV. Layer-to-Schema Mapping

| Lifecycle Stage | Primary Schema (Message) | Origin Layer | Destination Layer |
| :--- | :--- | :--- | :--- |
| **Strategy Output** | `UnsignedIntent` | Intelligence | Validation |
| **Policy Result** | `ValidationResponse`| Validation | Execution Kernel |
| **Execution Command**| `SignedIntent` | Execution Kernel | Action |
| **Execution Result** | `ExecutionReceipt` | Action | Audit Ledger |

## V. Prohibited Patterns
*   **Manual Structs:** Never manually define a Go struct for data interchange that should be defined in a `.proto` file.
*   **Implicit Units:** Never define a numeric field without a clear unit in its comment (e.g., `// price in USD cents`). Use standard Protobuf types like `google.type.Money` where applicable.
*   **Circular Imports:** Protobuf package definitions must be structured to avoid circular dependencies.
