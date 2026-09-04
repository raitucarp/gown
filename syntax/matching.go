package syntax

import (
	"strings"
)

// FindSubtrees finds all subtrees rooted with the specified constituent label.
func FindSubtrees(root *SyntaxNode, label string) []*SyntaxNode {
	if root == nil {
		return nil
	}
	var matches []*SyntaxNode
	if strings.EqualFold(root.Label, label) {
		matches = append(matches, root)
	}
	for _, child := range root.Children {
		matches = append(matches, FindSubtrees(child, label)...)
	}
	return matches
}

// MatchesRule checks if a node expands according to a specific production rule pattern:
// e.g. "VP -> V NP" or "NP -> Det N"
func MatchesRule(node *SyntaxNode, rulePattern string) bool {
	if node == nil || node.IsLeaf() {
		return false
	}
	parts := strings.Split(rulePattern, "->")
	if len(parts) != 2 {
		return false
	}

	lhs := strings.TrimSpace(parts[0])
	rhsTokens := strings.Fields(strings.TrimSpace(parts[1]))

	if !strings.EqualFold(node.Label, lhs) {
		return false
	}

	if len(node.Children) != len(rhsTokens) {
		return false
	}

	for i, token := range rhsTokens {
		if token == "*" || token == "_" {
			continue
		}
		if !strings.EqualFold(node.Children[i].Label, token) {
			return false
		}
	}

	return true
}

// FindByRule searches the entire tree for constituents matching a phrase structure rule.
func FindByRule(root *SyntaxNode, rulePattern string) []*SyntaxNode {
	if root == nil {
		return nil
	}
	var matches []*SyntaxNode
	if MatchesRule(root, rulePattern) {
		matches = append(matches, root)
	}
	for _, child := range root.Children {
		matches = append(matches, FindByRule(child, rulePattern)...)
	}
	return matches
}
