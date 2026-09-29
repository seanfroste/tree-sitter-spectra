# Tree-sitter CLI: parser development, offline command guide

Based on the official [CLI reference](https://github.com/tree-sitter/tree-sitter/tree/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli) at revision `790e6ad` ([MIT license](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/LICENSE)). Generation, focused testing, parse edits, queries, build, and fuzz were exercised with CLI 0.27.0; consult `tree-sitter <command> --help` for other installed versions. Run from a parser repository unless a command specifies `--grammar-path`. Quote paths/regexes for your shell (examples work in PowerShell). Commands marked **writes** modify the repo or build outputs.

## Contents

- [Start and generate](#start-and-generate)
- [Corpus and real-source debugging](#corpus-and-real-source-debugging)
- [Queries, highlights, and consumers](#queries-highlights-and-consumers)
- [Incremental, fuzz, release](#incremental-fuzz-release)

## Start and generate

| Command | Purpose and output |
| --- | --- |
| `tree-sitter --version` | Confirm CLI version before interpreting flags/default ABI. |
| `tree-sitter init-config` | **Writes** a personal CLI config (optional; CLI works without it). Set `parser-directories` there if language discovery/highlighting cannot find the grammar. Keep this machine-specific file out of source commits. |
| `tree-sitter init` | **Writes** scaffold: `grammar.js`, `tree-sitter.json`, manifests and enabled binding templates. `tree-sitter init -u` updates eligible generated scaffold files, not arbitrary manually edited ones. Do not run blindly in an existing parser repo. |
| `tree-sitter generate` | **Writes** `src/parser.c`, `src/grammar.json`, `src/node-types.json` and support headers, from `grammar.js` (or supplied structured grammar path). Run after grammar changes; never hand-edit these outputs. |
| `tree-sitter generate --abi 14` | **Writes** parser targeting ABI 14 for consumers that cannot load ABI 15. Check runtime compatibility first; current CLI default is 15. |
| `tree-sitter generate --log` | Expose generator decisions, recovery states, conflicts; use when grammar generation fails or performance regresses. `--report-states-for-rule RULE` reports rule states. |
| `tree-sitter build` | **Writes** native shared parser library (`.dll`/`.so`/`.dylib`). Test real packaging separately; `build --wasm` requires or downloads WASI SDK. |

[Init config](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/init-config.md) · [Init](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/init.md) · [Generate](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/generate.md) · [Build](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/build.md). `tree-sitter.json` supplies grammar scope, file types, query paths, bindings, and metadata. If parsing a file cannot select a language, inspect that metadata or use `--scope` where supported.

## Corpus and real-source debugging

1. Write a `test/corpus/<topic>.txt` case: name between `====` lines, input, `---`, expected **named-node** S-expression, optionally `field:` names. Add valid, malformed, EOF, multiline, CRLF, keyword-prefix, and nested-expression examples relevant to the language.
2. `tree-sitter generate` then `tree-sitter test`. For one failing test, run `tree-sitter test -i 'part of case name'` (`-i` takes a regex); `tree-sitter parse -n 3` inspects corpus test number 3. Use `tree-sitter test --show-fields` for field diffs; `tree-sitter test --debug` or `tree-sitter parse --debug` for lexer/parser logs.
3. Save actual source to a file, then `tree-sitter parse path/to/sample.ext`. Default output has node ranges; `--cst` shows full concrete syntax, `--no-ranges` makes named-node S-expressions easy to copy, `--stat` shows parse statistics, `--time` measures parsing, and `--json-summary` yields machine-readable results. With no file paths, `parse` reads stdin. It exits nonzero on parse errors; for malformed source, inspect location and subsequent nodes rather than suppressing errors.
4. `tree-sitter test --update` (**writes** corpus expectations) is a review aid after changing the tree contract, **not** a grammar fix. The CLI will **not update** cases whose current tree contains `ERROR` or `MISSING`; on clean cases, review the diff before accepting the new expected tree. `:error` corpus attribute checks invalidity, not desired recovery structure.

[Test](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/test.md) · [Parse](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/parse.md) · [Corpus format](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/5-writing-tests.md). Avoid `--debug-graph` unless you want generated `log.html`; leave transient logs/build outputs out of source commits.

## Queries, highlights, and consumers

- Inspect `src/node-types.json`, then write `queries/highlights.scm` or other required queries using actual node types and fields. `(ERROR)` and `(MISSING)` are distinct queryable forms. `tree-sitter query queries/highlights.scm path/to/sample.ext` prints matches; `--captures` orders by capture and `--test` runs query tests. For scoped queries, `--byte-range START:END` limits traversal; `--containing-byte-range START:END` retains only fully contained matches.
- `tree-sitter highlight path/to/sample.ext` exercises color captures; use `--html` for visual inspection or `--check` to validate strict capture names. To override default query location, `--query-paths queries/highlights.scm`. CLI `query` proves query matching, not that a particular editor supports the capture names.
- Compile/load the parser and queries through each **intended binding**. Check generated parser ABI against the installed runtime: the documented 0.24 runtime accepts parser ABI 13–14; 0.25+ accepts 13–15. Building a shared library with the CLI does not establish language-binding compatibility.

[Query](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/query.md) · [Highlight](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/highlight.md) · [Query syntax](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/using-parsers/queries/1-syntax.md) · [ABI table](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/using-parsers/7-abi-versions.md).

## Incremental, fuzz, release

- `tree-sitter parse path/to/sample.ext --edits '0,0 0 #' --time` applies an edit after the initial parse (example inserts `#` at row 0, column 0; adapt insertion to the language). Syntax: `row,col|byte_position removed_byte_count insertion_text`, with zero-indexed coordinates. Put the file path **before** `--edits` (its variable-length arguments otherwise consume the path). Compare with fresh parsing of the resulting bytes in the intended binding; CLI timing alone does not prove incremental equivalence.
- `tree-sitter fuzz` mutates corpus inputs and checks tree equality after undoing edits and changed-range consistency. Options `--iterations N`, `--edits N`, `--include REGEX`; fuzzing complements, not replaces, authored malformed-input and scanner-state tests.
- When releasable, `tree-sitter version 1.0.0` (**writes** versions across supported manifests/lockfiles); inspect the changes and regenerate/test. Only tag/publish after checking consumer-visible node/query changes and supported ABI; major versions for incompatible tree changes. Do not treat publishing as a verification step.

[Parse edits](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/parse.md) · [Fuzz](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/fuzz.md) · [Version](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/cli/version.md) · [Publishing](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/6-publishing.md).
