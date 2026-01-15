package gown

import (
	"slices"
	"strings"
)

type LexicalEntry struct {
	ID     string  `xml:"id,attr"`
	Forms  []Form  `xml:"Form"`
	Lemma  Lemma   `xml:"Lemma"`
	Senses []Sense `xml:"Sense"`

	resource *LexicalResource
}

type LexicalEntries []LexicalEntry

func (entries LexicalEntries) filterByPos(pos POS) (filteredEntries LexicalEntries) {
	for _, entry := range entries {
		if entry.Lemma.PartOfSpeech == string(pos) {
			filteredEntries = append(filteredEntries, entry)
		}
	}
	return
}

func (entries LexicalEntries) filterByLexFile(lexFile string) (filteredEntries LexicalEntries) {
	for _, entry := range entries {
		for _, sense := range entry.Senses {
			if sense.GetSynset().Lexfile == lexFile {
				filteredEntries = append(filteredEntries, entry)
			}
		}
	}
	return
}

func (resource *LexicalResource) filterSynsetsByLexFile(lexFile string) (synsets []Synset) {
	for _, synset := range resource.Lexicon.Synsets {
		if synset.Lexfile == lexFile {
			synsets = append(synsets, synset)
		}
	}

	return
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
