# SPECTRA parser completeness and LSP readiness

Assessed 2026-09-13 against `fc7276efec0be7715e14bb347b4b180dbf36df04`.
This is an assessment and a development goal, not a claim that the goal has been reached.

## Verdict

The parser is a useful foundation for an LSP prototype and syntax highlighting, but **not yet a complete documented-language parser or a reliable LSP syntax backend**. It recognizes common statements, parameters, numeric values, strings, function calls without intervening spaces, series, interpolation and basic conditional markers. Significant valid syntax, reference classification and error-recovery behavior remain incomplete.

Do not assign a single overall completeness percentage. Passing the current corpus, structurally accepting a documentation sample, recognizing the intended values and supporting editor operations are different measurements.

## Measured baseline

| Check | Observed result | Limitation |
| --- | --- | --- |
| `npm test` | 10/10 corpus cases pass | Only the Tree-sitter corpus runs. Documentation and fixture files are not discovered by this script. |
| All checked-in `.in` fixtures | 14/14 parse with zero `ERROR`, `MISSING` or `unparsed_line` nodes | Five bare references in three simulation fixtures are nevertheless classified as `character_value`. |
| Named node types mentioned in corpus trees | 30/31 | `expression_fragment` is absent. Node occurrence is not grammar branch or semantic coverage. |
| Statement-reference scenario screening | 89/131 structurally accepted | Some inputs contain transcription defects. Accepted inputs can still have incorrect classifications or argument boundaries. |
| General input-format examples | 3/7 structurally accepted | Includes mixed prose, mathematical markup and a positional `plot potential` example that needs scope clarification. |
| Combined documentation screening | 92/138 structurally accepted | This 66.7% structural acceptance rate is **not** a completeness score or the final valid-example denominator. |
| `go test -count=1 ./...` | Pass | The sole Go test constructs a language wrapper. It does not parse input or test incremental edits. |
| CLI incremental-versus-fresh smoke probes | 3/3 produce identical tree output | Numeric replacement, a replacement after a Unicode title with CRLF, and no-final-newline input only. No binding-level or randomized edit suite exists. |
| Unicode title smoke probe | Parses an emoji and Greek letter with CRLF | Does not establish correct LSP UTF-16 position conversion. |

The corpus contains 16 distinct leading statement spellings, including aliases and conditional markers. It is not a statement-by-statement example suite. The representative `test/fixtures/statement_docs.in` is curated coverage, not a transcription of all examples.

The [historical baseline inventory](completeness-baseline.json) records every screened scenario's input, source line numbers, normalizations and structural result, plus 19 focused parser probes. It also records source-document hashes and non-example fenced material. It is audit data, **not an automated regression test or an approved semantic oracle**.

### Documentation inventory and screening

There are 24 reference documents: 23 statement-reference files, including the combined IF/ELSE/ENDIF reference, and one general input-format reference. The inventory includes 27 marked example blocks, the inline TITLE and separator examples, and three unfenced DEFINE examples. Multi-statement IF is one scenario. Elsewhere, a new statement starts a new scenario and parameter continuation lines stay with it. Duplicate examples remain distinct source occurrences.

| Reference | Scenarios | Structural passes | Failing scenario start lines |
| --- | ---: | ---: | --- |
| 00-Input-File-Format | 7 | 3 | 8, 23, 29, 38 |
| BIAS | 3 | 3 | |
| CONSTANT | 2 | 2 | |
| CONVERGENCE | 1 | 1 | |
| DEFINE | 16 | 12 | 25, 42, 47, 50 |
| DEPOSIT | 1 | 1 | |
| DOPE | 26 | 18 | 157, 167, 168, 169, 176, 178, 183, 190 |
| ELECTRODE | 13 | 9 | 44, 58, 63, 65 |
| END | 0 | 0 | No separate example, only `[format] END`. Existing corpus tests cover END. |
| EXTRACT | 4 | 2 | 84, 90 |
| GRID | 2 | 2 | |
| IF-ELSE-ENDIF | 1 | 1 | |
| INSERT | 1 | 1 | |
| INTERFACE | 5 | 2 | 45, 50, 52 |
| LIGHT | 10 | 5 | 59, 60, 65, 66, 67 |
| MODEL | 1 | 1 | |
| NFERMI | 6 | 3 | 43, 49, 51 |
| PFERMI | 6 | 2 | 43, 44, 51, 53 |
| REGION | 18 | 11 | 70, 72, 76, 88, 93, 95, 99 |
| RESTART | 4 | 2 | 18, 21 |
| SAVE | 6 | 6 | |
| STRUCTURE | 1 | 1 | |
| SUBSTRATE | 3 | 3 | |
| TITLE | 1 | 1 | |

