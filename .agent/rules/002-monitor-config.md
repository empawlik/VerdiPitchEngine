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
created: "2025-12-26"
updated: 2026-05-09
body_hash: bfbcf42e63f109c7
tags: [dev-asset, antigravity-context, workspace-rule]
---

## 📂 AI Agent Rule: Sub-repo Monitor & Directive Protocol

This rule defines the operational logic for the Antigravity/Gemini agent when interacting with monitored documentation sub-repositories. 

---

# .agent/rules/monitor-config.md

## 🎯 Purpose
To provide the agent with a semantic framework for interpreting external markdown assets and executing autonomous directives found within them.

## 🔍 Monitoring & Indexing Logic
1.  **Continuous Audit:** Monitor all files within `.agent/rules/`, `.agent/knowledge/`, and custom `/docs` directories for filesystem changes.
2.  **Semantic Filtering:**
    * Compare the current file state against the previous Git hash.
    * Ignore changes identified as "Syntactic Noise" (e.g., typos, white space, header re-ordering).
    * Flag "Functional Deltas" (e.g., logic changes, requirement updates, or constant modifications) as Actionable Tasks.

## ⚡ Directive Interpretation
The agent must scan for the following tags within any markdown file in a monitored directory and adjust its execution policy accordingly:

* **``**:
    * **Action**: Elevate current execution policy to **Turbo** for this specific file.
    * **Requirement**: Implement suggested code changes immediately without a manual "Accept" gate.
* **``**:
    * **Action**: Insert derived tasks at the top of the Mission Control queue.
* **``**:
    * **Action**: Automatically create unit tests in the `/tests` directory that validate the requirements described in the spec file.

## 🛡️ Safety & Attribution Constraints
* **Author Verification**: Directives (specifically `@turbo`) are only valid if the commit author of the sub-repo is present in the `trusted_authors` list within `.agent/config.json`.
* **Impact Radius**: If a `@turbo` directive requires modification of more than 10 files, downgrade the policy to **Auto** and present an Implementation Plan for review.
* **Commit Metadata**: Every commit generated via this monitor must include the following footer:
    * `Source-Spec-Hash: [Sub-repo Commit Hash]`
    * `Co-authored-by: [Spec Author Name]`

---
*End of Protocol Configuration*
