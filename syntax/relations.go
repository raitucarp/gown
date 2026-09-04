package syntax

import (
	"strings"
)

// GrammaticalRole designates the syntactic relation of a constituent.
type GrammaticalRole string

const (
	RoleSubject       GrammaticalRole = "subject"
	RolePredicateVerb GrammaticalRole = "predicate_verb"
	RoleDirectObject  GrammaticalRole = "direct_object"
	RoleIndirectObject GrammaticalRole = "indirect_object"
	RolePrepObject    GrammaticalRole = "prepositional_object"
	RoleAdjunct       GrammaticalRole = "adjunct"
)

// GrammaticalRelation pairs a constituent subtree with its functional role.
type GrammaticalRelation struct {
	Role GrammaticalRole `json:"role"`
	Node *SyntaxNode     `json:"-"`
	Text string          `json:"text"`
	Head string          `json:"head"`
}

// ExtractRelations analyzes an S-level clause tree and returns its core grammatical relations.
func ExtractRelations(root *SyntaxNode) []GrammaticalRelation {
	if root == nil {
		return nil
	}

	var relations []GrammaticalRelation

	addRel := func(role GrammaticalRole, node *SyntaxNode) {
		if node == nil {
			return
		}
		headNode := FindHead(node)
		headText := ""
		if headNode != nil {
			headText = headNode.Terminal
		}
		relations = append(relations, GrammaticalRelation{
			Role: role,
			Node: node,
			Text: node.Yield(),
			Head: headText,
		})
	}

	// In standard S -> NP VP
	var subjectNode *SyntaxNode
	var vpNode *SyntaxNode

	for _, child := range root.Children {
		label := strings.ToUpper(child.Label)
		if strings.HasPrefix(label, "NP") && subjectNode == nil && vpNode == nil {
			subjectNode = child
		} else if strings.HasPrefix(label, "VP") {
			vpNode = child
		}
	}

	if subjectNode != nil {
		addRel(RoleSubject, subjectNode)
	}

	if vpNode != nil {
		// Predicate verb is head of VP
		if vHead := FindHead(vpNode); vHead != nil {
			relations = append(relations, GrammaticalRelation{
				Role: RolePredicateVerb,
				Node: vHead,
				Text: vHead.Terminal,
				Head: vHead.Terminal,
			})
		}

		// Look for objects and adjuncts inside VP
		var npObjects []*SyntaxNode
		for _, child := range vpNode.Children {
			cLabel := strings.ToUpper(child.Label)
			if strings.HasPrefix(cLabel, "NP") {
				npObjects = append(npObjects, child)
			} else if strings.HasPrefix(cLabel, "PP") {
				addRel(RoleAdjunct, child)
				// Within PP, find prepositional object
				for _, ppChild := range child.Children {
					if strings.HasPrefix(strings.ToUpper(ppChild.Label), "NP") {
						addRel(RolePrepObject, ppChild)
					}
				}
			}
		}

		if len(npObjects) == 1 {
			addRel(RoleDirectObject, npObjects[0])
		} else if len(npObjects) >= 2 {
			addRel(RoleIndirectObject, npObjects[0])
			addRel(RoleDirectObject, npObjects[1])
		}
	}

	return relations
}