Screening removes example labels and presentation indentation, omits identified explanatory prose, and unescapes copied Markdown escapes for `$`, `_`, `*`, `{` and `}`. It does **not** repair decimal spacing, filenames, missing separators or other apparent transcription errors. General mixed-prose samples are retained as candidates rather than silently rewritten. Each tested input receives a final LF. These choices are recorded and must be reviewed before turning the inventory into acceptance tests.

Format templates are not executable examples. Five other fenced blocks illustrate Fortran-like loop expansion, mathematical indexing or C programs for external data-file layouts. They are recorded separately, not claimed as SPECTRA input. External impurity/light/recombination data formats are not the `.in` grammar's syntax.

## Confirmed gaps

### P0: supported-looking syntax is not reliably recognized

1. **Space before function parentheses fails.** `REGION MASK=FILE (mask.txt,cell,3)` fails, whereas `REGION MASK=FILE(mask.txt,cell,3)` parses correctly. `grammar.js:167-187` contains an intended spaced-call alternative, but the observed parse takes `FILE` as a character value. This affects numerous examples, including `REGION.md:71-72`, `LIGHT.md:59` and `PFERMI.md:43`.
2. **Linked numeric and character series fail.** `DEFINE NAME=Width VALUE=#Size(0.1 1.0 0.4 1.8 3.0)` and the character-array equivalent produce errors. These are explicit examples at `DEFINE.md:42` and `DEFINE.md:47`.
3. **Grouped mask operations fail.** `&[ MASK=... ANDMASK=... ANDNEGA=... ]` produces fallback/error nodes. This is documented across DOPE, ELECTRODE, INTERFACE, LIGHT, NFERMI, PFERMI and REGION. Adding generic parameter parsing does not cover this grouping construct.
4. **Parenthesized expressions fail.** `DEFINE NAME=x VALUE={(@Width+1)*2}` produces errors. The general reference's nested power example at `00-Input-File-Format.md:38` fails after presentation unescaping. `_expression_item` at `grammar.js:252-262` lacks ordinary parenthesized expression grouping.
5. **Whitespace-only lines fail.** Lines containing only spaces or tabs produce `ERROR` nodes, unlike truly empty lines. This is an immediate editor-use problem, regardless of statement coverage.

### P0: an apparently successful parse can lose meaning

- **Bare references become strings.** `SUBS V=@Width`, `SAVE FILE=$INPUTFILE` and `DEFINE NAME=x OPTIMIZE=#Barrier` all parse with zero `variable_reference` nodes. `_value` at `grammar.js:126-136` omits `variable_reference`, although function arguments and series allow it. Existing simulation fixtures contain five such bare references. Highlighting them as strings is consistent with the wrong node classification, not proof of correctness.
- **Fallbacks mask unsupported text.** `???` becomes `unparsed_line` while the CLI exits successfully. Every valid-input gate must explicitly reject unexpected fallback nodes, not just check `HasError` or command success.
- **Expressions are token sequences, not validated expression trees.** `{1 + * 2}` is accepted without errors. There is no operator-precedence or operand structure to validate it. An expression parser or a deliberate semantic validation layer is needed before claiming expression diagnostics.

### P0/P1: error recovery can swallow the following statement

Both of these cases consume the next line's `SAVE` as the previous parameter's value, rather than preserving a separate SAVE statement:

```text
GRID XMIN=
SAVE FILE=after.dat
```

```text
DEFINE CHARACTER="unfinished
SAVE FILE=after.dat
```

