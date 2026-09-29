package tree_sitter_spectra_test

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	spectra "github.com/seanfroste/tree-sitter-spectra/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func analyzeText(t *testing.T, text string) spectra.Analysis {
	t.Helper()
	parser := tree_sitter.NewParser()
	defer parser.Close()
	if err := parser.SetLanguage(tree_sitter.NewLanguage(spectra.Language())); err != nil {
		t.Fatal(err)
	}
	source := []byte(text)
	tree := parser.Parse(source, nil)
	if tree == nil {
		t.Fatal("parse returned no tree")
	}
	defer tree.Close()
	return spectra.Analyze(tree.RootNode(), source)
}

func diagnosticCodes(a spectra.Analysis) []string {
	var codes []string
	for _, d := range a.Diagnostics {
		codes = append(codes, d.Code)
	}
	return codes
}

func TestAnalysisOwnershipAndTermination(t *testing.T) {
	text := "DEFINE NAME=Width VALUE=1\n# comment\n \t\n CHARACTER=wide\nSAVE FILE=$Width\nEND\nSAVE FILE=ignored"
	a := analyzeText(t, text)
	if len(a.Statements) != 4 || len(a.Diagnostics) != 0 {
		t.Fatalf("statements=%+v diagnostics=%+v", a.Statements, a.Diagnostics)
	}
	first := a.Statements[0]
	if first.Range != (spectra.ByteRange{Start: 0, End: 54}) || !reflect.DeepEqual(first.Parts, []spectra.ByteRange{{Start: 0, End: 25}, {Start: 39, End: 54}}) {
		t.Fatalf("logical ownership: %+v", first)
	}
	if len(first.Parameters) != 3 || first.Parameters[2].CanonicalName != "CHARACTER" {
		t.Fatalf("continuation parameters: %+v", first.Parameters)
	}
	if len(a.Declarations) != 1 || a.Declarations[0].Name != "Width" || a.Declarations[0].Range != (spectra.ByteRange{Start: 12, End: 17}) {
		t.Fatalf("declaration: %+v", a.Declarations)
	}
	if len(a.References) != 1 || a.References[0].Kind != spectra.CharacterReference || a.References[0].Name != "Width" || a.References[0].Range != (spectra.ByteRange{Start: 65, End: 71}) {
		t.Fatalf("reference: %+v", a.References)
	}
	if !a.Statements[2].Active || a.Statements[3].Active {
		t.Fatalf("END activity: %+v", a.Statements)
	}
}

func TestAnalysisCaseAliasesAndUnknownPolicy(t *testing.T) {
	a := analyzeText(t, "sTrUcT XB=P\nDEFINE NAME=Width CHAR=wide NAME=width VALUE=2\nEXTRACT PARA=VMAX NAME=Peak PARA=VMIN NAME=Floor\nSAVE FILE=$Width\nMystery FILE=a\nDEFINE MYSTERY=1\n")
	if a.Statements[0].Name != "sTrUcT" || a.Statements[0].CanonicalName != "STRUCTURE" || a.Statements[0].Parameters[0].CanonicalName != "XBOUNDARY" {
		t.Fatalf("documented aliases: %+v", a.Statements[0])
	}
	var names []string
	var groups []int
	for _, d := range a.Declarations {
		names = append(names, d.Name)
		groups = append(groups, d.Group)
	}
	if !reflect.DeepEqual(names, []string{"Width", "width", "Peak", "Floor"}) || !reflect.DeepEqual(groups, []int{0, 1, 0, 1}) {
		t.Fatalf("declarations=%+v", a.Declarations)
	}
	if a.Declarations[2].Kind != spectra.ExtractDeclaration || a.Statements[1].Parameters[1].Name != "CHAR" || a.Statements[1].Parameters[1].CanonicalName != "CHARACTER" {
		t.Fatalf("declaration kinds or aliases: %+v", a)
	}
	if !reflect.DeepEqual(diagnosticCodes(a), []string{"unknown-statement", "unknown-parameter"}) {
		t.Fatalf("diagnostics=%+v", a.Diagnostics)
	}
}

func TestAnalysisReferencesTraverseLinkedExpressionsAndMasks(t *testing.T) {
	a := analyzeText(t, "DEFINE NAME=Width VALUE=#Size(0.1 0.2) CHAR=#Size(a b)\nEXTRACT PARA=$$Label VALUE={atan2(@Y,@X)+#Peak}\nNFERMI &[MASK=FILE($Mask,cell,1) ANDMASK=POLYGON(0,0,@Width,1)]\n")
	var kinds []spectra.ReferenceKind
	var names []string
	var linked []bool
	for _, r := range a.References {
		kinds = append(kinds, r.Kind)
		names = append(names, r.Name)
		linked = append(linked, r.Linked)
	}
	if !reflect.DeepEqual(names, []string{"Size", "Size", "Label", "Y", "X", "Peak", "Mask", "Width"}) || !reflect.DeepEqual(kinds, []spectra.ReferenceKind{spectra.ExtractedReference, spectra.ExtractedReference, spectra.CharacterReference, spectra.NumericReference, spectra.NumericReference, spectra.ExtractedReference, spectra.CharacterReference, spectra.NumericReference}) {
		t.Fatalf("references=%+v diagnostics=%+v", a.References, a.Diagnostics)
	}
	if !reflect.DeepEqual(linked, []bool{true, true, false, false, false, false, false, false}) {
		t.Fatalf("linked=%v", linked)
	}
	if len(a.Statements[2].Parameters) != 2 || a.Statements[2].Parameters[1].CanonicalName != "ANDMASK" || a.References[7].Statement != 2 || a.References[7].Parameter != 1 {
		t.Fatalf("grouped mask ownership=%+v", a)
	}
}

