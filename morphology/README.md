# Package morphology

Package `morphology` provides data structures and algorithms for morphological analysis, including inflection classification, compound word decomposition, and derivational family discovery using WordNet relations and affix rules.

## Overview

Morphology is the branch of linguistics that studies the internal structure of words and their formation. Package `morphology` provides:

1. **Inflection Analysis**: Classifying words by morphological inflections (Plural, Past Tense, Progressive/Participle, Comparative, Superlative, Possessive).
2. **Compound Word Decomposition**: Decomposing closed and hyphenated compounds (e.g. "sunflower", "keyboard", "fire-engine") into constituent root morphemes validated against WordNet.
3. **Derivational Families**: Tracing word families through derivational morphology affixes (e.g. *act* -> *actor*, *action*, *activate*, *react*, *inactive*) and WordNet derivation links.

## Key Types and Functions

### Inflection

```go
type InflectionType string

const (
    InflectionPlural       InflectionType = "Plural"
    InflectionPast         InflectionType = "Past"
    InflectionProgressive  InflectionType = "Progressive"
    InflectionComparative  InflectionType = "Comparative"
    InflectionSuperlative  InflectionType = "Superlative"
    InflectionBase         InflectionType = "Base"
)
```

- `DetectInflection(word string) (InflectionType, string)`: Identifies the inflection type and extracts the inferred base lemma.

### Compound Decomposition

```go
type CompoundDecomposition struct {
    Word       string
    Morphemes  []string
    IsCompound bool
}
```

- `DecomposeCompound(res *gown.LexicalResource, word string) CompoundDecomposition`: Splits compound words into valid WordNet base morphemes.

### Derivational Families

```go
type DerivationalFamily struct {
    Root    string
    Members []string
}
```

- `FindDerivationalFamily(res *gown.LexicalResource, word string) DerivationalFamily`: Collects derivationally related lemmas using WordNet pointer links and affix rules.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/morphology"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    // 1. Inflection detection
    inf, base := morphology.DetectInflection("running")
    fmt.Printf("running -> Inflection: %s, Base: %s\n", inf, base)

    // 2. Compound decomposition
    decomp := morphology.DecomposeCompound(res, "sunflower")
    fmt.Printf("sunflower compound: %v (Morphemes: %v)\n", decomp.IsCompound, decomp.Morphemes)

    // 3. Derivational family
    family := morphology.FindDerivationalFamily(res, "act")
    fmt.Printf("Root: %s, Derivations: %v\n", family.Root, family.Members)
}
```
