package discourse

// DiscourseNodeType distinguishes entities in the discourse graph.
type DiscourseNodeType string

const (
	NodeSentence   DiscourseNodeType = "sentence"
	NodeEntity     DiscourseNodeType = "entity"
	NodeTopic      DiscourseNodeType = "topic"
)

// DiscourseNode represents a vertex in the discourse graph.
type DiscourseNode struct {
	ID    string            `json:"id"`
	Type  DiscourseNodeType `json:"type"`
	Label string            `json:"label"`
}

// DiscourseEdge represents a typed relation between discourse elements.
type DiscourseEdge struct {
	Source   string `json:"source"`
	Target   string `json:"target"`
	Relation string `json:"relation"` // "next_sentence", "mentions_entity", "topic_focus"
}

// DiscourseGraph models the inter-sentential and conceptual structure of a discourse.
type DiscourseGraph struct {
	Nodes map[string]DiscourseNode `json:"nodes"`
	Edges []DiscourseEdge          `json:"edges"`
}

// NewDiscourseGraph constructs an empty discourse graph.
func NewDiscourseGraph() *DiscourseGraph {
	return &DiscourseGraph{
		Nodes: make(map[string]DiscourseNode),
	}
}

// AddNode registers a discourse node.
func (dg *DiscourseGraph) AddNode(id string, nType DiscourseNodeType, label string) {
	dg.Nodes[id] = DiscourseNode{ID: id, Type: nType, Label: label}
}

// AddEdge registers a relation between two discourse nodes.
func (dg *DiscourseGraph) AddEdge(source, target, relation string) {
	dg.Edges = append(dg.Edges, DiscourseEdge{
		Source:   source,
		Target:   target,
		Relation: relation,
	})
}
