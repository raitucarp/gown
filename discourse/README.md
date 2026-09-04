# Package discourse

Package `discourse` provides data structures and algorithms for supra-sentential discourse analysis, including Elementary Discourse Unit (EDU) segmentation, Rhetorical Structure Theory (RST) parsing, pronominal/nominal coreference resolution, and topic progression modeling.

## Overview

Discourse analysis examines how sentences combine into coherent multi-sentence texts. The `discourse` package provides:

1. **Elementary Discourse Unit (EDU) Segmentation**: Partitioning raw text into clauses and discourse segments based on punctuation and discourse connective cues.
2. **Rhetorical Structure Theory (RST)**: Hierarchical decomposition of discourse into Nucleus and Satellite spans linked by rhetorical relations (Elaboration, Contrast, Cause, Condition, Attribution, Temporal).
3. **Coreference Resolution**: Linking referring expressions (pronouns, proper nouns, nominal phrases) to their antecedent entities across sentences.
4. **Discourse Knowledge Graphs**: Representing inter-sentence connections as a directed graph of entity and event nodes connected by coreference, temporal, and rhetorical edges.
5. **Thematic Progression (Daneš)**: Classifying how themes and rhemes transition across consecutive sentences (Linear Theme-Rheme progression, Constant Theme, Derived Theme).
6. **Centering-Inspired Coherence**: Evaluating local discourse transitions (Continue, Retain, Shift) to score overall text cohesion.

## Key Types and Functions

### Discourse Segmentation and RST

```go
type EDU struct {
    ID    int
    Text  string
    Start int
    End   int
}

type RSTNode struct {
    Relation  string    // Elaboration, Contrast, Cause, etc.
    IsNucleus bool
    EDU       *EDU
    Children  []*RSTNode
}
```

- `SegmentEDUs(documentText string) []EDU`: Splits text into discrete discourse segments.
- `BuildRSTTree(edus []EDU) *RSTNode`: Assembles segmented units into a hierarchical rhetorical structure tree.
- `(n *RSTNode) Render() string`: Formats the tree as an indented text visualization.

### Coreference Resolution

```go
type CoreferenceChain struct {
    EntityID int
    MainForm string
    Mentions []Mention
}
```

- `ResolveCoreference(documentText string) []CoreferenceChain`: Discovers entity mentions and groups them into coreference chains.

### Topic Progression and Coherence

- `TrackTopics(sentences []string) []SentenceTopic`: Extracts subject topics from consecutive sentences and computes continuity scores.
- `ClassifyThematicProgression(steps []ThemeRhemePair) []ThematicProgressionPattern`: Identifies Daneš progression patterns across sentence pairs.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/discourse"
)

func main() {
    text := "Although the rain was heavy, the match continued. The players remained focused."

    // Segment into EDUs
    edus := discourse.SegmentEDUs(text)
    fmt.Printf("Extracted %d EDUs\n", len(edus))
    for _, edu := range edus {
        fmt.Printf("  [%d] %s\n", edu.ID, edu.Text)
    }

    // Build and render RST Tree
    rst := discourse.BuildRSTTree(edus)
    fmt.Println("\nRST Structure:")
    fmt.Println(rst.Render())

    // Coreference resolution
    chains := discourse.ResolveCoreference(text)
    for _, chain := range chains {
        fmt.Printf("Entity: %s (Mentions: %d)\n", chain.MainForm, len(chain.Mentions))
    }
}
```
