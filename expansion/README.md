# Package expansion

Package `expansion` provides algorithms for recursive lexical and semantic tree expansion across WordNet relational links and definition glosses, with cycle detection and depth controls.

## Overview

Taxonomic exploration and lexical analysis often require expanding words or synsets recursively into hierarchical trees. Package `expansion` provides:

1. **Relational Expansion Trees**: Recursively traversing pointer links (such as `hypernym`, `hyponym`, `mero_part`, or `derivation`) to construct explicit constituent tree hierarchies.
2. **Recursive Definition Expansion**: Tokenizing definition glosses, filtering stopwords, looking up constituent content words, and recursively building definition dependency trees to analyze semantic depth.
3. **Cycle Detection & Depth Limiting**: Guarding against ontological loops and limiting traversal depth and total node counts.
4. **Visual Tree Rendering**: Formatting tree structures as indented ASCII diagrams.

## Key Types and Functions

### Expansion Trees

```go
type TreeNode struct {
    ID       string
    Label    string
    Depth    int
    Children []*TreeNode
}

type TreeOptions struct {
    MaxDepth  int
    MaxNodes  int
    StopWords []string
}
```

- `DefaultTreeOptions() TreeOptions`: Provides default limits (`MaxDepth: 5`, `MaxNodes: 100`).
- `ExpandRelationTree(res *gown.LexicalResource, synsetID string, relType gown.SynsetRelationType, opts ...TreeOptions) (*TreeNode, error)`: Builds a recursive hierarchy following the specified relation type.
- `ExpandDefinitionTree(res *gown.LexicalResource, word string, depth int) (*TreeNode, error)`: Expands a word through its definition content words.
- `(n *TreeNode) Render() string`: Renders the tree as an indented text diagram.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/expansion"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    dogSynID := res.LookupNoun("dog")[0].Senses[0].Synset

    // Expand hypernym tree up to depth 4
    opts := expansion.DefaultTreeOptions()
    opts.MaxDepth = 4
    tree, err := expansion.ExpandRelationTree(res, dogSynID, gown.SynsetRelationHypernym, opts)
    if err == nil {
        fmt.Println(tree.Render())
    }
}
```
