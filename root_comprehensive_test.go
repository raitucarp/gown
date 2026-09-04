package gown_test

import (
	"encoding/xml"
	"testing"

	"github.com/raitucarp/gown"
)

func TestRootFilterAndLexiconOperations(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("failed to load resource: %v", err)
	}

	t.Run("FilterByPos", func(t *testing.T) {
		entries, synsets := res.FilterByPos(gown.NounPos)
		if len(entries) == 0 || len(synsets) == 0 {
			t.Errorf("expected noun entries and synsets, got %d entries, %d synsets", len(entries), len(synsets))
		}
	})

	t.Run("FilterSynsetsByLexFile and SynsetByLexFile", func(t *testing.T) {
		lexFile := "noun.animal"
		synsets := res.FilterSynsetsByLexFile(lexFile)
		if len(synsets) == 0 {
			t.Fatalf("expected synsets for %s", lexFile)
		}
		pred := gown.SynsetByLexFile(lexFile)
		if !pred(synsets[0]) {
			t.Errorf("expected SynsetByLexFile to return true for %v", synsets[0].Lexfile)
		}
		filtered := gown.SynsetsByLexFile(synsets, lexFile)
		if len(filtered) != len(synsets) {
			t.Errorf("expected %d filtered synsets, got %d", len(synsets), len(filtered))
		}
	})

	t.Run("GroupEntryByPos and GroupSynsetsByPos", func(t *testing.T) {
		verbEntries := res.GroupEntryByPos(gown.VerbPos)
		if len(verbEntries) == 0 {
			t.Errorf("expected verb entries")
		}
		verbSynsets := res.GroupSynsetsByPos(gown.VerbPos)
		if len(verbSynsets) == 0 {
			t.Errorf("expected verb synsets")
		}
	})

	t.Run("SynsetsBySense", func(t *testing.T) {
		entries := res.LookupNoun("dog")
		if len(entries) == 0 || len(entries[0].Senses) == 0 {
			t.Fatalf("expected senses for dog")
		}
		mappedSynsets := res.SynsetsBySense(entries[0].Senses)
		if len(mappedSynsets) != len(entries[0].Senses) {
			t.Errorf("expected %d mapped synsets, got %d", len(entries[0].Senses), len(mappedSynsets))
		}
	})
}

func TestLexiconLazyInitialization(t *testing.T) {
	// Create a resource that has not had InitIndices called yet
	rawRes := &gown.LexicalResource{
		Lexicon: gown.Lexicon{
			LexicalEntries: []gown.LexicalEntry{
				{
					ID: "test-e1",
					Lemma: gown.Lemma{
						WrittenForm:  "testlemma",
						PartOfSpeech: "n",
					},
					Senses: []gown.Sense{
						{
							ID:     "test-s1",
							Synset: "test-syn1",
						},
					},
				},
			},
			Synsets: []gown.Synset{
				{
					ID:           "test-syn1",
					Ili:          "test-ili1",
					PartOfSpeech: "n",
					Members:      []string{"test-e1"},
					Definitions:  []string{"A test synset definition."},
				},
			},
		},
	}

	// Trigger lazy loading via SynsetsById, LexicalsById, SenseById
	synMap := rawRes.SynsetsById()
	if synMap["test-syn1"] == nil {
		t.Errorf("expected synset in lazy synsetsById")
	}

	lexMap := rawRes.LexicalsById()
	if lexMap["test-e1"] == nil {
		t.Errorf("expected lexical entry in lazy lexicalsById")
	}

	senseMap := rawRes.SenseById()
	if senseMap["test-s1"] == nil {
		t.Errorf("expected sense in lazy senseById")
	}
}

func TestSynsetUnmarshalXML(t *testing.T) {
	xmlData := `<Synset id="syn-1" members="w1 w2 w3">
		<Definition>A definition text.</Definition>
	</Synset>`

	var s gown.Synset
	err := xml.Unmarshal([]byte(xmlData), &s)
	if err != nil {
		t.Fatalf("xml unmarshal failed: %v", err)
	}
	if len(s.Members) != 3 || s.Members[0] != "w1" || s.Members[2] != "w3" {
		t.Errorf("unexpected members: %v", s.Members)
	}

	// Invalid XML
	var errSyn gown.Synset
	err = xml.Unmarshal([]byte("<Synset><unclosed>"), &errSyn)
	if err == nil {
		t.Errorf("expected error for invalid xml")
	}
}

