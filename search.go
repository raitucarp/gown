package gown

func (resource LexicalResource) SearchLemma(query string) (entries []LexicalEntry) {
	lexicalEntries := LexicalEntries(resource.Lexicon.LexicalEntries)

	return lexicalEntries.SearchLemma(query)
}

func (entries LexicalEntries) SearchLemma(query string) (filteredEntries []LexicalEntry) {
	for _, entry := range entries {
		if entry.Contains(query) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return
}

func (entries LexicalEntries) SearchLemmaByDefinition(query string) (filteredEntries []LexicalEntry) {
	for _, entry := range entries {
		if entry.HasDefinition(query) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return
}

func (resource LexicalResource) SearchLemmaByDefinition(query string) (entries []LexicalEntry) {
	lexicalEntries := LexicalEntries(resource.Lexicon.LexicalEntries)

	return lexicalEntries.SearchLemmaByDefinition(query)
}

func (nouns Nouns) SearchLemma(query string) (filteredNouns Nouns) {
	lexicalEntries := LexicalEntries(nouns)

	return Nouns(lexicalEntries.SearchLemma(query))
}

func (nouns Nouns) SearchLemmaByDefinition(query string) (filteredNouns Nouns) {
	lexicalEntries := LexicalEntries(nouns)

	return Nouns(lexicalEntries.SearchLemmaByDefinition(query))
}

func (verbs Verbs) SearchLemma(query string) (filteredVerbs Verbs) {
	lexicalEntries := LexicalEntries(verbs)

	return Verbs(lexicalEntries.SearchLemma(query))
}

func (verbs Verbs) SearchLemmaByDefinition(query string) (filteredVerbs Verbs) {
	lexicalEntries := LexicalEntries(verbs)

	return Verbs(lexicalEntries.SearchLemmaByDefinition(query))
}

func (adverbs Adverbs) SearchLemma(query string) (filteredAdverbs Adverbs) {
	lexicalEntries := LexicalEntries(adverbs)

	return Adverbs(lexicalEntries.SearchLemma(query))
}

func (adverbs Adverbs) SearchLemmaByDefinition(query string) (filteredAdverbs Adverbs) {
	lexicalEntries := LexicalEntries(adverbs)

	return Adverbs(lexicalEntries.SearchLemmaByDefinition(query))
}

func (adjectives Adjectives) SearchLemma(query string) (filteredAdjectives Adjectives) {
	lexicalEntries := LexicalEntries(adjectives)

	return Adjectives(lexicalEntries.SearchLemma(query))
}

func (adjectives Adjectives) SearchLemmaByDefinition(query string) (filteredAdjectives Adjectives) {
	lexicalEntries := LexicalEntries(adjectives)

	return Adjectives(lexicalEntries.SearchLemmaByDefinition(query))
}
