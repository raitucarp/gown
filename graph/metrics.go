package graph

// InDegree calculates the number of incoming edges for a node.
func (g *Graph) InDegree(nodeID string) int {
	count := 0
	for _, node := range g.nodes {
		for _, edge := range node.OutEdges {
			if edge.Target == nodeID {
				count++
			}
		}
	}
	return count
}

// OutDegree calculates the number of outgoing edges for a node.
func (g *Graph) OutDegree(nodeID string) int {
	if n := g.nodes[nodeID]; n != nil {
		return len(n.OutEdges)
	}
	return 0
}

// DegreeCentrality returns normalized degree centrality for all nodes in the graph.
func (g *Graph) DegreeCentrality() map[string]float64 {
	centrality := make(map[string]float64)
	n := len(g.nodes)
	if n <= 1 {
		for id := range g.nodes {
			centrality[id] = 0.0
		}
		return centrality
	}

	denom := float64(n - 1)
	for id, node := range g.nodes {
		in := g.InDegree(id)
		out := len(node.OutEdges)
		centrality[id] = float64(in+out) / (2.0 * denom)
	}
	return centrality
}

// ConnectedComponents identifies weakly connected components in the graph.
func (g *Graph) ConnectedComponents() [][]string {
	adj := make(map[string][]string)
	for id, node := range g.nodes {
		for _, edge := range node.OutEdges {
			adj[id] = append(adj[id], edge.Target)
			adj[edge.Target] = append(adj[edge.Target], id)
		}
	}

	visited := make(map[string]bool)
	var components [][]string

	for id := range g.nodes {
		if !visited[id] {
			var comp []string
			queue := []string{id}
			visited[id] = true

			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				comp = append(comp, curr)

				for _, neighbor := range adj[curr] {
					if !visited[neighbor] {
						visited[neighbor] = true
						queue = append(queue, neighbor)
					}
				}
			}
			components = append(components, comp)
		}
	}

	return components
}

// HasCycle checks if there is any directed cycle in the graph.
func (g *Graph) HasCycle() bool {
	visited := make(map[string]bool)
	recStack := make(map[string]bool)

	var isCyclic func(curr string) bool
	isCyclic = func(curr string) bool {
		visited[curr] = true
		recStack[curr] = true

		if node := g.nodes[curr]; node != nil {
			for _, edge := range node.OutEdges {
				if !visited[edge.Target] {
					if isCyclic(edge.Target) {
						return true
					}
				} else if recStack[edge.Target] {
					return true
				}
			}
		}

		recStack[curr] = false
		return false
	}

	for id := range g.nodes {
		if !visited[id] {
			if isCyclic(id) {
				return true
			}
		}
	}

	return false
}
