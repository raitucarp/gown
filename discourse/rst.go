package discourse

import (
	"fmt"
	"strings"
)

// Nuclearity designates whether an EDU is the central Nucleus or supportive Satellite.
type Nuclearity string

const (
	NuclearityNucleus   Nuclearity = "Nucleus"
	NuclearitySatellite Nuclearity = "Satellite"
)

// RhetoricalRelation specifies the rhetorical relation type in RST.
type RhetoricalRelation string

const (
	RelElaboration   RhetoricalRelation = "Elaboration"
	RelContrast      RhetoricalRelation = "Contrast"
	RelCause         RhetoricalRelation = "Cause"
	RelCondition     RhetoricalRelation = "Condition"
	RelBackground    RhetoricalRelation = "Background"
	RelSequence      RhetoricalRelation = "Sequence"
	RelJoint         RhetoricalRelation = "Joint"
)

// RSTNode represents a discourse constituent in a Rhetorical Structure Theory tree.
type RSTNode struct {
	ID         int                `json:"id"`
	Nuclearity Nuclearity         `json:"nuclearity"`
	Relation   RhetoricalRelation `json:"relation,omitempty"`
	Text       string             `json:"text,omitempty"`
	Children   []*RSTNode         `json:"children,omitempty"`
}

// Render formats the RST tree into an indented string.
func (node *RSTNode) Render() string {
	if node == nil {
		return "<empty rst>"
	}
	var sb strings.Builder
	node.renderIndent("", true, true, &sb)
	return sb.String()
}

func (node *RSTNode) renderIndent(prefix string, isLast bool, isRoot bool, sb *strings.Builder) {
	connector := "├── "
	if isLast {
		connector = "└── "
	}
	if isRoot {
		connector = ""
	}

	label := fmt.Sprintf("[%s]", node.Nuclearity)
	if node.Relation != "" {
		label += fmt.Sprintf(" (%s)", node.Relation)
	}
	if node.Text != "" {
		label += fmt.Sprintf(": \"%s\"", node.Text)
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

	for i, child := range node.Children {
		last := i == len(node.Children)-1
		child.renderIndent(childPrefix, last, false, sb)
	}
}

// BuildRSTTree constructs an RST discourse tree connecting sequential EDUs based on rhetorical markers.
func BuildRSTTree(edus []EDU) *RSTNode {
	if len(edus) == 0 {
		return nil
	}
	if len(edus) == 1 {
		return &RSTNode{
			ID:         edus[0].ID,
			Nuclearity: NuclearityNucleus,
			Text:       edus[0].Text,
		}
	}

	root := &RSTNode{
		ID:         0,
		Nuclearity: NuclearityNucleus,
		Relation:   RelSequence,
	}

	// First EDU is the primary nucleus
	nuc := &RSTNode{
		ID:         edus[0].ID,
		Nuclearity: NuclearityNucleus,
		Text:       edus[0].Text,
	}
	root.Children = append(root.Children, nuc)

	for i := 1; i < len(edus); i++ {
		edu := edus[i]
		lower := strings.ToLower(edu.Text)

		rel := RelElaboration
		nucType := NuclearitySatellite

		switch {
		case strings.Contains(lower, "because") || strings.Contains(lower, "since"):
			rel = RelCause
		case strings.Contains(lower, "but") || strings.Contains(lower, "however") || strings.Contains(lower, "although"):
			rel = RelContrast
		case strings.Contains(lower, "if ") || strings.Contains(lower, "unless"):
			rel = RelCondition
		case strings.Contains(lower, "then") || strings.Contains(lower, "after"):
			rel = RelSequence
			nucType = NuclearityNucleus
		}

		sat := &RSTNode{
			ID:         edu.ID,
			Nuclearity: nucType,
			Relation:   rel,
			Text:       edu.Text,
		}
		root.Children = append(root.Children, sat)
	}

	return root
}
