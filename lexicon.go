package gown

import (
	"encoding/xml"
	"strings"

	"github.com/samber/lo"
)

// SyntacticBehaviour represents a syntactic behaviour frame for lexical entries.
type SyntacticBehaviour struct {
	// ID uniquely identifies this syntactic behaviour.
	ID string `xml:"id,attr"`
	// SubcategorizationFrame describes the syntactic structure of the behaviour.
	SubcategorizationFrame string `xml:"subcategorizationFrame,attr"`
}

// Form represents an alternative written form of a lexical entry.
type Form struct {
	// WrittenForm is the textual representation of the form.
	WrittenForm string `xml:"writtenForm,attr"`
}

// Pronunciation represents a pronunciation variant of a word.
type Pronunciation struct {
	// Variety indicates the regional or linguistic variety (e.g., "en_US", "en_GB").
	Variety string `xml:"variety,attr"`
	// Text is the phonetic representation of the pronunciation.
	Text string `xml:",chardata"`
}

// SenseRelation represents a semantic relationship between two senses.
type SenseRelation struct {
	// RelType specifies the type of sense relation (e.g., "antonym", "derivation", "similar").
	RelType string `xml:"relType,attr"`
	// Target is the ID of the target sense.
	Target string `xml:"target,attr"`
	// Type specifies the Dublin Core type of the relation.
	Type string `xml:"type,attr"`
}

// Example represents a usage example for a synset.
type Example struct {
	// Source indicates the source of the example.
	Source string `xml:"source,attr"`
	// Text is the example sentence or phrase.
	Text string `xml:",chardata"`
}

// SynsetRelation represents a semantic relationship between two synsets.
type SynsetRelation struct {
	// RelType specifies the type of synset relation (e.g., "hypernym", "hyponym", "meronym").
	RelType string `xml:"relType,attr"`
	// Target is the ID of the target synset.
	Target string `xml:"target,attr"`
}

// Synset represents a set of synonymous words that share a common concept.
type Synset struct {
	// ID uniquely identifies this synset.
	ID string `xml:"id,attr"`
	// Ili is the Interlingual Index linking this synset to other languages.
	Ili string `xml:"ili,attr"`
	// Lexfile categorizes the synset by semantic field (e.g., "noun.person", "verb.motion").
	Lexfile string `xml:"lexfile,attr"`
	// Members is a list of lexical entry IDs that are members of this synset.
	Members []string `xml:"members,attr"`
	// PartOfSpeech is the part of speech tag (n, v, a, r, s).
	PartOfSpeech string `xml:"partOfSpeech,attr"`
	// Source indicates the data source of this synset.
	Source string `xml:"source,attr"`
	// Definitions contains one or more definitions of this concept.
	Definitions []string `xml:"Definition"`
	// Examples provides usage examples for this synset.
	Examples []Example `xml:"Example"`
	// ILIDefinition is the definition shared across all languages for this ILI.
	ILIDefinition string `xml:"ILIDefinition"`
	// SynsetRelations lists semantic relationships to other synsets.
	SynsetRelations []SynsetRelation `xml:"SynsetRelation"`
}

func (m *Synset) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type s Synset
	var aux struct {
		Members string `xml:"members,attr"`
		*s
	}
	aux.s = (*s)(m)

	if err := d.DecodeElement(&aux, &start); err != nil {
		return err
	}

	m.Members = strings.Fields(aux.Members)
	return nil
}

