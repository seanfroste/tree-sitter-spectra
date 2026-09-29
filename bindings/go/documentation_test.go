package tree_sitter_spectra_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"testing"

	spectra "github.com/seanfroste/tree-sitter-spectra/bindings/go"
	sitter "github.com/tree-sitter/go-tree-sitter"
)

type documentationParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Kind  string `json:"kind"`
}

type documentationCall struct {
	Name      string   `json:"name"`
	Arguments []string `json:"arguments"`
}

type documentationExpectation struct {
	Statements []string                 `json:"statements"`
	Parameters []documentationParameter `json:"parameters"`
	References []string                 `json:"references"`
	Calls      []documentationCall      `json:"calls"`
}

type documentationCase struct {
	ID          string                   `json:"id"`
	Document    string                   `json:"document"`
	SourceLines []int                    `json:"source_lines"`
	Original    string                   `json:"original"`
	Input       string                   `json:"input"`
	Disposition string                   `json:"disposition"`
	Reason      string                   `json:"reason"`
	Expect      documentationExpectation `json:"expect"`
}

func documentationRead(t *testing.T, path string) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "..", filepath.FromSlash(path)))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func documentationParser(t testing.TB) *sitter.Parser {
	t.Helper()
	parser := sitter.NewParser()
	t.Cleanup(parser.Close)
	if err := parser.SetLanguage(sitter.NewLanguage(spectra.Language())); err != nil {
		t.Fatal(err)
	}
	return parser
}

func documentationWalk(node *sitter.Node, visit func(*sitter.Node)) {
	visit(node)
	cursor := node.Walk()
	defer cursor.Close()
	if !cursor.GotoFirstChild() {
		return
	}
	for {
		child := cursor.Node()
		documentationWalk(child, visit)
		if !cursor.GotoNextSibling() {
			break
		}
	}
}

func documentationClean(t *testing.T, root *sitter.Node, source []byte) {
	t.Helper()
	documentationWalk(root, func(node *sitter.Node) {
		if node.IsError() || node.IsMissing() || node.Kind() == "unparsed_line" {
			t.Errorf("unexpected %s at %v: %q", node.Kind(), node.Range(), node.Utf8Text(source))
		}
		if node.StartByte() > node.EndByte() || node.EndByte() > uint(len(source)) {
			t.Fatalf("invalid node byte range: %v", node.Range())
		}
		for _, endpoint := range []struct {
			offset uint
			point  sitter.Point
		}{{node.StartByte(), node.StartPosition()}, {node.EndByte(), node.EndPosition()}} {
			prefix := source[:endpoint.offset]
			line := bytes.Count(prefix, []byte{'\n'})
			column := len(prefix) - bytes.LastIndexByte(prefix, '\n') - 1
			if endpoint.point != (sitter.Point{Row: uint(line), Column: uint(column)}) {
				t.Errorf("incorrect byte point %v at offset %d", endpoint.point, endpoint.offset)
			}
		}
	})
}