func TestLexicalEntryMethods(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("failed to load resource: %v", err)
	}

	entries := res.LookupNoun("dog")
	if len(entries) == 0 {
		t.Fatalf("expected noun dog entries")
	}
	dog := &entries[0]

	t.Run("POS and String", func(t *testing.T) {
		if dog.PartOfSpeech() != gown.NounPos {
			t.Errorf("expected NounPos, got %v", dog.PartOfSpeech())
		}
		if dog.String() != "dog" {
			t.Errorf("expected 'dog', got %s", dog.String())
		}
	})

	t.Run("Prefix and Suffix", func(t *testing.T) {
		if !dog.StartsWith("do") {
			t.Errorf("expected dog to start with 'do'")
		}
		if dog.StartsWith("cat") {
			t.Errorf("expected dog not to start with 'cat'")
		}
		if !dog.EndsWith("og") {
			t.Errorf("expected dog to end with 'og'")
		}
		if dog.EndsWith("cat") {
			t.Errorf("expected dog not to end with 'cat'")
		}
	})

	t.Run("Word and Collocation", func(t *testing.T) {
		if !dog.IsWord() {
			t.Errorf("expected dog to be a word")
		}
		if dog.IsCollocation() {
			t.Errorf("expected dog not to be a collocation")
		}

		// Collocation entry
		colls := res.Lookup("hot dog")
		if len(colls) > 0 {
			if !colls[0].IsCollocation() {
				t.Errorf("expected 'hot dog' to be a collocation")
			}
			patterns := colls[0].CVPatterns()
			if len(patterns) == 0 {
				t.Errorf("expected CV patterns for 'hot dog'")
			}
		}
	})

	t.Run("CVPatterns", func(t *testing.T) {
		patterns := dog.CVPatterns()
		if len(patterns) == 0 {
			t.Fatalf("expected CV patterns for dog")
		}
		if !dog.HasCVPattern("cvc") && !dog.HasCVPattern("CVC") {
			t.Logf("dog CV pattern: %v", patterns)
		}
	})

	t.Run("Definitions and Examples", func(t *testing.T) {
		defs := dog.Definitions()
		if len(defs) == 0 {
			t.Errorf("expected definitions for dog")
		}
		if !dog.HasDefinition("canine") && !dog.HasDefinition("domestic") && !dog.HasDefinition("animal") {
			t.Logf("dog definitions: %v", defs)
		}

		_ = dog.Examples()
		_ = dog.HasExample("pet")

		if !dog.Contains("og") {
			t.Errorf("expected dog to contain 'og'")
		}
		if !dog.HasLength(3) {
			t.Errorf("expected dog to have length 3")
		}
		if dog.HasLength(10) {
			t.Errorf("expected dog not to have length 10")
		}
	})

	t.Run("Collection Filter Helpers", func(t *testing.T) {
		allEntries := res.Lookup("run")
		nouns := allEntries.Nouns()
		verbs := allEntries.Verbs()
		adjs := allEntries.Adjectives()
		advs := allEntries.Adverbs()

		if len(nouns) == 0 && len(verbs) == 0 {
			t.Errorf("expected nouns or verbs for 'run'")
		}
		_ = adjs
		_ = advs
	})

	t.Run("Relation", func(t *testing.T) {
		rel := dog.Relation()
		if rel == nil {
			t.Errorf("expected relation object")
		}
	})
}

