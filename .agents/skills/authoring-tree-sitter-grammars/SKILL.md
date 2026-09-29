---
name: authoring-tree-sitter-grammars
description: Use when creating or modifying a Tree-sitter grammar, resolving generation conflicts, writing an external scanner, validating parser recovery, or preparing a parser for editor consumers.
---

# Authoring Tree-sitter grammars

## Core principle

Design the concrete syntax tree for its consumers, then prove each language construct and recovery boundary with source examples. A generated parser or green corpus alone does not establish editor readiness.

## Workflow

1. Identify the language specification, representative valid/invalid source, and consumer-visible nodes, fields, queries, and ABI requirements. For a new project, use `tree-sitter init`; preserve existing repository conventions when extending one.
2. Add a corpus case with the intended tree **before** changing `grammar.js`. Start with broad syntax categories; add fields to stable semantic children and hidden rules/supertypes for unhelpful wrappers. Read [official grammar reference](reference/grammar-guide.md) for the DSL, lexical decisions, ambiguity, and scanner contracts.
3. Change `grammar.js` (and `src/scanner.c` only when tokenization needs state or lookahead). Generate, inspect a real parse, and run focused and full corpus tests using [CLI workflow](reference/cli-workflow.md). Regenerate `src/parser.c`, `src/grammar.json`, and `src/node-types.json`; never hand-edit generated files.
4. Check valid samples for intended nodes/fields and no unexpected `ERROR`/`MISSING`; check malformed samples for localized recovery and intact following constructs. Compare incremental edits with fresh parses; exercise queries, intended language bindings, and ABI compatibility. Use the [upstream pattern atlas](reference/upstream-patterns.md) to choose patterns, not to copy another language's semantics.
5. Review changed node types and query/binding consumers before release. Update expected corpus trees only after inspecting the new tree; do not use snapshot updates as a substitute for deciding the contract.

## Decision quick reference

| Symptom | First check |
| --- | --- |
| Unexpected token or keyword prefix | `word`, `reserved`, token length, lexical precedence, and `extras` |
| `Unresolved conflict` | Intended associativity/static precedence vs genuine ambiguity (`conflicts`) |
| String/comment swallows next construct | Token boundaries and scanner recovery at newline/EOF |
| Syntax tree is noisy or query breaks | Hidden rules, fields, aliases, supertypes, `node-types.json` |
| Corpus passes but editor breaks | Incremental edits, `ERROR`/`MISSING` spans, queries, ABI, binding build |

**CLI:** [commands and end-to-end checks](reference/cli-workflow.md) · **Grammar/scanner:** [official reference](reference/grammar-guide.md) · **Real grammars:** [attributed upstream examples](reference/upstream-patterns.md).

## Common mistakes

- `prec` outside `token()` resolves parse decisions; `token(prec(...))` resolves competing tokens. `conflicts` is for real ambiguity, not a universal generator-error escape hatch.
- In scanners, `advance(lexer, true)` **skips** the current character from the token; `advance(lexer, false)` **includes** it. `mark_end` excludes subsequent lookahead from the token. Serialize all state affecting subsequent tokens.
- `tree-sitter test --update` does **not** update cases whose resulting trees contain `ERROR` or `MISSING`; it cannot validate expected semantics. Review the tree and test diff first.
