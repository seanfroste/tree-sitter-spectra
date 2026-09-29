# Reviewed documentation example policy

`test/documentation.json` accounts for 138 historical candidates plus an executable minimum `END`. Each record retains the original inventory input, document, source lines, reviewed executable input, disposition, reason, and independent parameter/reference/function expectations. All 139 are executable. Original `docs/*.md` content remains unchanged. The source hashes are computed over LF-normalized Markdown; repository files may use CRLF on Windows. Historical structural outcomes in `completeness-baseline.json` never serve as semantic expectations.

## Review decisions

- `docs/00-Input-File-Format.md:8,23`: prose attached to SUBSTRATE and mathematical `$0$` notation are presentation, not dotin input. The demonstrated assignments remain executable. At lines 26–29, `plot potential` is an executable local positional-argument policy. This is not proof that historical SPECTRA recognizes PLOT.
- `docs/DEFINE.md:25`: quote `Gate width`, following the quoted-space rule at `docs/00-Input-File-Format.md:26`. At lines 42 and 47, `#Size(...)` links a reference to ordered numeric or character series; both boundaries matter.
- `docs/DEFINE.md:51`, `docs/DOPE.md:157–158`, `docs/REGION.md:76–77,82–84`: restore visibly split decimals. In DOPE, restore damaged `X T` as `XJ` according to the statement's junction-depth table.
- `docs/DOPE.md:167–169,176`, `docs/RESTART.md:18,21`, `docs/NFERMI.md:43,51`, `docs/PFERMI.md:44,51,53`, `docs/EXTRACT.md:87`: join visibly separated filename extensions. `docs/LIGHT.md:65` preserves an intended internal space by quoting the filename. `docs/RESTART.md:21` restores a newline before ELECTRODE. `docs/DOPE.md:176` restores the required separator before XMIN.
- `docs/ELECTRODE.md:45`: compressed PWL digit stream does not define recoverable argument boundaries. The repository policy uses `(0,0), (2E-9,5), (4E-9,5), (6E-9,0)` as a reviewable example; historical waveform unknown. At line 51, FILE uses filename, cell, layer, data-type boundaries from its parameter description.
- `docs/EXTRACT.md:92–94`: `XVMAX 1` is consistently represented as `XVMAX_1`. `#Electron-1` is reference `#Electron` minus one, not a hyphenated symbol. Both are local policies, not recovered simulator intent.
- `docs/LIGHT.md:51`: `%` before a mask group contradicts its own syntax and examples plus other references; use `&[...]` for AND, `[...]` for OR. Keep `%` only in documented numeric formatting such as `docs/EXTRACT.md:19,75`.
- `docs/00-Input-File-Format.md:11`: retain source after column 511 in syntax trees; truncation belongs to simulator semantics. An unavailable appendix cannot establish the complete numerical-function catalogue. Syntax recognition does not imply evaluation.

Correction-specific reasons and source locations remain in each inventory record. Cases not corrected retain distinct identities even when source inputs duplicate. Format templates and external data-file descriptions remain separately listed as non-example fenced material. The minimum END example derives from `docs/END.md` format rather than a historical candidate.

## Coverage families

Corpus and binding tests cover case, names and prefix boundaries, whitespace and comments, assignments and continuations, quotations, numeric forms, references, linked series, spaced/nested calls, masks, interpolation/formatting, precedence, conditional markers, END trailing source, malformed recovery, incremental equality, and UTF-8/UTF-16 positions. `go test -count=1 ./...` verifies semantic expectations for every executable inventory case, every checked-in fixture, and query reference captures. Rust binding tests cover field and byte-range access required by `../spectrals`; that Rust server does not yet consume this crate.
