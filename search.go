package gown

func (resource LexicalResource) SearchLemma(query string) (entries []LexicalEntry) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.Contains(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (resource LexicalResource) SearchLemmaByDefinition(query string) (entries []LexicalEntry) {
	for _, entry := range resource.Lexicon.LexicalEntries {
		if entry.HasDefinition(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (nouns Nouns) SearchLemma(query string) (entries []Noun) {
	for _, entry := range nouns {
		if entry.Contains(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (nouns Nouns) SearchLemmaByDefinition(query string) (entries []Noun) {
	for _, entry := range nouns {
		if entry.HasDefinition(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (verbs Verbs) SearchLemma(query string) (entries Verbs) {
	for _, entry := range verbs {
		if entry.Contains(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (verbs Verbs) SearchLemmaByDefinition(query string) (entries Verbs) {
	for _, entry := range verbs {
		if entry.HasDefinition(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (adjectives Adjectives) SearchLemma(query string) (entries Adjectives) {
	for _, entry := range adjectives {
		if entry.Contains(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (adjectives Adjectives) SearchLemmaByDefinition(query string) (entries Adjectives) {
	for _, entry := range adjectives {
		if entry.HasDefinition(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (adverbs Adverbs) SearchLemma(query string) (entries Adverbs) {
	for _, entry := range adverbs {
		if entry.Contains(query) {
			entries = append(entries, entry)
		}
	}

	return
}

func (adverbs Adverbs) SearchLemmaByDefinition(query string) (entries Adverbs) {
	for _, entry := range adverbs {
		if entry.HasDefinition(query) {
			entries = append(entries, entry)
		}
	}

	return
}
