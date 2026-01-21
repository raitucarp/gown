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
	lexicalEntries := adjectives.LexicalEntries().
		filterByLexFile(string(kind))
	filteredAdjectives = Adjectives(lexicalEntries)
	return
}

func (adjectives Adjectives) LexicalEntries() LexicalEntries {
	return LexicalEntries(adjectives)
}

func (adjective Adjective) LexicalEntry() LexicalEntry {
	return LexicalEntry(adjective)
}

func (adjective Adjective) String() string {
	return adjective.Lemma.WrittenForm
}

func (adjectives Adjectives) AllKind() map[AdjectiveKind]Adjectives {
	return map[AdjectiveKind]Adjectives{
		AdjectiveAll:        adjectives.All(),
		AdjectivePertainym:  adjectives.Pertainym(),
		AdjectiveParticiple: adjectives.Participle(),
	}
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
