package gown

import (
	"github.com/samber/lo"
)

func (resource *LexicalResource) RandomLemma(n int) []LexicalEntry {
	shuffled := lo.Samples(resource.Lexicon.LexicalEntries, n)
	return shuffled
}

func (nouns Nouns) Random(n int) Nouns {
	shuffled := lo.Samples(nouns, n)
	return shuffled
}

func (verbs Verbs) Random(n int) Verbs {
	shuffled := lo.Samples(verbs, n)
	return shuffled
}
