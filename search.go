package gown

import "strings"

func (resource LexicalResource) SearchLemma(query string) (entries []LexicalEntry) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if strings.Contains(entry.Lemma.WrittenForm, query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (resource LexicalResource) SearchLemmaByDefinition(query string) (entries []LexicalEntry) {
	lexById := resource.LexicalsById()
	for _, synset := range resource.Lexicon.Synsets {
		for _, definition := range synset.Definitions {
			if !strings.Contains(definition, query) {
				continue
			}

			for _, member := range synset.Members {
				if _, ok := lexById[member]; ok {
					entries = append(entries, lexById[member])
				}
			}
		}
	}

	return
}

func (nouns Nouns) SearchLemma(query string) (entries []Noun) {
	for _, entry := range nouns {
		if strings.Contains(entry.Lemma.WrittenForm, query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (nouns Nouns) SearchLemmaByDefinition(query string) (entries []Noun) {
	for _, entry := range nouns {
		for _, definition := range entry.Definitions() {
			if strings.Contains(definition, query) {
				entries = append(entries, entry)
			}
		}
	}

	return
}

func (verbs Verbs) SearchLemma(query string) (entries Verbs) {
	for _, entry := range verbs {
		if strings.Contains(entry.Lemma.WrittenForm, query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (verbs Verbs) SearchLemmaByDefinition(query string) (entries Verbs) {
	for _, entry := range verbs {
		for _, definition := range entry.Definitions() {
			if strings.Contains(definition, query) {
				entries = append(entries, entry)
			}
		}
	}

	return
}
