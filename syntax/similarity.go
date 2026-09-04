package syntax

import (
	"strings"
)

// ExtractRules collects all context-free productions instantiated in the tree.
func ExtractRules(node *SyntaxNode) []string {
	if node == nil || node.IsLeaf() {
		return nil
	}
	var rhs []string
	for _, c := range node.Children {
		rhs = append(rhs, c.Label)
	}
	rule := node.Label + " -> " + strings.Join(rhs, " ")
	rules := []string{rule}

	for _, c := range node.Children {
		rules = append(rules, ExtractRules(c)...)
	}
	return rules
}

// TreeSimilarity computes the Jaccard similarity of production rules between two syntax trees.
func TreeSimilarity(t1, t2 *SyntaxNode) float64 {
	if t1 == nil && t2 == nil {
		return 1.0
	}
	if t1 == nil || t2 == nil {
		return 0.0
	}

	rules1 := ExtractRules(t1)
	rules2 := ExtractRules(t2)

	set1 := make(map[string]struct{}, len(rules1))
	for _, r := range rules1 {
		set1[r] = struct{}{}
	}
	set2 := make(map[string]struct{}, len(rules2))
	for _, r := range rules2 {
		set2[r] = struct{}{}
	}

	if len(set1) == 0 && len(set2) == 0 {
		// Both are leaves
		if t1.Label == t2.Label && t1.Terminal == t2.Terminal {
			return 1.0
		}
		return 0.0
	}

	intersection := 0
	for r := range set1 {
		if _, ok := set2[r]; ok {
			intersection++
		}
	}

	union := len(set1) + len(set2) - intersection
	if union == 0 {
		return 0.0
	}

	return float64(intersection) / float64(union)
}

// StructuralDistance computes the edit difference in size and height between two trees.
func StructuralDistance(t1, t2 *SyntaxNode) float64 {
	if t1 == nil && t2 == nil {
		return 0.0
	}
	if t1 == nil || t2 == nil {
		return 1.0
	}

	diffSize := float64(abs(t1.Size() - t2.Size()))
	diffDepth := float64(abs(t1.Depth() - t2.Depth()))
	maxSize := float64(max(t1.Size(), t2.Size()))

	if maxSize == 0 {
		return 0.0
	}
	return (diffSize + diffDepth) / (2.0 * maxSize)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
