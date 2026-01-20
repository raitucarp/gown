package gown

type AdverbKind string

const (
	AdverbAll AdverbKind = "adv.all"
)

type Adverb LexicalEntry
type Adverbs LexicalEntries

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

func (adverbs Adverbs) LexicalEntries() LexicalEntries {
	return LexicalEntries(adverbs)
}

func (adverb Adverb) LexicalEntry() LexicalEntry {
	return LexicalEntry(adverb)
}

func (adverb Adverb) String() string {
	return adverb.Lemma.WrittenForm
}

func (adverbs Adverbs) All() Adverbs {
	return adverbs.filteredByLexFile(AdverbAll)
}
