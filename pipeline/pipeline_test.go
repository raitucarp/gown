package pipeline_test

import (
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/pipeline"
)

func TestLinguisticPipeline(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	p := pipeline.NewPipeline(res)
	docText := "The dog chased the cat. Could you please help me? John stopped smoking."
	doc := p.Process(docText)

	if doc == nil {
		t.Fatalf("Expected non-nil LinguisticDocument")
	}

	if len(doc.Sentences) != 3 {
		t.Fatalf("Expected 3 sentences processed, got %d", len(doc.Sentences))
	}

	// 1. Sentence 1: "The dog chased the cat."
	s1 := doc.Sentences[0]
	if len(s1.Words) != 5 {
		t.Errorf("Expected 5 words in sentence 1, got %d", len(s1.Words))
	}
	if s1.SyntaxTree == nil {
		t.Errorf("Expected syntax tree for sentence 1")
	} else {
		t.Logf("Sentence 1 Syntax: %s", s1.SyntaxTree.BracketedString())
	}
	if s1.Roles.Predicate != "chased" {
		t.Errorf("Expected predicate 'chased', got '%s'", s1.Roles.Predicate)
	}
	t.Logf("Sentence 1 Semantic Roles: %s", s1.Roles.String())

	// 2. Sentence 2: "Could you please help me?"
	s2 := doc.Sentences[1]
	if len(s2.Deixis) == 0 {
		t.Errorf("Expected deixis in sentence 2")
	}
	t.Logf("Sentence 2 Politeness: %s (hedges: %v)", s2.Politeness.Strategy, s2.Politeness.MitigationTags)

	// 3. Sentence 3: "John stopped smoking."
	s3 := doc.Sentences[2]
	if len(s3.Presupposition) == 0 {
		t.Errorf("Expected presupposition triggered by 'stopped' in sentence 3")
	} else {
		t.Logf("Sentence 3 Presupposition: %s (trigger: %s)",
			s3.Presupposition[0].Presupposition, s3.Presupposition[0].Trigger)
	}

	// 4. Discourse & Semiotic layers
	if len(doc.EDUs) == 0 {
		t.Errorf("Expected EDUs segmented")
	}
	if len(doc.SemioticSquares) == 0 {
		t.Errorf("Expected semiotic squares generated")
	}
	for term, sq := range doc.SemioticSquares {
		t.Logf("Generated Semiotic Square for '%s': S1=%s, S2=%s", term, sq.S1, sq.S2)
	}
}
