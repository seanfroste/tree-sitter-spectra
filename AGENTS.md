# Repository agent instructions

## Local terminal environment

Machine-specific inventory: username `seanf`, hostname `GEYSER`, Windows, PowerShell (`pwsh`) 7.6.6. These observations apply only to this user/host, not to other clones.

Available CLI tools (confirmed on PATH): 7z.exe, bash, cargo, clangd, gcc, git, go, gofmt, node, npm, pwsh, rust-analyzer, rustc, tree-sitter, where.exe

Whenever another CLI tool is discovered and confirmed available, append its command name to this comma-separated list for the matching user/host. Record a separate user/host inventory on another machine; do not assume this list applies there. Invoke tools by command name, not absolute executable path.

Use Windows PowerShell quoting, environment variables, and command semantics. For compound PowerShell commands through a shell-neutral tool runner, invoke `pwsh -NoProfile -Command '...'` explicitly. `bash` here resolves to the Windows WSL launcher, not Git Bash. Read `$PROFILE` when profile aliases or initialization matter; the observed profile is `C:\Users\seanf\OneDrive\Documents\PowerShell\Microsoft.PowerShell_profile.ps1`.

## Completeness and consumer

The intended LSP consumer is the sibling `../spectrals` repository when present. Before changing the parser/adapter contract or claiming feature-completeness, inspect its current parser, analysis, position, and binding requirements. Keep this parser independently buildable; do not silently equate a Go-only smoke test with readiness for that consumer.

Read `docs/development/completeness-design.md` for the approved repository-specification boundary and `docs/development/completeness-plan.md` for the implementation sequence. Preserve source examples and distinguish local policy from historical simulator behavior.

## Commit hygiene

Keep test-run and completeness-verification output out of commits. Store transient logs and probes in ignored scratch space; commit source changes and durable contracts separately from execution output. Do not stage broad directory trees containing generated run artifacts.
