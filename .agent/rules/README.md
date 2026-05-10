---
project_name: VerdiPitchEngine
version: "v0.28.1"
status: "Active"
priority: "Medium"
dev_stage: "Production"
agent_role: "Core-Context"
agent_weight: 5
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell"]
dependencies: []
created: "2026-02-23"
updated: 2026-05-10
body_hash: 01be95363df8b27f
tags: [dev-asset, antigravity-context, workspace-rule]
---

# 📜 .agent/rules/ | Proactive Workspace Standards

This directory contains the **Mandatory Workspace Rules** that govern the architecture, implementation, and verification of this system. These rules are not suggestions; they are the "Laws of Physics" for the codebase.

## I. Purpose

The primary goal of these rules is to enforce **Decoupled Execution**. By separating logic, safety, and action, we minimize systemic risk and ensure that every state transition is deterministic and auditable.

---

## II. Agent Monitoring: Proactive (Hot)

This directory is the **"Hot" or "Proactive"** monitoring path for the AI agent. Files here are treated as the highest priority system constraints.

-   **Immediate Evaluation:** The agent is required to load and understand these rules before every action.
-   **Strict Enforcement:** A violation of a `critical` rule in this directory will cause the agent to halt and report a compliance failure.
-   **Autonomous Triggers:** This is the only location where `@turbo` directives are permitted, allowing for automated code implementation without manual review.

| Feature         | `.agent/rules/ (Hot)`         | `.agent/knowledge/ (Cold)` |
| :-------------- | :---------------------------- | :------------------------- |
| **Priority**    | **High (Proactive)**          | Low (Reactive)             |
| **Logic**       | **"Stop me if I violate this."** | "Inform me if I ask."      |
| **Index Mode**  | **System Prompt / Constraints** | Vector Search (RAG)        |
| **Directives**  | **Enabled (`@turbo` allowed)**  | Ignored (No `@turbo`)      |

---

## III. The Rule Hierarchy (000 - 1000)

|**Series**|**Name**|**Folder Target**|**Responsibility**|
|---|---|---|---|
|**000**|**Workspace Standards**|`.agent/rules`|Meta-rules and rule index definitions.|
|**020**|**Git Standards**|`.agent/rules`|Commit formats and history integrity.|
|**025**|**Changelog**|`.agent/rules`|Mandates synchronous history logging.|
|**050**|**Documentation**|`.agent/docs`|Live documentation and reference specifications.|
|**100**|**Core Architecture**|Root `/`|Enforces the global Decoupled Execution Pattern.|
|**500**|**Schema (Protobuf)**|`/api/proto`|Protects the central gRPC/Protobuf data contract.|
|**600**|**Audit**|`/pkg/audit`|Ensures forensic replayability via immutable event logging.|
|**700**|**Testing**|`/tests`|Defines the lifecycle from **SIMULATED** to **PRODUCTION**.|
|**800**|**Telemetry**|`/pkg/telemetry`|Observability standards and data ingestion.|
|**900**|**Go Standards**|`/(pkg\|cmd)`|Coding conventions, formatting, and style.|
|**1000**|**Antigravity Keys**|`.agent/rules`|Key management policies.|

---

## IV. How to Use These Rules

### For AI Agents

AI Agents are instructed to ingest these rules before generating code.

- **Scope-Awareness:** Rules with `scope: local` apply strictly to their designated folders.
    
    

### For Human Developers

- **Rule Invariants:** When refactoring, ensure that the boundaries defined in the `applies_to` frontmatter are not breached.
    
    

---

## V. Core Architecture Philosophy

> **"Decouple the decision from the consequence."**

By following this hierarchy, we ensure that:

1. The **Intelligence Layer** can be as complex as needed without risking the safety of the system.
    
2. The **Validation Layer** remains simple, robust, and absolute.
    
3. The **Action Layer** focuses solely on the "How" of execution, not the "Why."
    

---

## VI. Maintenance

These rules are versioned. Any changes to the core Antigravity folders must be preceded by a review of the corresponding rule file to ensure the architecture remains intact.

**Would you like me to initialize the actual `.agent/rules/` folder with these files in your environment?**