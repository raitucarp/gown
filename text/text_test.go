package text_test

import (
	"testing"

	"github.com/raitucarp/gown/text"
)

func TestTokenizeAndStopwordsComprehensive(t *testing.T) {
	// Empty string
	if tokens := text.Tokenize(""); len(tokens) != 0 {
		t.Errorf("Tokenize('') expected 0, got %d", len(tokens))
	}

	// Complex string with hyphens, quotes, apostrophes
	s := "The quick, brown-haired fox doesn't jump over the 'lazy' dog!"
	tokens := text.Tokenize(s)
	if len(tokens) < 8 {
		t.Errorf("Expected >= 8 tokens, got %d: %v", len(tokens), tokens)
	}

	// ExtractContentWords
	content := text.ExtractContentWords(s)
	for _, c := range content {
		if text.IsStopword(c) {
			t.Errorf("Unexpected stopword in content words: %s", c)
		}
	}

	// Non-stopword single character excluded by len(t) > 1
	singleCharText := "I a x y z"
	singleContent := text.ExtractContentWords(singleCharText)
	if len(singleContent) != 0 {
		t.Errorf("Expected 0 multi-character content words, got %d: %v", len(singleContent), singleContent)
	}

	// IsStopword case-insensitivity
	if !text.IsStopword("THE") || !text.IsStopword("the") {
		t.Errorf("IsStopword should be case-insensitive")
	}
	if text.IsStopword("computer") {
		t.Errorf("Expected 'computer' NOT to be stopword")
	}
}

func TestSentenceSegmentComprehensive(t *testing.T) {
	// Standard multi-sentence
	doc := "Hello world. How are you today? I am doing well!"
	sentences := text.SentenceSegment(doc)
	if len(sentences) != 3 {
		t.Errorf("Expected 3 sentences, got %d: %v", len(sentences), sentences)
	}

	// Trailing sentence without period
	docTrailing := "First sentence. Second without dot"
	sTrailing := text.SentenceSegment(docTrailing)
	if len(sTrailing) != 2 || sTrailing[1] != "Second without dot" {
		t.Errorf("Expected 2 sentences with trailing clause, got: %v", sTrailing)
	}

	// Empty text
	if sEmpty := text.SentenceSegment(""); len(sEmpty) != 0 {
		t.Errorf("SentenceSegment('') expected 0 sentences, got %d", len(sEmpty))
	}
}

func TestJaccardSimilarityComprehensive(t *testing.T) {
	s1 := []string{"cat", "dog", "animal"}
	s2 := []string{"cat", "dog", "pet"}

	// Intersection: 2, Union: 4 => 0.5
	j := text.JaccardSimilarity(s1, s2)
	if j != 0.5 {
		t.Errorf("Expected Jaccard 0.5, got %.4f", j)
	}

	// Both empty => 1.0
	if jEmpty := text.JaccardSimilarity(nil, nil); jEmpty != 1.0 {
		t.Errorf("Jaccard(nil, nil) expected 1.0, got %.2f", jEmpty)
	}

	// One empty => 0.0
	if jOneEmpty := text.JaccardSimilarity(s1, nil); jOneEmpty != 0.0 {
		t.Errorf("Jaccard(s1, nil) expected 0.0, got %.2f", jOneEmpty)
	}
	if jOneEmpty2 := text.JaccardSimilarity(nil, s2); jOneEmpty2 != 0.0 {
		t.Errorf("Jaccard(nil, s2) expected 0.0, got %.2f", jOneEmpty2)
	}

	// Disjoint sets => 0.0
	s3 := []string{"house", "building"}
	if jDisjoint := text.JaccardSimilarity(s1, s3); jDisjoint != 0.0 {
		t.Errorf("Jaccard disjoint expected 0.0, got %.2f", jDisjoint)
	}
}
