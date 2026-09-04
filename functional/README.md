# Package functional

Package `functional` provides computational structures and classification methods grounded in Systemic Functional Linguistics (SFL), founded by M. A. K. Halliday. It covers the ideational (transitivity), interpersonal (mood, modality, polarity), and textual (theme/rheme) metafunctions.

## Overview

Systemic Functional Linguistics models language as a social semiotic system structured around three simultaneous communicative functions (*metafunctions*):

1. **Ideational Metafunction (Transitivity)**:
   - Encodes our experience of the world into processes, participants, and circumstances.
   - Classifies verbal processes into Halliday's six primary types:
     - **Material**: Actions and happenings (*run*, *build*, *destroy*).
     - **Mental**: Cognition, perception, and emotion (*think*, *see*, *fear*).
     - **Relational**: Being, having, and identifying (*is*, *own*, *represent*).
     - **Verbal**: Saying and communicating (*tell*, *ask*, *explain*).
     - **Behavioral**: Physiological or psychological behavior (*breathe*, *sleep*, *cough*).
     - **Existential**: Existing or happening (*exist*, *occur*).
2. **Interpersonal Metafunction (Mood and Modality)**:
   - Encodes speech role relationships between speaker and listener.
   - Identifies clause mood (Declarative, Interrogative, Imperative), polarity (Positive, Negative), and modal operators (Epistemic, Deontic).
3. **Textual Metafunction (Theme and Rheme)**:
   - Decomposes clauses into the *Theme* (the point of departure, given information) and the *Rheme* (the core message or new information).
   - Classifies themes into Topical, Interpersonal, and Textual components.
4. **Lexical Cohesion**:
   - Quantifies repetition, synonymy, and collocational harmony to measure cohesive density across clauses.

## Key Types and Functions

### Transitivity

```go
type ProcessType string

const (
    ProcessMaterial    ProcessType = "Material"
    ProcessMental      ProcessType = "Mental"
    ProcessRelational  ProcessType = "Relational"
    ProcessVerbal      ProcessType = "Verbal"
    ProcessBehavioral  ProcessType = "Behavioral"
    ProcessExistential ProcessType = "Existential"
)
```

- `ClassifyProcess(verbLemma string) ProcessType`: Classifies a verb into its SFL transitivity process category.

### Interpersonal Analysis

```go
type InterpersonalProfile struct {
    Mood     string // Declarative, Interrogative, Imperative
    Polarity string // Positive, Negative
    Modality string // High, Medium, Low, None
}
```

- `AnalyzeInterpersonal(clauseText string) InterpersonalProfile`: Extracts mood, polarity, and modality ratings.

### Textual Analysis (Theme / Rheme)

```go
type ThemeRhemeAnalysis struct {
    Theme string
    Rheme string
}
```

- `AnalyzeThemeRheme(clauseText string) (theme, rheme string)`: Separates the starting point of the message (Theme) from its continuation (Rheme).
- `DecomposeTheme(clauseText string) ThemeComponents`: Separates textual (connectives), interpersonal (vocatives/adjuncts), and topical components of the Theme.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/functional"
)

func main() {
    // 1. Transitivity process classification
    fmt.Println("run:", functional.ClassifyProcess("run"))       // Material
    fmt.Println("ponder:", functional.ClassifyProcess("ponder")) // Mental
    fmt.Println("say:", functional.ClassifyProcess("say"))       // Verbal

    // 2. Theme / Rheme decomposition
    clause := "Yesterday evening, the committee finalized the report."
    theme, rheme := functional.AnalyzeThemeRheme(clause)
    fmt.Printf("Theme: %s\nRheme: %s\n", theme, rheme)

    // 3. Interpersonal profile
    interpersonal := functional.AnalyzeInterpersonal("Could you definitely verify the data?")
    fmt.Printf("Mood: %s, Modality: %s\n", interpersonal.Mood, interpersonal.Modality)
}
```
