---
project_name: "VerdiPitchEngine"
version: 0.1.0
status: "active"
priority: "high"
dev_stage: "development"
agent_role: "Core-Context"
agent_weight: 4.0
asset_scope: "Global"
platform: "CLI"
tech_stack: ["Go", "Shell", "ffmpeg"]
dependencies: []
created: "2026-05-10"
updated: 2026-05-10
body_hash: 665ea34e3c0afe28
tags: [dev-asset, memory, roon, inotify, deadlock]
---

# VPE-020: Roon Inotify Race Conditions & Go Pipe Deadlocks

## Context
During the implementation of the `metaflac` 1:1 parity injection pipeline to bypass `ffmpeg`'s buggy `-map_metadata` handling (which stripped MusicBrainz IDs and PICTURE blocks), two critical architectural edge cases were uncovered and resolved.

## Resolution 1: Go Concurrent `os/exec` Pipe Deadlocks
When bridging two processes (`metaflac --export` -> `metaflac --import`) concurrently in Go using `StdoutPipe()` to supply standard input to the second process, `Wait()` behavior can cause silent deadlocks. Calling `Wait()` on the exporter abruptly shuts down the underlying OS pipe upon process exit. If the downstream importer is still draining the buffer or starting up, it fails to receive a clean `EOF` and hangs the Goroutine indefinitely.
- **Fix:** Switched to a bounded `bytes.Buffer` execution. The export command fully flushes its payload to RAM, closes gracefully, and hands off the buffer sequentially to the importer. Memory overhead remains strictly minimal (< 5MB per FLAC file).

## Resolution 2: Media Scanner `inotify` Race Conditions
Roon relies heavily on `inotify` triggers and directory/file modification timestamps (`mtime`) and creation timestamps to infer "Recently Added" status.
- **The Bug:** `os.Rename(tmpOutPath, outPath)` followed by `os.Chtimes` caused a race condition where Roon scanned the renamed file **before** the historical `mtime` was restored, explicitly flagging the old album as newly added. Furthermore, `walker.go` neglected to preserve directory timestamps and non-FLAC artwork timestamps during copy operations.
- **Fix:** The `os.Chtimes` execution on the `tmpOutPath` is now strictly enforced **before** the atomic `os.Rename`. `walker.go` explicitly restores `mtime` onto all created subdirectories and supplemental assets via `io.Copy`, perfectly concealing the background processing from the media scanner's file-watcher logic.
