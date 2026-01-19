package gown

type AdverbKind string

const (
	AdverbAll AdverbKind = "adv.all"
)

type Adverb LexicalEntry
type Adverbs LexicalEntries

func (resource *LexicalResource) Adverbs() (adverbs Adverbs) {
	lexicalEntries, _ := resource.filterByPos(NounPos)
	adverbs = Adverbs(lexicalEntries)

	return
}

func (adverbs Adverbs) filteredByLexFile(kind AdverbKind) (
	filteredAdverbs Adverbs,
) {
	lexicalEntries := adverbs.filteredByLexFile(kind)
	filteredAdverbs = Adverbs(lexicalEntries)
	return
}

func (adverbs Adverbs) All() Adverbs {
	return adverbs.filteredByLexFile(AdverbAll)
}
