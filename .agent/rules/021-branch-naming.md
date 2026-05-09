---
project_name: VerdiPitchEngine
version: 1.0.0
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
updated: 2026-05-09
body_hash: bf984393e3f764b5
tags: [git, standards, process, branching]
---

# 021-BRANCH: Branch Naming Convention

## I. Purpose
To maintain a high-integrity repository history where branches are easily identifiable and linkable to tasks and logical domains. This standard ensures consistent traceability from branch creation through to the final merge.

## II. The Branch Structure
All branch names MUST follow the kebab-case pattern:

```text
<type>/<reference>-<subject>
```

### 1. The Type (Mandatory)
Branch types must align with the primary purpose of the work. Use the following prefixes:

- **feat/**: New functionality or high-level features.
- **fix/**: Bug fixes.
- **docs/**: Documentation-only changes.
- **refactor/**: Code restructuring without functional changes.
- **test/**: Adding or updating tests.
- **chore/**: Maintenance, build updates, or internal tooling changes.
- **security/**: Security-related updates or vulnerability fixes.
- **hotfix/**: Urgent production fixes launched from `main`.

### 2. The Reference (Mandatory)
Every branch must include a reference to either a task ID or a logical domain index:

- **CTX-XXX**: For specific tasks tracked in `BACKLOG.md` or JIRA (e.g., `CTX-001`).
- **DDD-Scope**: For general domain-level work using the 100-point indexing system (e.g., `100-core`, `200-int`).

### 3. The Subject (Mandatory)
A brief (2-4 words) description of the work in kebab-case format.

## III. Examples
- `feat/CTX-001-idempotency-checks`
- `fix/CTX-012-memory-leak-executor`
- `docs/050-update-api-specs`
- `refactor/100-core-decoding-logic`
- `security/patch-mtls-handshake`
- `hotfix/CTX-999-critical-order-rejection`

## IV. Exemptions
- **Dependabot**: Branches created automatically by GitHub Dependabot (e.g., `dependabot/go_modules/...`) are exempt from this naming convention.

## V. Agentic Rules
- **Automatic Compliance**: AI Agents must never create branches using names that do not follow this convention.
- **Task Verification**: Before creating a branch, the agent should verify the `Task-ID` exists in the backlog or equivalent tracking system.
- **No Generic Names**: Names like `temp`, `working`, or `fix-it` are strictly prohibited.
- **Branch Cleanup**: Once a PR is merged, the local and remote branches should be deleted to keep the workspace clean.
