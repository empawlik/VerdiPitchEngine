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
body_hash: bbe4d49599d1bfad
tags: [git, documentation, process]
---

# 025-CHANGELOG: Changelog Maintenance Policy

## I. The Mandate
To ensure a granular, human-readable history of the project's evolution, the `CHANGELOG.md` is updated **automatically** via CI/CD pipelines. Every single commit pushed to `main` must be tracked and logged with a GitHub link.

## II. CI/CD Server-Side Generation

The repository is configured with a GitHub Action (`.github/workflows/changelog.yml`) that triggers on every push to `main`.
- **Logic**: It automatically parses the commit history since the last release tag.
- **Mechanism**: The action gathers every commit, formats it into standard changelog sections (Added, Changed, Fixed, etc.), and generates a commit back to the `main` branch containing the updated `CHANGELOG.md`.
- **Developer Responsibility**: Developers do **NOT** need to manually edit `CHANGELOG.md` or use `{{full_hash}}` placeholders. Developers only need to write clear, conventional commit messages.

## III. Agent Instructions
*   **Commit Messages**: Ensure all commit messages strictly follow conventional formats (`feat`, `fix`, `docs`, `chore`, `refactor`).
*   **Ref-Rule / Task-ID**: Always include `Ref-Rule: <ID>` or `Task-ID: CTX-<NNN>` inside the body or footer of the git commit. The CI/CD parser relies on this to map changes correctly.
*   **Avoid Local Modifications**: Agents must not manually edit the `CHANGELOG.md` file during standard feature development, as the CI pipeline will overwrite or conflict with local changes upon merging.
*   **Agent Restriction**: The agent MUST NOT modify the `CHANGELOG.md` file or its headers. The CI pipeline will automatically format commits into standard changelog sections based on commit message structure.

## IV. Exceptions
*   **WIP Commits**: Intermediate work-in-progress commits on a feature branch.
*   **Trivial Assets**: Updating a binary image or simple non-semantic whitespace.
*   **Backfill Commits**: The Step 2 documentation commit (e.g. `docs: backfill git hash`) should NOT be logged in the changelog as this causes infinite recursion.

## V. Layered Versioning Standard
To decouple knowledge base evolution from engine maturity, all releases must follow the **Layered Versioning** scheme defined in ADR-001:

1.  **Vault Version (CalVer):** `YYYY.MM.DD` (e.g., `2026.02.22`).
2.  **Agent Version (SemVer):** `vX.Y.Z` (e.g., `v0.8.0`).

The primary version header in `00-System/CHANGELOG.md` must combine these:
`## [Vault YYYY.MM.DD] ~ [Agent vX.Y.Z]`

### Automation Note
While the `[Unreleased]` section uses a `{{date}}` placeholder for Step 1, the official version header created during a release must be manually or semi-automatically populated with the final commit hash metadata if the GitHub Action does not explicitly handle header transformation.
