# SPECTRA repository completeness implementation plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use subagent-driven-development or executing-plans. Steps use checkbox syntax for tracking.

**Goal:** Complete the reviewed repository language and its tested Go editor integration.

**Architecture:** One generic Tree-sitter grammar, a source-linked reviewed coverage oracle, and a small Go semantic/editor adapter. Parent owns integration, grammar generation, and final verification; independent workers own disjoint inventory and adapter files.

**Tech Stack:** Tree-sitter CLI 0.27.0; ABI 15; Node.js 24 or later; Go module and existing go-tree-sitter dependency.

This plan records the approved implementation sequence, not live task status. Its unchecked boxes remain historical planning text; see `HANDOFF.md` and `npm run verify` for current state.

**Spec:** `docs/development/completeness-design.md`

## Global Constraints

- Retain Tree-sitter CLI 0.27.0, parser ABI 15, Node.js 24 or later, and the existing Go binding dependency unless a demonstrated compatibility requirement forces a documented change.
- Preserve all 138 source candidate identities and original inputs; historical parse output is not an expected-result oracle.
- Original reference files are not rewritten to hide defects. Corrections and local policies are explicit in the reviewed inventory.
- Full LSP transport and simulation execution are out of scope.
- Use CLI executables by name, not absolute paths.
- Shared-checkout workers must not generate, build, lint, format, test, or commit mid-flight. Parent runs verification after integration.
- No production dependencies beyond the existing toolchain are planned.

## Task 1: M1 reviewed documentation oracle and integration gate

Files: create `test/documentation.json`, `bindings/go/documentation_test.go`, and update development coverage documentation. Parent owns the Go runner; inventory worker owns the JSON and its policy documentation.

Interface: inventory object has `version`, `source_sha256`, `cases`, and `other_fenced_material`. Each case has `id`, `document`, `source_lines`, `original`, `input`, `disposition` (`verbatim`, `corrected`, `local-policy`, `non-executable`), `reason`, and `expect`. `expect` contains literal ordered `statements` (source spellings), `parameters` (`name`, `value`, `kind`), `references` (source spelling), and `calls` (`name`, `arguments` as ordered source substrings). Expected node kinds are existing kinds plus `linked_series`, `mask_group`, `binary_expression`, `unary_expression`, `parenthesized_expression`, `formatted_expression`. Assertions cover actual parameter value node spans and field relationships, not only labels. All executable inputs receive a final LF in the runner.

- [ ] Review every original against its reference context; record evidence and independent expectations.
- [ ] Add END and syntax-family regressions with independently derived trees/values.
- [ ] Implement a runner that validates source hashes, candidate membership, explicit dispositions, expected semantic projections, clean trees, and fixture cleanliness/reference classification.
- [ ] Run the new gate before grammar changes; retain failure output as evidence, not an allowlist.

Example acceptance: `REGION MASK=FILE (mask.txt,cell,3)` has one parameter `MASK`, value kind `function_call`, one call `FILE`, and exactly three arguments `mask.txt`, `cell`, `3`. `DEFINE NAME=Width VALUE=#Size(0.1 1.0)` has a `linked_series` value and reference `#Size`.

## Task 2: M2 grammar completion and M4 recovery

Files: `grammar.js`, generated `src/*`, existing/new `test/corpus/*`, `queries/highlights.scm`.

Consumes Task 1 oracle. Produces existing generic statement fields plus explicit linked-series, mask-group, and structured-expression nodes. Preserve ordinary function/series argument node kinds.

- [ ] Diagnose lexing and parse precedence at failed syntax families from the baseline and new gate.
- [ ] Implement references, spaces before calls, linked series, mask groups, blank whitespace lines, and keyword-boundary handling.
- [ ] Replace permissive expression token sequences with operand/operator structure while retaining documented formatting and interpolation.
- [ ] Prevent newline consumption by unfinished quoted values; establish recoverable boundaries for incomplete groups and assignments.
- [ ] Regenerate ABI 15 parser; run corpus and the complete oracle. Fix behavior rather than accepting incidental parser output.

