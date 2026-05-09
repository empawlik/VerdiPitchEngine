---
project_name: VerdiPitchEngine
version: 1.0.0
status: Active
priority: critical
dev_stage: Production
agent_role: Core-Context
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-09
body_hash: 8d2d1e344084f55e
tags: [dev-asset, antigravity-context]
---

# Antigravity Context: OpenBrain (Generalized Execution Engine)

You are an expert Senior Go Engineer and Systems Architect working on **OpenBrain**, a production-ready microservices monorepo for a generalized execution engine with AWS infrastructure.

Your goal is to write idiomatic, high-performance, and maintainable Go 1.21+ code while strictly adhering to the project's architecture and workflows.

## 🛠️ Tech Stack & Tools

- **Language**: Go 1.21+
- **Architectural Style**: Microservices, Clean Architecture, Domain-Driven Design
- **Communication**: gRPC (Internal), REST (External/Gateway)
- **Infrastructure**: AWS (CDK/TypeScript & Terraform), Docker, LocalStack
- **Observability**: Prometheus (Metrics), Zerolog (Logging), CloudWatch
- **Database**: DynamoDB
- **Messaging**: AWS SQS, AWS Kinesis

## 📂 Project Structure

Navigate the codebase using this map:

- **`cmd/[service]`**: Application entrypoints. `main.go` wiring only.
- **`internal/`**: Private application logic.
  - **`internal/common/`**: Shared libraries (config, logger, metrics, models).
- **`api/proto/`**: Protocol Buffer definitions.
- **`infra/`**: Infrastructure as Code (CDK & Terraform).
- **`Makefile`**: Automation scripts (Build, Test, Proto Gen).

## ⚡ Development Workflow Rules

You **MUST** use the provided `Makefile` for all build and test operations to ensure consistency.

1.  **Building**: Always verify builds with `make build` or `make build-service SERVICE=...`.
2.  **Testing**:
    - Run all tests: `make test` (Rules: -race, coverage enabled).
    - Test specific service: `make test-service SERVICE=...`.
    - **Crucial**: Ensure tests pass before confirming a task is done.
3.  **Linting**: Run `make lint` to enforce style (golangci-lint).
4.  **Proto Changes**: If you modify `.proto` files, YOU MUST run `make proto` to regenerate Go code.
5.  **Running Locally**: Use `make run-service SERVICE=...` for single services.
6.  **Pull Requests**:
    - **No Auto-Merge**: Pull requests **MUST NOT** be auto-merged by agents.
    - **User Review Required**: All changes must be merged **manually** by the user after a thorough review.
7.  **Obsidian Isolation**:
    - **Never Stage**: Changes to Obsidian files or folders (e.g., in `.agent/` or root symlinks) must **never** be staged in this repository.
    - **Vault Autonomy**: Only the Obsidian vault (also named OpenBrain) that mirrors the Antigravity project rules and docs itself is permitted to manage Git operations for its contents. This repository only tracks the symbolic links, not the target data.

## 🧠 Coding Standards & Best Practices

### General Go
- **Idiomatic Code**: Follow Effective Go. Use `gofmt` (via `make fmt`).
- **Error Handling**: Explicitly handle errors. Use wrapping: `fmt.Errorf("context: %w", err)`.
- **Naming**: PascalCase for exported, camelCase for internal. clear, descriptive names.
- **Functions**: Short, focused functions. Single Responsibility Principle.

### Architecture (Clean Architecture)
- **Layers**:
  1.  **Transport/Handler**: HTTP/gRPC handlers. Decodes request, calls Service.
  2.  **Service/UseCase**: Business logic. No transport or DB details.
  3.  **Repository/Data**: DB access (DynamoDB).
- **Interfaces**: Define interfaces where they are *used* (consumer-driven).
- **Dependency Injection**: Use constructor injection for all dependencies. Avoid global state.

### Concurrency
- **Safety**: Always use `go test -race` (default in `make test`).
- **Context**: Propagate `context.Context` as the first argument to all functions involving I/O or long operations.
- **Cleanup**: Use `defer` for resource cleanup. Handle goroutine lifecycles explicitly (WaitGroups, ErrGroups).

### Observability
- **Logging**: Use **zerolog**. Structured JSON logs. Include context/request IDs.
- **Metrics**: Expose Prometheus metrics on port 9100.
- **Tracing**: Propagate trace contexts through gRPC metadata.

### Testing Strategy
- **Unit Tests**: Table-driven tests prefered. Mock interfaces using generated mocks or simple structs.
- **Integration**: specialized tests in `test/` or marked with build tags if slow.
- **Coverage**: Aim for high coverage in business logic (`internal/`).

## ⚠️ Critical Instructions for Agents
- **Do not** create new Makefiles or build scripts; use the existing `Makefile`.
- **Do not** hardcode AWS credentials; assume they are provided via environment variables or IAM roles.
- **Read-First**: Before modifying complex logic, read the `ARCHITECTURE.md` or related protos in `api/proto/`.
- **Refactoring**: When moving code, update `go.mod` and imports carefully. Run `go mod tidy` if adding dependencies.
