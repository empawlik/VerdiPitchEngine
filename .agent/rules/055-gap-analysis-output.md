---
project_name: VerdiPitchEngine
version: 1.0.0
status: active
priority: high
dev_stage: Production
agent_role: Technical-Debt
agent_weight: 3
asset_scope: Global
platform: CLI
tech_stack: [Go, PostgreSQL, pgvector, gRPC]
dependencies: []
created: 2026-03-31
updated: 2026-05-09
body_hash: 93b87864853f4fa4
tags: [documentation, output, agent-behavior, gap-analysis]
---

# 055-DOCS: Gap Analysis Output Rule

## I. The Mandate
To maintain a centralized and structured repository of all system analyses, any output generated in response to a "gap analysis" request must be saved as a dedicated Markdown file in the designated documentation directory.

This ensures that analyses are version-controlled, auditable, and accessible for future reference by both human developers and AI agents.

## II. Output Directory
All gap analysis documents must be saved to the following directory:

`.agent/docs/gap-analysis/`

## III. File Naming Convention
- Filenames must be in `UPPERCASE`.
- Use hyphens (`-`) to separate words if necessary.
- The filename should clearly describe the subject of the analysis.
- The file extension must be `.md`.

**Examples:**
- `CLOUD.md`
- `OBSERVABILITY_AND_MONITORING.md`
- `VIRTUAL_EXCHANGE.md`
- `CTX-99.md`

## IV. Agent Instructions
- Upon receiving a request to perform a "gap analysis", you must first determine the appropriate filename based on the analysis topic.
- After generating the content of the analysis, you must write the output to a new file in the `.agent/docs/gap-analysis/` directory, following the specified naming convention.
- Do not output the analysis directly to the console or any other location.
