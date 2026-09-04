# Package generative

Package `generative` implements tools for formal generative linguistics and Chomskyan Context-Free Grammars (CFG), including phrase structure parsing and generation, lexical insertion from WordNet, Penn Treebank bracket formatting, and syntactic feature structure unification.

## Overview

Generative grammar, pioneered by Noam Chomsky, models the human linguistic capacity as a formal rule-based computational system. Package `generative` provides:

1. **Context-Free Grammars (CFG)**: Parsing and representation of rewrite rules ($A \to \alpha$).
2. **Lexical Insertion**: Combining formal grammatical rules with WordNet lexical lookups to generate well-formed sentences.
3. **Penn Treebank Notation**: Parsing and rendering bracketed phrase structure trees.
4. **Feature Structures & Unification**: Attribute-value matrices representing agreement constraints (Person, Number, Gender, Case, Tense) with first-order unification.
5. **Subcategorization Frames**: Validation of verbal transitivity and subcategorization frames against syntactic dependents.

## Key Types and Functions

### Context-Free Grammar

```go
type ProductionRule struct {
    LHS string
    RHS []string
}

type CFG struct {
    Rules []ProductionRule
}
```

- `NewCFG(grammarText string) (*CFG, error)`: Compiles a string of production rules into a formal grammar.
- `GenerateSentence(cfg *CFG, res *gown.LexicalResource, startSymbol string) (*syntax.SyntaxNode, error)`: Recursively generates a constituent tree with lexical items inserted from WordNet.

### Feature Structure Unification

```go
type FeatureStructure map[string]string
```

- `UnifyFeatures(f1, f2 FeatureStructure) (FeatureStructure, bool)`: Performs feature unification. Returns the unified feature structure if compatible, or `false` if a feature clash occurs (e.g., singular vs plural).

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/generative"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    // 1. Define Context-Free Grammar
    grammar, err := generative.NewCFG(`
        S  -> NP VP
        NP -> Det N
        VP -> V NP
    `)
    if err != nil {
        log.Fatal(err)
    }

    // 2. Generate sentence with WordNet lexical insertion
    tree, err := generative.GenerateSentence(grammar, res, "S")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(tree.PennString())

    // 3. Feature Unification
    f1 := generative.FeatureStructure{"POS": "N", "NUM": "sg", "PERS": "3"}
    f2 := generative.FeatureStructure{"NUM": "sg", "CASE": "nom"}
    unified, ok := generative.UnifyFeatures(f1, f2)
    if ok {
        fmt.Println("Unified features:", unified)
    }
}
```