func TestAnalysisRejectsUnsupportedReferencePrefix(t *testing.T) {
	for _, input := range []string{
		"DEFINE VALUE=@@Width",
		"DEFINE VALUE=F(@@Width)",
		"DEFINE VALUE={@@Width}",
		"DEFINE VALUE=##Width",
		"DEFINE VALUE=$$$Width",
	} {
		t.Run(input, func(t *testing.T) {
			a := analyzeText(t, input+"\n")
			if len(a.References) != 0 {
				t.Fatalf("unsupported spelling became a reference: %+v", a.References)
			}
			if len(a.Diagnostics) == 0 || a.Diagnostics[0].Severity != "error" {
				t.Fatalf("unsupported spelling lacks syntax error: %+v", a.Diagnostics)
			}
		})
	}
}

func TestAnalysisMalformedBoundaryDoesNotStealContinuation(t *testing.T) {
	a := analyzeText(t, "DEFINE NAME=Good VALUE=1\n???\n VALUE=2\nSAVE FILE=ok\n")
	if len(a.Statements) != 3 || len(a.Statements[0].Parameters) != 2 || a.Statements[1].Kind != "continuation_line" || a.Statements[2].CanonicalName != "SAVE" {
		t.Fatalf("recovery ownership=%+v", a.Statements)
	}
	codes := diagnosticCodes(a)
	if !reflect.DeepEqual(codes, []string{"unparsed-line", "orphan-continuation"}) {
		t.Fatalf("diagnostics=%+v", a.Diagnostics)
	}
}

func TestAnalysisConditionalsAreStructuralNotEvaluated(t *testing.T) {
	a := analyzeText(t, "IF [@Flag]\nIF [1]\nELSE\nENDIF\nELSE\nELSE\nENDIF\nENDIF\nIF [0]\nEND\nIF [1]\nENDIF\n")
	want := []spectra.Conditional{
		{IfStatement: 0, ElseStatement: 4, EndIfStatement: 6},
		{IfStatement: 1, ElseStatement: 2, EndIfStatement: 3},
		{IfStatement: 8, ElseStatement: -1, EndIfStatement: -1},
		{IfStatement: 10, ElseStatement: -1, EndIfStatement: 11},
	}
	if !reflect.DeepEqual(a.Conditionals, want) {
		t.Fatalf("conditionals=%+v", a.Conditionals)
	}
	if !reflect.DeepEqual(diagnosticCodes(a), []string{"duplicate-else", "unmatched-endif", "unclosed-if"}) {
		t.Fatalf("diagnostics=%+v", a.Diagnostics)
	}
	if !a.Statements[1].Active || a.Statements[10].Active {
		t.Fatal("conditional expressions were evaluated or END ignored")
	}
}

func TestAnalysisIncludeResolution(t *testing.T) {
	a := analyzeText(t, "INSERT FILE='parts/child.in'\nDEFINE FILE=$INPUTFILE\nINSERT FILE=parts/{@Index}.in\nDEFINE FILE=missing.in\nSAVE FILE=output.in\nEND\nINSERT FILE=ignored.in\n")
	if len(a.Includes) != 5 || a.Includes[0].Path != "parts/child.in" || a.Includes[1].Status != spectra.IncludeSymbolic || a.Includes[2].Status != spectra.IncludeSymbolic || a.Includes[4].Active {
		t.Fatalf("includes=%+v", a.Includes)
	}
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "parts"), 0700); err != nil {
		t.Fatal(err)
	}
	file := filepath.Join(dir, "parts", "child.in")
	if err := os.WriteFile(file, []byte("END\n"), 0600); err != nil {
		t.Fatal(err)
	}
	resolved := spectra.ResolveIncludes(a.Includes, filepath.Join(dir, "main.in"))
	if resolved[0].Status != spectra.IncludeResolved || resolved[0].Path != file || resolved[0].Err != nil {
		t.Fatalf("relative include=%+v", resolved[0])
	}
	if resolved[1].Status != spectra.IncludeSymbolic || resolved[1].Path != "" || resolved[2].Path != "" {
		t.Fatalf("symbolic include guessed=%+v", resolved)
	}
	if resolved[3].Status != spectra.IncludeFileError || !errors.Is(resolved[3].Err, os.ErrNotExist) || resolved[4].Status != spectra.IncludeInactive {
		t.Fatalf("file errors or inactive includes=%+v", resolved)
	}
	if a.Includes[0].Status != spectra.IncludeLiteral {
		t.Fatal("resolution mutated analysis")
	}
}

func TestAnalysisRetainsSpellingAfterSourceAndTreeRelease(t *testing.T) {
	a := analyzeText(t, "DEFINE NAME='MixedCase' VALUE=1\nSAVE FILE=$MixedCase")
	if a.Declarations[0].Name != "MixedCase" || a.Declarations[0].Spelling != "'MixedCase'" || a.References[0].Spelling != "$MixedCase" {
		t.Fatalf("names=%+v references=%+v", a.Declarations, a.References)
	}
}
