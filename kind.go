package gown

import (
	"slices"
)

var collocationSeparator = []rune{' ', '-', '!', '/', '+', ':', ','}

func containSeparatedCollocation(r rune) bool {
	return slices.Contains(collocationSeparator, r)
}

func (resource LexicalResource) Words() (entries LexicalEntries) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.IsWord() {
			entries = append(entries, entry)
		}
	}
	return
}

func (resource LexicalResource) Collocations() (entries LexicalEntries) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.IsCollocation() {
			entries = append(entries, entry)
		}
	}
	return
}

func (nouns Nouns) Words() (wordNouns Nouns) {
	for _, noun := range nouns {
		if noun.IsWord() {
			wordNouns = append(wordNouns, noun)
		}
	}
	return
}

func (nouns Nouns) Collocations() (collocationNouns Nouns) {
	for _, noun := range nouns {
		if noun.IsCollocation() {
			collocationNouns = append(collocationNouns, noun)
		}
	}
	return
}

func (verbs Verbs) Words() (wordVerbs Verbs) {
	for _, verb := range verbs {
		if verb.IsWord() {
			wordVerbs = append(wordVerbs, verb)
		}
	}
	return
}

func (verbs Verbs) Collocations() (collocationVerbs Verbs) {
	for _, verb := range verbs {
		if verb.IsCollocation() {
			collocationVerbs = append(collocationVerbs, verb)
		}
	}
	return
}

func (adverbs Adverbs) Words() (wordAdverbs Adverbs) {
	for _, adverb := range adverbs {
		if adverb.IsWord() {
			wordAdverbs = append(wordAdverbs, adverb)
		}
	}
	return
}

func (adverbs Adverbs) Collocations() (collocationAdverbs Adverbs) {
	for _, adverb := range adverbs {
		if adverb.IsCollocation() {
			collocationAdverbs = append(collocationAdverbs, adverb)
		}
	}
	return
}

func (adjectives Adjectives) Words() (wordAdjectives Adjectives) {
	for _, adjective := range adjectives {
		if adjective.IsWord() {
			wordAdjectives = append(wordAdjectives, adjective)
		}
	}
	return
}

func (adjectives Adjectives) Collocations() (collocationAdjectives Adjectives) {
	for _, adjective := range adjectives {
		if adjective.IsCollocation() {
			collocationAdjectives = append(collocationAdjectives, adjective)
		}
	}
	return
}
