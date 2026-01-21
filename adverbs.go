package gown

// AdverbKind represents a semantic category of adverbs based on WordNet's lexfile classification.
type AdverbKind string

const (
	// AdverbAll is the semantic category for adverbs.
	AdverbAll AdverbKind = "adv.all"
)

// Adverb represents a single adverb lexical entry.
type Adverb LexicalEntry

// Adverbs represents a collection of adverbs.
type Adverbs LexicalEntries

// Adverbs returns all adverbs from this lexical resource.
func (resource *LexicalResource) Adverbs() (adverbs Adverbs) {
	entries := Adverbs(
		LexicalEntries(resource.Lexicon.LexicalEntries).
			filterByPos(AdverbPos),
	)

	return entries
}

func (adverbs Adverbs) filteredByLexFile(kind AdverbKind) (
	filteredAdverbs Adverbs,
) {
	lexicalEntries := adverbs.LexicalEntries().
		filterByLexFile(string(kind))
	filteredAdverbs = Adverbs(lexicalEntries)
	return
}

// LexicalEntries converts this Adverbs collection to a LexicalEntries collection.
func (adverbs Adverbs) LexicalEntries() LexicalEntries {
	return LexicalEntries(adverbs)
}

// LexicalEntry converts this Adverb to a LexicalEntry.
func (adverb Adverb) LexicalEntry() LexicalEntry {
	return LexicalEntry(adverb)
}

// String returns the written form of this adverb.
func (adverb Adverb) String() string {
	return adverb.Lemma.WrittenForm
}

func (adverbs Adverbs) All() Adverbs {
	return adverbs.filteredByLexFile(AdverbAll)
}
