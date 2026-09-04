# Package pragmatics

Package `pragmatics` provides computational abstractions for pragmatics and contextual language interpretation, including speech act classification, deictic expression resolution, presupposition trigger detection, scalar implicatures, and politeness analysis.

## Overview

Pragmatics analyzes how context contributes to meaning—what is meant beyond what is literally said. The `pragmatics` package provides:

1. **Common Ground & Communicative Context**: Tracking speaker, addressee, setting (time and place), shared assumptions, and discourse history.
2. **Deictic Reference Resolution**: Identifying and grounding Person (`I`, `you`, `we`), Spatial (`here`, `there`, `nearby`), and Temporal (`now`, `yesterday`, `tomorrow`) deixis to physical/discourse entities.
3. **Speech Act Classification (Austin / Searle)**: Classifying utterances by illocutionary force:
   - **Assertive**: Stating facts, claims, descriptions.
   - **Directive**: Requests, commands, questions, invitations.
   - **Commissive**: Promises, oaths, threats, pledges.
   - **Expressive**: Apologies, thanks, congratulations, greetings.
   - **Declarative**: Formal status changes (e.g. "You are hired").
4. **Presupposition Trigger Extraction**: Detecting presuppositions from factive predicates (*realize*, *know*), change-of-state verbs (*stop*, *begin*), iteratives (*again*), and cleft sentences, with a negation-survival projection test.
5. **Gricean Scalar Implicatures**: Deriving conversational inferences from quantity scales (e.g., `<some, most, all>`, `<warm, hot, boiling>`).
6. **Politeness Strategies (Brown & Levinson)**: Detecting positive politeness, negative politeness, indirect requests, and deference markers.

## Key Types and Functions

### Speech Acts

```go
type IllocutionaryForce struct {
    Primary    SpeechActType // Assertive, Directive, Commissive, Expressive, Declarative
    Confidence float64
    Mood       string
}
```

- `ClassifySpeechAct(sentence string) IllocutionaryForce`: Classifies the primary speech act of an utterance.

### Deixis

```go
type DeicticExpression struct {
    Type     DeixisType // DeixisPerson, DeixisSpatial, DeixisTemporal, DeixisDiscourse
    Form     string
    Index    int
    Resolved string
}
```

- `IdentifyDeixis(text string) []DeicticExpression`: Identifies all deictic tokens and their types.
- `ResolveDeixis(text string, ctx Context) []DeicticExpression`: Resolves identified deictic tokens to concrete referents within the communicative context.

### Presuppositions

```go
type Presupposition struct {
    Trigger     string
    Type        string // Factive, ChangeOfState, Iterative, Cleft
    Proposition string
}
```

- `ExtractPresuppositions(sentence string) []Presupposition`: Extracts presupposed propositions and identifies the triggering lexical item.

### Implicatures and Politeness

- `DetectScalarImplicatures(sentence string) []Implicature`: Detects Gricean quantity implicatures triggered by weak scalar terms.
- `AnalyzePoliteness(sentence string) PolitenessAnalysis`: Assesses honorifics, indirect speech acts, and politeness markers.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/pragmatics"
)

func main() {
    utterance := "Could you please pass the salt?"

    // Speech act classification
    force := pragmatics.ClassifySpeechAct(utterance)
    fmt.Printf("Speech Act: %s (Confidence: %.2f)\n", force.Primary, force.Confidence)

    // Politeness analysis
    polite := pragmatics.AnalyzePoliteness(utterance)
    fmt.Printf("Polite Strategy: %s, Score: %.2f\n", polite.Strategy, polite.Score)

    // Presupposition trigger
    pres := pragmatics.ExtractPresuppositions("She stopped smoking.")
    for _, p := range pres {
        fmt.Printf("Trigger: %s (%s) -> Presupposition: %s\n", p.Trigger, p.Type, p.Proposition)
    }
}
```
