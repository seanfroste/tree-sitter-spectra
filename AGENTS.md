# Repository agent instructions

## Terminal environment

Use CLI tools by command name rather than absolute executable path. On Windows,
use PowerShell quoting and command semantics; for compound commands through a
shell-neutral runner, invoke `pwsh -NoProfile -Command '...'` explicitly.

## Completeness and consumer

The intended LSP consumer is the sibling `../spectrals` repository when present. Before changing the parser/adapter contract or claiming feature-completeness, inspect its current parser, analysis, position, and binding requirements. Keep this parser independently buildable; do not silently equate a Go-only smoke test with readiness for that consumer.

Read `docs/development/completeness-design.md` for the approved repository-specification boundary and `docs/development/completeness-plan.md` for the implementation sequence. Preserve source examples and distinguish local policy from historical simulator behavior.

## Commit hygiene

Keep test-run and completeness-verification output out of commits. Store transient logs and probes in ignored scratch space; commit source changes and durable contracts separately from execution output. Do not stage broad directory trees containing generated run artifacts.
