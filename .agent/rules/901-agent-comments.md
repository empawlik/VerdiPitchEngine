---
project_name: VerdiPitchEngine
version: 1.0.0
status: Active
priority: Medium
dev_stage: Production
agent_role: Core-Context
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-10
body_hash: 04afcc3521e8f1d9
tags: [dev-asset, antigravity-context]
---

# 901-COMMENTS: Agent-Optimized Inline Comments

## I. Purpose
To provide explicit, machine-readable context directly within the source code to enhance AI agent reasoning (RAG), reduce hallucination, and enforce architectural boundaries (e.g., the Decoupled Execution Pattern).

## II. The Standard
Agent-optimized comments are highly structured metadata tags embedded in Go docstrings, specifically designed for LLM and AST parsers. These tags must be placed on exported types, interfaces, and functions that cross or enforce architectural boundaries (especially in `/pkg/guardrail`, `/pkg/strategy`, and `/pkg/executor`).

## III. Supported Tags
1. **`@layer: <LayerName>`**
   - **Description:** Explicitly identifies the Decoupled Execution layer the component resides in.
   - **Allowed Values:** `Intelligence`, `Validation`, `Action`, `Kernel`, `Telemetry`, `Audit`.
2. **`@ref-rule: <RuleID>`**
   - **Description:** Directly links the code block to a governing rule in `.agent/rules/`.
   - **Allowed Values:** Any valid rule index (e.g., `100-CORE`, `300-VALID`, `400-ACTION`).
3. **`@constraint: <Description>`**
   - **Description:** States a hard invariant or constraint for the logic within the block. 
   - **Example:** `Zero Side Effects`, `Must be Idempotent`.
4. **`@complexity: <Level>`**
   - **Description:** Indicates the logical density to help agents gauge how cautiously to proceed during modifications.
   - **Allowed Values:** `Low`, `Medium`, `High`, `Critical`.

## IV. Usage Example
```go
// ValidateIntent checks a proposed trading intent against hard compliance limits.
// @layer: Validation
// @ref-rule: 300-VALID
// @constraint: Fail-Closed. Must evaluate to an absolute ALLOW or REJECT.
// @complexity: High
func ValidateIntent(ctx context.Context, intent *pb.UnsignedIntent) error {
	// ... logic
}
```

## V. Enforcement
- **Automated AST Analysis:** Presence of these tags on critical pathways is enforced primarily via the project's native Go analysis plugin (`pkg/agentlint`). This standalone analyzer ensures the metadata is syntactically attached to exported identifiers across the defined architectural bounds.
- **CI/CD Pipeline:** The `agentlint` analysis runs seamlessly as the primary step of the `make lint` pipeline, halting the build if any tags are missing before falling back to `golangci-lint` style checks.
- **Agent Orchestration:** AI Agents traversing the codebase must parse and respect these embedded tags contextually prior to making modifications.
