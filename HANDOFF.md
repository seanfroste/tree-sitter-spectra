# Parser development handoff

## Goal and authority

Continue parser development against `docs/development/completeness-design.md` and `docs/development/completeness-plan.md`. The contract is the reviewed repository specification and its explicit local policies, not assumed compatibility with an unavailable simulator specification. Do not modify sibling `../spectrals`; its current LSP still uses Rust, while planned consumer is a new Go LSP consuming a release tag.

## Current implementation

One Tree-sitter grammar lives in `grammar.js`; hand-written recovery scanner lives in `src/scanner.c`. Generated ABI 15 files are `src/parser.c`, `src/grammar.json`, and `src/node-types.json`. Regenerate only with `tree-sitter generate --abi 15`.

Go API under `bindings/go` provides `Language`, `Analyze`, `ResolveIncludes`, `PositionAt`, and `ByteOffset`. Contract version 1 is documented in `docs/development/adapter-contract.md`. Reviewed executable documentation cases are in `test/documentation.json`; historical source screening is preserved in `docs/development/completeness-baseline.json`. The syntax matrix is `docs/development/syntax-coverage.md`.

Run `npm ci`, then `npm run verify`. The gate runs generation reproducibility, Tree-sitter corpus, Go tests, and `cargo test --locked`. Commit `Cargo.lock` for reproducible Rust test dependencies. Requirements and parser coverage are in `docs/development/completeness-design.md`, `docs/development/completeness-plan.md`, and `docs/development/completeness.md`.

## Verified here

On 2026-09-29, `npm run verify` passed on Windows/amd64 with Tree-sitter CLI 0.27.0, Node.js 24.18.0, Go 1.26.8, and Cargo 1.98.1. It reported 10/10 corpus cases, Go tests passing, three Rust unit tests, and one Rust doctest. These versions describe the prior workstation only; follow CI/toolchain declarations for another machine.

Direct CLI parses proved multiline expression structure, local recovery before a following TITLE, and `unparsed_line` for unsupported reference prefixes. New Go regressions cover these paths. `go test ./bindings/go -run '^$' -bench '^BenchmarkEditorParse$' -benchtime=5x -count=1` measured fresh parsing 45.0 ms/op and incremental middle edit 177.6 ms/op for 131,037 bytes. Incremental parsing was slower. No speedup claim.

## Next development steps

- Pull current `main`. Run `npm ci`, then `npm run verify` on the new machine. Investigate platform-specific failures before extending the grammar.
- Review any new documentation or consumer requirements against the source-linked oracle. Add a failing behavioral test before changes.
- Keep the generated ABI 15 parser reproducible. `scripts/verify.mjs` compares regenerated files byte for byte.
- Validate the future Go LSP in its own repository. This parser is not an integrated LSP or proof of full simulator compatibility.
## Future Go LSP migration

Planned migration identifiers are repository `seanfroste/tree-sitter-spectra`, parser tag `v0.1.0`, and sibling branch `migration/go-tree-sitter`. Preserve them unless a newer approved plan replaces them. Parser node names are public release contract.

Use historical Go and current Rust implementations as behavior oracles, not code to restore wholesale. Preserve SPECTRALS stdio, executable, configuration, capability, and public-name contracts. Keep protocol conversion and gating outside analysis. Future Go analysis consumes stable interfaces; spec generator owns documentation tables and supplemental metadata owns abbreviation minima. Use typed feature outputs and explicit document lifecycle. Remove old Rust implementation only after complete Go parity and release builds.

## Repository hygiene

Do not commit test logs, probes, local build outputs, `.superpowers/`, `.cache/`, `target/`, `node_modules/`, or local `*.skill` archives. The portable skill source is committed under `.agents/skills/authoring-tree-sitter-grammars/`; archive is copied separately by its owner. Inspect exact paths before staging. Keep this handoff synchronized with fresh verified results and remaining work.