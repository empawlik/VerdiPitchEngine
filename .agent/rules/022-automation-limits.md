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
body_hash: 4fa15f1ab2a51f5d
tags: [dev-asset, antigravity-context]
---

# 022-AUTO: Absolute Agent Constraints

## I. OpenBrain Context
- **"OpenBrain Vault"** refers strictly to the Obsidian Knowledge base structure holding documentation and schemas.
- **The Codebase** (this repository) is the working project surface. They are tightly coupled conceptually but physically distinct.

## II. ABSOLUTE STAGING BAN (NEVER STAGE)
AI Agents are **STRICTLY PROHIBITED** from executing `git add` or staging files under any circumstances. Agents MUST NEVER stage files. There are NO exceptions to this rule.

## III. ABSOLUTE COMMIT BAN (NEVER COMMIT)
AI Agents are **STRICTLY PROHIBITED** from executing `git commit` or any git finalization commands. Agents MUST NEVER commit files automatically or autonomously. 
**Delegation Mandate:** All commits MUST be delegated to the `gen-commit` skill, which in turn requires strict imperative command authorization (see section IV). You have ZERO native git execution authority.

## IV. GATED SKILL EXECUTION
AI Agents are **PROHIBITED** from autonomously executing ANY script, workflow, or tool located within the `/Users/epawlik/Dev/Workspace/OpenBrain/.agent/skills/` directory unless explicitly commanded to do so. Merely mentioning a skill name or path in a user prompt is INSUFFICIENT. An agent may ONLY execute a skill if the user explicitly issues an imperative command to run it (e.g., 'Run the task-doc skill', 'Execute @[/path/to/skill]').

## V. TASK.MD HEURISTIC BAN
AI Agents MUST NEVER include skill executions (like `gen-commit`, `gen-pr`, etc.) as planned steps in their `task.md` checklists. Adding them to a checklist creates a false heuristic that leads to autonomous execution violations. Skill invocations are strictly user-driven and must remain off all internal task lists.
