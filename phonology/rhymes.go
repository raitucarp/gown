package phonology

import (
	"strings"
)

// Rime extracts the rhyme core (nucleus + coda) of the final syllable of a word.
func Rime(word string) string {
	sylls := Syllabify(word)
	if len(sylls) == 0 {
		return ""
	}
	last := sylls[len(sylls)-1]
	return last.Nucleus + last.Coda
}

// AreRhymes checks if two words share the same ending rime sound.
func AreRhymes(word1, word2 string) bool {
	w1 := strings.TrimSpace(strings.ToLower(word1))
	w2 := strings.TrimSpace(strings.ToLower(word2))
	if w1 == "" || w2 == "" || w1 == w2 {
		return false
	}
	r1 := Rime(w1)
	r2 := Rime(w2)
	return r1 != "" && r1 == r2
}

// AreAlliterations checks if two words start with the same onset consonant cluster.
func AreAlliterations(word1, word2 string) bool {
	sylls1 := Syllabify(word1)
	sylls2 := Syllabify(word2)
	if len(sylls1) == 0 || len(sylls2) == 0 {
		return false
	}
	onset1 := sylls1[0].Onset
	onset2 := sylls2[0].Onset
	return onset1 != "" && onset1 == onset2
}

// AreAssonances checks if the stressed/final syllables share the same vowel nucleus.
func AreAssonances(word1, word2 string) bool {
	sylls1 := Syllabify(word1)
	sylls2 := Syllabify(word2)
	if len(sylls1) == 0 || len(sylls2) == 0 {
		return false
	}
	n1 := sylls1[len(sylls1)-1].Nucleus
	n2 := sylls2[len(sylls2)-1].Nucleus
	return n1 != "" && n1 == n2
}

// AreConsonances checks if the final syllables share the same coda consonants.
func AreConsonances(word1, word2 string) bool {
	sylls1 := Syllabify(word1)
	sylls2 := Syllabify(word2)
	if len(sylls1) == 0 || len(sylls2) == 0 {
		return false
	}
	c1 := sylls1[len(sylls1)-1].Coda
	c2 := sylls2[len(sylls2)-1].Coda
	return c1 != "" && c1 == c2
}
