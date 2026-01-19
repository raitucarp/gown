package gown

import (
	"fmt"
	"slices"
	"testing"
)

func TestNounsIdentity(t *testing.T) {
	nouns := resource.Nouns()

	t.Run("Nouns count", func(t *testing.T) {
		if len(nouns) < 100000 {
			t.Errorf("Expected less than 100000 nouns, got %d", len(nouns))
		}
	})

	t.Run("Nouns part of speech", func(t *testing.T) {
		for _, noun := range nouns {
			if noun.PartOfSpeech() != NounPos {
				t.Fatal("Not a noun")
			}
		}
	})

	for kind, nounsByKind := range nouns.AllKind() {
		t.Run(fmt.Sprintf("Test %s by lex kind %s", "noun", kind), func(t *testing.T) {
			if len(nounsByKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			for _, noun := range nounsByKind {
				if noun.PartOfSpeech() != NounPos {
					t.Fatal("Not a noun")
				}

				if !slices.ContainsFunc(noun.Synsets(),
					func(s *Synset) bool { return s.Lexfile == string(kind) }) {

					t.Errorf("Not %s, %s", kind,
						noun.Lemma.WrittenForm)
				}
			}
		})
	}
}

func TestNounsSearch(t *testing.T) {
	nouns := resource.Nouns()
	query := "fly"

	for kind, nounsKind := range nouns.AllKind() {
		t.Run(fmt.Sprintf("search %s", kind), func(t *testing.T) {
			if len(nounsKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := nounsKind.SearchLemma(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != NounPos {
					t.Fatal("Not a noun")
				}

				if !entry.Contains(query) {
					t.Errorf("Cannot find %s", query)
				}
			}
		})
	}
}

func TestNounsSearchByDefinition(t *testing.T) {
	nouns := resource.Nouns()
	query := "fly"

	for kind, nounsKind := range nouns.AllKind() {
		t.Run(fmt.Sprintf("search %s by definition contains %s", kind, query), func(t *testing.T) {
			if len(nounsKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := nounsKind.SearchLemmaByDefinition(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != NounPos {
					t.Fatal("Not a noun")
				}

				if !entry.HasDefinition(query) {
					t.Errorf("Cannot find in definition of %s", query)
				}

			}
		})
	}

}

func TestNounsRandom(t *testing.T) {
	nouns := resource.Nouns()
	randomNouns := nouns.Random(10)

	if len(randomNouns) < 10 {
		t.Errorf("Expected 10 nouns, got %d", len(randomNouns))
	}

	for _, noun := range randomNouns {
		if !slices.ContainsFunc(noun.Synsets(), func(s *Synset) bool {
			return s.PartOfSpeech == string(NounPos)
		}) {
			t.Errorf("Not a noun %s", noun.Lemma.WrittenForm)
		}
	}

}

func TestWordsOrCollocations(t *testing.T) {
	nouns := resource.Nouns()
	animals := nouns.Animal().Words()

	for _, entry := range animals {
		if entry.IsCollocation() {
			t.Errorf("Lemma is not word %s", entry.Lemma.WrittenForm)
		}
	}

	animalCollocations := nouns.Animal().Collocations()

	for _, entry := range animalCollocations {
		if entry.IsWord() {
			t.Errorf("Lemma is not collocation %s", entry.Lemma.WrittenForm)
		}
	}
}

func TestNounRelations(t *testing.T) {
	nouns := resource.Nouns()
	foodNouns := nouns.Food()

	for _, noun := range foodNouns {
		synonyms := noun.Relation().Synonyms()
		if len(synonyms) > 0 {
			for _, synonym := range synonyms {
				if synonym.PartOfSpeech() != NounPos {
					t.Errorf("Not a noun %s", synonym.Lemma.WrittenForm)
				}
			}
		}
	}
}
