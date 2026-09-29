package tree_sitter_spectra_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	spectra "github.com/seanfroste/tree-sitter-spectra/bindings/go"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
)

func editorParser(t testing.TB) *tree_sitter.Parser {
	t.Helper()
	parser := tree_sitter.NewParser()
	if err := parser.SetLanguage(tree_sitter.NewLanguage(spectra.Language())); err != nil {
		parser.Close()
		t.Fatal(err)
	}
	t.Cleanup(parser.Close)
	return parser
}

// Tree-sitter columns always count bytes, regardless of the editor's negotiated
// encoding. Derive these independently of the public position conversion API.
func editorPoint(source []byte, offset int) tree_sitter.Point {
	prefix := source[:offset]
	return tree_sitter.Point{
		Row:    uint(bytes.Count(prefix, []byte{'\n'})),
		Column: uint(offset - bytes.LastIndexByte(prefix, '\n') - 1),
	}
}

func editorInputEdit(old, next []byte, start, oldEnd, newEnd int) tree_sitter.InputEdit {
	return tree_sitter.InputEdit{
		StartByte:      uint(start),
		OldEndByte:     uint(oldEnd),
		NewEndByte:     uint(newEnd),
		StartPosition:  editorPoint(old, start),
		OldEndPosition: editorPoint(old, oldEnd),
		NewEndPosition: editorPoint(next, newEnd),
	}
}

type editorNodeShape struct {
	kind                   string
	startByte, endByte     uint
	startPoint, endPoint   tree_sitter.Point
	named, extra           bool
	missing, error, errors bool
	children               uint
}

func editorShape(node *tree_sitter.Node) editorNodeShape {
	return editorNodeShape{
		kind: node.Kind(), startByte: node.StartByte(), endByte: node.EndByte(),
		startPoint: node.StartPosition(), endPoint: node.EndPosition(),
		named: node.IsNamed(), extra: node.IsExtra(), missing: node.IsMissing(),
		error: node.IsError(), errors: node.HasError(), children: node.ChildCount(),
	}
}

// Compare every child, including anonymous punctuation, and its owning field.
// HasChanges and node IDs are deliberately excluded: they describe edit history,
// not the consumer-visible parse structure that a fresh parse must reproduce.
func editorAssertSameTree(t testing.TB, edited, fresh *tree_sitter.Node, path string) {
	t.Helper()
	got, want := editorShape(edited), editorShape(fresh)
	if got != want {
		t.Fatalf("tree differs at %s:\n incremental: %+v\n fresh:       %+v", path, got, want)
	}
	for i := range edited.ChildCount() {
		gotField := edited.FieldNameForChild(uint32(i))
		wantField := fresh.FieldNameForChild(uint32(i))
		if gotField != wantField {
			t.Fatalf("field differs at %s/%d: incremental %q, fresh %q", path, i, gotField, wantField)
		}
		editorAssertSameTree(t, edited.Child(i), fresh.Child(i), fmt.Sprintf("%s/%d", path, i))
	}
}

func editorHasProblem(node *tree_sitter.Node) bool {
	if node.HasError() || node.IsError() || node.IsMissing() || node.Kind() == "unparsed_line" {
		return true
	}
	for i := range node.ChildCount() {
		if editorHasProblem(node.Child(i)) {
			return true
		}
	}
	return false
}

func editorAssertLaterSave(t testing.TB, root *tree_sitter.Node, source []byte) {
	t.Helper()
	const statement = "SAVE FILE=\"later-β.dat\" TYPE=DATA"
	start := bytes.Index(source, []byte(statement))
	if start < 0 {
		t.Fatal("later SAVE fixture is absent")
	}
	var save *tree_sitter.Node
	for i := range root.NamedChildCount() {
		child := root.NamedChild(i)
		if child.Kind() != "statement" || child.StartByte() != uint(start) {
			continue
		}
		name := child.ChildByFieldName("name")
		if name != nil && string(source[name.StartByte():name.EndByte()]) == "SAVE" {
			save = child
			break
		}
	}
	if save == nil {
		t.Fatal("unrelated later SAVE was swallowed or lost its top-level statement/name field")
	}
	if editorHasProblem(save) {
		t.Fatal("unrelated later SAVE contains an error, missing node, or unparsed line")
	}
	end := start + len(statement)
	if save.EndByte() != uint(end) || save.StartPosition() != editorPoint(source, start) || save.EndPosition() != editorPoint(source, end) {
		t.Fatalf("later SAVE range = %+v, want byte range [%d,%d) with matching byte-column points", save.Range(), start, end)
	}
	values := map[string]string{}
	for i := range save.NamedChildCount() {
		parameter := save.NamedChild(i)
		if parameter.Kind() != "parameter" {
			continue
		}
		name, value := parameter.ChildByFieldName("name"), parameter.ChildByFieldName("value")
		if name == nil || value == nil {
			t.Fatal("later SAVE parameter lost name/value field")
		}
		values[string(source[name.StartByte():name.EndByte()])] = string(source[value.StartByte():value.EndByte()])
	}
	if len(values) != 2 || values["FILE"] != "\"later-β.dat\"" || values["TYPE"] != "DATA" {
		t.Fatalf("later SAVE values = %v; want FILE=\"later-β.dat\", TYPE=DATA", values)
	}
}

