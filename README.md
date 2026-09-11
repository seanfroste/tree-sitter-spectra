# tree-sitter-spectra

Private Tree-sitter grammar for the SPECTRA input language used by SPECTRA dotin (`.in`) files.

## Toolchain

- Tree-sitter CLI 0.27.0
- Parser ABI 15
- Node.js 24 or later for grammar generation
- Go for binding smoke tests

## Development

```sh
npm ci
tree-sitter generate --abi 15
tree-sitter test
go test ./...
```

Generated parser and binding files are committed. Regenerate them after every grammar change and verify the resulting diff before committing.

## Source material

The initial grammar is developed from the plain Markdown statement reference docs and example dotin files in the private `seanfroste/spectrals` repository. The grammar repository remains independently buildable and testable.

## License

Private and unlicensed. See `LICENSE`.
