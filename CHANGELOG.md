# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2026-05-09

### Added
* [feat(100-core): stabilize audio pipeline and implement ba...](https://github.com/empawlik/VerdiPitchEngine/commit/a78df48de5b04503ec9f9acfde9f1bcc1380d37b)
  * **feat(100-core)**: stabilize audio pipeline and implement batch orchestrator
* [feat(100-core): implement real-time progress bars and con...](https://github.com/empawlik/VerdiPitchEngine/commit/237a9b4fc4d559ff8a54fcd37558d80a12b1217b)
  * **feat(100-core)**: implement real-time progress bars and context timeouts

### Changed
* [test(100-core): implement robust CI environment mocks for...](https://github.com/empawlik/VerdiPitchEngine/commit/7ee8b8efd6fa4193390dbf61e7608d28e454d67b)
  * **test(100-core)**: implement robust CI environment mocks for metadata extraction
* [docs(050-docs): formalize batch orchestrator gap analysis...](https://github.com/empawlik/VerdiPitchEngine/commit/2d8b3bab7cc45493fa544299c2616d8a99f51247)
  * **docs(050-docs)**: formalize batch orchestrator gap analysis and runbooks
* [docs(055-docs): generate explorative post-mortem for VPE...](https://github.com/empawlik/VerdiPitchEngine/commit/c83b00b040c686616d274f2292019d5d71167e1a)
  * **docs(055-docs)**: generate explorative post-mortem for VPE-010
* [docs(055-docs): document UI and context safety features f...](https://github.com/empawlik/VerdiPitchEngine/commit/cf265d2848a1abdbd859daf65eda2bfed5b5435e)
  * **docs(055-docs)**: document UI and context safety features for VPE-010

## [0.1.0] - 2026-05-09

### Added
* [feat(engine): implement librubberband TSM and container o...](https://github.com/empawlik/VerdiPitchEngine/commit/c90498ffe63b6fe0dd78953039da3f4704cf47bc)
  * **feat(engine)**: implement librubberband TSM and container orchestration (Ref-Rule: 100-CORE, Task-ID: VPE-007)
* [refactor(VPE-007): implement asetrate optimizations and z...](https://github.com/empawlik/VerdiPitchEngine/commit/50881618e4ce696cda710adc0510a73d4b80c1f0)
  * **refactor(VPE-007)**: implement asetrate optimizations and zero-trust deployment (Ref-Rule: 1000, Task-ID: VPE-007)
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
* [test(converter): extract buildFFmpegArgs to increase test...](https://github.com/empawlik/VerdiPitchEngine/commit/7cf16b7d5cfae90fd404bfc452779d7f7ed6c982)
  * **test(converter)**: extract buildFFmpegArgs to increase test coverage (Task-ID: VPE-007)
* [chore(025-changelog): add automated release note generati...](https://github.com/empawlik/VerdiPitchEngine/commit/9d9809a31e6a1c19ec613ae5dd168dd84504af0e)
  * **chore(025-changelog)**: add automated release note generation workflow (Ref-Rule: 025-changelog)
* [chore(000-workspace): replicate core CI/CD workflows from...](https://github.com/empawlik/VerdiPitchEngine/commit/272e4de540afbe628aeb87a6f7a08c952c117c55)
  * **chore(000-workspace)**: replicate core CI/CD workflows from Cortex (Ref-Rule: 000-workspace)
* [docs(055-postmortem): generate VPE-007 explorative task p...](https://github.com/empawlik/VerdiPitchEngine/commit/18dc9731d0755b213ea17a5c048afe5db58ab5b9)
  * **docs(055-postmortem)**: generate VPE-007 explorative task post-mortem (Ref-Rule: 055-DOCS, Task-ID: VPE-007)
* [docs(050-formalization): formalize VPE-007 documentation ...](https://github.com/empawlik/VerdiPitchEngine/commit/acae4f7c05ca6caffc5e79d28702191e6cd6e313)
  * **docs(050-formalization)**: formalize VPE-007 documentation and runbooks (Ref-Rule: 050-DOCS, Task-ID: VPE-007)
* [docs(VPE-002): finalize task formalization and metadata u...](https://github.com/empawlik/VerdiPitchEngine/commit/2b43d76a5f950bb1d17c0912f982b00fa11854f9)
  * **docs(VPE-002)**: finalize task formalization and metadata updates (Task-ID: VPE-002)
* [docs(050-docs): formalize VPE-001 task completion and gen...](https://github.com/empawlik/VerdiPitchEngine/commit/aaeffe4e84bbf98f3cf55e5df7c7bef0574f708e)
  * **docs(050-docs)**: formalize VPE-001 task completion and generate session artifacts (Ref-Rule: 050-environmental-documentation, Task-ID: VPE-001)
* [Update README.md](https://github.com/empawlik/VerdiPitchEngine/commit/49fae44357368cd712ebf5306f8d5d7071a19588)
  * **chore**: Update README.md

### Fixed
* [chore(000-workspace): add release.sh to fix github action...](https://github.com/empawlik/VerdiPitchEngine/commit/3beeec06eacda8cfc81fd2049b1ef4bdb971b915)
  * **chore(000-workspace)**: add release.sh to fix github action pipeline (Ref-Rule: 000-workspace, Task-ID: VPE-007)
* [fix(025-changelog): initialize changelog and base tag](https://github.com/empawlik/VerdiPitchEngine/commit/629fa5dc630dd338961f58714d8e4f270e3f4e88)
  * **fix(025-changelog)**: initialize changelog and base tag (Ref-Rule: 025-CHANGELOG, Task-ID: VPE-001)
* [fix(000-workspace): resolve unknown target errors in CI w...](https://github.com/empawlik/VerdiPitchEngine/commit/15f9970e96c6eb4fd32d3e3d50a87aff9d8dd9cd)
  * **fix(000-workspace)**: resolve unknown target errors in CI workflow (Ref-Rule: 000-workspace)

### Dependencies
* [build(deps): make mage a direct dependency](https://github.com/empawlik/VerdiPitchEngine/commit/1942c860140ea6ee6ae110b428d78d3856519e08)
  * **build(deps)**: make mage a direct dependency (Ref-Rule: 900-go-standards, Task-ID: VPE-001)
