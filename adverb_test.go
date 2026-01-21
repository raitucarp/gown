package gown

import (
	"slices"
	"testing"
)

func TestAdverbsIdentity(t *testing.T) {
	adverbs := resource.Adverbs()

	t.Run("Adverbs count", func(t *testing.T) {
		if len(adverbs) < 1000 {
			t.Errorf("Expected less than 1000 Adverbs, got %d", len(adverbs))
		}
	})

	t.Run("Adverbs part of speech", func(t *testing.T) {
		for _, adverb := range adverbs {
			if adverb.PartOfSpeech() != AdverbPos {
				t.Fatal("Not a Adverb")
			}
		}
	})

	for _, adverb := range adverbs.All() {
		if adverb.PartOfSpeech() != AdverbPos {
			t.Fatal("Not a Adverb")
		}

		if !slices.ContainsFunc(adverb.Synsets(),
			func(s *Synset) bool { return s.Lexfile == string(AdverbAll) }) {

			for _, synset := range adverb.Synsets() {
				t.Errorf("%v", synset.Lexfile)
			}
			t.Errorf("Not %s, %s", AdverbAll,
				adverb.Lemma.WrittenForm)
		}
	}
}

func TestAdverbsSearch(t *testing.T) {
	adverbs := resource.Adverbs()
	query := "g"

	entries := adverbs.SearchLemma(query)
	for _, entry := range entries {
		if entry.PartOfSpeech() != AdverbPos {
			t.Fatal("Not a Adverb")
		}

		if !entry.Contains(query) {
			t.Errorf("Cannot find %s", query)
		}
	}
}

func TestAdverbsSearchByDefinition(t *testing.T) {
	adverbs := resource.Adverbs()
	query := "g"

	entries := adverbs.SearchLemmaByDefinition(query)
	for _, entry := range entries {
		if entry.PartOfSpeech() != AdverbPos {
			t.Fatal("Not a Adverb")
		}

		if !entry.HasDefinition(query) {
			t.Errorf("Cannot find in definition of %s", query)
		}

	}

}

func TestAdverbsRandom(t *testing.T) {
	Adverbs := resource.Adverbs()
	randomAdverbs := Adverbs.Random(10)

	if len(randomAdverbs) < 10 {
		t.Errorf("Expected 10 Adverbs, got %d", len(randomAdverbs))
	}

	for _, Adverb := range randomAdverbs {
		if !slices.ContainsFunc(Adverb.Synsets(), func(s *Synset) bool {
			return s.PartOfSpeech == string(AdverbPos)
		}) {
			t.Errorf("Not a Adverb %s", Adverb.Lemma.WrittenForm)
		}
	}

}

func TestAdverbWordsOrCollocations(t *testing.T) {
	adverbs := resource.Adverbs()
	allAdverbs := adverbs.All().Words()

	for _, entry := range allAdverbs {
		if entry.IsCollocation() {
			t.Errorf("Lemma is not word %s", entry.Lemma.WrittenForm)
		}
	}

	animalCollocations := adverbs.All().Collocations()

	for _, entry := range animalCollocations {
		if entry.IsWord() {
			t.Errorf("Lemma is not collocation %s", entry.Lemma.WrittenForm)
		}
	}
}

func TestAdverbRelations(t *testing.T) {
	Adverbs := resource.Adverbs()
	foodAdverbs := Adverbs.All()

	for _, Adverb := range foodAdverbs {
		synonyms := Adverb.Relation().Synonyms()
		if len(synonyms) > 0 {
			for _, synonym := range synonyms {
				if synonym.PartOfSpeech() != AdverbPos {
					t.Errorf("Not a Adverb %s", synonym.Lemma.WrittenForm)
				}
			}
		}
	}
}
