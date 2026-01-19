package gown

import (
	"strings"
	"testing"
)

var resource *LexicalResource

func TestMain(m *testing.M) {
	var err error
	resource, err = ReadLexicalResource()
	if err != nil {
		panic(err)
	}
	m.Run()
}

func TestReadLexicalResource(t *testing.T) {
	if len(resource.Lexicon.LexicalEntries) < 100000 {
		t.Errorf("Expected less than 100000 entries, got %d", len(resource.Lexicon.LexicalEntries))
	}
}

func TestSearch(t *testing.T) {
	query := "test"
	entries := resource.SearchLemma(query)
	for _, entry := range entries {
		if !strings.Contains(entry.Lemma.WrittenForm, query) {
			t.Errorf("Cannot find %s", query)
		}
	}
}

func TestSearchLemmaByDefinition(t *testing.T) {
	query := "test"
	entries := resource.SearchLemmaByDefinition(query)
	for _, entry := range entries {
		if !entry.HasDefinition(query) {
			t.Errorf("Cannot find in definition of %s", query)
		}
	}
}

func TestRandomLemma(t *testing.T) {
	entries := resource.RandomLemma(3)

	if len(entries) < 3 {
		t.Errorf("Expected at least 3 entries, got %d", len(entries))
	}

}

func TestRandomWords(t *testing.T) {
	entries := resource.Words().Random(5)
	for _, entry := range entries {
		if entry.IsCollocation() {
			t.Errorf("Lemma is not word = %s", entry.Lemma.WrittenForm)
		}
	}
}

func TestRandomCollocations(t *testing.T) {
	entries := resource.Collocations().Random(5)
	for _, entry := range entries {
		if !entry.IsCollocation() {
			t.Errorf("Lemma is not word = %s", entry.Lemma.WrittenForm)
		}
	}
}