func documentationProject(t *testing.T, root *sitter.Node, source []byte) documentationExpectation {
	t.Helper()
	got := documentationExpectation{Statements: []string{}, Parameters: []documentationParameter{}, References: []string{}, Calls: []documentationCall{}}
	documentationWalk(root, func(node *sitter.Node) {
		switch node.Kind() {
		case "statement", "title_statement", "if_statement", "else_statement", "endif_statement", "end_statement":
			name := node.ChildByFieldName("name")
			if name == nil || name.Kind() != "statement_name" {
				t.Errorf("statement missing name field: %s", node.ToSexp())
				return
			}
			got.Statements = append(got.Statements, name.Utf8Text(source))
		case "parameter":
			name, operator, value := node.ChildByFieldName("name"), node.ChildByFieldName("operator"), node.ChildByFieldName("value")
			if name == nil || operator == nil || value == nil {
				t.Errorf("parameter missing a field: %s", node.ToSexp())
				return
			}
			if operator.Utf8Text(source) != "=" || name.EndByte() > operator.StartByte() || operator.EndByte() > value.StartByte() || value.EndByte() != node.EndByte() {
				t.Errorf("parameter fields have incorrect boundaries: %s", node.ToSexp())
			}
			got.Parameters = append(got.Parameters, documentationParameter{name.Utf8Text(source), value.Utf8Text(source), value.Kind()})
		case "variable_reference":
			got.References = append(got.References, node.Utf8Text(source))
		case "function_call":
			name, arguments := node.ChildByFieldName("name"), node.ChildByFieldName("arguments")
			if name == nil {
				t.Errorf("function missing name: %s", node.ToSexp())
				return
			}
			call := documentationCall{Name: name.Utf8Text(source), Arguments: []string{}}
			if arguments != nil {
				for i := uint(0); i < arguments.NamedChildCount(); i++ {
					call.Arguments = append(call.Arguments, arguments.NamedChild(i).Utf8Text(source))
				}
			}
			got.Calls = append(got.Calls, call)
		}
	})
	return got
}

func TestDocumentationOracle(t *testing.T) {
	var inventory struct {
		Version int                 `json:"version"`
		Hashes  map[string]string   `json:"source_sha256"`
		Cases   []documentationCase `json:"cases"`
		Other   []json.RawMessage   `json:"other_fenced_material"`
	}
	if err := json.Unmarshal(documentationRead(t, "test/documentation.json"), &inventory); err != nil {
		t.Fatal(err)
	}
	var baseline struct {
		Samples []struct {
			ID, Document, Input string
			SourceLines         []int `json:"source_lines"`
		} `json:"samples"`
		Hashes map[string]string `json:"source_sha256"`
		Other  []json.RawMessage `json:"other_fenced_material"`
	}
	if err := json.Unmarshal(documentationRead(t, "docs/development/completeness-baseline.json"), &baseline); err != nil {
		t.Fatal(err)
	}
	if inventory.Version != 1 {
		t.Fatalf("unsupported oracle version %d", inventory.Version)
	}
	if !reflect.DeepEqual(inventory.Other, baseline.Other) {
		t.Error("non-example fenced material must remain accounted for")
	}
	for path, expected := range inventory.Hashes {
		// Git may materialize Markdown with CRLF on Windows; inventory hashes are canonical LF.
		data := bytes.ReplaceAll(documentationRead(t, path), []byte("\r\n"), []byte("\n"))
		hash := sha256.Sum256(data)
		if hex.EncodeToString(hash[:]) != expected {
			t.Errorf("source changed: %s; review its associated oracle cases", path)
		}
	}
	paths, err := filepath.Glob(filepath.Join("..", "..", "docs", "*.md"))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range paths {
		key := "docs/" + filepath.Base(path)
		if _, exists := inventory.Hashes[key]; !exists {
			t.Errorf("reference missing from inventory: %s", key)
		}
	}
	for path := range baseline.Hashes {
		if _, ok := inventory.Hashes[path]; !ok {
			t.Errorf("historical source omitted: %s", path)
		}
	}
	cases := make(map[string]documentationCase, len(inventory.Cases))
	for _, sample := range inventory.Cases {
		if _, exists := cases[sample.ID]; exists {
			t.Fatalf("duplicate inventory identity %s", sample.ID)
		}
		cases[sample.ID] = sample
	}
	for _, original := range baseline.Samples {
		sample, ok := cases[original.ID]
		if !ok {
			t.Errorf("candidate omitted: %s", original.ID)
			continue
		}
		if sample.Original != original.Input || sample.Document != original.Document || !reflect.DeepEqual(sample.SourceLines, original.SourceLines) {
			t.Errorf("lost historical source traceability: %s", original.ID)
		}
	}
	parser := documentationParser(t)
	for _, sample := range inventory.Cases {
		t.Run(sample.ID, func(t *testing.T) {
			if sample.Reason == "" || sample.ID == "" || sample.Document == "" || len(sample.SourceLines) == 0 {
				t.Fatal("incomplete source disposition")
			}
			if _, ok := inventory.Hashes[sample.Document]; !ok {
				t.Fatal("source not hashed")
			}
			lineCount := bytes.Count(documentationRead(t, sample.Document), []byte{'\n'}) + 1
			for _, line := range sample.SourceLines {
				if line < 1 || line > lineCount {
					t.Fatalf("source line out of range: %d", line)
				}
			}
			switch sample.Disposition {
			case "verbatim":
				if sample.Input != sample.Original {
					t.Fatal("verbatim input was changed")
				}
			case "corrected", "local-policy":
			case "non-executable":
				if sample.Input != "" {
					t.Fatal("non-executable entry must not silently contain executable input")
				}
				t.Logf("accounted-for non-executable source: %s", sample.Reason)
				return
			default:
				t.Fatalf("unknown disposition %q", sample.Disposition)
			}
			if sample.Input == "" {
				t.Fatal("executable input is empty")
			}
			source := []byte(sample.Input + "\n")
			tree := parser.Parse(source, nil)
			if tree == nil {
				t.Fatal("parse cancelled")
			}
			defer tree.Close()
			documentationClean(t, tree.RootNode(), source)
			got := documentationProject(t, tree.RootNode(), source)
			if !reflect.DeepEqual(got, sample.Expect) {
				actual, _ := json.MarshalIndent(got, "", "  ")
				want, _ := json.MarshalIndent(sample.Expect, "", "  ")
				t.Errorf("semantic projection mismatch\ngot: %s\nwant: %s", actual, want)
			}
		})
	}
}

