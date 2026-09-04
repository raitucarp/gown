# Package semantics

Package `semantics` provides algorithms and abstractions for computational lexical semantics, including Word Sense Disambiguation (WSD), polysemy and homonymy measurement, thematic role assignment, selectional restriction validation, and logical inference.

## Overview

Semantics studies meaning in language. The `semantics` package provides operations to:

1. **Word Sense Disambiguation (WSD)**: Resolve polysemous words to their most appropriate WordNet synset given sentential context using Simplified Lesk and Extended Lesk algorithms.
2. **Lexical Ambiguity & Entropy**: Calculate polysemy distributions, Shannon entropy of word senses, and distinguish homophones and homographs.
3. **Lexical Chains**: Form cohesive chains across discourse by linking terms with hypernym, hyponym, and sibling synset connections.
4. **Thematic Roles & Predicate-Argument Structure (PAS)**: Map clauses into semantic case roles: Agent, Patient, Experiencer, Stimulus, Instrument, Theme, Goal, and Location.
5. **Selectional Restrictions**: Validate semantic features (e.g. `[+animate]`, `[+food]`, `[+human]`) expected by verbs against noun arguments to detect semantic anomalies.
6. **Inference & First-Order Model Checking**: Represent predicates and entities in first-order logic, evaluate truth against a finite model, and test for semantic entailment and contradiction.

## Key Types and Functions

### Word Sense Disambiguation

- `SimplifiedLesk(res *gown.LexicalResource, word string, sentence string) (*gown.Synset, float64)`: Matches definition tokens of candidate synsets against the context sentence to identify the best sense.
- `ExtendedLesk(res *gown.LexicalResource, word string, sentence string) (*gown.Synset, float64)`: Expands gloss comparison to include hypernym, hyponym, and meronym definitions.

### Polysemy and Ambiguity

```go
type PolysemyAnalysis struct {
    Word       string
    SenseCount int
    Entropy    float64
    Senses     []SenseSummary
}
```

- `AnalyzePolysemy(res *gown.LexicalResource, word string) PolysemyAnalysis`: Computes sense count and Shannon entropy over sense distributions.
- `DetectHomonyms(res *gown.LexicalResource, word string) []HomonymGroup`: Groups senses into distinct etymological or semantic clusters.

### Thematic Roles & Selectional Restrictions

- `AssignThematicRoles(subject, verb, object string) PredicateArgumentStructure`: Assigns Agent, Theme, Patient, or Experiencer roles based on verb classes.
- `CheckSelectionalRestrictions(res *gown.LexicalResource, verb, role, noun string) (bool, string)`: Validates whether a noun satisfies the semantic features required by the predicate.

### Logic and Inference

- `CheckEntailment(res *gown.LexicalResource, premiseSynsetID, hypothesisSynsetID string) bool`: Determines if a premise concept semantically entails a hypothesis concept via WordNet hypernymy.
- `CheckContradiction(res *gown.LexicalResource, term1, term2 string) bool`: Checks if two concepts are antonyms or contradictory opposites.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/semantics"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    // Disambiguate "bank" in financial context
    synset, score := semantics.SimplifiedLesk(res, "bank", "He deposited cash into his bank account.")
    if synset != nil {
        fmt.Printf("Selected Synset: %s (Confidence: %.2f)\n", synset.ID, score)
        fmt.Printf("Definition: %s\n", synset.PrimaryDefinition())
    }

    // Polysemy analysis
    poly := semantics.AnalyzePolysemy(res, "head")
    fmt.Printf("Senses: %d, Entropy: %.2f\n", poly.SenseCount, poly.Entropy)
}
```
