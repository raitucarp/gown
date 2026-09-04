package graph_test

import (
	"context"
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/graph"
)

func ExampleSemanticPath() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	path, err := graph.SemanticPath(res, "dog", "animal")
	if err != nil {
		log.Fatalf("failed to find semantic path: %v", err)
	}

	// The semantic path traces the hypernym taxonomy up to animal
	fmt.Printf("Found path from dog to animal: %t\n", len(path) > 0)
	// Output:
	// Found path from dog to animal: true
}

func ExampleLowestCommonHypernym() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	dogEntries := res.LookupExact("dog", gown.NounPos)
	catEntries := res.LookupExact("cat", gown.NounPos)
	if len(dogEntries) == 0 || len(catEntries) == 0 {
		return
	}

	sDog := dogEntries[0].Synsets()[0]
	sCat := catEntries[0].Synsets()[0]

	lcs, depth := graph.LowestCommonHypernym(res, sDog, sCat)
	fmt.Printf("LCS found: %t, depth >= 0: %t\n", lcs != nil, depth >= 0)
	// Output:
	// LCS found: true, depth >= 0: true
}

func ExampleGraph_BreadthFirstSearch() {
	g := graph.NewGraph()
	g.AddEdge("root", "child1", "branch")
	g.AddEdge("root", "child2", "branch")
	g.AddEdge("child1", "leaf", "branch")

	var visited []string
	g.BreadthFirstSearch(context.Background(), "root", func(node *graph.Node, depth int) bool {
		visited = append(visited, node.ID)
		return true
	})

	fmt.Printf("BFS visited root first: %t, total visited: %d\n", visited[0] == "root", len(visited))
	// Output:
	// BFS visited root first: true, total visited: 4
}
