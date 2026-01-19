package gown

import (
	"slices"
	"strings"
)

type Lemma struct {
	PartOfSpeech   string          `xml:"partOfSpeech,attr"`
	WrittenForm    string          `xml:"writtenForm,attr"`
	Pronunciations []Pronunciation `xml:"Pronunciation"`
}

type LexicalEntry struct {
	ID     string  `xml:"id,attr"`
	Forms  []Form  `xml:"Form"`
	Lemma  Lemma   `xml:"Lemma"`
	Senses []Sense `xml:"Sense"`

	resource *LexicalResource
}

type LexicalEntries []LexicalEntry

func (entries LexicalEntries) filterByPos(pos ...POS) (filteredEntries LexicalEntries) {
	for _, entry := range entries {
		for _, _pos := range pos {
			if entry.PartOfSpeech() == _pos {
				filteredEntries = append(filteredEntries, entry)
			}
		}
	}
	return
}

func (entries LexicalEntries) filterByLexFile(lexFile string) (filteredEntries LexicalEntries) {
	for _, entry := range entries {
		if slices.ContainsFunc(entry.Synsets(),
			func(s *Synset) bool { return s.Lexfile == lexFile }) {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return
}

func (entries LexicalEntries) Nouns() Nouns {
	lexicalEntries := entries.filterByPos(NounPos)
	return Nouns(lexicalEntries)
}

func (entries LexicalEntries) Verbs() Verbs {
	lexicalEntries := entries.filterByPos(VerbPos)
	return Verbs(lexicalEntries)
}

func (entries LexicalEntries) Adjectives() Adjectives {
	lexicalEntries := entries.filterByPos(AdjectivePos, AdjectiveSatellitePos)
	return Adjectives(lexicalEntries)
}

func (entries LexicalEntries) Adverbs() Adverbs {
	lexicalEntries := entries.filterByPos(AdverbPos)
	return Adverbs(lexicalEntries)
}

func (entry *LexicalEntry) PartOfSpeech() POS {
	return POS(entry.Lemma.PartOfSpeech)
}

func (entry *LexicalEntry) String() string {
	return entry.Lemma.WrittenForm
}

func (entry *LexicalEntry) Synsets() (synsets []*Synset) {
	for _, sense := range entry.Senses {
		synsets = append(synsets, sense.GetSynset())
	}
	return
}

func (entry *LexicalEntry) filterBySenseRelationType(relType SenseRelationKind) (filteredEntries LexicalEntries) {

	senseById := entry.resource.SenseById()

	for _, sense := range entry.Senses {
		for _, senseRel := range sense.SenseRelations {
			if senseRel.RelType == string(relType) {
				targetSense := senseById[senseRel.Target]

				filteredEntries = append(filteredEntries, *targetSense.lexicalEntry)

			}
		}
	}

	return
}

func (entry *LexicalEntry) filterBySenseDublinCoreType(dcType DublinCoreRelType) (filteredEntries LexicalEntries) {

	senseById := entry.resource.SenseById()

	for _, sense := range entry.Senses {
		for _, senseRel := range sense.SenseRelations {
			if senseRel.Type == string(dcType) {
				targetSense := senseById[senseRel.Target]

				filteredEntries = append(filteredEntries, *targetSense.lexicalEntry)

			}
		}
	}

	return
}

func (entry *LexicalEntry) filterBySynsetRelation(relTyppe SynsetRelationType) (filteredEntries LexicalEntries) {
	synsetsById := entry.resource.SynsetsById()
	lexicalsById := entry.resource.LexicalsById()

	for _, synset := range entry.Synsets() {
		for _, synsetRel := range synset.SynsetRelations {
			if synsetRel.RelType == string(relTyppe) {
				targetSynset := synsetsById[synsetRel.Target]
				for _, member := range targetSynset.Members {
					if lexEntry, ok := lexicalsById[member]; ok {
						filteredEntries = append(filteredEntries, *lexEntry)
					}
				}

			}

		}
	}

	return
}

func (entry *LexicalEntry) Relation() *LexicalRelation {
	return &LexicalRelation{entry: entry}
}

func (entry *LexicalEntry) StartsWith(s string) bool {
	return strings.HasPrefix(entry.Lemma.WrittenForm, s)
}

func (entry *LexicalEntry) EndsWith(s string) bool {
	return strings.HasSuffix(entry.Lemma.WrittenForm, s)
}

func (entry *LexicalEntry) IsWord() bool {
	return !strings.ContainsFunc(entry.Lemma.WrittenForm, containSeparatedCollocation)
}

func (entry *LexicalEntry) IsCollocation() bool {
	return strings.ContainsFunc(entry.Lemma.WrittenForm, containSeparatedCollocation)
}

func wordToPattern(word string, v string, c string) string {
	var pattern string
	for _, r := range word {
		if slices.Contains([]rune{'a', 'i', 'u', 'e', 'o'}, r) {
			pattern += v
		} else {
			pattern += c
		}
	}
	return pattern
}

func (entry *LexicalEntry) CVPatterns() (patterns []string) {
	if !entry.IsWord() {
		pattern := wordToPattern(entry.Lemma.WrittenForm, "v", "c")
		patterns = append(patterns, pattern)

		return
	}

	words := strings.FieldsFunc(entry.Lemma.WrittenForm,
		func(r rune) bool {
			return slices.Contains(collocationSeparator, r)
		})

	for _, word := range words {
		pattern := wordToPattern(word, "v", "c")
		patterns = append(patterns, pattern)
	}

	return
}

func (entry *LexicalEntry) HasCVPattern(pattern string) bool {
	for _, p := range entry.CVPatterns() {
		if p == pattern {
			return true
		}
	}
	return false
}

func (entry *LexicalEntry) Definitions() (definitions []string) {
	for _, sense := range entry.Senses {
		definitions = append(definitions, sense.Definitions()...)

	}
	return
}

func (entry *LexicalEntry) Examples() (examples []string) {
	for _, sense := range entry.Senses {
		examples = append(examples, sense.Examples()...)

	}
	return
}

func (entry *LexicalEntry) Contains(s string) bool {
	return strings.Contains(entry.Lemma.WrittenForm, s)
}

func (entry *LexicalEntry) HasDefinition(s string) bool {
	return slices.ContainsFunc(entry.Definitions(), func(definition string) bool {
		return strings.Contains(definition, s)
	})
}

func (entry *LexicalEntry) HasExample(s string) bool {
	return slices.ContainsFunc(entry.Examples(), func(example string) bool {
		return strings.Contains(example, s)
	})
}

func (entry *LexicalEntry) HasLength(length int) bool {
	return len(entry.Lemma.WrittenForm) == length
}
