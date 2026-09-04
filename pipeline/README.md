# Package pipeline

Package `pipeline` provides an integrated, multi-layer computational linguistics pipeline that orchestrates tokenization, phonology, morphology, syntax, semantics, pragmatics, discourse analysis, and semiotics on continuous text in a single pass.

## Overview

Complex language processing applications typically require coordinating operations across multiple distinct branches of linguistics. Package `pipeline` unifies all Gown modules into an end-to-end processing pipeline backed by the embedded WordNet database:

$$\text{Raw Text} \longrightarrow \begin{cases}
\text{Orthography \& Phonology} & \text{(syllables, CV patterns)} \\
\text{Morphology} & \text{(lemmatization, POS tags)} \\
\text{Syntax} & \text{(constituent parsing, grammatical relations)} \\
\text{Semantics} & \text{(thematic roles, selectional restrictions)} \\
\text{Pragmatics} & \text{(speech acts, deixis, presupposition)} \\
\text{Discourse} & \text{(EDUs, RST tree, coreference chains)} \\
\text{Semiotics} & \text{(sign modes, connotation, semiotic squares)}
\end{cases}$$

## Key Types and Functions

### Pipeline

```go
type Pipeline struct {
    res *gown.LexicalResource
}
```

- `NewPipeline(res *gown.LexicalResource) *Pipeline`: Constructs a pipeline instance backed by WordNet.
- `(p *Pipeline) Process(documentText string) *LinguisticDocument`: Executes full multi-layer analysis across the document.

### Output Document

```go
type LinguisticDocument struct {
    RawText           string
    Sentences         []SentenceAnalysis
    EDUs              []discourse.EDU
    RSTTree           *discourse.RSTNode
    CoreferenceChains []discourse.CoreferenceChain
    TopicTracking     []discourse.SentenceTopic
    SemioticSquares   map[string]semiotics.SemioticSquare
}

type SentenceAnalysis struct {
    ID             int
    Raw            string
    Words          []WordAnalysis
    Relations      []syntax.GrammaticalRelation
    Roles          semantics.PredicateArgumentStructure
    SpeechAct      pragmatics.IllocutionaryForce
    Deixis         []pragmatics.DeicticExpression
    Presupposition []pragmatics.Presupposition
    Implicatures   []pragmatics.Implicature
    Politeness     pragmatics.PolitenessAnalysis
}

type WordAnalysis struct {
    Surface     string
    Lemma       string
    POS         string
    CVPattern   string
    Syllables   int
    SignMode    semiotics.SignMode
    Connotation semiotics.ConnotationAnalysis
}
```

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/pipeline"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    pipe := pipeline.NewPipeline(res)
    doc := pipe.Process("The friendly dog barked joyfully. The owner smiled.")

    fmt.Printf("Document contains %d sentences and %d EDUs\n", len(doc.Sentences), len(doc.EDUs))

    for _, sent := range doc.Sentences {
        fmt.Printf("\nSentence %d: %s\n", sent.ID, sent.Raw)
        fmt.Printf("  Speech Act: %s\n", sent.SpeechAct.Primary)
        fmt.Printf("  Politeness: %s\n", sent.Politeness.Strategy)
        for _, w := range sent.Words {
            fmt.Printf("    [%s] Lemma=%s POS=%s Syllables=%d Connotation=%s\n",
                w.Surface, w.Lemma, w.POS, w.Syllables, w.Connotation.Valence)
        }
    }
}
```
