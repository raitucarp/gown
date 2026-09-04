package syntax

import (
	"strings"
)

// FindHead returns the lexical head leaf of a given constituent node using Collins-style head rules.
func FindHead(node *SyntaxNode) *SyntaxNode {
	if node == nil {
		return nil
	}
	if node.IsLeaf() {
		return node
	}

	label := strings.ToUpper(node.Label)

	switch {
	case strings.HasPrefix(label, "NP"):
		// Head of NP: rightmost noun or NP child
		for i := len(node.Children) - 1; i >= 0; i-- {
			child := node.Children[i]
			cLabel := strings.ToUpper(child.Label)
			if strings.HasPrefix(cLabel, "N") || strings.HasPrefix(cLabel, "PRP") || cLabel == "PROPERNOUN" {
				return FindHead(child)
			}
		}
		// Fallback to rightmost child
		if len(node.Children) > 0 {
			return FindHead(node.Children[len(node.Children)-1])
		}

	case strings.HasPrefix(label, "VP"):
		// Head of VP: leftmost verb or auxiliary
		for _, child := range node.Children {
			cLabel := strings.ToUpper(child.Label)
			if strings.HasPrefix(cLabel, "V") || strings.HasPrefix(cLabel, "MD") {
				return FindHead(child)
			}
		}
		// Fallback to leftmost child
		if len(node.Children) > 0 {
			return FindHead(node.Children[0])
		}

	case strings.HasPrefix(label, "PP"):
		// Head of PP: leftmost preposition
		for _, child := range node.Children {
			cLabel := strings.ToUpper(child.Label)
			if cLabel == "P" || cLabel == "IN" || cLabel == "TO" {
				return FindHead(child)
			}
		}
		if len(node.Children) > 0 {
			return FindHead(node.Children[0])
		}

	case strings.HasPrefix(label, "S"):
		// Head of S: VP child
		for _, child := range node.Children {
			if strings.HasPrefix(strings.ToUpper(child.Label), "VP") {
				return FindHead(child)
			}
		}
		if len(node.Children) > 0 {
			return FindHead(node.Children[len(node.Children)-1])
		}

	case strings.HasPrefix(label, "ADJ") || label == "AP":
		for _, child := range node.Children {
			cLabel := strings.ToUpper(child.Label)
			if strings.HasPrefix(cLabel, "ADJ") || strings.HasPrefix(cLabel, "JJ") {
				return FindHead(child)
			}
		}

	case strings.HasPrefix(label, "ADV"):
		for _, child := range node.Children {
			cLabel := strings.ToUpper(child.Label)
			if strings.HasPrefix(cLabel, "ADV") || strings.HasPrefix(cLabel, "RB") {
				return FindHead(child)
			}
		}
	}

	// Default fallback: first child's head
	if len(node.Children) > 0 {
		return FindHead(node.Children[0])
	}

	return node
}
