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
updated: 2026-05-10
body_hash: 8098f6324720a76d
tags: [git, standards, process, llm-friendly]
---

# 020-GIT: Git Standards & History Integrity

## I. purpose
To ensure the git log remains a searchable, semantic, and machine-parseable database of intent. Random or lazy commit messages break the chain of reasoning required for advanced Agentic coding.

## II. The Commit Structure
All commits must follow the **Antigravity v1** format:

```text
<type>(<scope>): <subject>

<body>

<footer>
```

### 1. The Header (Mandatory)
- **Type**: Must be one of `feat`, `fix`, `docs`, `style`, `refactor`, `perf`, `test`, `build`, `ci`, `chore`, `revert`.
- **Scope**: The numbered domain or module affected (e.g., `200-int`, `payment-api`).
- **Subject**: Imperative, lowercase, no period (e.g., `add recursive validation logic`).

### 2. The Body (Mandatory for Features/Fixes)
- Must explain **WHY** the change was made, not **WHAT** (the diff shows strictly what).
- Must reference specific constraints or requirements.

### 3. The Footer (Context Linking)
To support LLM cross-referencing, use these tokens:
- **Ref-Rule**: `<ID>` (e.g., `Ref-Rule: 200-VALIDATION`)
- **Task-ID**: `CTX-<NNN>` (If applicable)
- **Signed-off-by**: (Developer)

### 4. Exemptions
- **Dependabot**: Commits authored by `dependabot[bot]` are exempt from the mandatory `Ref-Rule` or `Task-ID` footer requirements. Dependabot PRs may be merged without these footers.

## III. Agentic Rules

- **Strict Delegation**: AI Agents MUST NEVER directly create or execute commit messages via standard `git commit` commands. All commits MUST be delegated to the `gen-commit` skill (located at `.agent/skills/gen-commit`).
- **PR Creation**: AI Agents MUST NEVER directly create Pull Requests using the `gh pr create` command. All PR creation MUST be delegated to the `gen-pr` skill (located at `.agent/skills/gen-pr`) to ensure correct formatting and repository merge instructions.
- **Atomic Commits**: Agents must commit *per task step*, not *per session*.
- **No Force Push**: On shared branches (main/develop).
- **Branch Naming**: Must comply with `021-BRANCH` standards for all development and feature branches.
- **Empty Commits**: Allowed only for triggering CI pipelines (`git commit --allow-empty`).
- **PR Merge Strategy**: Pull Requests MUST be merged using **"Create a merge commit"**. "Squash and merge" and "Rebase and merge" are strictly prohibited, as they destroy atomic commit histories (Rule 020) and invalidate Step-2 Changelog Hashes (Rule 025).

## IV. Cryptographic Integrity & Attribution
### 1. Verified Signatures
To ensure non-repudiation and prevent history injection, **EVERY** commit must be cryptographically signed (GPG, SSH, or S/MIME).
- **Enforcement**: Commits without a "Good Signature" from a verified committer are considered non-compliant.
- **Agentic Constraint**: AI Agents must never use `--no-gpg-sign` to bypass local environment errors.

### 2. Environment Verification
If a commit fails due to a signing error (e.g. `gpg-agent` prompt timeout):
1.  **Stop**: Do not attempt to re-commit without signing.
2.  **Verify**: Check GPG installation and agent health (`gpg-agent --version`).
3.  **Repair**: Ensure the local environment is configured for non-interactive signing or request user intervention.

## V. Execution Authority & Automation Limits
- **Pull Requests**: Pull Request creation and branch finalization are **State-Altering Operations** that reside exclusively under human authority. An AI agent is strictly prohibited from invoking the `gen-pr` skill autonomously without explicit user instruction. Task completion heuristics MUST NOT trigger PR generation.
- **Commits**: Execute atomic commits per task step ONLY AFTER explicit user authorization. You MUST NOT execute the `gen-commit` skill automatically as part of a task completion heuristic unless explicitly commanded by the user.
