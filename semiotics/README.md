# Package semiotics

Package `semiotics` provides computational models for structural semiotics, sign theory, and cultural connotation analysis, including Saussurean dyadic signs, Peircean triadic signs, Greimas semiotic squares, and affective valence evaluation.

## Overview

Semiotics is the study of signs and sign-using behavior (*semiosis*). The `semiotics` package provides computational implementations of classical semiotic frameworks:

1. **Saussurean Dyadic Model**:
   - Models the relationship between the *Signifier* (the acoustic or written form) and the *Signified* (the mental concept or WordNet synset).
   - Computes structural linguistic value through differential opposition against related lexical items.
2. **Peircean Triadic Model**:
   - Models the triad of *Representamen* (sign vehicle), *Object* (referent), and *Interpretant* (mental effect).
   - Classifies signs into C. S. Peirce's three fundamental sign modes:
     - **Icon**: Sign signifies via physical resemblance or acoustic mimicry (e.g. onomatopoeia like "buzz", "moo").
     - **Index**: Sign signifies via direct existential or causal connection (e.g. "smoke" -> fire).
     - **Symbol**: Sign signifies purely by arbitrary convention (e.g. "justice", "tree").
3. **Greimas Semiotic Square (*Carré Sémiotique*)**:
   - Maps the elementary structure of signification across four relational poles:
     - $S_1$: Primary assertion (e.g. "good")
     - $S_2$: Contrary opposite (e.g. "bad")
     - $\sim S_1$: Contradictory negation of $S_1$ ("not good")
     - $\sim S_2$: Contradictory negation of $S_2$ ("not bad")
   - Generates ASCII diagrams displaying contrariety, contradiction, and implication.
4. **Denotation vs Connotation**:
   - Separates literal dictionary glosses from affective valence (Positive, Neutral, Negative) and social register (Formal, Informal, Slang).

## Key Types and Functions

### Saussurean and Peircean Models

```go
type SaussureanSign struct {
    Signifier string
    Signified string
    Value     []string
}

type PeirceanTriad struct {
    Representamen string
    Object        string
    Interpretant  string
    Mode          SignMode // ModeIcon, ModeIndex, ModeSymbol
}
```

- `CreateSaussureanSign(res *gown.LexicalResource, word string) SaussureanSign`: Builds a dyadic sign with structural value from WordNet.
- `ClassifySignMode(word string) SignMode`: Determines whether a word functions primarily as an icon, index, or symbol.
- `CreatePeirceanTriad(rep, obj, interp string) PeirceanTriad`: Constructs a triadic sign representation.

### Semiotic Square

```go
type SemioticSquare struct {
    S1    string
    S2    string
    NotS1 string
    NotS2 string
}
```

- `GenerateSemioticSquare(res *gown.LexicalResource, term string) SemioticSquare`: Generates a 4-pole Greimas square using WordNet antonyms.
- `(sq SemioticSquare) Render() string`: Formats the square as an ASCII diagram showing structural relations.

### Connotation and Valence

```go
type ConnotationAnalysis struct {
    Word         string
    Denotation   string
    Valence      Valence // ValencePositive, ValenceNeutral, ValenceNegative
    Register     string
    Associations []string
}
```

- `AnalyzeConnotation(res *gown.LexicalResource, word string) ConnotationAnalysis`: Analyzes the evaluative sentiment and cultural associations of a word.

## Example

```go
package main

import (
    "fmt"
    "log"

    "github.com/raitucarp/gown"
    "github.com/raitucarp/gown/semiotics"
)

func main() {
    res, err := gown.ReadLexicalResource()
    if err != nil {
        log.Fatal(err)
    }

    // Generate Greimas Semiotic Square
    square := semiotics.GenerateSemioticSquare(res, "good")
    fmt.Println(square.Render())

    // Classify Peircean sign modes
    fmt.Println("buzz:", semiotics.ClassifySignMode("buzz"))     // Icon
    fmt.Println("smoke:", semiotics.ClassifySignMode("smoke"))   // Index
    fmt.Println("justice:", semiotics.ClassifySignMode("justice")) // Symbol

    // Connotation analysis
    con := semiotics.AnalyzeConnotation(res, "noble")
    fmt.Printf("Word: %s, Valence: %s, Register: %s\n", con.Word, con.Valence, con.Register)
}
```
