package graph_test

import (
	"context"
	"math"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/graph"
)

func TestGraphCoreAndTraversals(t *testing.T) {
	g := graph.NewGraph()
	g.AddNode(&graph.Node{ID: "A", Label: "Node A", Type: "synset", POS: "n", Gloss: "first"})
	g.AddEdge("A", "B", "hypernym", 2.0)
	g.AddEdge("B", "C", "hypernym", 1.5)
	g.AddEdge("B", "D", "hypernym", 3.0)
	g.AddEdge("C", "E", "hypernym", 1.0)

	// NodeCount & EdgeCount
	if g.NodeCount() != 5 {
		t.Errorf("Expected 5 nodes, got %d", g.NodeCount())
	}
	if g.EdgeCount() != 4 {
		t.Errorf("Expected 4 edges, got %d", g.EdgeCount())
	}

	// GetNode
	if nA := g.GetNode("A"); nA == nil || nA.Label != "Node A" {
		t.Errorf("GetNode(A) failed or unexpected: %v", nA)
	}
	if nMissing := g.GetNode("Z"); nMissing != nil {
		t.Errorf("GetNode(Z) expected nil, got %v", nMissing)
	}

	// Adding existing node does not duplicate or overwrite
	g.AddNode(&graph.Node{ID: "A", Label: "Duplicate A"})
	if g.GetNode("A").Label != "Node A" {
		t.Errorf("Duplicate AddNode overwrote existing node")
	}

	// BFS traversal from A
	var bfsVisited []string
	g.BreadthFirstSearch(context.Background(), "A", func(n *graph.Node, depth int) bool {
		bfsVisited = append(bfsVisited, n.ID)
		return true
	})
	if len(bfsVisited) != 5 {
		t.Errorf("BFS expected 5 visited, got %d", len(bfsVisited))
	}

	// BFS early termination
	var bfsEarly []string
	g.BreadthFirstSearch(context.Background(), "A", func(n *graph.Node, depth int) bool {
		bfsEarly = append(bfsEarly, n.ID)
		return len(bfsEarly) < 2 // stop after 2
	})
	if len(bfsEarly) != 2 {
		t.Errorf("BFS early termination expected 2, got %d", len(bfsEarly))
	}

	// BFS missing start
	g.BreadthFirstSearch(context.Background(), "MISSING", func(n *graph.Node, depth int) bool {
		t.Errorf("BFS on missing start should not invoke callback")
		return true
	})

	// BFS context canceled
	ctxCanceled, cancel := context.WithCancel(context.Background())
	cancel()
	g.BreadthFirstSearch(ctxCanceled, "A", func(n *graph.Node, depth int) bool {
		return true
	})

	// DFS traversal from A
	var dfsVisited []string
	g.DepthFirstSearch(context.Background(), "A", func(n *graph.Node, depth int) bool {
		dfsVisited = append(dfsVisited, n.ID)
		return true
	})
	if len(dfsVisited) != 5 {
		t.Errorf("DFS expected 5 visited, got %d", len(dfsVisited))
	}

	// DFS early termination
	var dfsEarly []string
	g.DepthFirstSearch(context.Background(), "A", func(n *graph.Node, depth int) bool {
		dfsEarly = append(dfsEarly, n.ID)
		return len(dfsEarly) < 3
	})
	if len(dfsEarly) != 3 {
		t.Errorf("DFS early stop expected 3, got %d", len(dfsEarly))
	}

	// DFS missing start
	g.DepthFirstSearch(context.Background(), "MISSING", func(n *graph.Node, depth int) bool {
		t.Errorf("DFS on missing start should not invoke callback")
		return true
	})

	// DFS context canceled
	ctxCanceled2, cancel2 := context.WithCancel(context.Background())
	cancel2()
	g.DepthFirstSearch(ctxCanceled2, "A", func(n *graph.Node, depth int) bool {
		return true
	})

	// SemanticNeighborhood (k=1 from B)
	subg := g.SemanticNeighborhood("B", 1)
	if subg == nil || subg.NodeCount() < 3 {
		t.Errorf("SemanticNeighborhood(B, 1) expected at least 3 nodes, got %d", subg.NodeCount())
	}
	// SemanticNeighborhood with depth filter
	subg2 := g.SemanticNeighborhood("A", 0)
	if subg2.NodeCount() != 1 {
		t.Errorf("SemanticNeighborhood(A, 0) expected 1 node, got %d", subg2.NodeCount())
	}
}

