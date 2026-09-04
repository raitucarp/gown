package phonology_test

import (
	"testing"

	"github.com/raitucarp/gown/phonology"
)

func TestSyllableCountingComprehensive(t *testing.T) {
	tests := []struct {
		word     string
		expected int
	}{
		{"", 0},
		{"   ", 0},
		{"a", 1},
		{"cat", 1},
		{"make", 1},
		{"late", 1},
		{"little", 2},
		{"water", 2},
		{"syllable", 3},
		{"linguistics", 3},
		{"banana", 3},
		{"rhythm", 1},
		{"cat123!?", 1},
	}

	for _, tt := range tests {
		cnt := phonology.CountSyllables(tt.word)
		if cnt != tt.expected {
			t.Errorf("CountSyllables(%q) = %d; expected %d", tt.word, cnt, tt.expected)
		}
	}
}

func TestSyllabifyComprehensive(t *testing.T) {
	// 1. Empty string
	if sylls := phonology.Syllabify(""); sylls != nil {
		t.Errorf("Syllabify('') expected nil, got %+v", sylls)
	}

	// 2. Single syllable
	syllsCat := phonology.Syllabify("cat")
	if len(syllsCat) != 1 {
		t.Fatalf("Expected 1 syllable for 'cat', got %d", len(syllsCat))
	}
	s := syllsCat[0]
	if s.Onset != "c" || s.Nucleus != "a" || s.Coda != "t" {
		t.Errorf("Expected onset 'c', nucleus 'a', coda 't', got %+v", s)
	}
	if s.Text() != "cat" {
		t.Errorf("Syllable.Text() expected 'cat', got '%s'", s.Text())
	}

	// 3. Multi-syllable word
	syllsWater := phonology.Syllabify("water")
	if len(syllsWater) < 2 {
		t.Errorf("Expected >= 2 syllables for 'water', got %d", len(syllsWater))
	}
}

func TestRhymesAndAlliterationComprehensive(t *testing.T) {
	// Rime function
	if r := phonology.Rime(""); r != "" {
		t.Errorf("Rime('') expected empty, got '%s'", r)
	}
	if r := phonology.Rime("cat"); r != "at" {
		t.Errorf("Rime('cat') expected 'at', got '%s'", r)
	}

	// AreRhymes edge cases
	if phonology.AreRhymes("", "") {
		t.Errorf("AreRhymes('', '') should be false")
	}
	if phonology.AreRhymes("cat", "cat") {
		t.Errorf("AreRhymes identical words should be false")
	}
	if !phonology.AreRhymes("cat", "hat") {
		t.Errorf("Expected 'cat' and 'hat' to rhyme")
	}
	if phonology.AreRhymes("cat", "dog") {
		t.Errorf("Expected 'cat' and 'dog' NOT to rhyme")
	}

	// AreAlliterations edge cases
	if phonology.AreAlliterations("", "peter") {
		t.Errorf("AreAlliterations with empty should be false")
	}
	if !phonology.AreAlliterations("peter", "piper") {
		t.Errorf("Expected 'peter' and 'piper' to alliterate")
	}
	if phonology.AreAlliterations("cat", "dog") {
		t.Errorf("Expected 'cat' and 'dog' NOT to alliterate")
	}

	// AreAssonances edge cases
	if phonology.AreAssonances("", "late") {
		t.Errorf("AreAssonances with empty should be false")
	}
	if !phonology.AreAssonances("bake", "late") {
		t.Errorf("Expected 'bake' and 'late' to have assonance")
	}
	if phonology.AreAssonances("cat", "dog") {
		t.Errorf("Expected 'cat' and 'dog' NOT to have assonance")
	}

	// AreConsonances edge cases
	if phonology.AreConsonances("", "hat") {
		t.Errorf("AreConsonances with empty should be false")
	}
	if !phonology.AreConsonances("cat", "bat") {
		t.Errorf("Expected 'cat' and 'bat' to have consonance")
	}
	if phonology.AreConsonances("cat", "dog") {
		t.Errorf("Expected 'cat' and 'dog' NOT to have consonance")
	}
}

func TestIPAParsingComprehensive(t *testing.T) {
	// Standard word /kæt/
	phonemes := phonology.IPAToPhonemes("kæt")
	if len(phonemes) != 3 {
		t.Fatalf("Expected 3 phonemes for /kæt/, got %d", len(phonemes))
	}

	cv := phonology.IPAToCV("kæt")
	if cv != "CVC" {
		t.Errorf("Expected CV for /kæt/ to be CVC, got %s", cv)
	}

	// With stress marks, length marks, brackets
	complexIPA := "[ˈtʃeɪndʒ]" // "change": affricate tʃ, diphthong eɪ, affricate ndʒ/dʒ
	pComplex := phonology.IPAToPhonemes(complexIPA)
	if len(pComplex) == 0 {
		t.Errorf("Expected phonemes for '%s', got 0", complexIPA)
	}

	// Unknown character in IPA
	pUnknown := phonology.IPAToPhonemes("x")
	if len(pUnknown) != 1 || pUnknown[0].Symbol != "x" {
		t.Errorf("Unexpected unknown phoneme parsing: %+v", pUnknown)
	}
}
