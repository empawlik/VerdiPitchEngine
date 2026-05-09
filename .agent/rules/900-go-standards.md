---
project_name: VerdiPitchEngine
version: 1.0.0
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
body_hash: 02fc1131820e2def
tags: [go, standards, style, formatting, linting]
---

# 900-GO: Go Coding Standards

## 900.1 Code Formatting & Linting
- **Standard**: All Go code **must** be formatted with `gofmt` and `goimports`.
- **Enforcement**: The CI pipeline must run `golangci-lint` with a standard ruleset. A failed lint check must fail the build.
- **Line Length**: Adhere to the community standard of keeping lines to a reasonable length (soft limit around 80-100 characters), but clarity is more important than a strict rule.
- **Imports**: `goimports` handles this automatically. It groups standard library, third-party, and project packages.

## 900.2 Naming Conventions
- **Packages**: Use short, concise, all-lowercase names (e.g., `strategy`, `guardrail`). Avoid `_` or `camelCase`.
- **Variables & Functions**: Use `camelCase`. Exported identifiers must start with an uppercase letter (`PascalCase`).
- **Interfaces**: Do NOT prefix with `I`. Good: `io.Reader`. Bad: `IReader`. Interfaces should be named for what they do (e.g., `Stringer`).
- **Structs**: No special prefix or suffix. Name them for what they represent.
- **Test Files**: Must end with `_test.go` (e.g., `strategy_test.go`).

## 900.3 Idiomatic Go
- **Error Handling**: Errors are values. Check for `err != nil` immediately after a function call that can return an error. Provide context to errors before returning up the stack (e.g., `fmt.Errorf("failed to do X: %w", err)`).
- **Concurrency**: Use goroutines and channels for concurrent operations. Avoid sharing memory by communicating. Use `context.Context` for cancellation and deadlines.
- **Structs**: Embed structs for composition, not inheritance.
- **Zero Values**: Design structs so that their zero value is useful.

## 900.4 Project Layout
Adhere to the [Standard Go Project Layout](https://github.com/golang-standards/project-layout).
- `/cmd`: Main application entry points.
- `/pkg`: Public library code, safe for others to import.
- `/internal`: Private application and library code. It's code you don't want others importing in their applications.
- `/api`: Protobuf definitions, Swagger docs, etc.

## 900.5 Code Quality & Best Practices
- **Immutability**: Pass large structs by pointer, but be mindful of who can modify them. Prefer passing values where feasible to avoid side effects.
- **`interface{}`**: Avoid empty interfaces. Be explicit about types.
- **Comments**: Public functions and types must have doc comments explaining their purpose. Explain *WHY*, not *WHAT*.
- **Logging**: Use a structured logger like the standard library's `slog`.
- **Testing**:
    - **Framework**: Use the standard `testing` package.
    - **Assertions**: Use `t.Errorf` or `t.Fatalf`. For more complex assertions, `stretchr/testify` is permitted.
    - **Subtests**: Use `t.Run` to organize tests into logical groups.
    - **Table-Driven Tests**: Use table-driven tests for testing multiple cases of the same function.

## 900.6 Performance & Concurrency Heuristics
When writing Go code, you must proactively evaluate the following performance and concurrency bounds:
1.  **Goroutine Lifecycles:** Enforce strict goroutine lifecycle correctness and cancellation propagation via `context.Context`.
2.  **Channel Safety:** Restrict channel buffering assumptions to avoid deadlocks or panics on closed channels.
3.  **Allocation Efficiency:** Dictate explicit allocation and copy behavior rules on performance-sensitive backend paths to minimize GC pressure.
4.  **Interface Cohesion:** Validate interface boundaries to ensure package-level cohesion.
