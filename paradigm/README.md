# Package paradigm

Package `paradigm` provides data structures and adapters for working across diverse modern theoretical paradigms in computational linguistics: Universal Dependencies (UD / CoNLL-U), FrameNet-inspired semantic frames, and Categorial Grammar (CG / CCG).

## Overview

Modern computational linguistics encompasses several distinct grammatical frameworks. Package `paradigm` provides unified representations for:

1. **Universal Dependencies (UD)**:
   - Dependency graph representation where heads are directly connected to dependents via typed grammatical relations (e.g. `nsubj`, `obj`, `amod`, `punct`).
   - Parses and serializes to the standard CoNLL-U tabular file format.
2. **Frame Semantics (FrameNet)**:
   - Encodes events and scenarios as semantic frames with specific frame elements (core and non-core roles).
   - Built-in models for primary frames such as *Ingestion* (Ingestor, Ingestibles) and *Motion* (Theme, Source, Path, Goal).
3. **Categorial Grammar (CG)**:
   - Models syntactic types as directional functions: basic types ($N$, $NP$, $S$) and complex functor types ($X/Y$ forward slash, $X\backslash Y$ backward slash).
   - Performs syntactic type reduction via directional function application ($NP$ followed by $NP\backslash S$ yields $S$).

## Key Types and Functions

### Universal Dependencies (CoNLL-U)

```go
type UDToken struct {
    ID     int
    Form   string
    Lemma  string
    UPOS   string
    XPOS   string
    Feats  string
    Head   int
    Deprel string
}

type UDSentence struct {
    Tokens []UDToken
}
```

- `ParseCoNLLU(conlluText string) ([]UDSentence, error)`: Parses CoNLL-U text into structured sentence trees.
- `(s UDSentence) FormatCoNLLU() string`: Serializes a sentence into standard CoNLL-U 10-column tabular format.

### Frame Semantics

```go
type SemanticFrame struct {
    Name     string
    Elements map[string]string // Role -> Token
}
```

- `InstantiateFrame(frameName string, bindings map[string]string) SemanticFrame`: Creates a populated semantic frame instance with validated element bindings.

### Categorial Grammar

```go
type CGType struct {
    IsPrimitive bool
    Base        string // "S", "NP", "N"
    Left        *CGType
    Right       *CGType
    Slash       string // "/", "\"
}
```

- `ParseCGType(typeStr string) (*CGType, error)`: Parses complex category strings (e.g. `(S\NP)/NP` for transitive verbs).
- `ApplyCG(functor, argument *CGType) (*CGType, bool)`: Performs forward or backward function application.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/paradigm"
)

func main() {
    // 1. Categorial Grammar type parsing & application
    transVerb, _ := paradigm.ParseCGType("(S\\NP)/NP")
    objNP, _ := paradigm.ParseCGType("NP")

    // Forward application: (S\NP)/NP applied to NP yields S\NP
    vp, ok := paradigm.ApplyCG(transVerb, objNP)
    if ok {
        fmt.Println("Result of forward application:", vp.String()) // S\NP
    }

    // 2. Semantic Frame instantiation
    frame := paradigm.InstantiateFrame("Ingestion", map[string]string{
        "Ingestor":    "child",
        "Ingestibles": "apple",
    })
    fmt.Printf("Frame: %s, Ingestor: %s\n", frame.Name, frame.Elements["Ingestor"])
}
```