type editorReplacement struct {
	name      string
	anchor    string // The edit begins immediately after this unique text.
	remove    string
	insert    string
	malformed bool
}

func TestEditorIncrementalMatchesFresh(t *testing.T) {
	tests := []struct {
		name   string
		source string
		edits  []editorReplacement
	}{
		{
			name:   "unicode CRLF and no final newline",
			source: "TITLE Device 😀\r\nDEFINE NAME=Width VALUE=1\r\nSAVE FILE=\"later-β.dat\" TYPE=DATA",
			edits: []editorReplacement{
				{name: "insert numeric suffix", anchor: "VALUE=1", insert: ".5"},
				{name: "delete value", anchor: "VALUE=", remove: "1.5", malformed: true},
				{name: "repair missing value", anchor: "VALUE=", insert: "2"},
				{name: "unfinished quote", anchor: "VALUE=", remove: "2", insert: "\"open", malformed: true},
				{name: "repair quote", anchor: "VALUE=", remove: "\"open", insert: "\"closed\""},
				{name: "replace supplementary rune", anchor: "TITLE Device ", remove: "😀", insert: "界"},
				{name: "insert Unicode comment line", anchor: "VALUE=\"closed\"\r\n", insert: "# note μ\r\n"},
				{name: "delete comment line", anchor: "VALUE=\"closed\"\r\n", remove: "# note μ\r\n"},
				{name: "replace CRLF with LF", anchor: "TITLE Device 界", remove: "\r\n", insert: "\n"},
				{name: "append final newline", anchor: "TYPE=DATA", insert: "\n"},
				{name: "remove final newline", anchor: "TYPE=DATA", remove: "\n"},
			},
		},
		{
			name:   "LF nested expression repair",
			source: "DEFINE NAME=Width VALUE={1+(2*3)}\nSAVE FILE=\"later-β.dat\" TYPE=DATA",
			edits: []editorReplacement{
				{name: "remove closing delimiters", anchor: "VALUE={1+(2*3", remove: ")}", malformed: true},
				{name: "repair nested value", anchor: "VALUE={1+(2*3", insert: ")}"},
				{name: "replace nested operand", anchor: "VALUE={1+(", remove: "2", insert: "20"},
				{name: "insert leading empty line", insert: "\n"},
				{name: "delete leading empty line", remove: "\n"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, freshParser := editorParser(t), editorParser(t)
			source := []byte(tt.source)
			tree := parser.Parse(source, nil)
			if tree == nil {
				t.Fatal("initial parse returned nil")
			}
			defer func() { tree.Close() }()
			if editorHasProblem(tree.RootNode()) {
				t.Fatal("literal valid initial document did not parse cleanly")
			}
			editorAssertLaterSave(t, tree.RootNode(), source)
			for _, step := range tt.edits {
				t.Run(step.name, func(t *testing.T) {
					start := 0
					if step.anchor != "" {
						if bytes.Count(source, []byte(step.anchor)) != 1 {
							t.Fatalf("edit anchor %q is not unique", step.anchor)
						}
						start = bytes.Index(source, []byte(step.anchor)) + len(step.anchor)
					}
					if !bytes.HasPrefix(source[start:], []byte(step.remove)) {
						t.Fatalf("edit removal %q does not match source at %d", step.remove, start)
					}
					oldEnd := start + len(step.remove)
					next := make([]byte, 0, len(source)-len(step.remove)+len(step.insert))
					next = append(next, source[:start]...)
					next = append(next, step.insert...)
					next = append(next, source[oldEnd:]...)
					edit := editorInputEdit(source, next, start, oldEnd, start+len(step.insert))
					tree.Edit(&edit)
					edited := parser.Parse(next, tree)
					if edited == nil {
						t.Fatal("incremental parse returned nil")
					}
					tree.Close()
					tree, source = edited, next
					fresh := freshParser.Parse(source, nil)
					if fresh == nil {
						t.Fatal("fresh parse returned nil")
					}
					defer fresh.Close()
					editorAssertSameTree(t, tree.RootNode(), fresh.RootNode(), "root")
					if got := editorHasProblem(tree.RootNode()); got != step.malformed {
						t.Fatalf("syntax problem = %v, want %v for %q", got, step.malformed, source)
					}
					editorAssertLaterSave(t, tree.RootNode(), source)
				})
				if t.Failed() {
					return // Later steps require the exact previous document state.
				}
			}
		})
	}
}

// This independently exercises recovery without relying on fresh/incremental
// equality: two equally wrong trees must not hide a swallowed later statement.
func TestEditorRecoveryRetainsLaterSave(t *testing.T) {
	for _, newline := range []string{"\n", "\r\n"} {
		for _, broken := range []string{"DEFINE NAME=Width VALUE=", "DEFINE NAME=Width VALUE=\"unfinished", "DEFINE NAME=Width VALUE={1+(2"} {
			t.Run(fmt.Sprintf("%q/%q", newline, broken), func(t *testing.T) {
				source := []byte(broken + newline + "SAVE FILE=\"later-β.dat\" TYPE=DATA")
				tree := editorParser(t).Parse(source, nil)
				if tree == nil {
					t.Fatal("parse returned nil")
				}
				defer tree.Close()
				if !editorHasProblem(tree.RootNode()) {
					t.Fatal("incomplete value was silently accepted")
				}
				editorAssertLaterSave(t, tree.RootNode(), source)
			})
		}
	}
}

const editorBenchmarkBlock = "# cell μ\r\n" +
	"DEFINE NAME=Width VALUE=1.25\r\n" +
	"DOPE XMIN=0 XMAX=10 PROFILE=Gaussian\r\n" +
	"SAVE FILE=\"device-β.dat\" TYPE=DATA TSTEP=(0,5,10)\r\n"

// The budget covers a representative 131 KiB input on CI-class hardware.
// Measured on seanf@GEYSER: fresh 45 ms, incremental 183 ms per edit.
// It is deliberately loose to avoid treating workstation speed as a language contract.
func TestEditorLargeFileLatency(t *testing.T) {
	const budget = 2 * time.Second
	const header = "TITLE Incremental 😀 benchmark\r\n"
	const blocks = 1000
	original := []byte(header + strings.Repeat(editorBenchmarkBlock, blocks) + "END")
	modified := bytes.Clone(original)
	start := len(header) + (blocks/2)*len(editorBenchmarkBlock) + strings.Index(editorBenchmarkBlock, "1.25")
	copy(modified[start:start+4], "2.75")
	parser := editorParser(t)
	began := time.Now()
	old := parser.Parse(original, nil)
	if old == nil {
		t.Fatal("fresh parse returned nil")
	}
	defer old.Close()
	if elapsed := time.Since(began); elapsed > budget {
		t.Errorf("fresh 131 KiB parse took %s; budget %s", elapsed, budget)
	}
	edit := editorInputEdit(original, modified, start, start+4, start+4)
	old.Edit(&edit)
	began = time.Now()
	next := parser.Parse(modified, old)
	if next == nil {
		t.Fatal("incremental parse returned nil")
	}
	defer next.Close()
	if elapsed := time.Since(began); elapsed > budget {
		t.Errorf("incremental 131 KiB edit took %s; budget %s", elapsed, budget)
	}
}

// Both modes alternate the same middle-document numeric edit. Setup, document
// copies, full validation, and final cleanup are excluded from timing. Timed
// work includes Tree.Edit in incremental mode, Parse, and replacing/closing the
// previous tree. ReportAllocs covers Go allocations, not Tree-sitter's C heap.
func BenchmarkEditorParse(b *testing.B) {
	const header = "TITLE Incremental 😀 benchmark\r\n"
	const blocks = 1000
	original := []byte(header + strings.Repeat(editorBenchmarkBlock, blocks) + "END")
	modified := bytes.Clone(original)
	start := len(header) + (blocks/2)*len(editorBenchmarkBlock) + strings.Index(editorBenchmarkBlock, "1.25")
	copy(modified[start:start+4], "2.75")
	forward := editorInputEdit(original, modified, start, start+4, start+4)
	backward := editorInputEdit(modified, original, start, start+4, start+4)
	for _, incremental := range []bool{false, true} {
		name := "Fresh"
		if incremental {
			name = "Incremental"
		}
		b.Run(name, func(b *testing.B) {
			parser := editorParser(b)
			tree := parser.Parse(original, nil)
			if tree == nil {
				b.Fatal("initial parse returned nil")
			}
			defer func() { tree.Close() }()
			if editorHasProblem(tree.RootNode()) {
				b.Fatal("benchmark document did not parse cleanly")
			}
			b.ReportAllocs()
			b.SetBytes(int64(len(original)))
			b.ResetTimer()
			for i := range b.N {
				source, edit := modified, &forward
				if i%2 != 0 {
					source, edit = original, &backward
				}
				var oldTree *tree_sitter.Tree
				if incremental {
					tree.Edit(edit)
					oldTree = tree
				}
				next := parser.Parse(source, oldTree)
				if next == nil {
					b.Fatal("parse returned nil")
				}
				tree.Close()
				tree = next
			}
			b.StopTimer()
			b.ReportMetric(float64(len(original)), "fixture-bytes")
			if editorHasProblem(tree.RootNode()) {
				b.Fatal("edited benchmark document did not parse cleanly")
			}
		})
	}
}
