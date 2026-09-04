package phonology

import (
	"strings"
	"unicode"
)

// Syllable represents a single syllable with onset, nucleus, and coda.
type Syllable struct {
	Onset   string `json:"onset"`
	Nucleus string `json:"nucleus"`
	Coda    string `json:"coda"`
	Stress  string `json:"stress,omitempty"` // "primary", "secondary", "none"
}

// Text returns the surface text of the syllable.
func (s Syllable) Text() string {
	return s.Onset + s.Nucleus + s.Coda
}

// CountSyllables estimates the number of syllables in an English word.
func CountSyllables(word string) int {
	word = strings.TrimSpace(strings.ToLower(word))
	if len(word) == 0 {
		return 0
	}

	// Filter non-letters
	var sb strings.Builder
	for _, r := range word {
		if unicode.IsLetter(r) {
			sb.WriteRune(r)
		}
	}
	w := sb.String()
	if len(w) <= 3 {
		return 1
	}

	// Remove silent trailing 'e' (e.g. "make", "late") unless preceded by 'l' (e.g. "little")
	if strings.HasSuffix(w, "e") && !strings.HasSuffix(w, "le") {
		w = w[:len(w)-1]
	}

	vowels := "aeiouy"
	count := 0
	prevIsVowel := false

	for _, r := range w {
		isVowel := strings.ContainsRune(vowels, r)
		if isVowel && !prevIsVowel {
			count++
		}
		prevIsVowel = isVowel
	}

	if count == 0 {
		count = 1
	}
	return count
}

// Syllabify decomposes a simple English word into heuristic Onset-Nucleus-Coda syllables.
func Syllabify(word string) []Syllable {
	word = strings.TrimSpace(strings.ToLower(word))
	if len(word) == 0 {
		return nil
	}

	numSyllables := CountSyllables(word)
	if numSyllables <= 1 {
		// Single syllable: split into onset, nucleus, coda
		return []Syllable{splitSingleSyllable(word)}
	}

	// Multi-syllable simple division based on vowel cores
	vowels := "aeiouy"
	var syllables []Syllable
	runes := []rune(word)

	var currentVowelCore strings.Builder
	var currentConsonants strings.Builder

	type segment struct {
		isVowel bool
		str     string
	}
	var segments []segment

	inVowel := strings.ContainsRune(vowels, runes[0])
	var cur strings.Builder
	for _, r := range runes {
		isV := strings.ContainsRune(vowels, r)
		if isV == inVowel {
			cur.WriteRune(r)
		} else {
			segments = append(segments, segment{isVowel: inVowel, str: cur.String()})
			cur.Reset()
			cur.WriteRune(r)
			inVowel = isV
		}
	}
	if cur.Len() > 0 {
		segments = append(segments, segment{isVowel: inVowel, str: cur.String()})
	}

	_ = currentVowelCore
	_ = currentConsonants

	// Build syllables from vowel centers
	for i := 0; i < len(segments); i++ {
		if segments[i].isVowel {
			onset := ""
			if i > 0 && !segments[i-1].isVowel {
				onset = segments[i-1].str
			}
			nucleus := segments[i].str
			coda := ""
			if i+1 < len(segments) && !segments[i+1].isVowel {
				// If next is not last segment, divide consonants
				if i+2 < len(segments) {
					cons := segments[i+1].str
					if len(cons) > 1 {
						coda = string(cons[0])
						segments[i+1].str = cons[1:]
					}
				} else {
					coda = segments[i+1].str
				}
			}
			syllables = append(syllables, Syllable{
				Onset:   onset,
				Nucleus: nucleus,
				Coda:    coda,
			})
		}
	}

	if len(syllables) == 0 {
		syllables = append(syllables, splitSingleSyllable(word))
	}

	return syllables
}

func splitSingleSyllable(word string) Syllable {
	vowels := "aeiouy"
	runes := []rune(word)

	firstVowel := -1
	for i, r := range runes {
		if strings.ContainsRune(vowels, r) {
			firstVowel = i
			break
		}
	}

	if firstVowel == -1 {
		return Syllable{Onset: word}
	}

	endVowel := firstVowel
	for endVowel+1 < len(runes) && strings.ContainsRune(vowels, runes[endVowel+1]) {
		endVowel++
	}

	onset := string(runes[:firstVowel])
	nucleus := string(runes[firstVowel : endVowel+1])
	coda := string(runes[endVowel+1:])

	return Syllable{
		Onset:   onset,
		Nucleus: nucleus,
		Coda:    coda,
	}
}
