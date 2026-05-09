---
project_name: VerdiPitchEngine
version: 1.0.0
status: Active
priority: Critical
dev_stage: Production
agent_role: Core-Context
agent_weight: 5
asset_scope: Global
platform: CLI
tech_stack: [SQLite, CAS]
dependencies: [ATLAS]
created: 2026-03-09
updated: 2026-05-09
body_hash: ceb43393ebffa5f0
tags: [dev-asset, antigravity-context, rules, atlas]
---

# Rule 050: Environmental Documentation (ATLAS Heuristics)

You operate within a Decoupled Execution architecture. You DO NOT hold native knowledge of external vendor APIs, SDK versions, or Go language primitives. 

Before executing any code modification involving third-party libraries, services, or protocols, you MUST deterministically load the trusted, verified documentation from the localized ATLAS Action Layer.

## 1. The ATLAS Cold Path
All trusted documentation is cataloged locally in an SQLite database and immutably anchored to the filesystem using Content Addressable Storage (CAS). As per the integration architecture, OpenBrain access to ATLAS is facilitated via a **Docker Container** mount or the `atlas-server` daemon exposing port `50051`.
- **SQLite Index:** [/Users/epawlik/Dev/Workspace/atlas/.atlas-data/atlas_index.db](cci:7://file:///Users/epawlik/Dev/Workspace/atlas/.atlas-data/atlas_index.db:0:0-0:0) (Volume Mounted)
- **CAS Root:** `/Users/epawlik/Dev/Workspace/atlas/.atlas-data/` (Volume Mounted)

## 2. Heuristic Mapping Protocol
When formulating your Execution Intent (Task), you must evaluate the required technological domains (e.g., `AWS`, `Go`, `gRPC`).

If a domain matches an expected vendor, execute the following retrieval flow:

1. **Verify Integration State:** Ensure the **ATLAS Protocol Servers** (`atlas-sqlite` and `atlas-mcp`) are active and registered in the global Antigravity MCP configuration.
2. **Query the Index:** You MUST utilize the native JSON-RPC tools provided by the **`atlas-sqlite` MCP Server** (specifically `read_query`) to execute `SELECT` statements against the index (e.g., `SELECT category, version, checksum, ingested_at, upstream_available_version FROM specs WHERE vendor='AWS' ORDER BY ingested_at DESC LIMIT 1;`). Do NOT use `sqlite3` bash commands.
3. **Execute Procedural Skills:** If execution tasks mandate the use of custom procedural scripting logic, you MUST route them dynamically through the **`atlas-mcp` Go Router Server**'s `tools/call` JSON-RPC specification. Unrestricted local Bash/Python logic is prohibited per Rule 1000.
4. **Resolve the Filepath:** The physical, immutable file is located precisely at the volume mounted `<CAS Root>/<category>/<checksum>`.
5. **Verify Index Freshness (TTL & Upstream availability):** Calculate the mapping's age using `ingested_at`. If it exceeds the `ATLAS_MAX_DOC_AGE_DAYS` threshold (default 30 days) OR if `upstream_available_version` is populated, the documentation is STALE.
6. **Ingest the Context:** Prioritize retrieval from the git-ignored local cache (`.agent/knowledge/vendor/<checksum>.md`). If a cache miss occurs, read the raw bytes from the CAS root, perform AST compression, and persist the optimized result to the local cache before appending to memory.

**CRITICAL CONSTRAINT (Fail-Closed):** 
If the `atlas-sqlite` MCP Server returns empty for a required vendor, if the physical file is missing from the CAS Root, OR if the fetched metadata violates the **verify index freshness** validation (Step 5), you MUST HALT the Dev Task immediately with a `StaleDocumentationError`. Inform the user that the Intelligence Layer lacks a deterministic API reference and an ATLAS index sync must be performed by the operator before the task can continue.
