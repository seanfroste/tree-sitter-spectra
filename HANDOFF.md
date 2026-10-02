# Parser handoff

## Scope and authority

Continue against `docs/development/completeness-design.md` and the reviewed
source dispositions in `test/documentation.json`. The contract is the checked-in
repository specification and explicit local policies, not assumed compatibility
with an unavailable simulator specification. Full LSP transport, simulation
execution, and the sibling `../spectrals` repository are outside this parser
work; that LSP still uses Rust.

## Parser and consumer contract

The grammar is `grammar.js`; recovery scanner is `src/scanner.c`. Generated ABI
15 files are `src/parser.c`, `src/grammar.json`, and `src/node-types.json`; run
`tree-sitter generate --abi 15` after grammar changes. The Go package under
`bindings/go` exposes `Language`, `Analyze`, `ResolveIncludes`, `PositionAt`,
and `ByteOffset`. Contract version 1 is in
`docs/development/adapter-contract.md`; node names are public API.

Run `npm ci`, then `npm run verify`. The gate covers parser reproducibility,
corpus, documentation and fixture checks, queries, Go editor/adapter tests, and
Rust binding tests. Preserve original examples and historical inventory; the
approved design, plan, assessment, source policies, and syntax matrix remain
maintained project evidence, not disposable handoff artifacts.

## Release work in progress

Target: `v0.1.0`; selected license: MIT. The release candidate is in progress;
it is not published until the verified source, tag, public visibility, and
GitHub release exist. The parser is not a completed Go LSP or proof of full
historical simulator compatibility.

The sibling migration identifiers remain repository
`seanfroste/tree-sitter-spectra`, parser tag `v0.1.0`, and branch
`migration/go-tree-sitter` unless a newer approved plan replaces them. Preserve
SPECTRALS contracts during any future Go LSP work; remove the Rust implementation
only after complete Go parity and release builds.

## Repository hygiene

Do not commit test logs, probes, local build outputs, `.superpowers/`, `.cache/`,
`target/`, `node_modules/`, or local `*.skill` archives. Keep the portable
Tree-sitter grammar skill source under `.agents/skills/`.
