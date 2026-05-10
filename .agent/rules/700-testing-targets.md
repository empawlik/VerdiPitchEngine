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
body_hash: 113203dbc1b824a0
tags: [testing, simulation, sandbox, production, validation, go-test]
---

# 700-TESTING: Execution Targets & Verification

## I. The Environment Hierarchy
The system must explicitly support three execution modes. Logic must behave identically across all modes.

1.  **SIMULATED (Local/CI):**
    *   **Authority:** Zero.
    *   **Goal:** Verify logic via Go's standard `testing` package, using table-driven tests and mocking external dependencies where necessary.

2.  **SANDBOX (Testnet/Dev API):**
    *   **Authority:** Limited.
    *   **Execution:** Run with build tags or environment variables (e.g., `GO_ENV=sandbox`).
    *   **Goal:** Verify the **Action Layer** integration.

3.  **PRODUCTION (Live):**
    *   **Authority:** High.
    *   **Goal:** High-assurance execution.

## II. Mandatory Testing Principles
1.  **Environmental Parity:** The **Validation Layer** rules must be exactly the same in `SIMULATED` as they are in `PRODUCTION`.
2.  **Deterministic Verification:** Every logic change in Intelligence must be accompanied by a simulation test.
3.  **Failure Injection:** Tests must include error scenarios (e.g., network errors from mocked clients).
4.  **Mutagenic Testing:** All high-assurance packages (e.g., `/pkg/guardrail`, `/pkg/strategy`) must undergo mutation testing to prove safety invariant enforcement.

## III. Verification Gates

| Target | Requirement | Gatekeeper |
| :--- | :--- | :--- |
| **SIMULATED** | High test coverage, all tests passing, acceptable mutation score. | CI/CD |
| **SANDBOX** | Connectivity check, Idempotency verification. | QA Lead |
| **PRODUCTION** | Multi-sig approval, Audit log readiness. | Architect |

## IV. Prohibited Patterns
*   **Live-Testing:** Never test new **Guardrails** directly in `PRODUCTION`.
*   **Mocks in Strategy:** Avoid mocking the strategy logic itself; mock only the inputs.
