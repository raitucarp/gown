package syntax

import (
	"fmt"
	"strings"
)

// SyntaxNode represents a constituent node in a syntactic parse tree.
type SyntaxNode struct {
	Label    string        `json:"label"`              // e.g. "S", "NP", "VP", "N", "V"
	Terminal string        `json:"terminal,omitempty"`  // surface word form if leaf
	Parent   *SyntaxNode   `json:"-"`                  // parent pointer
	Children []*SyntaxNode `json:"children,omitempty"` // child constituents
	Features map[string]string `json:"features,omitempty"`
}

// NewNode constructs a non-terminal syntax node.
func NewNode(label string, children ...*SyntaxNode) *SyntaxNode {
	node := &SyntaxNode{
		Label:    label,
		Features: make(map[string]string),
	}
	for _, child := range children {
		node.AddChild(child)
	}
	return node
}

// NewLeaf constructs a terminal syntax leaf node.
func NewLeaf(label, terminal string) *SyntaxNode {
	return &SyntaxNode{
		Label:    label,
		Terminal: terminal,
		Features: make(map[string]string),
	}
}

// AddChild attaches a child node, maintaining parent references.
func (n *SyntaxNode) AddChild(child *SyntaxNode) {
	if child == nil {
		return
	}
	child.Parent = n
	n.Children = append(n.Children, child)
}

// IsLeaf returns true if the node represents a terminal lexical word.
func (n *SyntaxNode) IsLeaf() bool {
	return len(n.Children) == 0 && n.Terminal != ""
}

// Siblings returns other children of the same parent.
func (n *SyntaxNode) Siblings() []*SyntaxNode {
	if n.Parent == nil {
		return nil
	}
	var sibs []*SyntaxNode
	for _, child := range n.Parent.Children {
		if child != n {
			sibs = append(sibs, child)
		}
	}
	return sibs
}

// Ancestors returns all ancestor nodes from parent up to root.
func (n *SyntaxNode) Ancestors() []*SyntaxNode {
	var ancs []*SyntaxNode
	curr := n.Parent
	for curr != nil {
		ancs = append(ancs, curr)
		curr = curr.Parent
	}
	return ancs
}

// Descendants traverses and collects all descendant nodes.
func (n *SyntaxNode) Descendants() []*SyntaxNode {
	var descs []*SyntaxNode
	for _, child := range n.Children {
		descs = append(descs, child)
		descs = append(descs, child.Descendants()...)
	}
	return descs
}

// Leaves returns all terminal leaf nodes in left-to-right yield order.
func (n *SyntaxNode) Leaves() []*SyntaxNode {
	if n.IsLeaf() {
		return []*SyntaxNode{n}
	}
	var leaves []*SyntaxNode
	for _, child := range n.Children {
		leaves = append(leaves, child.Leaves()...)
	}
	return leaves
}

// Yield returns the full sentence string formed by the leaf terminals.
func (n *SyntaxNode) Yield() string {
	leaves := n.Leaves()
	words := make([]string, len(leaves))
	for i, l := range leaves {
		words[i] = l.Terminal
	}
	return strings.Join(words, " ")
}

// Depth returns the maximum height of the subtree.
func (n *SyntaxNode) Depth() int {
	if len(n.Children) == 0 {
		return 1
	}
	maxD := 0
	for _, child := range n.Children {
		d := child.Depth()
		if d > maxD {
			maxD = d
		}
	}
	return maxD + 1
}

// Size returns the total count of nodes in the subtree.
func (n *SyntaxNode) Size() int {
	count := 1
	for _, child := range n.Children {
		count += child.Size()
	}
	return count
}

// Clone creates a deep copy of the syntax tree.
func (n *SyntaxNode) Clone() *SyntaxNode {
	if n == nil {
		return nil
	}
	cloned := &SyntaxNode{
		Label:    n.Label,
		Terminal: n.Terminal,
		Features: make(map[string]string),
	}
	for k, v := range n.Features {
		cloned.Features[k] = v
	}
	for _, child := range n.Children {
		cloned.AddChild(child.Clone())
	}
	return cloned
}

// ReplaceChild substitutes an existing child with a new subtree.
func (n *SyntaxNode) ReplaceChild(oldChild, newChild *SyntaxNode) bool {
	for i, child := range n.Children {
		if child == oldChild {
			newChild.Parent = n
			n.Children[i] = newChild
			return true
		}
	}
	return false
}

// BracketedString returns Penn Treebank format: (S (NP (Det The) (N dog)) (VP (V barks)))
func (n *SyntaxNode) BracketedString() string {
	if n.IsLeaf() {
		return fmt.Sprintf("(%s %s)", n.Label, n.Terminal)
	}
	var parts []string
	for _, child := range n.Children {
		parts = append(parts, child.BracketedString())
	}
	return fmt.Sprintf("(%s %s)", n.Label, strings.Join(parts, " "))
}

// Render returns an indented ASCII representation of the tree.
func (n *SyntaxNode) Render() string {
	var sb strings.Builder
	n.renderNode("", true, true, &sb)
	return sb.String()
}

func (n *SyntaxNode) renderNode(prefix string, isLast bool, isRoot bool, sb *strings.Builder) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if isRoot {
		connector = ""
	}

	label := n.Label
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
		child.renderNode(childPrefix, last, false, sb)
	}
}
