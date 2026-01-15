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
