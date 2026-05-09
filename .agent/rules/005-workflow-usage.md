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
body_hash: 163f6d791b24236b
tags: [workflows, process, agent-behavior]
---

# 005-WORKFLOWS: Workflow Usage Policy

## I. Purpose

Files within the `.agent/workflows/` directory are **procedural guides and instructional manuals**. They serve as standardized "runbooks" for both AI agents and human developers to ensure critical, multi-step tasks are performed consistently and correctly.

## II. Workflow Categories

Workflows are categorized into two primary types:

### 1. Agent-Governing Workflows
These files define internal, meta-processes for the AI agent itself. They dictate how the agent should document its own work and the knowledge it acquires.
- **Examples**: `logs.md`, `memory.md`
- **Function**: To ensure a consistent and auditable trail of the agent's activities and decisions.

### 2. Operational Guides
These files document the "how-to" for core development and operational tasks specific to the OpenBrain project. They provide shell commands and step-by-step instructions.
- **Examples**: `setup.md`, `paper-trading.md`, `release.md`
- **Function**: To provide a reliable checklist for complex procedures like setting up an environment, running the application in a specific mode, or executing a software release.

## III. Agent Instructions

- **Primary Interpretation**: Treat all files in this directory as executable checklists or runbooks.
- **Procedural Adherence**: When tasked with an operation that matches a workflow file (e.g., "set up the project"), you must retrieve and follow the steps in the corresponding file precisely.
- **Task Finalization**: Upon finalizing any complex task or technical resolution, you MUST automatically follow the `memory.md` workflow to generate a technical memory file documenting the design decisions, patterns, and rationale. Do not wait for user prompting if a significant architectural or technical decision was made.
- **Automation Directives**: Scan for automation tags (e.g., `// turbo`). When a directive is found within a command block, you are authorized to execute that command block to fulfill the relevant step.
