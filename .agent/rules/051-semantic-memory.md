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
created: "2026-04-01"
updated: 2026-05-10
body_hash: 87937d67f14528f8
tags: [dev-asset, antigravity-context, workspace-rule]
---

# 051-MEM: Semantic Memory & Knowledge Ingestion (OpenBrain)

## I. The Mandate
This workspace integrates with the **OpenBrain Semantic Memory Layer**, a local-first, multi-tenant `pgvector` database and distributed file system (QNAP). OpenBrain serves as the primary "Active Knowledge" backend for the agent, maintaining historical context, architectural decisions, external literature (e.g., books), and cross-project knowledge.

While **ATLAS** (Rule 050) is used to load static, immutable third-party vendor API specifications deterministically, **OpenBrain** is responsible for all dynamic, unstructured, conversational, and semantic knowledge retrieval.

## II. Knowledge Retrieval (The Pull)
When an agent requires domain context, external literature, or past system decisions, it MUST retrieve that information via the `openbrain` skill.

1. **Mandatory Interface:** Agents are strictly prohibited from manually executing raw database commands (e.g., `docker exec openbrain-pgvector psql ...`), schema dumps, or arbitrary filesystem sweeps to read OpenBrain internal state.
2. **Authorized Tools:** Only the formalized wrapper scripts located in `.agent/skills/openbrain/scripts/` may be executed:
   - `search-qnap.sh`: Used to locate and verify the presence of documents across authorized NAS volume mounts.
   - `search-pgvector.sh`: Used for semantic relationship querying and memory retrieval directly from the PostgreSQL backend.

## III. Knowledge Ingestion (The Push)
Agents must respect the boundary between knowledge retrieval and knowledge ingestion to preserve system isolation.

1. **Passive Ingestion (Standard Mode):** The agent does NOT manually insert memory vectors into the database. OpenBrain relies on an event-driven `fsnotify` watcher (the Delta Hydration Strategy) to detect new files and automatically process them into the `pgvector` tables.
2. **Agent Ingestion Vector:** To "ingest" external knowledge, web research, or internal context, the agent must simply write a well-formatted Markdown file into a directory monitored by OpenBrain (e.g., `00-System/00-Resources/90-Inbox/`, `.agent/docs/gap-analysis/`, or the `/share/OpenBrain_Books_*` mounts). The background pipeline will automatically handle extraction and vectorization.
3. **Strict Decoupling:** The agent MUST NOT attempt to execute `INSERT` statements into the `geetwee_memories` table or force manual indexing pipelines.

## IV. Error Handling & Constraints
- **Zero-Trust Boundary:** The agent can inject knowledge by creating files but cannot control the embedding parameters. 
- **Staleness/Missing Data:** If an agent recently created a markdown file (ingestion) but a subsequent `openbrain` query cannot find its semantic context, the agent must assume the OpenBrain indexing service is lagging or offline. The agent should inform the user rather than attempting to debug the OpenBrain Postgres container.
