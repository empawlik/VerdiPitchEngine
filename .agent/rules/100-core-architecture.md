---
project_name: VerdiPitchEngine
version: 1.1.0
status: active
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
body_hash: 6cb584c81818f594
tags: [architecture, isolating, pgvector, orchestration, rbac]
---

# 100-CORE: The Semantic Memory Orchestration Pattern

## I. Identity & Mission
All development in this workspace must adhere to the **OpenBrain Architecture**. The primary objective is the secure, isolated translation of static markdown into vector embeddings, enabling rich semantic retrieval across distinct data silos. This project is implemented using **Go 1.22+**.

## II. Mandatory Layer Isolation
The system is divided into three non-overlapping functional areas. No cross-silo logic leaks are permitted.

1.  **The Indexing Layer (`/pkg/indexer`):**
    *   **Role:** Reads static source material (e.g., Obsidian Vaults, PDFs, EPUBs) exclusively from QNAP shared folder mounts.
    *   **Output:** Generates normalized Documents with extracted metadata.
    *   **Constraint (Immutability Exception):** Read-only access to source material is the default constraint. The ONLY exception is for automated wiki-link generation, which is strictly restricted to authorized Obsidian vault markdown files as explicitly defined in the global OpenBrain configuration file.
    *   **Constraint (Access Control):** System access to the QNAP mounts MUST be authenticated exclusively through a dedicated `OpenBrain` user account on the NAS configured with least-privilege read access.
    *   **Constraint (Vault Structure):** Every candidate Obsidian vault MUST contain a root directory named `00-System`. The Indexing layer must ensure the contents of this system folder are always ingested.
    *   **Constraint (Exclusion Rules):** The Indexer MUST respect local `.openbrainignore` files present within any shared folder, skipping all specified paths and patterns during ingestion.
    *   **Constraint (Source Restriction):** ONLY QNAP shared folders are valid candidates for OpenBrain ingestion. Ingestion from local host directories, unapproved cloud drives, or other peripheral sources is strictly prohibited, EXCEPT when the system is explicitly launched in Disaster Recovery (DR) / Local Sandbox mode via the global configuration.

2.  **The Core API & Orchestrator (`/pkg/api` & `/pkg/orchestrator`):**
    *   **Role:** Center of truth for routing logic, API interactions, embeddings extraction, and MCP hosting.
    *   **Logic:** Receives normalized documents or user queries and routes them to the correct silo.
    *   **Constraint:** Must enforce **Rule 1000 Zero-Trust Isolation**. Queries against `brain_personal` cannot bleed into `brain_geetwee`.

3.  **The Storage Layer (`/pkg/storage`):**
    *   **Role:** Protocol-specific implementation to interface with PostgreSQL (`pgvector`).
    *   **Capability:** State mutation and semantic execution (K-Nearest Neighbors searches).
    *   **Constraint:** Must map directly to siloed databases or schemas with explicitly scoped credentials.
    *   **Constraint (Naming):** Tenant database names must strictly map to their semantic resources without redundant `open_brain_` prefixes (e.g., `vault_geetwee_db`).

## III. Execution Principles
*   **Asymmetric Immutability:** The Markdown files hold the ultimate source-of-truth. OpenBrain serves primarily as a downstream Read Replica. The only permitted upstream deviation from this rule is the authorized injection of wiki-links back into Obsidian markdown files (driven strictly by the global configuration).
*   **Zero-Trust Boundaries:** Data silos are treated as physically separated. Cross-silo database JOINs are strictly forbidden. 
*   **Audit-First:** Context retrieved or modified by AI Agents must be forensically replayable.

## IV. Prohibited Actions
*   **NEVER** bypass the Orchestrator routing to directly query `pgvector`.
*   **NEVER** use loosely typed data (`interface{}`, `map[string]interface{}`) for the ingestion pipeline or GraphQL/gRPC schema.
*   **NEVER** assume tenant ID context is global. Every API request must carry explicit authentication/tenant context to maintain Rules of Isolation.
*   **NEVER** ingest, index, or monitor file paths that do not explicitly originate from a designated QNAP shared folder volume mount.

