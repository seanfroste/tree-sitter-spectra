# tree-sitter-spectra

Tree-sitter grammar for SPECTRA `.in` files, with Go analysis/editor helpers
and Rust and Node bindings.

## Install

The Go module is `github.com/seanfroste/tree-sitter-spectra`:

```sh
go get github.com/seanfroste/tree-sitter-spectra@v0.1.0
```

Go consumers need a C compiler and Tree-sitter's `go-tree-sitter` v0.25.0.
The parser's Go package is `github.com/seanfroste/tree-sitter-spectra/bindings/go`.
It exposes `Language`, `Analyze`, `ResolveIncludes`, `PositionAt`, and
`ByteOffset`; see the [adapter contract](docs/development/adapter-contract.md).

To vendor the parser source as a Git submodule:

```sh
git submodule add https://github.com/seanfroste/tree-sitter-spectra.git \
  third_party/tree-sitter-spectra
git -C third_party/tree-sitter-spectra checkout v0.1.0
```

Tree-sitter CLI and Node.js are needed for grammar generation, not ordinary
Go module consumption. The optional Node binding requires Node.js and a C/C++
build toolchain. The Rust binding is available as the `tree-sitter-spectra`
crate when published to crates.io; until then, consume it from this repository.

## Source material
The grammar is developed from the statement references and `.in` examples
checked into this repository.

## License

MIT. See [LICENSE](LICENSE).

## Development toolchain

- Tree-sitter CLI 0.27.0
- Parser ABI 15
- Node.js 24 or later for grammar generation
- Go 1.23 or later with C compiler for the `go-tree-sitter v0.25.0` binding
- Rust toolchain and C compiler for the Rust binding

## Development

```sh
npm ci
npm run verify
```

Generated parser and binding files are committed. Regenerate them after every
grammar change. `npm run verify` checks reproducible generation, corpus, 139
reviewed executable cases, all checked-in `.in` fixtures, highlight query
recognition, Go adapter/editor tests, and Rust binding tests. Do not commit
test-run output or local build directories.

## Completeness and LSP readiness

The [reviewed repository specification](docs/development/completeness-design.md),
[documentation example policies](docs/development/documentation-policies.md),
[syntax coverage matrix](docs/development/syntax-coverage.md), and
[versioned Go adapter contract](docs/development/adapter-contract.md) define
what this repository verifies. The sibling `../spectrals` LSP still uses its
own Rust lexer/parser. Full LSP readiness requires Go server integration and
end-to-end tests there. Compatibility with the original simulator beyond the
checked-in references and explicit local policies remains unverified.
