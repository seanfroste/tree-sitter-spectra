package tree_sitter_spectra_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_spectra "github.com/seanfroste/tree-sitter-spectra/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_spectra.Language())
	if language == nil {
		t.Errorf("Error loading Spectra grammar")
	}
}
