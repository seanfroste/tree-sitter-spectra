# Reviewed SPECTRA repository specification: completion design

Approved scope: focused grammar completion plus a Go semantic adapter, against the checked-in references and explicitly documented local policies. Implement in the current checkout. Full LSP transport, simulation execution, and compatibility with unavailable original appendices are outside scope.

## Authority and traceability

The source material is `docs/*.md`, the historical inventory in `completeness-baseline.json`, and checked-in `.in` fixtures. Historical structural acceptance is not a semantic oracle. Account for all 138 candidates with stable source-linked identities, source hashes, original input, explicit disposition, and evidence-backed corrections or local-policy rationale. Preserve originals. Separate executable examples from prose, format templates, and external data formats. Add a minimum executable END example. No silent exclusions or unexplained expected failures.

The user selected a reviewed repository specification rather than strict compatibility with an unavailable original specification. Resolve contradictions from repository evidence where possible; otherwise explicitly label the selected behavior as local policy, not verified historical simulator behavior. Record corrected inputs separately from their sources. A policy decision is not permission to accept malformed syntax indiscriminately.

## Grammar

Retain generic statements and named parameter fields. Implement whitespace-only lines, spaced function calls, bare reference classification, linked numeric/character series, grouped mask operations, and nested expressions. Expressions have operand/operator structure and precedence rather than arbitrary token sequences. Preserve interpolation and formatting forms documented by DEFINE and EXTRACT. Preserve full input after END and after column 511; runtime ignoring/truncation is not destructive parsing.

Keep unknown statement names intact, including keyword prefixes. Recovery must retain unrelated subsequent statements after missing values, unfinished strings, and incomplete nested values. The syntax-family matrix covers aliases/case, separators, comments, strings, numeric formats, references, linked series, interpolation, nested expressions, function spacing, mask groups, continuations, conditionals, and END trailing content, including invalid/incomplete variants.

## Go adapter and public contract

Use the existing Go binding as the selected integration boundary. Define a versioned adapter contract for logical statement ownership, continuation lines across comments and blank lines, END termination, malformed boundaries, DEFINE/EXTRACT declarations, reference kinds, conditional matching, documented aliases, syntax diagnostics, and include paths. Use tree fields and byte ranges, not a second parser over opaque source text. Unknown-name validation and simulation semantics remain outside grammar productions. Resolve includes relative to their containing document with explicit filesystem errors; do not execute or evaluate the language.

Document exact ownership and name/case policies from the references. Preserve source spelling and original byte ranges even where semantic lookup canonicalizes names. Do not claim a complete unavailable numerical-function catalogue. Optional locals/tags queries are unnecessary if the adapter establishes and tests equivalent direct-field traversal behavior.

## Editor contract

Test through Go: actual parsing, source ranges, declaration/reference discovery, local recovery, and incremental/fresh equivalence. Position conversion covers negotiated UTF-8 and UTF-16, LF/CRLF, tabs, multibyte and supplementary characters, and no final newline; invalid boundaries return errors rather than splitting code points. Incremental edits include token changes and transitions into and out of malformed syntax. Measure recovery and edit latency on a representative larger input before setting an enforceable budget.

## Verification and enforcement

One documented command runs corpus tests, the source inventory gate, executable example assertions, all checked-in fixtures, query checks, Go integration/editor tests, and generated-parser reproducibility. CI runs the same command. Assertions reject unexpected ERROR, MISSING, and unparsed_line nodes and verify values, arguments, reference kinds, fields, ownership, and source ranges. Expected results are reviewed against language sources, not copied blindly from parser output.

Retain Tree-sitter CLI 0.27.0, parser ABI 15, Node.js 24 or later, and the existing Go binding dependency unless a demonstrated compatibility requirement forces a documented change. Generated parser artifacts remain committed. No second grammar or new production dependency is planned.

## Implementation sequence and completion condition

1. M1: reviewed inventory, explicit source/local-policy decisions, semantic expectations, fixture/doc runner, and syntax-family matrix before grammar changes.
2. M2: confirmed syntax/classification and expression gaps, with targeted behavioral regressions.
3. M3: adapter and versioned syntax/semantic contract, with consumer-visible behavior tests.
4. M4: local recovery, positions/ranges, binding-level incremental equivalence, and measured performance.
5. M5: one command and CI; complete executable-example coverage; reproducible parser artifacts; updated documentation that precisely states the supported repository specification.

Completion requires every criterion in `completeness.md` section "Goal and acceptance criteria", interpreted against this approved repository-specification authority. A passing corpus alone is insufficient. Unresolved implementation gaps remain failures, not skips. Report only exercised verification; keep historical evidence distinct from current results.
