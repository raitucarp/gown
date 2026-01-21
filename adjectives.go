package gown

// AdjectiveKind represents a semantic category of adjectives based on WordNet's lexfile classification.
type AdjectiveKind string

const (
	// AdjectiveAll is the semantic category for general adjectives.
	AdjectiveAll AdjectiveKind = "adj.all"
	// AdjectivePertainym is the semantic category for pertainym adjectives.
	AdjectivePertainym AdjectiveKind = "adj.pert"
	// AdjectiveParticiple is the semantic category for participle adjectives.
	AdjectiveParticiple AdjectiveKind = "adj.ppl"
)

// Adjective represents a single adjective lexical entry.
type Adjective LexicalEntry

// Adjectives represents a collection of adjectives.
type Adjectives LexicalEntries

// Adjectives returns all adjectives from this lexical resource.
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

// LexicalEntries converts this Adjectives collection to a LexicalEntries collection.
func (adjectives Adjectives) LexicalEntries() LexicalEntries {
	return LexicalEntries(adjectives)
}

// LexicalEntry converts this Adjective to a LexicalEntry.
func (adjective Adjective) LexicalEntry() LexicalEntry {
	return LexicalEntry(adjective)
}

// String returns the written form of this adjective.
func (adjective Adjective) String() string {
	return adjective.Lemma.WrittenForm
}

// AllKind returns a map of all adjective kinds with their corresponding adjective collections.
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
