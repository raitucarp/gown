package gown

import (
	"fmt"
	"slices"
	"testing"
)

func TestAdjectivesIdentity(t *testing.T) {
	Adjectives := resource.Adjectives()

	t.Run("Adjectives count", func(t *testing.T) {
		if len(Adjectives) < 10000 {
			t.Errorf("Expected less than 10000 Adjectives, got %d", len(Adjectives))
		}
	})

	t.Run("Adjectives part of speech", func(t *testing.T) {
		for _, Adjective := range Adjectives {
			if Adjective.PartOfSpeech() != AdjectivePos && Adjective.PartOfSpeech() != AdjectiveSatellitePos {
				t.Fatal("Not a Adjective")
			}
		}
	})

	for kind, AdjectivesByKind := range Adjectives.AllKind() {
		t.Run(fmt.Sprintf("Test %s by lex kind %s", "Adjective", kind), func(t *testing.T) {
			if len(AdjectivesByKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			for _, Adjective := range AdjectivesByKind {
				if Adjective.PartOfSpeech() != AdjectivePos && Adjective.PartOfSpeech() != AdjectiveSatellitePos {
					t.Fatal("Not a Adjective")
				}

				if !slices.ContainsFunc(Adjective.Synsets(),
					func(s *Synset) bool { return s.Lexfile == string(kind) }) {

					t.Errorf("Not %s, %s", kind,
						Adjective.Lemma.WrittenForm)
				}
			}
		})
	}
}

func TestAdjectivesSearch(t *testing.T) {
	adjectives := resource.Adjectives()
	query := "fly"

	for kind, AdjectivesKind := range adjectives.AllKind() {
		t.Run(fmt.Sprintf("search %s", kind), func(t *testing.T) {
			if len(AdjectivesKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := AdjectivesKind.SearchLemma(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != AdjectivePos {
					t.Fatal("Not a Adjective")
				}

				if !entry.Contains(query) {
					t.Errorf("Cannot find %s", query)
				}
			}
		})
	}
}

func TestAdjectivesSearchByDefinition(t *testing.T) {
	adjectives := resource.Adjectives()
	query := "f"

	for kind, AdjectivesKind := range adjectives.AllKind() {
		t.Run(fmt.Sprintf("search %s by definition contains %s", kind, query), func(t *testing.T) {
			if len(AdjectivesKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := AdjectivesKind.SearchLemmaByDefinition(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != AdjectivePos && entry.PartOfSpeech() != AdjectiveSatellitePos {
					t.Fatal("Not a Adjective")
				}

				if !entry.HasDefinition(query) {
					t.Errorf("Cannot find in definition of %s", query)
				}

			}
		})
	}

}

func TestAdjectivesRandom(t *testing.T) {
	adjectives := resource.Adjectives()
	randomAdjectives := adjectives.Random(10)

	if len(randomAdjectives) < 10 {
		t.Errorf("Expected 10 Adjectives, got %d", len(randomAdjectives))
	}

	for _, Adjective := range randomAdjectives {
		if !slices.ContainsFunc(Adjective.Synsets(), func(s *Synset) bool {
			return s.PartOfSpeech == string(AdjectivePos) || s.PartOfSpeech == string(AdjectiveSatellitePos)
		}) {
			t.Errorf("Not a Adjective %s", Adjective.Lemma.WrittenForm)
		}
	}

}

func TestAdjectiveWordsOrCollocations(t *testing.T) {
	adjectives := resource.Adjectives()
	pertainyms := adjectives.Pertainym().Words()

	for _, entry := range pertainyms {
		if entry.IsCollocation() {
			t.Errorf("Lemma is not word %s", entry.Lemma.WrittenForm)
		}
	}

	animalCollocations := adjectives.Pertainym().Collocations()

	for _, entry := range animalCollocations {
		if entry.IsWord() {
			t.Errorf("Lemma is not collocation %s", entry.Lemma.WrittenForm)
		}
	}
}

func TestAdjectiveRelations(t *testing.T) {
	adjectives := resource.Adjectives()
	pertainyms := adjectives.Pertainym()

	for _, Adjective := range pertainyms {
		synonyms := Adjective.Relation().Synonyms()
		if len(synonyms) > 0 {
			for _, synonym := range synonyms {
				if synonym.PartOfSpeech() != AdjectivePos && synonym.PartOfSpeech() != AdjectiveSatellitePos {
					t.Errorf("Not a Adjective %s, expected %s, got %s", synonym.Lemma.WrittenForm, AdjectivePos, synonym.PartOfSpeech())
				}
			}
		}
	}
}