This undermines completion context, document symbols and diagnostic stability during ordinary edits. Recovery tests must assert the shape and ranges of unaffected later statements, not merely the presence of an error.

### P1: the LSP tree/adapter contract is not defined

- Parameters have useful `name`, `operator` and `value` fields, but DEFINE declaration names are ordinary character values. A semantic adapter can interpret them by statement and parameter name.
- `continuation_line` is a sibling of statements in `source_file`, not a child of its owning logical statement. Ownership across comments, blank lines, END and malformed input needs an explicit rule and tests.
- IF, ELSE and ENDIF are flat independent nodes. An unmatched ELSE parses successfully. Nesting and matching can be represented by grammar nodes or an adapter, but cannot be assumed from the current tree.
- Generic statement/parameter names are intentional and can support aliases and extensions. `NOT_A_STATEMENT X=1` parses. Recognition of valid names, aliases, required parameters, types and repeated parameter groups needs a versioned semantic catalogue. One grammar rule per statement is not inherently necessary.
- Keyword-prefix handling needs an editor regression test: `ENDPOINT X=1` is split at `END` and errors, rather than becoming a whole unknown-name statement. This is an incomplete-name/unknown-token concern, not a claim that ENDPOINT is a valid SPECTRA statement.
- END trailing content remains in the tree, which is useful for editing. The semantic layer must apply the documented "ignored after END" behavior rather than deleting source text.
- There are no locals/tags queries or tests establishing definition/reference behavior. Those queries are optional if the LSP uses direct field traversal, but the corresponding contract and integration tests are not optional.

### P1: bindings and editor guarantees are mostly untested

Go and Node binding tests are load-only. Rust's explicit unit test is also load-only, with an empty-input documentation example. No checked-in suite proves source ranges, incremental/fresh equivalence, declaration/reference discovery, included-file resolution, edit recovery or LSP position conversion. The three successful CLI edit probes do not replace those tests.

Tree-sitter reports byte-oriented offsets/columns. The LSP boundary must test negotiated position encoding, particularly UTF-16, with multibyte and supplementary characters. Include LF/CRLF, tabs, no final newline, edits inside tokens and malformed input. Establish measured performance thresholds on representative larger inputs before making latency claims.

## Documentation ambiguity is a separate workstream

Not every failed documentation candidate demonstrates a grammar defect. Keep the original source and record reviewed decisions:

- `DEFINE.md:25` uses unquoted `CHAR=Gate width`, conflicting with the quoted-space rule in the general reference.
- `DEFINE.md:51`, `DOPE.md:158` and `REGION.md:83` contain separated decimal fragments. Some parse into the wrong number of series elements instead of failing.
- `DOPE.md:167-169,176`, `LIGHT.md:65` and `RESTART.md:18,21` contain filename spaces or missing separators. `RESTART.md:21` also puts ELECTRODE on the same line.
- `ELECTRODE.md:45` contains the compressed `PWL(002E-954E-956E-90)`. Line 51 contains a dot-separated FILE argument string. Syntax acceptance alone cannot establish intended argument boundaries.
- `EXTRACT.md:87,92-94` contains spaced filename/symbol fragments and copied escapes. Do not guess whether a hyphen in a reference is subtraction or part of a name without authoritative evidence.
- `LIGHT.md:51` describes `%` before a group while its format and examples use `&`. This discrepancy requires a documented resolution.
- The general reference includes a positional `plot potential` form without a corresponding statement reference, a mathematical `$0$` presentation, and a 511-column truncation rule. Decide and document whether each is syntax, legacy compatibility, semantic policy or a reference defect. An editor parser should retain the complete source even if runtime semantics ignore later columns.
- Referenced appendices, including the complete numerical function catalogue, are not in this repository. Repository-document completeness must not be represented as completeness against an unavailable full language specification.

The target is **100% accounted-for source examples and 100% correct recognition of approved executable examples**, not 100% indiscriminate acceptance of prose or damaged samples. Ambiguous cases remain visible and block the relevant completeness claim until resolved. There must be no silent exclusions or unexplained expected failures.

## Goal and acceptance criteria

