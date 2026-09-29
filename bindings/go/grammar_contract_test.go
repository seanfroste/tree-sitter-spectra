package tree_sitter_spectra_test

import (
	"strings"
	"testing"

	sitter "github.com/tree-sitter/go-tree-sitter"
)

func expressionShape(node *sitter.Node, source []byte) string {
	if node == nil {
		return "missing"
	}
	switch node.Kind() {
	case "binary_expression":
		return node.ChildByFieldName("operator").Utf8Text(source) + "(" + expressionShape(node.ChildByFieldName("left"), source) + "," + expressionShape(node.ChildByFieldName("right"), source) + ")"
	case "unary_expression":
		return node.ChildByFieldName("operator").Utf8Text(source) + "(" + expressionShape(node.ChildByFieldName("operand"), source) + ")"
	case "braced_expression", "parenthesized_expression":
		return expressionShape(node.NamedChild(0), source)
	default:
		return node.Utf8Text(source)
	}
}

func TestExpressionStructure(t *testing.T) {
	parser := documentationParser(t)
	for _, test := range []struct{ input, shape string }{
		{"1+2*3", "+(1,*(2,3))"},
		{"1 + 2 * 3", "+(1,*(2,3))"},
		{"1 * 2 + 3", "+(*(1,2),3)"},
		{"8 - 3 - 1", "-(-(8,3),1)"},
		{"2 ^ 3 ^ 2", "^(2,^(3,2))"},
		{"(1+2)*3", "*(+(1,2),3)"},
		{"@Width-0.5", "-(@Width,0.5)"},
		{"PI-1", "-(PI,1)"},
		{"-2*3", "*(-(2),3)"},
		{"1 < 2 && 3 != 4", "&&(<(1,2),!=(3,4))"},
	} {
		t.Run(test.input, func(t *testing.T) {
			source := []byte("DEFINE VALUE={" + test.input + "}\n")
			tree := parser.Parse(source, nil)
			defer tree.Close()
			documentationClean(t, tree.RootNode(), source)
			value := tree.RootNode().NamedChild(0).NamedChild(1).ChildByFieldName("value")
			if got := expressionShape(value, source); got != test.shape {
				t.Errorf("expression shape %s; want %s", got, test.shape)
			}
		})
	}
}

func TestMultilineExpressionPreservesOperandStructure(t *testing.T) {
	parser := documentationParser(t)
	source := []byte("DEFINE VALUE={1+\n2}\n")
	tree := parser.Parse(source, nil)
	defer tree.Close()
	documentationClean(t, tree.RootNode(), source)
	value := tree.RootNode().NamedChild(0).NamedChild(1).ChildByFieldName("value")
	if got := expressionShape(value, source); got != "+(1,2)" {
		t.Fatalf("multiline expression shape %s; want +(1,2)", got)
	}
}

func TestMalformedSyntaxPreservesLaterStatement(t *testing.T) {
	parser := documentationParser(t)
	for _, line := range []string{
		"GRID XMIN=", "DEFINE CHARACTER=\"unfinished", "DEFINE VALUE={1 + * 2}",
		"DEFINE VALUE={1+(2*3", "DEFINE VALUE={}", "DEFINE VALUE={1 2}",
		"REGION MASK=FILE(mask.txt,cell,1", "DEFINE VALUE=(1 2", "REGION &[MASK=FILE(a,b,1)",
		"GRID X=1Y=2", "DEFINE VALUE={1+}",
	} {
		t.Run(line, func(t *testing.T) {
			source := []byte(line + "\nSAVE FILE=after.dat\n")
			tree := parser.Parse(source, nil)
			defer tree.Close()
			if !editorHasProblem(tree.RootNode()) {
				t.Error("malformed input accepted")
			}
			var save *sitter.Node
			for i := uint(0); i < tree.RootNode().NamedChildCount(); i++ {
				n := tree.RootNode().NamedChild(i)
				if n.Kind() == "statement" && n.StartByte() == uint(len(line)+1) {
					save = n
				}
			}
			if save == nil || save.HasError() || save.Utf8Text(source) != "SAVE FILE=after.dat" {
				t.Errorf("later SAVE lost: %s", tree.RootNode().ToSexp())
			}
		})
	}
}

func TestUnclosedSeriesPreservesFollowingTitle(t *testing.T) {
	parser := documentationParser(t)
	source := []byte("DEFINE VALUE=(1\nTITLE Device\n")
	tree := parser.Parse(source, nil)
	defer tree.Close()
	root := tree.RootNode()
	if root.NamedChildCount() != 2 {
		t.Fatalf("top-level nodes = %s; want malformed statement and separate TITLE", root.ToSexp())
	}
	if root.NamedChild(0).Kind() != "statement" || root.NamedChild(1).Kind() != "title_statement" {
		t.Fatalf("following TITLE was not preserved: %s", root.ToSexp())
	}
	title := root.NamedChild(1).ChildByFieldName("title")
	if title == nil || title.Utf8Text(source) != "Device" {
		t.Fatalf("TITLE value lost: %s", root.ToSexp())
	}
}

func TestLinkedSeriesMaskGroupsAndPositionalValue(t *testing.T) {
	parser := documentationParser(t)
	source := []byte("DEFINE VALUE=#Size(0.1 0.2)\nREGION &[MASK=FILE (a,b,1) [ANDMASK=POLYGON(0,0 1,1)]]\nplot potential dim=3\nENDPOINT X=1\n \t\n")
	tree := parser.Parse(source, nil)
	defer tree.Close()
	documentationClean(t, tree.RootNode(), source)
	linked := tree.RootNode().NamedChild(0).NamedChild(1).ChildByFieldName("value")
	if linked.Kind() != "linked_series" || linked.ChildByFieldName("reference").Utf8Text(source) != "#Size" {
		t.Fatalf("linked series %s", linked.ToSexp())
	}
	series := linked.ChildByFieldName("values")
	if series.NamedChildCount() != 2 || series.NamedChild(0).Utf8Text(source) != "0.1" || series.NamedChild(1).Utf8Text(source) != "0.2" {
		t.Error("linked series boundaries changed")
	}
	groups := []string{}
	documentationWalk(tree.RootNode(), func(n *sitter.Node) {
		if n.Kind() == "mask_group" {
			op := n.ChildByFieldName("operator")
			if op == nil {
				groups = append(groups, "OR")
			} else {
				groups = append(groups, op.Utf8Text(source))
			}
		}
	})
	if strings.Join(groups, ",") != "&,OR" {
		t.Errorf("mask grouping = %v", groups)
	}
	plot := tree.RootNode().NamedChild(2)
	if plot.NamedChild(1).Kind() != "positional_argument" || plot.NamedChild(1).Utf8Text(source) != "potential" {
		t.Error("positional value lost")
	}
	unknown := tree.RootNode().NamedChild(3)
	if unknown.ChildByFieldName("name").Utf8Text(source) != "ENDPOINT" {
		t.Error("keyword prefix split")
	}
}
