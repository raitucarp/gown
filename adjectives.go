package gown

type AdjectiveKind string

const (
	AdjectiveAll        AdjectiveKind = "adj.all"
	AdjectivePertainym  AdjectiveKind = "adj.pert"
	AdjectiveParticiple AdjectiveKind = "adj.ppl"
)

type Adjective LexicalEntry
type Adjectives LexicalEntries

func (resource *LexicalResource) Adjectives() (adjectives Adjectives) {
	entries := Adjectives(
		LexicalEntries(resource.Lexicon.LexicalEntries).
			filterByPos(AdjectivePos, AdjectiveSatellitePos),
	)

	return entries
}

func (adjectives Adjectives) filteredByLexFile(kind AdjectiveKind) (
	filteredAdjectives Adjectives,
) {
	lexicalEntries := adjectives.filteredByLexFile(kind)
	filteredAdjectives = Adjectives(lexicalEntries)
	return
}

func (adjectives Adjectives) All() Adjectives {
	return adjectives.filteredByLexFile(AdjectiveAll)
}

func (adjectives Adjectives) Pertainym() Adjectives {
	return adjectives.filteredByLexFile(AdjectivePertainym)
}

func (adjectives Adjectives) Participle() Adjectives {
	return adjectives.filteredByLexFile(AdjectiveParticiple)
}