func TestGraphMetrics(t *testing.T) {
	// Single-node graph edge case
	gSingle := graph.NewGraph()
	gSingle.AddNode(&graph.Node{ID: "isolated"})
	dcSingle := gSingle.DegreeCentrality()
	if dcSingle["isolated"] != 0.0 {
		t.Errorf("DegreeCentrality single node expected 0.0, got %.2f", dcSingle["isolated"])
	}

	g := graph.NewGraph()
	g.AddEdge("A", "B", "rel", 1.0)
	g.AddEdge("B", "C", "rel", 1.0)
	g.AddEdge("C", "D", "rel", 1.0)
	g.AddEdge("A", "C", "rel", 1.0)

	// InDegree and OutDegree
	if inB := g.InDegree("B"); inB != 1 {
		t.Errorf("InDegree(B) expected 1, got %d", inB)
	}
	if inC := g.InDegree("C"); inC != 2 {
		t.Errorf("InDegree(C) expected 2, got %d", inC)
	}
	if outA := g.OutDegree("A"); outA != 2 {
		t.Errorf("OutDegree(A) expected 2, got %d", outA)
	}
	if outMissing := g.OutDegree("MISSING"); outMissing != 0 {
		t.Errorf("OutDegree(MISSING) expected 0, got %d", outMissing)
	}

	// Degree Centrality
	dc := g.DegreeCentrality()
	if len(dc) != 4 {
		t.Errorf("Expected centrality for 4 nodes, got %d", len(dc))
	}
	if dc["C"] <= 0.0 {
		t.Errorf("Expected positive centrality for C, got %.4f", dc["C"])
	}

	// ConnectedComponents
	comps := g.ConnectedComponents()
	if len(comps) != 1 {
		t.Errorf("Expected 1 connected component, got %d", len(comps))
	}

	// Add disconnected component
	g.AddEdge("X", "Y", "rel", 1.0)
	comps2 := g.ConnectedComponents()
	if len(comps2) != 2 {
		t.Errorf("Expected 2 connected components, got %d", len(comps2))
	}

	// Cycle Detection
	if g.HasCycle() {
		t.Errorf("Graph without back-edges should not have cycle")
	}

	g.AddEdge("D", "A", "back")
	if !g.HasCycle() {
		t.Errorf("Graph with back-edge D->A should have cycle")
	}
}

func TestGraphShortestPaths(t *testing.T) {
	g := graph.NewGraph()
	g.AddEdge("A", "B", "link", 1.0)
	g.AddEdge("B", "C", "link", 2.0)
	g.AddEdge("A", "C", "link", 5.0)
	g.AddEdge("C", "D", "link", 1.0)
	g.AddNode(&graph.Node{ID: "isolated"})

	// Unweighted ShortestPath
	// Same source and target
	selfPath := g.ShortestPath("A", "A")
	if len(selfPath) != 1 || selfPath[0] != "A" {
		t.Errorf("ShortestPath(A, A) expected [A], got %v", selfPath)
	}
	// Direct hop A -> C (1 hop vs 2 hops via B)
	pathAC := g.ShortestPath("A", "C")
	if len(pathAC) != 2 || pathAC[0] != "A" || pathAC[1] != "C" {
		t.Errorf("ShortestPath(A, C) expected [A, C], got %v", pathAC)
	}
	// Path to isolated
	unreachable := g.ShortestPath("A", "isolated")
	if unreachable != nil {
		t.Errorf("ShortestPath to isolated node should be nil, got %v", unreachable)
	}

	// Dijkstra ShortestPath (Weighted)
	// Same source and target
	dSelf, dSelfCost := g.DijkstraShortestPath("A", "A")
	if len(dSelf) != 1 || dSelfCost != 0.0 {
		t.Errorf("Dijkstra(A, A) expected [A] with cost 0, got %v, cost %.2f", dSelf, dSelfCost)
	}

	// Weighted path A -> C: A->B (1.0) + B->C (2.0) = 3.0, vs direct A->C (5.0)
	dPath, dCost := g.DijkstraShortestPath("A", "C")
	if len(dPath) != 3 || dCost != 3.0 {
		t.Errorf("Dijkstra(A, C) expected [A, B, C] with cost 3.0, got %v, cost %.2f", dPath, dCost)
	}

	// Dijkstra unreachable
	dUnreach, dUnreachCost := g.DijkstraShortestPath("A", "isolated")
	if dUnreach != nil || !math.IsInf(dUnreachCost, 1) {
		t.Errorf("Dijkstra to isolated expected nil with +Inf cost, got %v, cost %.2f", dUnreach, dUnreachCost)
	}
}

