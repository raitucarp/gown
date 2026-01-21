package gown

import (
	"slices"
)

// containSeparatedCollocation checks if a rune is a collocation separator.
var collocationSeparator = []rune{' ', '-', '!', '/', '+', ':', ','}

// containSeparatedCollocation returns true if the rune is a collocation separator.
func containSeparatedCollocation(r rune) bool {
	return slices.Contains(collocationSeparator, r)
}

// Words returns all entries that are single words (not collocations) from this resource.
func (resource LexicalResource) Words() (entries LexicalEntries) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.IsWord() {
			entries = append(entries, entry)
		}
	}
	return
}

// Collocations returns all entries that are multi-word collocations from this resource.
func (resource LexicalResource) Collocations() (entries LexicalEntries) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.IsCollocation() {
			entries = append(entries, entry)
		}
	}
	return
}

// Words returns all nouns that are single words (not collocations) from this collection.
func (nouns Nouns) Words() (wordNouns Nouns) {
	for _, noun := range nouns {
		if noun.IsWord() {
			wordNouns = append(wordNouns, noun)
		}
	}
	return
}

// Collocations returns all nouns that are multi-word collocations from this collection.
func (nouns Nouns) Collocations() (collocationNouns Nouns) {
	for _, noun := range nouns {
		if noun.IsCollocation() {
			collocationNouns = append(collocationNouns, noun)
		}
	}
	return
}

// Words returns all verbs that are single words (not collocations) from this collection.
func (verbs Verbs) Words() (wordVerbs Verbs) {
	for _, verb := range verbs {
		if verb.IsWord() {
			wordVerbs = append(wordVerbs, verb)
		}
	}
	return
}

// Collocations returns all verbs that are multi-word collocations from this collection.
func (verbs Verbs) Collocations() (collocationVerbs Verbs) {
	for _, verb := range verbs {
		if verb.IsCollocation() {
			collocationVerbs = append(collocationVerbs, verb)
		}
	}
	return
}

// Words returns all adverbs that are single words (not collocations) from this collection.
func (adverbs Adverbs) Words() (wordAdverbs Adverbs) {
	for _, adverb := range adverbs {
		if adverb.IsWord() {
			wordAdverbs = append(wordAdverbs, adverb)
		}
	}
	return
}

// Collocations returns all adverbs that are multi-word collocations from this collection.
func (adverbs Adverbs) Collocations() (collocationAdverbs Adverbs) {
	for _, adverb := range adverbs {
		if adverb.IsCollocation() {
			collocationAdverbs = append(collocationAdverbs, adverb)
		}
	}
	return
}

// Words returns all adjectives that are single words (not collocations) from this collection.
func (adjectives Adjectives) Words() (wordAdjectives Adjectives) {
	for _, adjective := range adjectives {
		if adjective.IsWord() {
			wordAdjectives = append(wordAdjectives, adjective)
		}
	}
	return
}

// Collocations returns all adjectives that are multi-word collocations from this collection.
func (adjectives Adjectives) Collocations() (collocationAdjectives Adjectives) {
	for _, adjective := range adjectives {
		if adjective.IsCollocation() {
			collocationAdjectives = append(collocationAdjectives, adjective)
		}
	}
	return
}