func TestCheckedInFixtures(t *testing.T) {
	parser := documentationParser(t)
	root := filepath.Join("..", "..", "test", "fixtures")
	err := filepath.WalkDir(root, func(path string, entry fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() || filepath.Ext(path) != ".in" {
			return nil
		}
		t.Run(filepath.ToSlash(strings.TrimPrefix(path, root+string(filepath.Separator))), func(t *testing.T) {
			source, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			tree := parser.Parse(source, nil)
			if tree == nil {
				t.Fatal("parse cancelled")
			}
			defer tree.Close()
			documentationClean(t, tree.RootNode(), source)
			documentationWalk(tree.RootNode(), func(node *sitter.Node) {
				if node.Kind() == "character_value" && documentationReference.MatchString(node.Utf8Text(source)) {
					t.Errorf("reference misclassified as string at %v: %q", node.Range(), node.Utf8Text(source))
				}
			})
		})
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

var documentationReference = regexp.MustCompile(`^[@#$][A-Za-z_][A-Za-z0-9_.]*$`)

func TestHighlightQueryRecognizesReferences(t *testing.T) {
	parser := documentationParser(t)
	source := []byte("SAVE FILE=$INPUTFILE\n")
	tree := parser.Parse(source, nil)
	defer tree.Close()
	query, err := sitter.NewQuery(sitter.NewLanguage(spectra.Language()), string(documentationRead(t, "queries/highlights.scm")))
	if err != nil {
		t.Fatal(err)
	}
	defer query.Close()
	cursor := sitter.NewQueryCursor()
	defer cursor.Close()
	matches := cursor.Matches(query, tree.RootNode(), source)
	got := []string{}
	for match := matches.Next(); match != nil; match = matches.Next() {
		for _, capture := range match.Captures {
			got = append(got, fmt.Sprintf("%s:%s", query.CaptureNames()[capture.Index], capture.Node.Utf8Text(source)))
		}
	}
	want := []string{"keyword:SAVE", "property:FILE", "operator:=", "variable:$INPUTFILE"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("highlight captures = %v; want %v", got, want)
	}
}