// Lexicon represents a complete WordNet lexicon including entries and synsets.
type Lexicon struct {
	// Email is the contact email for the lexicon maintainer.
	Email string `xml:"email,attr"`
	// ID uniquely identifies this lexicon.
	ID string `xml:"id,attr"`
	// Label is a human-readable label for the lexicon.
	Label string `xml:"label,attr"`
	// Language is the ISO language code (e.g., "en" for English).
	Language string `xml:"language,attr"`
	// License specifies the license under which the lexicon is distributed.
	License string `xml:"license,attr"`
	// Url is the URL to the lexicon's homepage or repository.
	Url string `xml:"url,attr"`
	// Version indicates the version number of this lexicon.
	Version int `xml:"version,attr"`
	// LexicalEntries contains all lexical entries (words) in the lexicon.
	LexicalEntries []LexicalEntry `xml:"LexicalEntry"`
	// Synsets contains all synsets (concept groups) in the lexicon.
	Synsets []Synset `xml:"Synset"`
	// SyntacticBehaviours lists syntactic behaviour frames for verbs.
	SyntacticBehaviours []SyntacticBehaviour `xml:"SyntacticBehaviour"`
}

// LexicalResource represents the complete WordNet database with all lexicons and indices.
type LexicalResource struct {
	// Lexicon contains the main lexicon data.
	Lexicon Lexicon `xml:"Lexicon"`

	// synsetsById is an internal cache mapping synset IDs to synset objects.
	synsetsById map[string]*Synset
	// lexicalsById is an internal cache mapping lexical entry IDs to their objects.
	lexicalsById map[string]*LexicalEntry
	// senseById is an internal cache mapping sense IDs to their objects.
	senseById map[string]*Sense
}

// POS represents a Part of Speech tag.
type POS string

const (
	// NounPos is the part-of-speech tag for nouns.
	NounPos POS = "n"
	// VerbPos is the part-of-speech tag for verbs.
	VerbPos POS = "v"
	// AdjectivePos is the part-of-speech tag for adjectives.
	AdjectivePos POS = "a"
	// AdverbPos is the part-of-speech tag for adverbs.
	AdverbPos POS = "r"
	// AdjectiveSatellitePos is the part-of-speech tag for satellite adjectives.
	AdjectiveSatellitePos POS = "s"
)

// SynsetsById returns a cached map of all synsets indexed by their IDs.
func (resource *LexicalResource) SynsetsById() map[string]*Synset {
	if len(resource.synsetsById) > 0 {
		return resource.synsetsById
	}

	resource.synsetsById = make(map[string]*Synset)
	for _, synset := range resource.Lexicon.Synsets {
		resource.synsetsById[synset.ID] = &synset
	}
	return resource.synsetsById
}

// LexicalsById returns a cached map of all lexical entries indexed by their IDs.
func (resource *LexicalResource) LexicalsById() map[string]*LexicalEntry {
	if len(resource.lexicalsById) > 0 {
		return resource.lexicalsById
	}

	resource.lexicalsById = make(map[string]*LexicalEntry)
	for _, lexicalEntry := range resource.Lexicon.LexicalEntries {
		resource.lexicalsById[lexicalEntry.ID] = &lexicalEntry
	}

	return resource.lexicalsById
}

// SenseById returns a cached map of all senses indexed by their IDs.
func (resource *LexicalResource) SenseById() map[string]*Sense {
	if len(resource.lexicalsById) > 0 {
		return resource.senseById
	}

	resource.senseById = make(map[string]*Sense)
	for _, lexicalEntry := range resource.Lexicon.LexicalEntries {
		for _, sense := range lexicalEntry.Senses {
			resource.senseById[sense.ID] = &sense
		}
	}

	return resource.senseById
}

func (resource *LexicalResource) groupEntryByPos(pos POS) (
	entries []LexicalEntry,
) {
	return lo.GroupBy(resource.Lexicon.LexicalEntries,
		func(entry LexicalEntry) string {
			return entry.Lemma.PartOfSpeech
		})[string(pos)]
}

func (resource *LexicalResource) groupSynsetsByPos(pos POS) (
	entries []Synset,
) {
	return lo.GroupBy(resource.Lexicon.Synsets,
		func(synset Synset) string {
			return synset.PartOfSpeech
		})[string(pos)]
}

func (resource *LexicalResource) synsetsBySense(senses []Sense) []*Synset {
	synsetsById := resource.SynsetsById()
	return lo.Map(senses, func(sense Sense, i int) *Synset {
		return synsetsById[sense.Synset]
	})
}
