package gown

import (
	"slices"
	"strings"
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
		if strings.ContainsFunc(
			noun.Lemma.WrittenForm,
			containSeparatedCollocation) {
			collocationNouns = append(collocationNouns, noun)
		}
	}
	return
}

func (verbs Verbs) Words() (wordVerbs Verbs) {
	for _, verb := range verbs {
		if !strings.ContainsFunc(
			verb.Lemma.WrittenForm,
			containSeparatedCollocation) {
			wordVerbs = append(wordVerbs, verb)
		}
	}
	return
}

func (verbs Verbs) Collocations() (collocationVerbs Verbs) {
	for _, verb := range verbs {
		if strings.ContainsFunc(
			verb.Lemma.WrittenForm,
			containSeparatedCollocation) {
			collocationVerbs = append(collocationVerbs, verb)
		}
	}
	return
}
