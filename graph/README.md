# Package graph

Package `graph` provides a graph-theoretical representation and traversal engine for WordNet lexical and conceptual networks, including Breadth-First Search (BFS), Depth-First Search (DFS), shortest path algorithms, weighted Dijkstra search, Lowest Common Hypernym (LCS), and centrality calculations.

## Overview

WordNet is inherently a directed property graph where nodes represent synsets or lexical entries and edges represent typed semantic relations. Package `graph` provides:

1. **Graph Construction**: Transforming WordNet synsets and relations into an in-memory graph structure with edge weights.
2. **Path Finding**: Unweighted BFS shortest path and weighted Dijkstra shortest path between arbitrary synsets.
3. **Lowest Common Subsumer (LCS)**: Finding the lowest common ancestor in the hypernym taxonomy, fundamental for semantic distance calculation.
4. **Neighborhoods & Subgraphs**: Extracting $k$-hop ego networks around a target synset.
5. **Structural Metrics**: Node degree centrality (in-degree, out-degree) and cycle detection.

## Key Types and Functions

### Graph and Builder

```go
type Graph struct {
    // Directed graph with adjacency lists and edge attributes
}

type Builder struct {
    res *gown.LexicalResource
}
```

- `NewBuilder(res *gown.LexicalResource) *Builder`: Creates a graph builder.
- `(b *Builder) BuildSynsetGraph(relations ...gown.SynsetRelationType) *Graph`: Constructs a directed graph filtered by specific relation types (e.g. `gown.SynsetRelationHypernym`).

### Algorithms

- `(g *Graph) ShortestPath(sourceID, targetID string) ([]string, bool)`: Computes the unweighted shortest path using BFS.
- `(g *Graph) Dijkstra(sourceID, targetID string) ([]string, float64, bool)`: Computes the shortest weighted path.
- `(g *Graph) LowestCommonHypernym(id1, id2 string) (string, int, error)`: Identifies the lowest common ancestor node and its depth.
- `(g *Graph) Degree(id string) (inDegree, outDegree int)`: Computes directional node connectivity.
- `(g *Graph) HasCycle() bool`: Checks if the graph contains directed cycles.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/graph"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    builder := graph.NewBuilder(res)
    g := builder.BuildSynsetGraph(gown.SynsetRelationHypernym)

    // Lookup synsets
    dogEntries := res.LookupNoun("dog")
    catEntries := res.LookupNoun("cat")
    dogSynID := dogEntries[0].Senses[0].Synset
    catSynID := catEntries[0].Senses[0].Synset

    // Find Lowest Common Subsumer
    lcs, depth, err := g.LowestCommonHypernym(dogSynID, catSynID)
    if err == nil {
        fmt.Printf("Lowest Common Hypernym: %s (Depth: %d)\n", lcs, depth)
    }

    // Shortest path
    path, found := g.ShortestPath(dogSynID, catSynID)
    if found {
        fmt.Println("Path:", path)
    }
}
```
