# tree-sitter-spectra

Private Tree-sitter grammar for the SPECTRA input language used by SPECTRA dotin (`.in`) files.

## Toolchain

- Tree-sitter CLI 0.27.0
- Parser ABI 15
- Node.js 24 or later for grammar generation
- Go 1.23 or later with C compiler for the ABI 15-compatible `go-tree-sitter v0.25.0` binding
- Rust toolchain and C compiler for the existing Rust binding tests

## Development

```sh
npm ci
npm run verify
```

Generated parser and binding files are committed. Regenerate them after every grammar change. `npm run verify` checks reproducible generation, corpus, 139 reviewed executable cases, all checked-in `.in` fixtures, highlight query recognition, Go adapter/editor tests, and Rust binding tests. Do not commit test-run output or local build directories.

## Completeness and LSP readiness

The [reviewed repository specification](docs/development/completeness-design.md), [documentation example policies](docs/development/documentation-policies.md), [syntax coverage matrix](docs/development/syntax-coverage.md), and [versioned Go adapter contract](docs/development/adapter-contract.md) define what this repository verifies. The [historical assessment](docs/development/completeness.md) and [historical source inventory](docs/development/completeness-baseline.json) preserve earlier evidence; neither is a current passing oracle.
The sibling `../spectrals` LSP still uses its own Rust lexer/parser. Future work will migrate that server into a Go LSP, consuming a release tag of this module. Go adapter and editor tests therefore define the target integration contract; existing Rust binding tests remain cross-binding checks. No sibling source changes occur here. Full LSP readiness requires Go server integration and end-to-end tests in that repository. Original simulator compatibility beyond checked-in references and explicit local policies remains unverified.

## Source material

The initial grammar is developed from the plain Markdown statement reference docs and example dotin files in the private `seanfroste/spectrals` repository. The grammar repository remains independently buildable and testable.

## License

Private and unlicensed. See `LICENSE`.
