# Package phonology

Package `phonology` provides tools for computational phonology, prosody, and articulatory linguistics, including syllable counting, Sonority Sequencing Principle syllabification, rhyme detection, poetic meter analysis, and International Phonetic Alphabet (IPA) parsing.

## Overview

Phonology investigates the sound patterns and systematic organization of sounds in language. Package `phonology` provides:

1. **Syllable Counting**: Deterministic count of syllables using diphthong merging and silent vowel handling.
2. **Syllabification (Onset-Nucleus-Coda)**: Decomposing words into syllables following the Maximal Onset Principle and Sonority Sequencing Principle.
3. **Phonetic Rhyme Detection**: Classifying rhyme types between word pairs:
   - **Perfect Rhyme**: Exact nucleus and coda match with differing onsets (*cat* / *hat*).
   - **Slant / Half Rhyme**: Similar vowel or final consonant sounds (*shape* / *keep*).
   - **Identical Rhyme**: Exact repetition.
4. **Alliteration, Assonance, and Consonance**: Detecting acoustic repetition patterns across word sequences.
5. **Prosodic Meter**: Analyzing stress patterns (0 = unstressed, 1 = stressed) and classifying classical poetic feet (Iamb, Trochee, Anapest, Dactyl, Spondee).
6. **IPA Articulatory Features**: Mapping IPA phonemes to articulatory features (vowel height, backness, roundedness; consonant place, manner, voicing).

## Key Types and Functions

### Syllabification

```go
type Syllable struct {
    Onset   string
    Nucleus string
    Coda    string
}
```

- `CountSyllables(word string) int`: Returns the number of syllables in a word.
- `Syllabify(word string) []Syllable`: Deconstructs a word into constituent syllables.

### Rhyme and Sound Devices

- `CheckRhyme(w1, w2 string) RhymeType`: Evaluates rhyme relationship (RhymePerfect, RhymeSlant, RhymeNone).
- `IsAlliterative(words []string) bool`: Returns `true` if words share initial consonant sounds.
- `IsAssonant(w1, w2 string) bool`: Returns `true` if words share primary vowel nucleus sounds.

### Poetic Meter

- `DetectMeterFoot(stressPattern []int) string`: Identifies poetic foot (e.g. `[0, 1]` -> "Iamb", `[1, 0]` -> "Trochee").

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/phonology"
)

func main() {
    word := "anticipation"
    fmt.Printf("Syllables in '%s': %d\n", word, phonology.CountSyllables(word))

    sylls := phonology.Syllabify("linguist")
    for i, s := range sylls {
        fmt.Printf("  Syllable %d: Onset=[%s] Nucleus=[%s] Coda=[%s]\n", i+1, s.Onset, s.Nucleus, s.Coda)
    }

    fmt.Println("Rhyme night/bright:", phonology.CheckRhyme("night", "bright"))
    fmt.Println("Meter [0, 1]:", phonology.DetectMeterFoot([]int{0, 1}))
}
```
