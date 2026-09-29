# Tree-sitter parser authoring: offline guide

Adapted technical notes from Tree-sitter's [Creating Parsers guide](https://github.com/tree-sitter/tree-sitter/tree/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers) (revision `790e6ad`, [MIT license](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/LICENSE)); links below identify each primary source. Use the installed CLI's `--help` when its version differs. This is a reference, not a language specification.

## Contents

- [Structure and DSL](#structure-and-dsl)
- [Lexing and ambiguity](#lexing-and-ambiguity)
- [External scanners](#external-scanners)
- [Corpus, recovery, consumers](#corpus-recovery-consumers)

## Structure and DSL

[Getting started](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/1-getting-started.md) · [Grammar DSL](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/2-the-grammar-dsl.md) · [Writing the grammar](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/3-writing-the-grammar.md).

- `tree-sitter init` creates `grammar.js`, `tree-sitter.json`, bindings, and manifests. A JavaScript runtime interprets `grammar.js`; a C compiler builds the generated parser. The **first** rule in `rules` is the start symbol. Grammar output is a **concrete** syntax tree: choose named nodes/fields for constructs consumers need, not a one-to-one transcription of a language-specification CFG.
- Each rule is a function of `$`; `$.foo` refers to a rule. String and regex literals describe tokens. JavaScript regex syntax is not the runtime lexer: Tree-sitter generates matching code, and lookahead/lookbehind assertions are unsupported. `RustRegex('(?i)...')` supports suitable inline flags (only the supported regex subset).

| DSL element | Meaning / use |
| --- | --- |
| `seq(a, b)`, `choice(a, b)` | Sequence or alternative; `choice` order does not encode priority. |
| `repeat(a)`, `repeat1(a)`, `optional(a)` | Zero+, one+, or zero/one occurrences. Avoid nullable recursion. |
| `token(seq(...))` | One terminal, with no named children; accepts terminals, not `$.nonterminal`. |
| `token.immediate(a)` | Require no intervening extras before this token (e.g. escapes within a string). |
| `field('name', rule)` | Stable child relationship exposed to queries/bindings. |
| `alias(rule, $.name)` / `alias(rule, 'name')` | Expose alternative named / anonymous node. |
| `_internal` rule, `inline`, `supertypes` | Hide wrapper, substitute rule definition, or expose abstract subtype group without a visible wrapper. Supertypes remain queryable. |
| `extras: $ => [...]` | Tokens permitted anywhere; default whitespace unless replaced. If newline delimits statements, exclude newline from extras; use explicit separators. Define complex comment patterns as rules referenced from extras. |
| `word: $ => $.identifier`, `reserved` | Keyword extraction/boundaries and reserved/contextual word sets. `word` must refer to a unique token; reserved tokens must also occur in the grammar's rules. |
| `eof()` | Match end of input; only final symbol of a rule, never inside `token`. |

**Runnable keyword-boundary example** (ESM; use `module.exports` in a CommonJS project):

```js
export default grammar({
  name: 'keyword_example',
  extras: _ => [/[ \t\r]/], // Keep line separators significant.
  word: $ => $.identifier,
  reserved: { global: _ => ['if'] },
  rules: {
    source_file: $ => repeat(seq(choice($.if_statement, $.identifier), '\n')),
    if_statement: _ => 'if', // Reserved token is actually used.
    identifier: _ => /[a-z]+/,
  },
});
```

`if` is an `if_statement`; `iffy` and `gift` remain identifiers. Check line separators and EOF against the actual language. This example was generated and parsed with CLI 0.27.0.

**Before / after:** `seq('func', $.identifier, $.block)` has positional consumers; `seq('func', field('name', $.identifier), field('body', $.block))` gives consumers stable roles. Verify the actual tree and `src/node-types.json`.

## Lexing and ambiguity

[Precedence, conflicts, extras, keywords](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/3-writing-the-grammar.md).

Tree-sitter lexes **contextually**, then parses tokens. When valid tokens overlap, selection is: higher **lexical** precedence, longer match, string over regex at equal length, then earlier grammar declaration. A keyword matching a prefix of an identifier needs `word`/reserved handling; don't assume the enclosing parse state enforces the boundary. Don't hide significant newlines in extras.

| Tool | Choose when |
| --- | --- |
| `prec(n, rule)` | Static parse precedence between derivations. |
| `prec.left(n, rule)` / `.right` | Static precedence plus associativity, e.g. `a-b-c` or exponentiation. |
| `token(prec(n, pattern))` | Lexical precedence between tokens competing for the same characters. |
| `conflicts: $ => [[...]]` | **Genuine** ambiguity where both parses must be explored by GLR; not a default response to generation failure. |
| `prec.dynamic(n, rule)` | Runtime choice among surviving ambiguous parses declared with conflicts. |
| `precedences` | Named parse-precedence ordering, not lexical precedence. |

For `a + b * c`, make multiplication's parse precedence higher and use appropriate associativity. For an `Unresolved conflict`, read both generator interpretations and the bullet position: distinguish intended syntax from an ambiguity that only extra lookahead could resolve. Parse real samples to prove grouping; generator success alone does not do so.

## External scanners

[Official scanner contract](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/4-external-scanners.md). Introduce one for layout/indentation, heredocs, context-dependent delimiter matching, or other lexical behavior not practical as regular tokens; ordinary syntax belongs in the DSL. Put the hand-written implementation in `src/scanner.c`. `externals: $ => [...]` order **must match** the C token enum. Export exactly these five `tree_sitter_<name>_external_scanner_*` functions: `create`, `destroy`, `serialize`, `deserialize`, `scan`.

- `scan(payload, lexer, valid_symbols)`: emit a token only when its `valid_symbols[ENUM]` is true; set `lexer->result_symbol` and return `true`. `lookahead` is a Unicode code point; use `lexer->eof(lexer)` for EOF (NUL input can also have `lookahead == 0`). `get_column` counts code points, not bytes.
- `lexer->advance(lexer, false)` **includes** a character in the token. `advance(lexer, true)` skips text *before* the token. `mark_end` records the token end when later lookahead should not be included; avoid skipping after `mark_end` because it can shift the token start. A false return must not leave scanner state mutated.
- `serialize` writes the **complete** state that affects future lexing into at most `TREE_SITTER_SERIALIZATION_BUFFER_SIZE` bytes; `deserialize` resets state then restores it (including the zero-length case). Keep it compact: trees store scanner state for edits and ambiguity. Return zero for a stateless scanner. Test EOF, partial tokens, error recovery, and fresh-vs-incremental trees.

## Corpus, recovery, consumers

[Writing tests](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/creating-parsers/5-writing-tests.md) · [Query syntax](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/using-parsers/queries/1-syntax.md) · [ABI versions](https://github.com/tree-sitter/tree-sitter/blob/790e6ad9c5a26dfcf96785fe4c9cc7ad6e24d6d5/docs/src/using-parsers/7-abi-versions.md).

Corpus files in `test/corpus/` have `==== name ====`, source, `---`, then expected named-node S-expression (optionally with `field:`). Add positive and negative/boundary cases **before** changing behavior. `:error` asserts an error without spelling out the tree; do not use it where recovery shape matters. `tree-sitter test --update` will not update tests whose trees contain `ERROR` or `MISSING`. Inspect generated outputs and edit expectations only when the tree itself is right.

Malformed input should recover at documented boundaries without swallowing later declarations. Check error/missing spans, multiline tokens, comments, EOF and CRLF. Query `(ERROR)` and `(MISSING)` separately; fields and supertype/subtype selectors require corresponding generated nodes. Compare incremental parse after edits against fresh parse; test actual binding load and query captures. A parser's ABI must be supported by its runtime; CLI 0.27 docs list `generate --abi` and the published compatibility table lists runtime 0.24 supporting ABI 13–14, 0.25+ supporting 13–15. Confirm the targeted installed runtime before claiming compatibility.
