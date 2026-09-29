# Syntax coverage matrix

The repository specification is `completeness-design.md` plus reviewed dispositions in `test/documentation.json`. This matrix names executable verification, not a count of grammar branches. `npm run verify` runs all checks listed here.

| Family | Positive proof | Negative/editor proof |
| --- | --- | --- |
| Case, names, aliases | `test/corpus/syntax.txt` case-insensitive names; `bindings/go/analysis_test.go` alias lookup | `bindings/go/grammar_contract_test.go` ENDPOINT prefix; `analysis_test.go` unknown-name warning |
| Separators, continuations, comments | `test/corpus/syntax.txt` assignment separators and multiline arguments; `analysis_test.go` ownership | `analysis_test.go` orphan continuation and malformed boundary |
| Blank and whitespace-only lines | `grammar_contract_test.go` spaced blank line | `editor_test.go` insertion/deletion and ranges |
| Strings, numeric formats | `test/corpus/syntax.txt` quoted and unquoted, signed decimals, exponents | `grammar_contract_test.go` unfinished string, missing values |
| Bare `@`, `$`, `#` references | `documentation_test.go` all sources, fixture character/reference classification, query capture; `analysis_test.go` kind and range | `grammar_contract_test.go` malformed expression and recovered statement; `analysis_test.go` rejects unsupported repeated/mixed sigils |
| Linked numeric/character series | `documentation_test.go` DEFINE:42, DEFINE:47; `grammar_contract_test.go` value boundaries | `editor_test.go` malformed series following SAVE |
| Spaced and nested function calls | `documentation_test.go` all approved examples, ordered arguments and value nodes | `grammar_contract_test.go` unclosed FILE call with following SAVE |
| Mask AND/OR groups | `documentation_test.go` DOPE, ELECTRODE, INTERFACE, LIGHT, NFERMI, PFERMI, REGION; `grammar_contract_test.go` nested group | `grammar_contract_test.go` unclosed mask group |
| Expressions and precedence | `grammar_contract_test.go` binary, unary, grouping, comparisons, power associativity, multiline operators; `test/corpus/syntax.txt` nested functions | `grammar_contract_test.go` operator without operand, empty expression, adjacent operands |
| Conditional control flow | `test/corpus/syntax.txt` IF/ELSE/ENDIF; `analysis_test.go` matching and grouping | `analysis_test.go` duplicate/unmatched/unclosed markers; `grammar_contract_test.go` incomplete series preserves following TITLE |
| END and trailing source | `test/corpus/syntax.txt`; `documentation.json` END:minimum; `analysis_test.go` inactive trailing statements | `grammar_contract_test.go` keyword-prefix test |
| Byte ranges and positions | `bindings/go/positions_test.go` LF/CRLF, tabs, Unicode, UTF-8/UTF-16; `bindings/rust/consumer_tests.rs` source ranges | `positions_test.go` invalid boundaries; `editor_test.go` malformed edit recovery, fresh/incremental equality |
| Fixture and documentation examples | `documentation_test.go` 139 reviewed cases and all checked-in `.in` fixtures | hash check fails on changed reference source; unexpected ERROR/MISSING/fallback fails |

Queries: `TestHighlightQueryRecognizesReferences` runs the shipped highlights query against real parsed input. The tree/adapter split and sibling Rust server status are in `adapter-contract.md`.
