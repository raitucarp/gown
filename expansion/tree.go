package expansion

import (
	"fmt"
	"strings"
)

// TraversalStrategy determines the order of exploration.
type TraversalStrategy int

const (
	// StrategyBFS performs breadth-first expansion.
	StrategyBFS TraversalStrategy = iota
	// StrategyDFS performs depth-first expansion.
	StrategyDFS
)

// NodeType identifies the entity represented by an expansion node.
type NodeType string

const (
	NodeWord       NodeType = "word"
	NodeSynset     NodeType = "synset"
	NodeDefinition NodeType = "definition"
	NodeToken      NodeType = "token"
)

// Node represents a vertex in an expansion tree.
type Node struct {
	ID         string   `json:"id"`
	Word       string   `json:"word"`
	Type       NodeType `json:"type"`
	POS        string   `json:"pos,omitempty"`
	Definition string   `json:"definition,omitempty"`
	Relation   string   `json:"relation,omitempty"`
	Depth      int      `json:"depth"`
	Score      float64  `json:"score"`
	Children   []*Node  `json:"children,omitempty"`
}

// Tree represents a hierarchical lexical or definition expansion.
type Tree struct {
	Root       *Node `json:"root"`
	TotalNodes int   `json:"total_nodes"`
}

// Render returns a formatted ASCII tree representation.
func (t *Tree) Render() string {
	if t == nil || t.Root == nil {
		return "<empty tree>"
	}
	var sb strings.Builder
	t.renderNode(t.Root, "", true, true, &sb)
	return sb.String()
}

func (t *Tree) renderNode(n *Node, prefix string, isLast bool, isRoot bool, sb *strings.Builder) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if isRoot {
		connector = ""
	}

	label := n.Word
	if n.Relation != "" {
		label = fmt.Sprintf("[%s] %s", n.Relation, label)
	}
	if n.Definition != "" {
		label = fmt.Sprintf("%s: \"%s\"", label, n.Definition)
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
		t.renderNode(child, childPrefix, last, false, sb)
	}
}
