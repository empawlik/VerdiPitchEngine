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
tech_stack: []
dependencies: []
created: 2026-04-19
updated: 2026-05-09
body_hash: 7828dea5d36f7e19
tags: [architecture, messaging, standards, pubsub, routing]
---

# 026-TOPIC: Pub/Sub Topic Naming Convention

## I. Purpose
To maintain a high-integrity, heavily decoupled messaging architecture (`Rule 100-CORE`). This standard ensures that Pub/Sub event topics (like those used in RiverQueue) explicitly separate the domains defining *what happened* from the downstream configurations determining *where it is routed*.

## II. The Core Topic Structure (Dot-Notation)
All primary event broadcast topics MUST follow the dot-notation pattern:

```text
<domain>.<action>
```
OR
```text
<domain>.<resource>.<state>
```

### 1. The Domain (Mandatory)
The subsystem or logical boundary generating the event.
- **Examples**: `chat`, `system`, `trading`, `indexer`, `auth`

### 2. The Action or State (Mandatory)
A concise description of the lifecycle event. It should reflect an observation, not a commanded destination.
- **Examples**: `incoming`, `alerts`, `completed`, `failed`

### 3. Valid Dot-Notation Examples
- `chat.incoming`
- `system.alerts`
- `indexer.vault.synced`
- `auth.login.failed`

> [!IMPORTANT]
> **Decoupled Purity**: Dot-notation topics MUST NOT contain routing information (e.g., `chat.post_to_mattermost` is prohibited). They must strictly describe the event context.

## III. The Alias Mapping Structure (Kebab-Case)
The only exception to the Dot-Notation pattern is when referencing hardcoded, endpoint-specific routing aliases assigned within the subscription mapping registry (e.g., the `mattermost_subscriptions` table). These use the kebab-case pattern:

```text
<target>-<context>
```

### 1. Alias Examples
- `devops-alerts`
- `sandbox-ad-hoc`
- `test-topic`

> [!WARNING]
> Kebab-case generic aliases should only be used as physical endpoints registered within the router layer or fallback testing targets. The core intelligence layer MUST emit events using the semantic dot-notation (Domain.Action).

## IV. Agentic Rules
- **Automatic Compliance**: AI Agents proposing new queue architectures, Webhooks, or RiverQueue pipelines MUST formulate topics using the `domain.action` convention.
- **Refactoring Mandate**: If an agent encounters a business logic event hardcoded with a destination name (e.g., `TargetTopic: "mattermost"`), it should highlight this as a violation of `Rule 100-CORE` and `026-TOPIC`.
