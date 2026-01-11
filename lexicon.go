package gown

import (
	"encoding/xml"
	"strings"
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

type Lemma struct {
	PartOfSpeech   string          `xml:"partOfSpeech,attr"`
	WrittenForm    string          `xml:"writtenForm,attr"`
	Pronunciations []Pronunciation `xml:"Pronunciation"`
}

type SenseRelation struct {
	RelType string `xml:"relType,attr"`
	Target  string `xml:"target,attr"`
	Type    string `xml:"type,attr"`
}

type Sense struct {
	Adjposition    string          `xml:"adjposition,attr"`
	ID             string          `xml:"id,attr"`
	Subcat         string          `xml:"subcat,attr"`
	Synset         string          `xml:"synset,attr"`
	SenseRelations []SenseRelation `xml:"SenseRelation"`

	resource *LexicalResource
}

func (sense *Sense) Definitions() (definitions []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		definitions = append(definitions, synsetsById[sense.Synset].Definitions...)
	}
	return
}

func (sense *Sense) Examples() (examples []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		for _, example := range synsetsById[sense.Synset].Examples {
			examples = append(examples, example.Text)
		}
	}
	return
}

type LexicalEntry struct {
	ID     string  `xml:"id,attr"`
	Forms  []Form  `xml:"Form"`
	Lemma  Lemma   `xml:"Lemma"`
	Senses []Sense `xml:"Sense"`

	resource *LexicalResource
}

func (entry *LexicalEntry) Definitions() (definitions []string) {
	synsetsById := entry.resource.SynsetsById()
	for _, sense := range entry.Senses {
		if _, ok := synsetsById[sense.Synset]; ok {
			definitions = append(definitions, synsetsById[sense.Synset].Definitions...)
		}
	}
	return
}

func (entry *LexicalEntry) Examples() (examples []string) {
	synsetsById := entry.resource.SynsetsById()
	for _, sense := range entry.Senses {
		if _, ok := synsetsById[sense.Synset]; ok {
			for _, example := range synsetsById[sense.Synset].Examples {
				examples = append(examples, example.Text)
			}
		}
	}
	return
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

	synsetsById  map[string]Synset
	lexicalsById map[string]LexicalEntry
}

type POS string

const (
	NounPos               POS = "n"
	VerbPos               POS = "v"
	AdjectivePos          POS = "a"
	AdverbPos             POS = "r"
	AdjectiveSattelitePos POS = "s"
)

func (resource *LexicalResource) SynsetsById() map[string]Synset {
	if len(resource.synsetsById) > 0 {
		return resource.synsetsById
	}

	resource.synsetsById = make(map[string]Synset)
	for _, synset := range resource.Lexicon.Synsets {
		resource.synsetsById[synset.ID] = synset
	}
	return resource.synsetsById
}

func (resource *LexicalResource) LexicalsById() map[string]LexicalEntry {
	if len(resource.lexicalsById) > 0 {
		return resource.lexicalsById
	}

	resource.lexicalsById = make(map[string]LexicalEntry)
	for _, lexicalEntry := range resource.Lexicon.LexicalEntries {
		resource.lexicalsById[lexicalEntry.ID] = lexicalEntry
	}

	return resource.lexicalsById
}

func (resource *LexicalResource) groupEntryByPos(pos POS) (entries []LexicalEntry) {
	for _, lexicalEntry := range resource.Lexicon.LexicalEntries {
		if lexicalEntry.Lemma.PartOfSpeech == string(pos) {
			entries = append(entries, lexicalEntry)
		}
	}
	return
}
