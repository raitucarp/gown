package generative

import (
	"fmt"
	"strings"
)

// ParseNode represents a constituent node in a phrase-structure parse tree.
type ParseNode struct {
	Symbol   string       `json:"symbol"`             // e.g. "S", "NP", "VP", "N", "V"
	Terminal string       `json:"terminal,omitempty"` // surface word if leaf
	Children []*ParseNode `json:"children,omitempty"`
}

// IsLeaf returns true if this node is a terminal lexical item.
func (n *ParseNode) IsLeaf() bool {
	return len(n.Children) == 0 && n.Terminal != ""
}

// Sentence returns the concatenated surface yield of the parse tree.
func (n *ParseNode) Sentence() string {
	if n.IsLeaf() {
		return n.Terminal
	}
	var words []string
	for _, child := range n.Children {
		words = append(words, child.Sentence())
	}
	return strings.Join(words, " ")
}

// BracketedString serializes the parse tree into Penn Treebank bracketed style:
// (S (NP (Det The) (N dog)) (VP (V chased) (NP (Det the) (N cat))))
func (n *ParseNode) BracketedString() string {
	if n.IsLeaf() {
		return fmt.Sprintf("(%s %s)", n.Symbol, n.Terminal)
	}
	var parts []string
	for _, child := range n.Children {
		parts = append(parts, child.BracketedString())
	}
	return fmt.Sprintf("(%s %s)", n.Symbol, strings.Join(parts, " "))
}

// RenderTree returns an indented ASCII representation of the constituency tree.
func (n *ParseNode) RenderTree() string {
	var sb strings.Builder
	n.renderIndent("", true, true, &sb)
	return sb.String()
}

func (n *ParseNode) renderIndent(prefix string, isLast bool, isRoot bool, sb *strings.Builder) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if isRoot {
		connector = ""
	}

	label := n.Symbol
	if n.Terminal != "" {
		label += ": " + n.Terminal
	}

	sb.WriteString(prefix + connector + label + "\n")

	childPrefix := prefix
	if !isRoot {
		if isLast {
			childPrefix += "    "
		} else {
			childPrefix += "│   "
		}
	}

	for i, child := range n.Children {
		last := i == len(n.Children)-1
		child.renderIndent(childPrefix, last, false, sb)
	}
}
