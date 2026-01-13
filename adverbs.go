package gown

import "github.com/samber/lo"

type AdverbKind string

const (
	AdverbAll AdverbKind = "adv.all"
)

type Adverb struct {
	*LexicalEntry
	synsets []Synset
}

type Adverbs []Adverb

func (resource *LexicalResource) Adverbs() (adverbs Adverbs) {
	entries, _ := resource.filterByPos(VerbPos)

	for _, entry := range entries {
		adverb := Adverb{LexicalEntry: &entry}
		adverb.synsets = resource.synsetsBySense(adverb.Senses)
		adverbs = append(adverbs, adverb)
	}

	return
}

func (adverbs Adverbs) filteredByLexFile(kind AdverbKind) (
	filteredVerbs Adverbs,
) {

	filteredByLex := lo.Filter(adverbs, func(r Adverb, index int) bool {
		return lo.SomeBy(r.synsets, synsetByLexFile(string(kind)))
	})

	onlyByLex := lo.Map(filteredByLex, func(r Adverb, index int) Adverb {
		r.synsets = synsetsByLexFile(r.synsets, string(kind))
		return r
	})

	return onlyByLex
}

func (adverbs Adverbs) All() Adverbs {
	return adverbs.filteredByLexFile(AdverbAll)
}
