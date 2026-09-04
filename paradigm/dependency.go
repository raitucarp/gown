package paradigm

import (
	"fmt"
	"strings"
)

// DependencyNode represents a token in a dependency syntax tree (Universal Dependencies format).
type DependencyNode struct {
	ID     int    `json:"id"`
	Form   string `json:"form"`
	Lemma  string `json:"lemma"`
	UPOS   string `json:"upos"`
	Head   int    `json:"head"` // 0 if root
	DepRel string `json:"deprel"`
}

// DependencyTree represents a parsed sentence dependency structure.
type DependencyTree struct {
	Nodes []DependencyNode `json:"nodes"`
}

// Root returns the root node of the dependency tree (where Head == 0).
func (dt *DependencyTree) Root() *DependencyNode {
	for i := range dt.Nodes {
		if dt.Nodes[i].Head == 0 {
			return &dt.Nodes[i]
		}
	}
	return nil
}

// DependentsOf finds all immediate dependent children of a given node ID.
func (dt *DependencyTree) DependentsOf(nodeID int) []DependencyNode {
	var deps []DependencyNode
	for _, n := range dt.Nodes {
		if n.Head == nodeID {
			deps = append(deps, n)
		}
	}
	return deps
}

// CoNLLU formats the dependency tree into standard CoNLL-U format.
func (dt *DependencyTree) CoNLLU() string {
	var sb strings.Builder
	for _, n := range dt.Nodes {
		sb.WriteString(fmt.Sprintf("%d\t%s\t%s\t%s\t_\t_\t%d\t%s\t_\t_\n",
			n.ID, n.Form, n.Lemma, n.UPOS, n.Head, n.DepRel))
	}
	return sb.String()
}
