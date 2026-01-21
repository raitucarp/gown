package gown

import (
	"github.com/samber/lo"
)

// RandomLemma returns n random lexical entries from this resource.
func (resource LexicalResource) RandomLemma(n int) LexicalEntries {
	shuffled := lo.Samples(resource.Lexicon.LexicalEntries, n)
	return shuffled
}

// Random returns n random entries from this collection.
func (entries LexicalEntries) Random(n int) LexicalEntries {
	shuffled := lo.Samples(entries, n)
	return shuffled
}

// Random returns n random nouns from this collection.
func (nouns Nouns) Random(n int) Nouns {
	shuffled := lo.Samples(nouns, n)
	return shuffled
}

// Random returns n random verbs from this collection.
func (verbs Verbs) Random(n int) Verbs {
	shuffled := lo.Samples(verbs, n)
	return shuffled
}

// Random returns n random adjectives from this collection.
func (adjectives Adjectives) Random(n int) Adjectives {
	shuffled := lo.Samples(adjectives, n)
	return shuffled
}

// Random returns n random adverbs from this collection.
func (adverbs Adverbs) Random(n int) Adverbs {
	shuffled := lo.Samples(adverbs, n)
	return shuffled
}
