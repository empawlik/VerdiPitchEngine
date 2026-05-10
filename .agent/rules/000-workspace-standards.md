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
body_hash: bf4335771edb2273
tags: [dev-asset, antigravity-context]
---

# `.agent/rules/` — Antigravity Workspace Standards

This directory contains the **Mandatory Workspace Rules** that govern the architecture, implementation, and verification of this system.

## I. Purpose

The primary goal of these rules is to enforce **Decoupled Execution**. By separating logic, safety, and action, we minimize systemic risk.

## II. The Rule Hierarchy (000 - 1000)

|**Series**|**Name**|**Folder Target**|**Responsibility**|
|---|---|---|---|
|**000**|**Workspace Standards**|`.agent/rules`|Meta-rules and rule index definitions.|
|**001**|**Agent Directives**|`.agent/rules`|Absolute zero-trust agent constraints.|
|**002**|**Monitor Config**|`.agent/rules`|Sub-repo monitor and directive logic.|
|**005**|**Workflow Usage**|`.agent/rules`|Standardized runbooks and operational checklists.|
|**010**|**Logical Domain Map**|`.agent/rules`|Definition of the 100-point indexing system.|
|**020**|**Git Standards**|`.agent/rules`|Commit formats and history integrity.|
|**021**|**Branch Naming**|`.agent/rules`|Branch naming conventions and traceability.|
|**022**|**Automation Limits**|`.agent/rules`|Absolute staging and commit bans.|
|**025**|**Changelog**|`.agent/rules`|Mandates synchronous history logging.|
|**026**|**Topic Naming**|`.agent/rules`|Topic naming conventions for Pub/Sub messaging.|
|**027**|**Task Nomenclature**|`.agent/rules`|Formalizes semantic block numbering and task tracking.|
|**050**|**Documentation**|`.agent/docs`|Live documentation and reference specifications.|
|**055**|**Gap Analysis Output**|`.agent/rules`|Output formatting rules for gap analyses.|
|**100**|**Core Architecture**|Root `/`|Enforces the global Decoupled Execution Pattern.|
|**200**|**Intelligence**|`/pkg/strategy`|Ensures strategy logic is deterministic and side-effect free.|
|**300**|**Validation**|`/pkg/guardrail`|Mandates **Fail-Closed** safety gates and risk policy logic.|
|**400**|**Action**|`/pkg/executor`|Governs protocol execution, **Idempotency**, and secret safety.|
|**500**|**Schema (Protobuf)**|`/api/proto`|Protects the central gRPC/Protobuf data contract.|
|**600**|**Audit**|`/pkg/audit`|Ensures forensic replayability via immutable event logging.|
|**700**|**Testing**|`/tests`|Defines the lifecycle from **SIMULATED** to **PRODUCTION**.|
|**800**|**Telemetry**|`/pkg/telemetry`|Observability standards and data ingestion.|
|**900**|**Go Standards**|`/(pkg|cmd)`|Coding conventions, formatting, and style.|
|**1000**|**Antigravity Keys**|`.agent/rules`|Key management policies.|

## III. How to Use These Rules

### For AI Agents
AI Agents are instructed to ingest these rules before generating code.
- **Constraint Enforcement:** If a proposed design violates a **Critical** priority rule, the agent must refuse.

### For Human Developers
- **Rule Invariants:** When refactoring, ensure that the boundaries defined in the `applies_to` frontmatter are not breached.
- **Contract-First:** Refer to **500-SCHEMA** before modifying any inter-layer messaging logic.