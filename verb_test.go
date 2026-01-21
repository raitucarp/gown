package gown

import (
	"fmt"
	"slices"
	"testing"
)

func TestVerbsIdentity(t *testing.T) {
	verbs := resource.Verbs()

	t.Run("Verbs count", func(t *testing.T) {
		if len(verbs) < 10000 {
			t.Errorf("Expected less than 10000 Verbs, got %d", len(verbs))
		}
	})

	t.Run("Verbs part of speech", func(t *testing.T) {
		for _, Verb := range verbs {
			if Verb.PartOfSpeech() != VerbPos {
				t.Fatal("Not a Verb")
			}
		}
	})

	for kind, VerbsByKind := range verbs.AllKind() {
		t.Run(fmt.Sprintf("Test %s by lex kind %s", "Verb", kind), func(t *testing.T) {
			if len(VerbsByKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			for _, Verb := range VerbsByKind {
				if Verb.PartOfSpeech() != VerbPos {
					t.Fatal("Not a Verb")
				}

				if !slices.ContainsFunc(Verb.Synsets(),
					func(s *Synset) bool { return s.Lexfile == string(kind) }) {

					t.Errorf("Not %s, %s", kind,
						Verb.Lemma.WrittenForm)
				}
			}
		})
	}
}

func TestVerbsSearch(t *testing.T) {
	verbs := resource.Verbs()
	query := "run"

	for kind, verbsKind := range verbs.AllKind() {
		t.Run(fmt.Sprintf("search %s", kind), func(t *testing.T) {
			if len(verbsKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := verbsKind.SearchLemma(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != VerbPos {
					t.Fatal("Not a Verb")
				}

				if !entry.Contains(query) {
					t.Errorf("Cannot find %s", query)
				}
			}
		})
	}
}

func TestVerbsSearchByDefinition(t *testing.T) {
	Verbs := resource.Verbs()
	query := "r"

	for kind, VerbsKind := range Verbs.AllKind() {
		t.Run(fmt.Sprintf("search %s by definition contains %s", kind, query), func(t *testing.T) {
			if len(VerbsKind) <= 0 {
				t.Errorf("No %s found", kind)
			}

			entries := VerbsKind.SearchLemmaByDefinition(query)
			for _, entry := range entries {
				if entry.PartOfSpeech() != VerbPos {
					t.Fatal("Not a Verb")
				}

				if !entry.HasDefinition(query) {
					t.Errorf("Cannot find in definition of %s", query)
				}

			}
		})
	}

}

func TestVerbsRandom(t *testing.T) {
	verbs := resource.Verbs()
	randomVerbs := verbs.Random(10)

	if len(randomVerbs) < 10 {
		t.Errorf("Expected 10 Verbs, got %d", len(randomVerbs))
	}

	for _, Verb := range randomVerbs {
		if !slices.ContainsFunc(Verb.Synsets(), func(s *Synset) bool {
			return s.PartOfSpeech == string(VerbPos)
		}) {
			t.Errorf("Not a Verb %s", Verb.Lemma.WrittenForm)
		}
	}

}

func TestVerbWordsOrCollocations(t *testing.T) {
	verbs := resource.Verbs()
	perceptions := verbs.Perception().Words()

	for _, entry := range perceptions {
		if entry.IsCollocation() {
			t.Errorf("Lemma is not word %s", entry.Lemma.WrittenForm)
		}
	}

	perceptionCollocations := verbs.Perception().Collocations()

	for _, entry := range perceptionCollocations {
		if entry.IsWord() {
			t.Errorf("Lemma is not collocation %s", entry.Lemma.WrittenForm)
		}
	}
}

func TestVerbRelations(t *testing.T) {
	verbs := resource.Verbs()
	foodVerbs := verbs.Perception()

	for _, Verb := range foodVerbs {
		synonyms := Verb.Relation().Synonyms()
		if len(synonyms) > 0 {
			for _, synonym := range synonyms {
				if synonym.PartOfSpeech() != VerbPos {
					t.Errorf("Not a Verb %s", synonym.Lemma.WrittenForm)
				}
			}
		}
	}
}
