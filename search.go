package gown

// SearchLemma searches for lexical entries whose written form contains the query string.
func (resource LexicalResource) SearchLemma(query string) (entries []LexicalEntry) {
	lexicalEntries := LexicalEntries(resource.Lexicon.LexicalEntries)

	return lexicalEntries.SearchLemma(query)
}

// SearchLemma searches for entries whose written form contains the query string.
func (entries LexicalEntries) SearchLemma(query string) (filteredEntries []LexicalEntry) {
	for _, entry := range entries {
		if entry.Contains(query) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return
}

// SearchLemmaByDefinition searches for entries whose synset definitions contain the query string.
func (entries LexicalEntries) SearchLemmaByDefinition(query string) (filteredEntries []LexicalEntry) {
	for _, entry := range entries {
		if entry.HasDefinition(query) {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	return
}

// SearchLemmaByDefinition searches for entries whose synset definitions contain the query string.
func (resource LexicalResource) SearchLemmaByDefinition(query string) (entries []LexicalEntry) {
	lexicalEntries := LexicalEntries(resource.Lexicon.LexicalEntries)

	return lexicalEntries.SearchLemmaByDefinition(query)
}

// SearchLemma searches for nouns whose written form contains the query string.
func (nouns Nouns) SearchLemma(query string) (filteredNouns Nouns) {
	lexicalEntries := LexicalEntries(nouns)

	return Nouns(lexicalEntries.SearchLemma(query))
}

// SearchLemmaByDefinition searches for nouns whose synset definitions contain the query string.
func (nouns Nouns) SearchLemmaByDefinition(query string) (filteredNouns Nouns) {
	lexicalEntries := LexicalEntries(nouns)

	return Nouns(lexicalEntries.SearchLemmaByDefinition(query))
}

// SearchLemma searches for verbs whose written form contains the query string.
func (verbs Verbs) SearchLemma(query string) (filteredVerbs Verbs) {
	lexicalEntries := LexicalEntries(verbs)

	return Verbs(lexicalEntries.SearchLemma(query))
}

// SearchLemmaByDefinition searches for verbs whose synset definitions contain the query string.
func (verbs Verbs) SearchLemmaByDefinition(query string) (filteredVerbs Verbs) {
	lexicalEntries := LexicalEntries(verbs)

	return Verbs(lexicalEntries.SearchLemmaByDefinition(query))
}

// SearchLemma searches for adverbs whose written form contains the query string.
func (adverbs Adverbs) SearchLemma(query string) (filteredAdverbs Adverbs) {
	lexicalEntries := LexicalEntries(adverbs)

	return Adverbs(lexicalEntries.SearchLemma(query))
}

// SearchLemmaByDefinition searches for adverbs whose synset definitions contain the query string.
func (adverbs Adverbs) SearchLemmaByDefinition(query string) (filteredAdverbs Adverbs) {
	lexicalEntries := LexicalEntries(adverbs)

	return Adverbs(lexicalEntries.SearchLemmaByDefinition(query))
}

// SearchLemma searches for adjectives whose written form contains the query string.
func (adjectives Adjectives) SearchLemma(query string) (filteredAdjectives Adjectives) {
	lexicalEntries := LexicalEntries(adjectives)

	return Adjectives(lexicalEntries.SearchLemma(query))
}

// SearchLemmaByDefinition searches for adjectives whose synset definitions contain the query string.
func (adjectives Adjectives) SearchLemmaByDefinition(query string) (filteredAdjectives Adjectives) {
	lexicalEntries := LexicalEntries(adjectives)

	return Adjectives(lexicalEntries.SearchLemmaByDefinition(query))
}
