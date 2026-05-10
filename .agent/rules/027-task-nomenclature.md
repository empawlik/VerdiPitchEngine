---
project_name: VerdiPitchEngine
version: 1.0.0
status: active
priority: high
dev_stage: Production
agent_role: Core-Context
agent_weight: 4
asset_scope: Global
platform: CLI
tech_stack: []
dependencies: []
created: 2026-04-25
updated: 2026-05-10
body_hash: e3d80ba276ddf1e6
tags: [architecture, tracking, standards, nomenclature]
---

# 027-TASK: Semantic Block Numbering & Task Nomenclature

## I. Purpose
The jump from `CTX-00004` to `CTX-10001` is a deliberate architectural pattern known as **Semantic Block Numbering**, which is essential for "Audit-First Engineering" and maintaining strict referential integrity in git commit histories (Rule 020-GIT).

In a high-assurance environment, task IDs are not just sequential counters; they are telemetry markers. By assigning specific number blocks to specific priorities or architectural phases, an engineer can instantly identify the risk profile and context of a task just by looking at its ID.

## II. The C.O.R.T.E.X. Backlog Mapping
Here is the deterministic mapping used in the C.O.R.T.E.X. backlog:

*   **CTX-0XXXX Series (00001 - 09999): Priority 0 / Foundation / Critical Blockers.**
    *   *Reasoning:* If a developer sees a `0XXXX` ticket, they immediately know it is a core architectural prerequisite or a critical security blocker that halts all other work. 
*   **CTX-1XXXX Series (10001 - 19999): Priority 1 / Active Development (MVP).**
    *   *Reasoning:* These are the core features required for the first operational milestone.
*   **CTX-2XXXX Series (20001 - 29999): Priority 2 / Refinement.**
    *   *Reasoning:* Features that enhance the MVP but are not strict execution blockers.
*   **CTX-3XXXX & 4XXXX Series: Icebox Epics.**
    *   *Reasoning:* Separated by epic (e.g., 300s for the B2C Funnel, 400s for B2B SaaS Prep).

## III. The Engineering Advantage (Audit Collision Mitigation)
If we used strict sequential numbering (e.g., Priority 0 ends at `00004`, so Priority 1 starts at `00005`), we would run into a severe **Audit Collision Risk**. 

If tomorrow we discover a massive security vulnerability in the RiverQueue implementation that must be fixed immediately (Priority 0), we would have to inject it into the middle of the backlog. 
*   In sequential numbering, we'd either have to name it `CTX-00018` (which makes a critical blocker look like a low-priority Icebox task) or shift all the numbers, which breaks existing `git commit` history references.
*   With Semantic Block Numbering, we simply assign the new critical blocker `CTX-00005`. It correctly groups with the foundation tasks, and the `CTX-1XXXX` block remains perfectly intact.
