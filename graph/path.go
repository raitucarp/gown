package graph

import (
	"container/heap"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/raitucarp/gown"
)

// ShortestPath finds the minimum-hop path between two node IDs in the graph.
func (g *Graph) ShortestPath(sourceID, targetID string) []string {
	if sourceID == targetID {
		return []string{sourceID}
	}

	queue := []string{sourceID}
	visited := map[string]bool{sourceID: true}
	parent := make(map[string]string)

	found := false
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == targetID {
			found = true
			break
		}

		node := g.nodes[curr]
		if node == nil {
			continue
		}

		for _, edge := range node.OutEdges {
			if !visited[edge.Target] {
				visited[edge.Target] = true
				parent[edge.Target] = curr
				queue = append(queue, edge.Target)
			}
		}
	}

	if !found {
		return nil
	}

	var path []string
	curr := targetID
	for curr != "" {
		path = append(path, curr)
		curr = parent[curr]
	}
	slices.Reverse(path)
	return path
}

type priorityItem struct {
	nodeID string
	dist   float64
	index  int
}

type priorityQueue []*priorityItem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *priorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*priorityItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *priorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

// DijkstraShortestPath computes the weighted shortest path between sourceID and targetID.
func (g *Graph) DijkstraShortestPath(sourceID, targetID string) ([]string, float64) {
	if sourceID == targetID {
		return []string{sourceID}, 0.0
	}

	dist := make(map[string]float64)
	parent := make(map[string]string)
	for id := range g.nodes {
		dist[id] = math.Inf(1)
	}
	dist[sourceID] = 0.0

	pq := make(priorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &priorityItem{nodeID: sourceID, dist: 0.0})

	for pq.Len() > 0 {
		top := heap.Pop(&pq).(*priorityItem)
		u := top.nodeID

		if u == targetID {
			break
		}
		if top.dist > dist[u] {
			continue
		}

		node := g.nodes[u]
		if node == nil {
			continue
		}

		for _, edge := range node.OutEdges {
			alt := dist[u] + edge.Weight
			if alt < dist[edge.Target] {
				dist[edge.Target] = alt
				parent[edge.Target] = u
				heap.Push(&pq, &priorityItem{nodeID: edge.Target, dist: alt})
			}
		}
	}

	if math.IsInf(dist[targetID], 1) {
		return nil, math.Inf(1)
	}

	var path []string
	curr := targetID
	for curr != "" {
		path = append(path, curr)
		curr = parent[curr]
	}
	slices.Reverse(path)
	return path, dist[targetID]
}

// HypernymAncestors finds all hypernym ancestors of a synset along with their shortest depths.
func HypernymAncestors(res *gown.LexicalResource, synset *gown.Synset) map[string]int {
	ancestors := make(map[string]int)
	if synset == nil {
		return ancestors
	}

	ancestors[synset.ID] = 0
	queue := []string{synset.ID}
	depths := map[string]int{synset.ID: 0}
	synsetsById := res.SynsetsById()

	for len(queue) > 0 {
		currID := queue[0]
		queue = queue[1:]
		currDepth := depths[currID]

		curr := synsetsById[currID]
		if curr == nil {
			continue
		}

		for _, rel := range curr.SynsetRelations {
			if rel.RelType == string(gown.SynsetRelationTypeHypernym) ||
				rel.RelType == string(gown.SynsetRelationTypeInstanceHypernym) {
				if d, ok := depths[rel.Target]; !ok || currDepth+1 < d {
					depths[rel.Target] = currDepth + 1
					ancestors[rel.Target] = currDepth + 1
					queue = append(queue, rel.Target)
				}
			}
		}
	}

	return ancestors
}

// LowestCommonHypernym finds the lowest common hypernym (subsumer) between two synsets.
// Returns the LCS synset and its shortest distance (depth from root) or nil if no common ancestor.
func LowestCommonHypernym(res *gown.LexicalResource, s1, s2 *gown.Synset) (*gown.Synset, int) {
	if s1 == nil || s2 == nil {
		return nil, 0
	}

	if s1.ID == s2.ID {
		return s1, synsetDepth(res, s1)
	}

	anc1 := HypernymAncestors(res, s1)
	anc2 := HypernymAncestors(res, s2)

	synsetsById := res.SynsetsById()
	var bestSynset *gown.Synset
	bestDepth := -1

	for id := range anc1 {
		if _, exists := anc2[id]; exists {
			s := synsetsById[id]
			if s != nil {
				depth := synsetDepth(res, s)
				if depth > bestDepth {
					bestDepth = depth
					bestSynset = s
				}
			}
		}
	}

	return bestSynset, bestDepth
}

// SynsetDepth returns the minimum path length from the root (e.g. entity.n.01) to this synset.
func SynsetDepth(res *gown.LexicalResource, synset *gown.Synset) int {
	return synsetDepth(res, synset)
}

func synsetDepth(res *gown.LexicalResource, synset *gown.Synset) int {
	if synset == nil {
		return 0
	}
	anc := HypernymAncestors(res, synset)
	maxDepth := 0
	for _, d := range anc {
		if d > maxDepth {
			maxDepth = d
		}
	}
	return maxDepth
}

// SemanticPath discovers a path between two words through their hypernym hierarchy.
// Example: "dog" -> "canine" -> "carnivore" -> "placental mammal" -> "mammal" -> "vertebrate" -> "animal".
func SemanticPath(res *gown.LexicalResource, word1, word2 string) ([]string, error) {
	entries1 := res.Lookup(word1)
	entries2 := res.Lookup(word2)
	if len(entries1) == 0 {
		return nil, fmt.Errorf("word not found: %s", word1)
	}
	if len(entries2) == 0 {
		return nil, fmt.Errorf("word not found: %s", word2)
	}

	// Build a bidirectional hypernym graph for path finding
	g := NewGraph()
	synsetsById := res.SynsetsById()

	for _, s := range synsetsById {
		g.AddNode(&Node{ID: s.ID, Label: s.PrimaryDefinition()})
		for _, rel := range s.SynsetRelations {
			if rel.RelType == string(gown.SynsetRelationTypeHypernym) ||
				rel.RelType == string(gown.SynsetRelationTypeInstanceHypernym) {
				g.AddEdge(s.ID, rel.Target, "hypernym", 1.0)
				g.AddEdge(rel.Target, s.ID, "hyponym", 1.0)
			}
		}
	}

	var bestPath []string
	lexicalsById := res.LexicalsById()

	for _, e1 := range entries1 {
		for _, s1 := range e1.Synsets() {
			if s1 == nil {
				continue
			}
			for _, e2 := range entries2 {
				for _, s2 := range e2.Synsets() {
					if s2 == nil {
						continue
					}
					p := g.ShortestPath(s1.ID, s2.ID)
					if len(p) > 0 && (len(bestPath) == 0 || len(p) < len(bestPath)) {
						bestPath = p
					}
				}
			}
		}
	}

	if len(bestPath) == 0 {
		return nil, fmt.Errorf("no semantic path between %s and %s", word1, word2)
	}

	// Convert synset IDs into readable member words
	var readablePath []string
	for _, synID := range bestPath {
		s := synsetsById[synID]
		if s != nil {
			var words []string
			for _, m := range s.Members {
				if le, ok := lexicalsById[m]; ok {
					words = append(words, le.Lemma.WrittenForm)
				}
			}
			if len(words) > 0 {
				readablePath = append(readablePath, strings.Join(words, "/"))
			} else {
				readablePath = append(readablePath, synID)
			}
		} else {
			readablePath = append(readablePath, synID)
		}
	}

	return readablePath, nil
}
