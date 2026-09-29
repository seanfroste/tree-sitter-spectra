# Official grammar pattern atlas (offline)

Patterns below are **short attributed adaptations**, not copied grammars and not substitute language specifications. The four grammars live in the [`tree-sitter` GitHub organization](https://github.com/tree-sitter); upstream code is MIT licensed. Each file link is pinned so an offline agent can apply the explanation without fetching a moving branch; follow links only when a full implementation or current upstream change is required.

## Contents

- [JSON: small grammar, stable fields](#json-small-grammar-stable-fields)
- [Go: precedence and editor-friendly files](#go-precedence-and-editor-friendly-files)
- [JavaScript: explicit ambiguity and queries](#javascript-explicit-ambiguity-and-queries)
- [Python: indentation and scanner state](#python-indentation-and-scanner-state)
- [Pattern selection](#pattern-selection)

## JSON: small grammar, stable fields

Source: [tree-sitter-json `grammar.js`](https://github.com/tree-sitter/tree-sitter-json/blob/254c42a6476413b776221e03982ac8ae159eeb72/grammar.js), [license](https://github.com/tree-sitter/tree-sitter-json/blob/254c42a6476413b776221e03982ac8ae159eeb72/LICENSE), revision `254c42a` (Max Brunsfeld, Amaan Qureshi; MIT).

- `document` starts with repeated `_value`; `_value` is a supertype of object/array/scalars. `pair` assigns `field('key', $.string)` and `field('value', $._value)`: queries can address semantic roles without depending on child positions.
- `string_content` and `escape_sequence` use `token.immediate(...)` so whitespace/extras cannot separate an escape from surrounding string content. `number` uses `token(decimalLiteral)` to produce one token rather than a tree of digit pieces.
- A non-strict `document: repeat(_value)` is an **upstream choice**, not proof every language permits multiple top-level values. Pick a start rule from your own language specification.

## Go: precedence and editor-friendly files

Source: [tree-sitter-go `grammar.js`](https://github.com/tree-sitter/tree-sitter-go/blob/2346a3ab1bb3857b48b29d779a1ef9799a248cd7/grammar.js), [license](https://github.com/tree-sitter/tree-sitter-go/blob/2346a3ab1bb3857b48b29d779a1ef9799a248cd7/LICENSE), revision `2346a3a` (Max Brunsfeld, Amaan Qureshi; MIT).

- A `PREC` table names operator levels (`multiplicative` > `additive` > `comparative` > `and` > `or`); grouping is expressed in parsing rules, not by making `*` a more important lexical token. Use test trees for `a + b * c`, `a - b - c`, and unary/binary overlap.
- `word: $ => $.identifier`, explicit reserved keywords, and `field('name', ...)` on declarations keep keyword boundaries and consumer-facing names predictable. Helper rules like `_top_level_declaration` can remain hidden.
- Its `source_file` intentionally accepts some top-level statements to parse documentation fragments. That is a deliberate recovery/editor tradeoff, **not** permission to silently accept malformed programs in a stricter language. Verify with its [corpus tests](https://github.com/tree-sitter/tree-sitter-go/tree/2346a3ab1bb3857b48b29d779a1ef9799a248cd7/test/corpus) when transferring the idea.

## JavaScript: explicit ambiguity and queries

Sources: [tree-sitter-javascript `grammar.js`](https://github.com/tree-sitter/tree-sitter-javascript/blob/58404d8cf191d69f2674a8fd507bd5776f46cb11/grammar.js), [scanner.c](https://github.com/tree-sitter/tree-sitter-javascript/blob/58404d8cf191d69f2674a8fd507bd5776f46cb11/src/scanner.c), [queries/highlights.scm](https://github.com/tree-sitter/tree-sitter-javascript/blob/58404d8cf191d69f2674a8fd507bd5776f46cb11/queries/highlights.scm), [license](https://github.com/tree-sitter/tree-sitter-javascript/blob/58404d8cf191d69f2674a8fd507bd5776f46cb11/LICENSE), revision `58404d8` (Max Brunsfeld, Amaan Qureshi; MIT).

- `precedences` lists named parsing levels such as member, call, unary, and binary. `conflicts` explicitly names genuinely ambiguous expression/pattern pairs; `prec.dynamic` should only be introduced if ambiguity persists at runtime. If one interpretation is always correct, choose static precedence/associativity instead of GLR branching.
- `externals` includes automatic semicolon, template, regexp, JSX and other tokens that need lexical context. The C scanner supplements `grammar.js`; its enum order must match `externals`. Do not copy the whole scanner into languages without these token families.
- Highlight queries use **grammar fields** rather than arbitrary child positions, e.g. `(call_expression function: (identifier) @function)`; anonymous punctuation and keywords have separate captures. A grammar rename requires query updates and testing, not just regeneration.

## Python: indentation and scanner state

Sources: [tree-sitter-python `grammar.js`](https://github.com/tree-sitter/tree-sitter-python/blob/26855eabccb19c6abf499fbc5b8dc7cc9ab8bc64/grammar.js), [src/scanner.c](https://github.com/tree-sitter/tree-sitter-python/blob/26855eabccb19c6abf499fbc5b8dc7cc9ab8bc64/src/scanner.c), [license](https://github.com/tree-sitter/tree-sitter-python/blob/26855eabccb19c6abf499fbc5b8dc7cc9ab8bc64/LICENSE), revision `26855ea` (Max Brunsfeld; MIT).

- `_newline`, `_indent`, and `_dedent` are external tokens. The scanner maintains indentation and string delimiter state, handles blank/comment lines and bracket contexts, and serializes/deserializes state so incremental parsing can resume correctly. Its `externals` also lists comments and closing delimiters for scanner coordination; these choices are language-specific.
- `valid_symbols` controls which token the scanner may emit. The C implementation uses `lexer->mark_end` before scanning farther, and `advance(..., true)` for skipped whitespace vs `advance(..., false)` for token text. `get_column()` counts code points, not bytes. Compare fresh/incremental parse trees after indentation edits and EOF changes; a passing normal corpus cannot verify saved scanner state.
- `PREC` names expression precedence, while `conflicts` handles genuinely ambiguous pattern/expression forms. Never use a conflict array as a blanket fix for generator warnings.

## Pattern selection

| Need | Primary example | Transfer test |
| --- | --- | --- |
| Stable query/binding API | JSON fields, JavaScript highlight query | Query captures and `node-types.json` consumer compatibility. |
| Operator precedence | Go / JavaScript `PREC` | Distinct associativity and grouping corpus cases. |
| Context-sensitive lexical boundary | JavaScript scanner | Malformed input followed by intact next construct; EOF. |
| Layout-sensitive edits | Python scanner | Incremental parse vs fresh parse after indent/dedent and blank-line edits. |
| Permissive editor fragments | Go source-file rule | Decide explicitly whether valid source, recovery, or unparsed fragment. |
