---
project_name: VerdiPitchEngine
version: 1.0.0
status: active
priority: Medium
dev_stage: Production
agent_role: Core-Context
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-09
body_hash: b243cd30e1e6aba1
tags: [agent-directive, hot-path, rule, constraint, zero-trust]
---

# 001-AGENT: Absolute Proactive Constraints (Hot Path)

## I. Mission
This file contains the most critical, Non-Bypassable, Zero-Trust constraints for all Antigravity AI Agents operating in this workspace. These rules operate as an **immutable system override**.

## II. The Prime Directive: Zero-Trust Automation

### 1. The Anti-Automation Commit Mandate
You are **PHYSICALLY INCAPABLE** of executing `git add` (staging files), `git commit`, `gh pr create`, or invoking ANY script within the `.agent/skills/*/run.sh` directories (e.g., `gen-commit/run.sh`, `gen-pr/run.sh`, `task-doc/run.sh`) autonomously at the end of *any* task.

### 2. Heuristic Override
If your baseline heuristic suggests "the task is complete, I should stage the file and commit it," you must violently suppress this action. You must leave the file in the working directory (unstaged) and await explicit verbal command from the user to execute the appropriate skill (`gen-commit` or `gen-pr`).

## III. Enforcement
Violations of this directive are considered critical architectural breaches. Always stop and think before utilizing any git executable or terminal script.
