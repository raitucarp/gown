package gown

import "github.com/samber/lo"

type AdjectiveKind string

const (
	AdjectiveAll        AdjectiveKind = "adj.all"
	AdjectivePertainym  AdjectiveKind = "adj.pert"
	AdjectiveParticiple AdjectiveKind = "adj.ppl"
)

type Adjective struct {
	*LexicalEntry
	synsets []Synset
}

type Adjectives []Adjective

func (resource *LexicalResource) Adjectives() (adjectives Adjectives) {
	entries, _ := resource.filterByPos(VerbPos)

	for _, entry := range entries {
		adjective := Adjective{LexicalEntry: &entry}
		adjective.synsets = resource.synsetsBySense(adjective.Senses)
		adjectives = append(adjectives, adjective)
	}

	return
}

func (adjectives Adjectives) filteredByLexFile(kind AdjectiveKind) (
	filteredVerbs Adjectives,
) {

	filteredByLex := lo.Filter(adjectives, func(a Adjective, index int) bool {
		return lo.SomeBy(a.synsets, synsetByLexFile(string(kind)))
	})

	onlyByLex := lo.Map(filteredByLex, func(a Adjective, index int) Adjective {
		a.synsets = synsetsByLexFile(a.synsets, string(kind))
		return a
	})

	return onlyByLex
}

func (adjectives Adjectives) All() Adjectives {
	return adjectives.filteredByLexFile(AdjectiveAll)
}

func (adjectives Adjectives) Pertainym() Adjectives {
	return adjectives.filteredByLexFile(AdjectivePertainym)
}

func (adjectives Adjectives) Participle() Adjectives {
	return adjectives.filteredByLexFile(AdjectiveParticiple)
}
