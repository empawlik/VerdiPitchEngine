---
project_name: VerdiPitchEngine
version: 1.0.0
status: "Active"
priority: "Medium"
dev_stage: "Production"
agent_role: "Core-Context"
agent_weight: 5
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell"]
dependencies: []
created: "2026-04-26"
updated: 2026-05-10
body_hash: fc79d42d0bf3108b
tags: [dev-asset, antigravity-context, workspace-rule]
---

# 1000-KEYS: Antigravity Keys & Metadata
This template maximizes the **YAML frontmatter** to ensure that Obsidian’s "Properties" UI handles the selection of values while keeping the note body clean and ready for your specific architectural notes.

By focusing on the `antigravity_keys`, you provide the AI with explicit metadata it can use to prioritize this file within its context window.

---

## Dev Project Metadata Template

Markdown

```
---
project_name: VerdiPitchEngine
version: 1.0.0
status: # Options: [Backlog, Discovery, Specification, Development, QA, Active, Archive]
priority: # Options: [Low, Medium, High, Critical, Immediate]
dev_stage: # Options: [PoC, MVP, Alpha, Beta, Production]

# Antigravity Specific Metadata
agent_role: # Options: [Core-Context, Feature-Spec, Technical-Debt, API-Contract, UI-Logic]
agent_weight: # Options: [1, 2, 3, 4, 5] (1 = Low priority context, 5 = Critical Requirement)
asset_scope: # Options: [Global, Component-Specific, Module-Specific]

# Technical Stack
platform: # Options: [Web, Mobile, Desktop, Cross-Platform, CLI]
tech_stack: []
dependencies: []

# Timestamps
created: <% tp.date.now("YYYY-MM-DD") %>
updated: <% tp.date.now("YYYY-MM-DD") %>
body_hash: "e95dc114e3f86314"
tags: [dev-asset, antigravity-context]
---

# Project: {{title}}

> [!ABSTRACT] Core Concept
> Brief one-sentence summary of the project or feature for Antigravity's high-level reasoning.

## 📝 Specifications
- 

## 🔗 References
- 
```

---

### Why these fields matter for Antigravity

- **`agent_role`**: This tells the AI how to treat the information. If set to `API-Contract`, the AI knows it cannot deviate from the defined endpoints. If set to `Core-Context`, it treats the file as general background knowledge.
    
- **`agent_weight`**: When the Obsidian vault (also named OpenBrain) that mirrors the Antigravity project rules and docs grows large, you can use this field to filter which files are sent to the AI's context window. High-weight files (4 or 5) should always be included in the prompt.
    
- **`asset_scope`**: Helps the AI understand if the logic in this file applies to the entire application or just a specific subdirectory/module.
    

### Metadata Update Policy

> [!CAUTION]
> **Update Mandate:** You must **NEVER** update the `updated` date or other metadata fields in the YAML frontmatter of a documentation file unless the **body** of that file has also been modified. "Meaningless" frontmatter-only updates destroy the semantic history of when the knowledge surface actually evolved.

### Managing Drop-Downs in Obsidian

To ensure these fields act as drop-downs in the Obsidian UI:

1. Go to **Settings > Colors & Appearance > Properties**.
    
2. Set the **Property view** to **Visible**.
    
3. When you click a field like `status`, Obsidian will automatically suggest values you have used in other notes.
    
4. _Tip:_ Create a single "Template Seed" note where you fill in all the options once; Obsidian will then index those values for the drop-down list across the entire Obsidian vault (also named OpenBrain) that mirrors the Antigravity project rules and docs.
    

**Todo: Python script or a Dataview query that exports these files into a JSON format optimized for Antigravity's ingestion?**