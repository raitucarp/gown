package gown_test

import (
	"testing"

	"github.com/raitucarp/gown"
)

func TestAllLexicalRelations(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Test on multiple representative words (dog, run, fast, good, snore, eat)
	words := []string{"dog", "run", "fast", "good", "snore", "eat", "give", "car"}

	for _, w := range words {
		entries := res.Lookup(w)
		for _, e := range entries {
			rel := e.Relation()
			_ = rel.Synonyms()
			_ = rel.Also()
			_ = rel.Antonyms()
			_ = rel.Derivations()
			_ = rel.Exemplifies()
			_ = rel.IsExemplifiedBy()
			_ = rel.Others()
			_ = rel.Participles()
			_ = rel.Pertainyms()
			_ = rel.Similars()
			_ = rel.Agents()
			_ = rel.BodyParts()
			_ = rel.ByMeansOf()
			_ = rel.Destinations()
			_ = rel.Events()
			_ = rel.Instruments()
			_ = rel.Locations()
			_ = rel.Materials()
			_ = rel.Properties()
			_ = rel.Results()
			_ = rel.States()
			_ = rel.Undergoers()
			_ = rel.Uses()
			_ = rel.Vehicles()
			_ = rel.Attributes()
			_ = rel.Causes()
			_ = rel.DomainRegions()
			_ = rel.DomainTopics()
			_ = rel.Entails()
			_ = rel.HasDomainRegions()
			_ = rel.HasDomainTopics()
			_ = rel.HoloMembers()
			_ = rel.HoloParts()
			_ = rel.HoloSubstances()
			_ = rel.Hypernyms()
			_ = rel.Hyponyms()
			_ = rel.InstanceHypernyms()
			_ = rel.InstanceHyponyms()
			_ = rel.IsCausedBy()
			_ = rel.IsEntailedBy()
			_ = rel.MeroMembers()
			_ = rel.MeroParts()
			_ = rel.MeroSubstances()
		}
	}
}

func TestSynsetMethods(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	entries := res.LookupExact("dog", gown.NounPos)
	if len(entries) == 0 || len(entries[0].Synsets()) == 0 {
		t.Fatalf("dog synset not found")
	}

	s := entries[0].Synsets()[0]
	if def := s.PrimaryDefinition(); len(def) == 0 {
		t.Errorf("Expected non-empty primary definition for dog synset")
	}

	lexEntries := s.LexicalEntries(res)
	if len(lexEntries) == 0 {
		t.Errorf("Expected lexical entries for dog synset")
	}

	hyps := s.Hypernyms(res)
	if len(hyps) == 0 {
		t.Errorf("Expected hypernyms for dog synset")
	}

	hypos := s.Hyponyms(res)
	if len(hypos) == 0 {
		t.Errorf("Expected hyponyms for dog synset")
	}

	related := s.RelatedSynsets(res, gown.SynsetRelationTypeHypernym)
	if len(related) == 0 {
		t.Errorf("Expected related hypernym synsets for dog")
	}

	// Senses and examples
	for _, sense := range entries[0].Senses {
		_ = sense.Examples()
		_ = sense.Definitions()
		_ = sense.GetSynset()
	}
}

func TestPOSWrappersAndStringMethods(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Noun wrapper
	nouns := res.Nouns()
	if len(nouns) > 0 {
		n := gown.Noun(nouns[0])
		_ = n.String()
		_ = n.LexicalEntry()
	}

	// Verb wrapper
	verbs := res.Verbs()
	if len(verbs) > 0 {
		v := gown.Verb(verbs[0])
		_ = v.String()
		_ = v.LexicalEntry()
	}

	// Adjective wrapper
	adjs := res.Adjectives()
	if len(adjs) > 0 {
		a := gown.Adjective(adjs[0])
		_ = a.String()
		_ = a.LexicalEntry()
	}

	// Adverb wrapper
	advs := res.Adverbs()
	if len(advs) > 0 {
		adv := gown.Adverb(advs[0])
		_ = adv.String()
		_ = adv.LexicalEntry()
	}
}