func TestBuildSynsetGraphAndSemanticPath(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// BuildSynsetGraph with relation filter
	sg := graph.BuildSynsetGraph(res, string(gown.SynsetRelationTypeHypernym))
	if sg == nil || sg.NodeCount() == 0 {
		t.Fatalf("BuildSynsetGraph returned empty graph")
	}
	if sg.EdgeCount() == 0 {
		t.Errorf("BuildSynsetGraph expected hypernym edges, got 0")
	}

	// SemanticPath tests
	// 1. Existing path: dog to animal
	path, err := graph.SemanticPath(res, "dog", "animal")
	if err != nil {
		t.Fatalf("SemanticPath(dog, animal) error: %v", err)
	}
	if len(path) == 0 {
		t.Errorf("SemanticPath(dog, animal) returned empty path")
	}

	// 2. Missing word1 error
	_, errMissing1 := graph.SemanticPath(res, "nonexistentwordxyz123", "animal")
	if errMissing1 == nil {
		t.Errorf("SemanticPath with missing word1 should error")
	}

	// 3. Missing word2 error
	_, errMissing2 := graph.SemanticPath(res, "dog", "nonexistentwordxyz123")
	if errMissing2 == nil {
		t.Errorf("SemanticPath with missing word2 should error")
	}

	// LowestCommonHypernym edge cases
	dogEntries := res.LookupExact("dog", gown.NounPos)
	catEntries := res.LookupExact("cat", gown.NounPos)
	if len(dogEntries) > 0 && len(catEntries) > 0 {
		sDog := dogEntries[0].Synsets()[0]
		sCat := catEntries[0].Synsets()[0]

		// Nil inputs
		if lcs, _ := graph.LowestCommonHypernym(res, nil, sCat); lcs != nil {
			t.Errorf("LCS(nil, sCat) should be nil")
		}
		if lcs, _ := graph.LowestCommonHypernym(res, sDog, nil); lcs != nil {
			t.Errorf("LCS(sDog, nil) should be nil")
		}

		// Same synset
		if lcs, depth := graph.LowestCommonHypernym(res, sDog, sDog); lcs == nil || lcs.ID != sDog.ID || depth <= 0 {
			t.Errorf("LCS(sDog, sDog) should return sDog with positive depth")
		}

		// Find animal sense of cat to check LCS with dog
		var sCatAnimal *gown.Synset
		for _, e := range catEntries {
			for _, s := range e.Synsets() {
				if s.Lexfile == "noun.animal" {
					sCatAnimal = s
					break
				}
			}
			if sCatAnimal != nil {
				break
			}
		}
		if sCatAnimal == nil {
			sCatAnimal = sCat
		}

		// LowestCommonHypernym between dog and animal cat
		lcs, depth := graph.LowestCommonHypernym(res, sDog, sCatAnimal)
		if lcs == nil || depth <= 0 {
			t.Errorf("Expected common hypernym for dog and animal cat, got nil or depth <= 0")
		}

		// SynsetDepth
		if d := graph.SynsetDepth(res, nil); d != 0 {
			t.Errorf("SynsetDepth(nil) expected 0, got %d", d)
		}
		if d := graph.SynsetDepth(res, sDog); d <= 0 {
			t.Errorf("SynsetDepth(sDog) expected > 0, got %d", d)
		}

		// HypernymAncestors nil
		ancNil := graph.HypernymAncestors(res, nil)
		if len(ancNil) != 0 {
			t.Errorf("HypernymAncestors(nil) expected empty map, got %d items", len(ancNil))
		}
		ancDog := graph.HypernymAncestors(res, sDog)
		if len(ancDog) <= 1 {
			t.Errorf("HypernymAncestors(sDog) expected multiple ancestors, got %d", len(ancDog))
		}
	}
}
