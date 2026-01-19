package gown

import (
	"encoding/xml"
	"strings"

	"github.com/samber/lo"
)

type SyntacticBehaviour struct {
	ID                     string `xml:"id,attr"`
	SubcategorizationFrame string `xml:"subcategorizationFrame,attr"`
}

type Form struct {
	WrittenForm string `xml:"writtenForm,attr"`
}

type Pronunciation struct {
	Variety string `xml:"variety,attr"`
	Text    string `xml:",chardata"`
}

type SenseRelation struct {
	RelType string `xml:"relType,attr"`
	Target  string `xml:"target,attr"`
	Type    string `xml:"type,attr"`
}

type Example struct {
	Source string `xml:"source,attr"`
	Text   string `xml:",chardata"`
}

type SynsetRelation struct {
	RelType string `xml:"relType,attr"`
	Target  string `xml:"target,attr"`
}

type Synset struct {
	ID              string           `xml:"id,attr"`
	Ili             string           `xml:"ili,attr"`
	Lexfile         string           `xml:"lexfile,attr"`
	Members         []string         `xml:"members,attr"`
	PartOfSpeech    string           `xml:"partOfSpeech,attr"`
	Source          string           `xml:"source,attr"`
	Definitions     []string         `xml:"Definition"`
	Examples        []Example        `xml:"Example"`
	ILIDefinition   string           `xml:"ILIDefinition"`
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

type Lexicon struct {
	Email               string               `xml:"email,attr"`
	ID                  string               `xml:"id,attr"`
	Label               string               `xml:"label,attr"`
	Language            string               `xml:"language,attr"`
	License             string               `xml:"license,attr"`
	Url                 string               `xml:"url,attr"`
	Version             int                  `xml:"version,attr"`
	LexicalEntries      []LexicalEntry       `xml:"LexicalEntry"`
	Synsets             []Synset             `xml:"Synset"`
	SyntacticBehaviours []SyntacticBehaviour `xml:"SyntacticBehaviour"`
}

type LexicalResource struct {
	Lexicon Lexicon `xml:"Lexicon"`

	synsetsById  map[string]*Synset
	lexicalsById map[string]*LexicalEntry
	senseById    map[string]*Sense
}

type POS string

const (
	NounPos               POS = "n"
	VerbPos               POS = "v"
	AdjectivePos          POS = "a"
	AdverbPos             POS = "r"
	AdjectiveSatellitePos POS = "s"
)

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