Durable project initiative: `lsp-ready-spectra-parser-with-complete-documented-example-coverage`.

1. **Source traceability:** every example in every reference has a stable inventory identity, source location, and explicit disposition. Source changes invalidate or refresh associated tests. Add executable minimum examples for statements with only format templates, including END.
2. **Complete executable coverage:** every approved documentation example and all simulation/reference fixtures run in the standard verification command. Require no unexpected `ERROR`, `MISSING`, `unparsed_line` or semantic misclassification. Check values, reference kinds, function arguments, field names, ownership and source ranges, not only node existence.
3. **Corpus breadth:** map every documented syntax family and variant to expected-tree tests. Cover aliases/case, separators, whitespace-only lines, comments, strings, numeric formats, references, linked series, interpolation, nested expressions, function spacing, mask groups, continuations, control flow and trailing END content. Test invalid and incomplete variants too. No aggregate "coverage" claim without the per-requirement matrix.
4. **LSP contract:** document and test the grammar/adapter split for declarations, references, logical statements, conditional structure, aliases, diagnostics and include paths. Unknown-name validation and simulation semantics do not have to be grammar productions.
5. **Editor behavior:** malformed edits preserve unrelated later statements. Incremental and fresh trees agree under representative edit sequences through the intended binding. Validate byte/position conversions and all supported newline/encoding boundaries. Measure recovery and incremental latency before setting an enforceable performance budget.
6. **Reproducibility and enforcement:** one documented command covers corpus, source inventory, fixtures, queries and the selected binding. CI enforces it and checks generated parser reproducibility. Any unsupported or ambiguous documented form is visible in the report, not silently skipped.

## Recommended milestones

| Order | Outcome | Acceptance evidence |
| --- | --- | --- |
| M0 | Audit and preserve the current baseline | This report and the source-linked historical inventory. |
| M1 | Build a reviewed coverage oracle before changing behavior | Classify all 138 candidates, resolve/track source defects, add expected-tree assertions and a fixture/doc runner. Known syntax gaps are explicit failing regressions, not swallowed fallbacks. |
| M2 | Fix confirmed syntax and classification gaps | Whitespace-only lines, spaced function calls, references, linked series, grouping and mask operations pass targeted tests plus all accepted examples. |
| M3 | Stabilize the LSP syntax/semantic boundary | Versioned tree/adapter contract and tests for statement ownership, declaration/reference discovery, control flow and aliases. |
| M4 | Harden editing and integration | Local recovery, binding-level parsing and incremental equivalence, ranges/position conversion, and measured larger-file performance. |
| M5 | Enforce the complete gate | One repeatable command and CI, zero unresolved executable-example gaps, reproducible parser artifacts and a reviewed coverage matrix. |

### Approach options

- **Minimal adapter over the current tree:** quickest for an LSP experiment, but cannot responsibly claim complete recognition while valid syntax and recovery failures remain.
- **Focused grammar completion plus a semantic adapter (recommended):** fix syntax/classification/recovery in this repo and keep statement catalogues, name resolution and simulation rules outside grammar productions. This avoids reparsing opaque text in the LSP while keeping grammar complexity controlled.
- **Statement-specific grammar and a fully structured AST:** stronger built-in structure, but greater coupling to the statement/alias catalogue and a larger compatibility change. Use only where the LSP contract demonstrates a need.

This assessment does not implement any of these approaches. M1 and agreement on the syntax/semantic boundary are the next work items. Full LSP transport/server implementation is outside this parser-completeness goal.

## Rechecking the evidence

Run `npm test`, `go test -count=1 ./...`, and parse all fixture paths with the installed Tree-sitter CLI. The fixture check must also inspect tree output for unexpected fallback nodes and assert recognition, which the current npm script does not do.

For each entry in `completeness-baseline.json`, write its `input` plus LF to a temporary `.in` file, run `node_modules/tree-sitter-cli/tree-sitter.exe parse <file>` from this repository, and compare only structural outcomes to this historical baseline. The original probe process did not change grammar, generated parser or tests. A future regression runner must use reviewed expected semantics rather than blessing this baseline's behavior.
