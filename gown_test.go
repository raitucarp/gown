package gown

import (
	"slices"
	"strings"
	"testing"
)

func TestReadLexicalResource(t *testing.T) {
	resource, err := ReadLexicalResource()
	if err != nil {
		t.Fatal(err)
	}

	if len(resource.Lexicon.LexicalEntries) < 100000 {
		t.Errorf("Expected less than 100000 entries, got %d", len(resource.Lexicon.LexicalEntries))
	}

	t.Run("Search anything", func(t *testing.T) {
		query := "test"
		entries := resource.SearchLemma(query)
		for _, entry := range entries {
			if !strings.Contains(entry.Lemma.WrittenForm, query) {
				t.Errorf("Cannot find %s", query)
			}
		}
	})

	t.Run("Search lemma by definition", func(t *testing.T) {
		query := "test"
		entries := resource.SearchLemmaByDefinition(query)
		for _, entry := range entries {
			definitions := entry.Definitions()

			exists := slices.ContainsFunc(definitions, func(definition string) bool {
				return strings.Contains(definition, query)
			})

			if !exists {
				t.Errorf("Cannot find in definition of %s", query)
			}
		}
	})

	t.Run("Random lemma", func(t *testing.T) {
		entries := resource.RandomLemma(3)

		if len(entries) < 3 {
			t.Errorf("Expected at least 3 entries, got %d", len(entries))
		}
	})
}

func TestNouns(t *testing.T) {
	resource, err := ReadLexicalResource()
	if err != nil {
		t.Fatal(err)
	}

	nouns := resource.Nouns()
	if len(nouns) < 100000 {
		t.Errorf("Expected less than 100000 nouns, got %d", len(nouns))
	}

	for _, noun := range nouns {
		if noun.Lemma.PartOfSpeech != string(NounPos) {
			t.Fatal("Not a noun")
		}
	}

	for _, animal := range nouns.Animal() {
		if animal.Lemma.PartOfSpeech != string(NounPos) {
			t.Fatal("Not anoun")
		}

		for _, synset := range animal.synsets {
			if synset.Lexfile != string(NounAnimal) {
				t.Errorf("Not an animal, %s %s", synset.Lexfile, animal.Lemma.WrittenForm)
			}
		}
	}

	t.Run("Search animal", func(t *testing.T) {
		animal := nouns.Animal()
		query := "fly"
		entries := animal.SearchLemma(query)
		for _, entry := range entries {
			if entry.Lemma.PartOfSpeech != string(NounPos) {
				t.Fatal("Not a noun")
			}

			if !strings.Contains(entry.Lemma.WrittenForm, query) {
				t.Errorf("Cannot find %s", query)
			}
		}
	})

	t.Run("Search plant with fly definition", func(t *testing.T) {
		plant := nouns.Animal()
		query := "fly"
		entries := plant.SearchLemmaByDefinition(query)
		for _, entry := range entries {
			if entry.Lemma.PartOfSpeech != string(NounPos) {
				t.Fatal("Not a noun")
			}

			definitions := entry.Definitions()
			exists := slices.ContainsFunc(definitions, func(definition string) bool {
				return strings.Contains(definition, query)
			})

			if !exists {
				t.Errorf("Cannot find in definition of %s", query)
			}

		}
	})

	t.Run("Random nouns", func(t *testing.T) {
		randomNouns := nouns.Random(10)

		if len(randomNouns) < 10 {
			t.Errorf("Expected 10 nouns, got %d", len(randomNouns))
		}
	})

	t.Run("Random animal", func(t *testing.T) {
		animal := nouns.Animal().Random(10)

		if len(animal) < 10 {
			t.Errorf("Expected 10 nouns, got %d", len(animal))
		}

		for _, a := range animal {
			for _, synset := range a.synsets {
				if synset.Lexfile != string(NounAnimal) {
					t.Errorf("Not an animal, %s %s", synset.Lexfile, a.Lemma.WrittenForm)
				}
			}
		}
	})
}
