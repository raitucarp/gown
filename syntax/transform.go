package syntax

import (
	"fmt"
	"strings"
)

// ActiveToPassive transforms a transitive active parse tree (S -> NP1 VP[V NP2]) into passive voice:
// NP2 was V-ed by NP1.
func ActiveToPassive(root *SyntaxNode) (*SyntaxNode, error) {
	if root == nil {
		return nil, fmt.Errorf("root node is nil")
	}

	cloned := root.Clone()
	var np1 *SyntaxNode
	var vp *SyntaxNode

	for _, child := range cloned.Children {
		label := strings.ToUpper(child.Label)
		if strings.HasPrefix(label, "NP") && np1 == nil && vp == nil {
			np1 = child
		} else if strings.HasPrefix(label, "VP") {
			vp = child
		}
	}

	if np1 == nil || vp == nil {
		return nil, fmt.Errorf("clause lacks NP subject or VP predicate")
	}

	var verbNode *SyntaxNode
	var np2 *SyntaxNode
	var remainingChildren []*SyntaxNode

	for _, child := range vp.Children {
		cLabel := strings.ToUpper(child.Label)
		if strings.HasPrefix(cLabel, "V") && verbNode == nil {
			verbNode = child
		} else if strings.HasPrefix(cLabel, "NP") && np2 == nil {
			np2 = child
		} else {
			remainingChildren = append(remainingChildren, child)
		}
	}

	if verbNode == nil || np2 == nil {
		return nil, fmt.Errorf("VP lacks transitive verb or direct object NP")
	}

	// Build passive tree:
	// S -> NP2 VP[Aux(was) V(verb) PP[P(by) NP1]]
	newRoot := NewNode("S")
	newRoot.AddChild(np2)

	newVP := NewNode("VP")
	newVP.AddChild(NewLeaf("Aux", "was"))
	newVP.AddChild(verbNode)

	byPP := NewNode("PP", NewLeaf("P", "by"), np1)
	newVP.AddChild(byPP)

	for _, rem := range remainingChildren {
		newVP.AddChild(rem)
	}

	newRoot.AddChild(newVP)
	return newRoot, nil
}

// PassiveToActive transforms a passive parse tree (NP2 was V-ed by NP1) into active voice:
// NP1 V-ed NP2.
func PassiveToActive(root *SyntaxNode) (*SyntaxNode, error) {
	if root == nil {
		return nil, fmt.Errorf("root node is nil")
	}

	cloned := root.Clone()
	var np2 *SyntaxNode
	var vp *SyntaxNode

	for _, child := range cloned.Children {
		label := strings.ToUpper(child.Label)
		if strings.HasPrefix(label, "NP") && np2 == nil && vp == nil {
			np2 = child
		} else if strings.HasPrefix(label, "VP") {
			vp = child
		}
	}

	if np2 == nil || vp == nil {
		return nil, fmt.Errorf("clause lacks passive subject or VP")
	}

	var verbNode *SyntaxNode
	var byPP *SyntaxNode
	var np1 *SyntaxNode

	for _, child := range vp.Children {
		cLabel := strings.ToUpper(child.Label)
		if strings.HasPrefix(cLabel, "V") {
			verbNode = child
		} else if strings.HasPrefix(cLabel, "PP") {
			if len(child.Children) >= 2 && strings.EqualFold(child.Children[0].Terminal, "by") {
				byPP = child
				np1 = child.Children[1]
			}
		}
	}

	if verbNode == nil || np1 == nil {
		return nil, fmt.Errorf("passive VP lacks main verb or 'by' agent phrase")
	}

	newRoot := NewNode("S")
	newRoot.AddChild(np1)

	newVP := NewNode("VP")
	newVP.AddChild(verbNode)
	newVP.AddChild(np2)

	for _, child := range vp.Children {
		if child != verbNode && child != byPP && strings.ToUpper(child.Label) != "AUX" {
			newVP.AddChild(child)
		}
	}

	newRoot.AddChild(newVP)
	return newRoot, nil
}

// NegateClause inserts auxiliary negation into the main VP of the clause.
func NegateClause(root *SyntaxNode) (*SyntaxNode, error) {
	if root == nil {
		return nil, fmt.Errorf("root node is nil")
	}

	cloned := root.Clone()
	var vp *SyntaxNode
	for _, child := range cloned.Children {
		if strings.HasPrefix(strings.ToUpper(child.Label), "VP") {
			vp = child
			break
		}
	}

	if vp == nil {
		return nil, fmt.Errorf("clause lacks VP")
	}

	// Insert "did not" before the main verb if no auxiliary exists
	hasAux := false
	for _, child := range vp.Children {
		if strings.ToUpper(child.Label) == "AUX" || strings.ToUpper(child.Label) == "MD" {
			hasAux = true
			break
		}
	}

	var newVPChildren []*SyntaxNode
	if hasAux {
		for _, child := range vp.Children {
			newVPChildren = append(newVPChildren, child)
			if strings.ToUpper(child.Label) == "AUX" || strings.ToUpper(child.Label) == "MD" {
				newVPChildren = append(newVPChildren, NewLeaf("Neg", "not"))
			}
		}
	} else {
		newVPChildren = append(newVPChildren, NewLeaf("Aux", "did"), NewLeaf("Neg", "not"))
		newVPChildren = append(newVPChildren, vp.Children...)
	}

	vp.Children = newVPChildren
	return cloned, nil
}

// ToInterrogative transforms declarative S -> NP VP into yes/no question structure with auxiliary inversion:
// Aux NP VP...
func ToInterrogative(root *SyntaxNode) (*SyntaxNode, error) {
	if root == nil {
		return nil, fmt.Errorf("root node is nil")
	}

	cloned := root.Clone()
	var np *SyntaxNode
	var vp *SyntaxNode

	for _, child := range cloned.Children {
		label := strings.ToUpper(child.Label)
		if strings.HasPrefix(label, "NP") && np == nil && vp == nil {
			np = child
		} else if strings.HasPrefix(label, "VP") {
			vp = child
		}
	}

	if np == nil || vp == nil {
		return nil, fmt.Errorf("clause lacks NP or VP")
	}

	// Check if VP has an auxiliary
	var auxNode *SyntaxNode
	var remainingVP []*SyntaxNode

	for _, child := range vp.Children {
		if strings.ToUpper(child.Label) == "AUX" && auxNode == nil {
			auxNode = child
		} else {
			remainingVP = append(remainingVP, child)
		}
	}

	if auxNode == nil {
		auxNode = NewLeaf("Aux", "Did")
	} else {
		// Capitalize initial auxiliary
		auxNode.Terminal = strings.Title(auxNode.Terminal)
	}

	vp.Children = remainingVP

	newRoot := NewNode("SQ")
	newRoot.AddChild(auxNode)
	newRoot.AddChild(np)
	newRoot.AddChild(vp)

	return newRoot, nil
}
