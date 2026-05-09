# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
* [feat(core): initialize VerdiPitchEngine workspace with Go...](https://github.com/empawlik/VerdiPitchEngine/commit/c6744793eef814ee4264d00efbe189dd0e238eba)
  * **feat(core)**: initialize VerdiPitchEngine workspace with Go pipeline and Antigravity standards (Ref-Rule: 100-CORE, Task-ID: TASK-001)
* [test(converter): implement end-to-end ffmpeg pitch shift ...](https://github.com/empawlik/VerdiPitchEngine/commit/15cd97aadf5e46c7b44fa2d6c58e3bc3aa7b3ef9)
  * **test(converter)**: implement end-to-end ffmpeg pitch shift validation (Task-ID: VPE-002)

### Changed
* [docs(VPE-002): execute post-mortem gap analysis and task ...](https://github.com/empawlik/VerdiPitchEngine/commit/b74772e21f28c4414011f366252b07d74f4d3810)
  * **docs(VPE-002)**: execute post-mortem gap analysis and task generation (Task-ID: VPE-002)
* [docs(task-validate): validate VPE-002 and transition to a...](https://github.com/empawlik/VerdiPitchEngine/commit/b301e0d7d3498d21081a804083b5af94f63a28e4)
  * **docs(task-validate)**: validate VPE-002 and transition to active backlog (Task-ID: VPE-002)
* [docs(055-gap-analysis-output): generate explorative task ...](https://github.com/empawlik/VerdiPitchEngine/commit/21bd8b72c79a8b068bbcfe30cf725bf593f6d020)
  * **docs(055-gap-analysis-output)**: generate explorative task post-mortem for VPE-001 (Ref-Rule: 055-gap-analysis-output, Task-ID: VPE-001)
* [chore(025-changelog): add automated release note generati...](https://github.com/empawlik/VerdiPitchEngine/commit/9d9809a31e6a1c19ec613ae5dd168dd84504af0e)
  * **chore(025-changelog)**: add automated release note generation workflow (Ref-Rule: 025-changelog)
* [chore(000-workspace): replicate core CI/CD workflows from...](https://github.com/empawlik/VerdiPitchEngine/commit/272e4de540afbe628aeb87a6f7a08c952c117c55)
  * **chore(000-workspace)**: replicate core CI/CD workflows from Cortex (Ref-Rule: 000-workspace)
* [docs(VPE-002): finalize task formalization and metadata u...](https://github.com/empawlik/VerdiPitchEngine/commit/2b43d76a5f950bb1d17c0912f982b00fa11854f9)
  * **docs(VPE-002)**: finalize task formalization and metadata updates (Task-ID: VPE-002)
* [docs(050-docs): formalize VPE-001 task completion and gen...](https://github.com/empawlik/VerdiPitchEngine/commit/aaeffe4e84bbf98f3cf55e5df7c7bef0574f708e)
  * **docs(050-docs)**: formalize VPE-001 task completion and generate session artifacts (Ref-Rule: 050-environmental-documentation, Task-ID: VPE-001)
* [Update README.md](https://github.com/empawlik/VerdiPitchEngine/commit/49fae44357368cd712ebf5306f8d5d7071a19588)
  * **chore**: Update README.md

### Fixed
* [fix(025-changelog): initialize changelog and base tag](https://github.com/empawlik/VerdiPitchEngine/commit/629fa5dc630dd338961f58714d8e4f270e3f4e88)
  * **fix(025-changelog)**: initialize changelog and base tag (Ref-Rule: 025-CHANGELOG, Task-ID: VPE-001)
* [fix(000-workspace): resolve unknown target errors in CI w...](https://github.com/empawlik/VerdiPitchEngine/commit/15f9970e96c6eb4fd32d3e3d50a87aff9d8dd9cd)
  * **fix(000-workspace)**: resolve unknown target errors in CI workflow (Ref-Rule: 000-workspace)

### Dependencies
* [build(deps): make mage a direct dependency](https://github.com/empawlik/VerdiPitchEngine/commit/1942c860140ea6ee6ae110b428d78d3856519e08)
  * **build(deps)**: make mage a direct dependency (Ref-Rule: 900-go-standards, Task-ID: VPE-001)
