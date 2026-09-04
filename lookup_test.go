package gown_test

import (
	"testing"

	"github.com/raitucarp/gown"
)

func TestLookupExact(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	entries := res.LookupExact("dog")
	if len(entries) == 0 {
		t.Errorf("Expected entries for 'dog', got 0")
	}

	// Verify exact match
	for _, e := range entries {
		if e.Lemma.WrittenForm != "dog" {
			t.Errorf("Expected 'dog', got '%s'", e.Lemma.WrittenForm)
		}
	}
}

func TestLookupCaseInsensitive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	lower := res.Lookup("cat")
	upper := res.Lookup("CAT")

	if len(lower) == 0 {
		t.Errorf("Expected entries for 'cat'")
	}
	if len(upper) == 0 {
		t.Errorf("Expected entries for 'CAT' via case-insensitive lookup")
	}
}

func TestLookupWithPOS(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	nouns := res.LookupNoun("run")
	verbs := res.LookupVerb("run")

	if len(nouns) == 0 {
		t.Errorf("Expected 'run' as noun")
	}
	if len(verbs) == 0 {
		t.Errorf("Expected 'run' as verb")
	}

	for _, n := range nouns {
		if n.Lemma.PartOfSpeech != "n" {
			t.Errorf("Expected noun POS 'n', got %s", n.Lemma.PartOfSpeech)
		}
	}
	for _, v := range verbs {
		if v.Lemma.PartOfSpeech != "v" {
			t.Errorf("Expected verb POS 'v', got %s", v.Lemma.PartOfSpeech)
		}
	}
}

func TestLookupMorphy(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Irregular noun plural
	mice := res.Lookup("mice", gown.WithPOS(gown.NounPos))
	if len(mice) == 0 {
		t.Errorf("Expected Morphy to resolve 'mice' to 'mouse'")
	} else {
		found := false
		for _, e := range mice {
			if e.Lemma.WrittenForm == "mouse" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find lemma 'mouse' for 'mice'")
		}
	}

	// Regular verb inflection
	running := res.Lookup("running", gown.WithPOS(gown.VerbPos))
	if len(running) == 0 {
		t.Errorf("Expected Morphy to resolve 'running' to 'run'")
	} else {
		found := false
		for _, e := range running {
			if e.Lemma.WrittenForm == "run" {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected to find lemma 'run' for 'running'")
		}
	}
}

func TestSynsetByIDAndILI(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	dogEntries := res.LookupExact("dog", gown.NounPos)
	if len(dogEntries) == 0 || len(dogEntries[0].Senses) == 0 {
		t.Fatalf("Could not find senses for dog")
	}

	synsetID := dogEntries[0].Senses[0].Synset
	synset := res.SynsetByID(synsetID)
	if synset == nil {
		t.Fatalf("Expected to find synset by ID %s", synsetID)
	}
	if synset.ID != synsetID {
		t.Errorf("ID mismatch: got %s, expected %s", synset.ID, synsetID)
	}

	if synset.Ili != "" {
		iliSynset := res.SynsetByILI(synset.Ili)
		if iliSynset == nil {
			t.Errorf("Expected to find synset by ILI %s", synset.Ili)
		} else if iliSynset.ID != synset.ID {
			t.Errorf("ILI synset ID mismatch: %s != %s", iliSynset.ID, synset.ID)
		}
	}
}

func TestReverseLookup(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	results := res.ReverseLookup("domesticated canine", gown.WithReverseLimit(5))
	if len(results) == 0 {
		t.Logf("Note: exact phrase 'domesticated canine' not in definitions, trying 'canine'")
		results = res.ReverseLookup("canine", gown.WithReverseLimit(5))
	}
	if len(results) == 0 {
		t.Errorf("Expected reverse lookup results for 'canine'")
	}
}
