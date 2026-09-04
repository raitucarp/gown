# Package syntax

Package `syntax` provides data structures, algorithms, and utilities for modeling and manipulating syntactic constituent trees, determining lexical heads, extracting grammatical relations, and performing syntactic transformations.

## Overview

In theoretical and computational linguistics, syntax concerns the rules and principles that govern sentence structure and word order. The `syntax` package provides:

1. **Syntactic Category Classification**: Categorization of words into parts of speech and phrase types (Noun Phrase, Verb Phrase, Prepositional Phrase, Adjective Phrase, etc.).
2. **Constituency Trees**: Hierarchical representation of sentences as nested phrase structures with bidirectional tree navigation (parent, children, siblings, ancestors).
3. **Lexical Head Detection**: Collins-style head percolation rules to identify head words and head children within phrases.
4. **Grammatical Relations**: Identification of grammatical roles (Subject, Direct Object, Indirect Object, Oblique/Adjunct) from constituent geometry.
5. **Structural Transformations**: Rule-based rewriting including active-to-passive voice conversion, auxiliary inversion for question formation, and clausal negation.
6. **Syntactic Tree Comparison**: Tree edit distance, subtree isomorphism, and production rule extraction.

## Key Types and Functions

### SyntaxNode

```go
type SyntaxNode struct {
    Category  string
    Word      string
    Children  []*SyntaxNode
    Parent    *SyntaxNode
}
```

A node in a constituent parse tree. Terminal nodes contain `Word` and a part-of-speech category (e.g. `N`, `V`, `Det`). Non-terminal nodes contain phrasal categories (e.g. `S`, `NP`, `VP`) and a slice of child nodes.

### Tree Construction and Navigation

- `NewNode(category string, word ...string) *SyntaxNode`: Creates a new node.
- `(n *SyntaxNode) AddChild(child *SyntaxNode)`: Adds a child node and sets its parent pointer.
- `(n *SyntaxNode) FindHeadChild() *SyntaxNode`: Determines the head child of a phrase using Collins head percolation rules.
- `(n *SyntaxNode) HeadWord() string`: Returns the lexical head of the constituent.
- `(n *SyntaxNode) PennString() string`: Formats the tree in Penn Treebank bracketed notation (e.g. `(S (NP (N dogs)) (VP (V bark)))`).
- `(n *SyntaxNode) Leaves() []*SyntaxNode`: Returns all terminal leaf nodes in left-to-right order.

### Grammatical Relations

```go
type GrammaticalRelation struct {
    Type     RelationType // RelSubject, RelDirectObject, RelIndirectObject, RelAdjunct
    Head     *SyntaxNode
    Dependent *SyntaxNode
}
```

- `ExtractRelations(root *SyntaxNode) []GrammaticalRelation`: Traverses a parse tree and extracts grammatical dependencies between heads and dependents.

### Transformations

- `PassiveTransform(root *SyntaxNode) (*SyntaxNode, error)`: Transforms an active-voice constituent tree into passive voice (e.g. "The dog chased the cat" -> "The cat was chased by the dog").
- `NegateSentence(root *SyntaxNode) (*SyntaxNode, error)`: Negates a declarative sentence by inserting appropriate auxiliary verbs and negative particles.
- `InvertQuestion(root *SyntaxNode) (*SyntaxNode, error)`: Converts a declarative sentence into a Yes/No question via auxiliary-subject inversion.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/syntax"
)

func main() {
    // Construct constituent tree: (S (NP (Det The) (N dog)) (VP (V barked)))
    s := syntax.NewNode("S")
    np := syntax.NewNode("NP")
    np.AddChild(syntax.NewNode("Det", "The"))
    np.AddChild(syntax.NewNode("N", "dog"))

    vp := syntax.NewNode("VP")
    vp.AddChild(syntax.NewNode("V", "barked"))

    s.AddChild(np)
    s.AddChild(vp)

    fmt.Println(s.PennString())
    fmt.Printf("Subject head: %s\n", np.HeadWord())
}
```
