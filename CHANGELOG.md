# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.1] - 2026-10-03

### Fixed

- Declare native Node build and loader dependencies directly so isolated,
  non-hoisted package installations work.

### Added

- Node consumer tests for declaration fields, END source retention, and
  shipped highlight-query captures, included in the canonical verification gate.

## [0.1.0] - 2026-10-03

### Added

- First public release under the MIT license.
- Tree-sitter grammar for the reviewed SPECTRA repository specification, with
  recoverable syntax, expressions, references, linked series, mask groups,
  and ABI 15 generated parser artifacts.
- Go adapter contract version 1 for analysis, include resolution, and UTF-8 and
  UTF-16 position conversion; Go, Rust, and Node language bindings.
- Reproducible verification across documented examples, fixtures, corpus,
  queries, Go editor tests, and Rust binding tests.

[Unreleased]: https://github.com/seanfroste/tree-sitter-spectra/compare/v0.1.1...HEAD
[0.1.1]: https://github.com/seanfroste/tree-sitter-spectra/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/seanfroste/tree-sitter-spectra/releases/tag/v0.1.0
