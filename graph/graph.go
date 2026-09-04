package graph

import (
	"context"

	"github.com/raitucarp/gown"
)

// Edge represents a directed, typed relationship between two nodes.
type Edge struct {
	Source string
	Target string
	Type   string
	Weight float64
}

// Node represents a vertex in the WordNet knowledge graph.
type Node struct {
	ID       string
	Label    string
	Type     string // "synset", "lemma", "sense"
	POS      string
	Gloss    string
	OutEdges []Edge
}

// Graph represents a directed knowledge graph over WordNet.
type Graph struct {
	nodes map[string]*Node
}

// NewGraph constructs an empty graph.
func NewGraph() *Graph {
	return &Graph{
		nodes: make(map[string]*Node),
	}
}

// AddNode registers a node in the graph if not already present.
func (g *Graph) AddNode(node *Node) {
	if _, exists := g.nodes[node.ID]; !exists {
		g.nodes[node.ID] = node
	}
}

// AddEdge registers a directed edge between two existing or newly added nodes.
func (g *Graph) AddEdge(sourceID, targetID, relType string, weight ...float64) {
	w := 1.0
	if len(weight) > 0 {
		w = weight[0]
	}

	src, ok := g.nodes[sourceID]
	if !ok {
		src = &Node{ID: sourceID, Label: sourceID}
		g.nodes[sourceID] = src
	}
	if _, ok := g.nodes[targetID]; !ok {
		g.nodes[targetID] = &Node{ID: targetID, Label: targetID}
	}

	src.OutEdges = append(src.OutEdges, Edge{
		Source: sourceID,
		Target: targetID,
		Type:   relType,
		Weight: w,
	})
}

// NodeCount returns the total number of vertices in the graph.
func (g *Graph) NodeCount() int {
	return len(g.nodes)
}

// EdgeCount returns the total number of edges in the graph.
func (g *Graph) EdgeCount() int {
	total := 0
	for _, n := range g.nodes {
		total += len(n.OutEdges)
	}
	return total
}

// GetNode looks up a node by its ID.
func (g *Graph) GetNode(id string) *Node {
	return g.nodes[id]
}

// BuildSynsetGraph constructs a knowledge graph linking synsets via their semantic relations.
func BuildSynsetGraph(resource *gown.LexicalResource, relTypes ...string) *Graph {
	g := NewGraph()
	relFilter := make(map[string]bool)
	for _, rt := range relTypes {
		relFilter[rt] = true
	}

	for i := range resource.Lexicon.Synsets {
		s := &resource.Lexicon.Synsets[i]
		g.AddNode(&Node{
			ID:    s.ID,
			Label: s.PrimaryDefinition(),
			Type:  "synset",
			POS:   s.PartOfSpeech,
			Gloss: s.PrimaryDefinition(),
		})
	}

	for i := range resource.Lexicon.Synsets {
		s := &resource.Lexicon.Synsets[i]
		for _, rel := range s.SynsetRelations {
			if len(relFilter) > 0 && !relFilter[rel.RelType] {
				continue
			}
			g.AddEdge(s.ID, rel.Target, rel.RelType, 1.0)
		}
	}

	return g
}

// BreadthFirstSearch traverses the graph starting from startID using BFS.
func (g *Graph) BreadthFirstSearch(ctx context.Context, startID string, visit func(node *Node, depth int) bool) {
	start := g.nodes[startID]
	if start == nil {
		return
	}

	visited := make(map[string]bool)
	type qItem struct {
		node  *Node
		depth int
	}

	queue := []qItem{{node: start, depth: 0}}
	visited[start.ID] = true

	for len(queue) > 0 {
		select {
		case <-ctx.Done():
			return
		default:
		}

		curr := queue[0]
		queue = queue[1:]

		if !visit(curr.node, curr.depth) {
			return
		}

		for _, edge := range curr.node.OutEdges {
			target := g.nodes[edge.Target]
			if target != nil && !visited[target.ID] {
				visited[target.ID] = true
				queue = append(queue, qItem{node: target, depth: curr.depth + 1})
			}
		}
	}
}

// DepthFirstSearch traverses the graph starting from startID using DFS.
func (g *Graph) DepthFirstSearch(ctx context.Context, startID string, visit func(node *Node, depth int) bool) {
	start := g.nodes[startID]
	if start == nil {
		return
	}

	visited := make(map[string]bool)
	var dfs func(n *Node, depth int) bool
	dfs = func(n *Node, depth int) bool {
		select {
		case <-ctx.Done():
			return false
		default:
		}

		visited[n.ID] = true
		if !visit(n, depth) {
			return false
		}

		for _, edge := range n.OutEdges {
			target := g.nodes[edge.Target]
			if target != nil && !visited[target.ID] {
				if !dfs(target, depth+1) {
					return false
				}
			}
		}
		return true
	}

	dfs(start, 0)
}

// SemanticNeighborhood extracts the k-hop subgraph around a given node.
func (g *Graph) SemanticNeighborhood(startID string, k int) *Graph {
	subgraph := NewGraph()
	ctx := context.Background()

	g.BreadthFirstSearch(ctx, startID, func(node *Node, depth int) bool {
		if depth > k {
			return true
		}
		subgraph.AddNode(&Node{
			ID:    node.ID,
			Label: node.Label,
			Type:  node.Type,
			POS:   node.POS,
			Gloss: node.Gloss,
		})
		if depth < k {
			for _, edge := range node.OutEdges {
				subgraph.AddEdge(edge.Source, edge.Target, edge.Type, edge.Weight)
			}
		}
		return true
	})

	return subgraph
}
