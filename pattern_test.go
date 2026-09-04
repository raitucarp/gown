package gown_test

import (
	"testing"

	"github.com/raitucarp/gown"
)

func TestOrthographicCV(t *testing.T) {
	tests := []struct {
		word     string
		expected string
	}{
		{"cat", "CVC"},
		{"eat", "VVC"},
		{"bread", "CCVVC"},
		{"strike", "CCCVCV"},
		{"rhythm", "CCVCCC"}, // Y medial treated as vowel
	}

	for _, tt := range tests {
		cv := gown.OrthographicCV(tt.word)
		if cv != tt.expected {
			t.Errorf("OrthographicCV(%q) = %s; expected %s", tt.word, cv, tt.expected)
		}
	}
}

func TestPatternQuery(t *testing.T) {
	pq, err := gown.CompilePattern("CVC")
	if err != nil {
		t.Fatalf("Failed to compile pattern: %v", err)
	}

	if !pq.MatchesWord("cat") {
		t.Errorf("Expected 'cat' to match CVC")
	}
	if !pq.MatchesWord("dog") {
		t.Errorf("Expected 'dog' to match CVC")
	}
	if pq.MatchesWord("bread") {
		t.Errorf("Expected 'bread' NOT to match CVC")
	}

	// Wildcard pattern
	wildcard, err := gown.CompilePattern("C?C")
	if err != nil {
		t.Fatalf("Failed to compile wildcard: %v", err)
	}
	if !wildcard.Matches("CVC") || !wildcard.Matches("CCC") {
		t.Errorf("Expected C?C to match CVC and CCC")
	}
}

func TestFindByPattern(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	matches, err := res.FindByPattern("CVC", 5)
	if err != nil {
		t.Fatalf("FindByPattern error: %v", err)
	}

	if len(matches) == 0 {
		t.Errorf("Expected at least one CVC match")
	}

	for _, m := range matches {
		cv := gown.OrthographicCV(m.Lemma.WrittenForm)
		if cv != "CVC" {
			t.Errorf("Expected match %s to have CV pattern CVC, got %s", m.Lemma.WrittenForm, cv)
		}
	}
}