Regression inputs include `SUBS V=@Width`, `SAVE FILE=$INPUTFILE`, `DEFINE NAME=x VALUE={(@Width+1)*2}`, `DEFINE NAME=x VALUE={1 + * 2}`, `GRID XMIN=\nSAVE FILE=after.dat`, unfinished strings followed by SAVE, and unknown `ENDPOINT X=1`.

## Task 3: M3 Go semantic adapter

Files: create `bindings/go/analysis.go`, `bindings/go/analysis_test.go`, and `docs/development/adapter-contract.md`. Only this worker edits these files.

Consumes Tree-sitter source root and source bytes. Public API: `Analyze(root *tree_sitter.Node, source []byte) Analysis`; exported value records have byte ranges and retained source spelling. `Analysis` exposes statements, declarations, references, conditionals, includes, and diagnostics. Fields and exact helper APIs are owned by this worker and documented for consumers. No tree pointers survive in returned values.

- [ ] Define literal behavior tests for logical ownership across blank/comment lines, END, malformed lines, aliases, declarations/reference kinds, mismatched conditional markers, and include extraction.
- [ ] Implement direct named-field traversal without reparsing source syntax.
- [ ] Resolve include files relative to the containing path with explicit file errors; symbolic/interpolated filenames produce explicit unresolved results, not guessed paths.
- [ ] Document contract version and case sensitivity, ownership, diagnostics, source ranges, and catalogue limits.

Example behavior: `DEFINE NAME=Width VALUE=1\n# comment\n CHARACTER=wide\nSAVE FILE=$Width\nEND\nSAVE FILE=ignored` yields one logical DEFINE with its continuation, a character reference to Width, and the last SAVE retained but inactive.

## Task 4: M4 Go positions and incremental editor proof

Files: create `bindings/go/positions.go`, `bindings/go/positions_test.go`, `bindings/go/editor_test.go`. Only this worker edits these files.

Public API: `PositionEncoding` with `UTF8` and `UTF16`; `Position{Line, Character uint}`; `PositionAt(source []byte, offset uint, encoding PositionEncoding) (Position, error)` and `ByteOffset(source []byte, position Position, encoding PositionEncoding) (uint, error)`. Invalid UTF-8, unsupported encoding, offsets inside encoded runes/CRLF, out-of-range positions, and UTF-16 surrogate splits return errors. Line terminators are not editable character positions; EOF is valid.

- [ ] Add literal round-trip cases for Unicode, supplementary characters, CRLF/LF, tabs, EOF, and malformed boundaries before implementation.
- [ ] Implement explicit byte/position conversion using Go standard library only.
- [ ] Add binding-level incremental versus fresh tree comparisons over insertions, deletions, malformed transitions, ranges, and clean recovery to later statements.
- [ ] Add deterministic larger-document benchmarks; parent records measurements and sets a conservative enforceable smoke budget from evidence.

Example: source `a😀β\r\nZ` has UTF-16 position `(0,3)` at byte 5; `(0,2)` is invalid because it splits the emoji. Byte 8 is invalid because it lies between CR and LF.

## Task 5: M5 integration, CI, review, and runtime proof

Files: `package.json`, new `scripts/verify.mjs`, new `.github/workflows/verify.yml`, `README.md`, `docs/development/completeness.md`, syntax coverage matrix; any targeted integration fixes remain owned by parent.

- [ ] Integrate all slices and run generation, corpus, Go tests, document/fixture gates, query checks, and incremental benchmarks.
- [ ] Implement one verification command with nonzero exit on every failing subcommand and byte-for-byte generated artifact comparison after regeneration.
- [ ] Add CI for the same toolchain and verification command.
- [ ] Exercise CLI parsing and a throwaway Go consumer over changed syntax, malformed recovery, analysis, and position conversion; remove throwaway artifacts.
- [ ] Obtain independent spec/quality review; correct findings and run focused verification of each amended path.
- [ ] Update public documentation with measured evidence, explicit local policies, contract references, and exact completeness boundary. Do not replace historical evidence with current results.

Completion: every acceptance criterion in the approved design is exercised. No ignored executable examples, unexpected fallbacks, missing callers, or unsupported local-policy gaps remain.