func TestLookupVariantsAndEdgeCases(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("failed to load resource: %v", err)
	}

	t.Run("Lookup Empty", func(t *testing.T) {
		if res.Lookup("") != nil {
			t.Errorf("expected nil for empty lookup")
		}
	})

	t.Run("Lookup CaseSensitive", func(t *testing.T) {
		matches := res.Lookup("dog", gown.WithCaseSensitive())
		if len(matches) == 0 {
			t.Errorf("expected matches for lowercase 'dog' with case sensitivity")
		}
		matchesUpper := res.Lookup("DOG", gown.WithCaseSensitive())
		_ = matchesUpper
	})

	t.Run("Lookup ExactOnly", func(t *testing.T) {
		exact := res.Lookup("dog", gown.WithExactOnly())
		if len(exact) == 0 {
			t.Errorf("expected exact matches for dog")
		}
	})

	t.Run("LookupLemma", func(t *testing.T) {
		entries := res.LookupLemma("dog")
		if len(entries) == 0 {
			t.Errorf("expected entries from LookupLemma")
		}
	})

	t.Run("LookupPOSVariants", func(t *testing.T) {
		v := res.LookupVerb("run")
		if len(v) == 0 {
			t.Errorf("expected verb entries for 'run'")
		}
		adj := res.LookupAdjective("quick")
		if len(adj) == 0 {
			t.Errorf("expected adjective entries for 'quick'")
		}
		adv := res.LookupAdverb("quickly")
		if len(adv) == 0 {
			t.Errorf("expected adverb entries for 'quickly'")
		}
	})

	t.Run("SenseByID and SenseByKey", func(t *testing.T) {
		entries := res.LookupNoun("dog")
		if len(entries) > 0 && len(entries[0].Senses) > 0 {
			sense := &entries[0].Senses[0]
			found := res.SenseByID(sense.ID)
			if found == nil || found.ID != sense.ID {
				t.Errorf("expected to find sense by ID %s", sense.ID)
			}
			foundKey := res.SenseByKey(sense.ID)
			if foundKey == nil {
				t.Errorf("expected to find sense by Key %s", sense.ID)
			}
		}
		if res.SenseByID("non-existent-sense-id") != nil {
			t.Errorf("expected nil for non-existent sense ID")
		}
		if res.SenseByKey("non-existent-key-xyz") != nil {
			t.Errorf("expected nil for non-existent sense key")
		}
	})

	t.Run("ReverseLookup", func(t *testing.T) {
		if res.ReverseLookup("") != nil {
			t.Errorf("expected nil for empty reverse lookup")
		}

		synsets := res.ReverseLookup("domestic canine", gown.WithReversePOS(gown.NounPos), gown.WithReverseLimit(5))
		if len(synsets) == 0 {
			t.Logf("no synsets found with 'domestic canine'")
		}

		// PrimaryDefinition empty check
		emptySyn := gown.Synset{}
		if emptySyn.PrimaryDefinition() != "" {
			t.Errorf("expected empty string for synset without definitions")
		}
	})
}

func TestMorphyAndPatternEdgeCases(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("failed to load resource: %v", err)
	}

	t.Run("Morphy Empty and Unknown", func(t *testing.T) {
		if res.Morphy("", gown.NounPos) != nil {
			t.Errorf("expected nil for empty word")
		}
		if res.MorphyAll("") != nil {
			t.Errorf("expected nil for empty word")
		}
		unknown := res.Morphy("unrecognizedwordxyz12345", gown.NounPos)
		if len(unknown) == 0 {
			t.Errorf("expected fallback result for unknown word")
		}
	})

	t.Run("Morphy All POS and Irregulars", func(t *testing.T) {
		adjLemmas := res.Morphy("better", gown.AdjectivePos)
		if len(adjLemmas) == 0 {
			t.Errorf("expected lemmas for 'better'")
		}
		advLemmas := res.Morphy("best", gown.AdverbPos)
		if len(advLemmas) == 0 {
			t.Errorf("expected lemmas for 'best'")
		}
		verbLemmas := res.Morphy("ran", gown.VerbPos)
		if len(verbLemmas) == 0 {
			t.Errorf("expected lemmas for 'ran'")
		}
		nounLemmas := res.Morphy("oxen", gown.NounPos)
		if len(nounLemmas) == 0 {
			t.Errorf("expected lemmas for 'oxen'")
		}
	})

	t.Run("Pattern Edge Cases", func(t *testing.T) {
		if gown.OrthographicCV("") != "" {
			t.Errorf("expected empty string for empty input")
		}
		cv := gown.OrthographicCV("rhythm")
		if cv == "" {
			t.Errorf("expected CV string for rhythm")
		}
		cvNoMedial := gown.OrthographicCV("rhythm", gown.PatternClassifierConfig{TreatYAsVowelWhenMedial: false})
		if cvNoMedial == "" {
			t.Errorf("expected CV string for rhythm with no medial Y")
		}

		// CompilePattern errors
		if _, err := gown.CompilePattern("[invalid regex ("); err == nil {
			t.Errorf("expected regex compile error")
		}

		// FindByPattern with limit
		entries, err := res.FindByPattern("CVC", 5)
		if err != nil {
			t.Fatalf("FindByPattern failed: %v", err)
		}
		if len(entries) > 5 {
			t.Errorf("expected max 5 entries, got %d", len(entries))
		}

		// FindByPattern invalid
		if _, err := res.FindByPattern("[invalid regex ("); err == nil {
			t.Errorf("expected error from FindByPattern with invalid pattern")
		}
	})
}
