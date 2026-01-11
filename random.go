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
