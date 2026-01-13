package gown

import "github.com/samber/lo"

func (resource *LexicalResource) filterByPos(pos POS) (entries []LexicalEntry, synsets []Synset) {
	entries = resource.groupEntryByPos(pos)
	synsets = resource.groupSynsetsByPos(pos)

	return
}

func (resource *LexicalResource) filterSynsetsByLexFile(lexFile string) (synsets []Synset) {
	for _, synset := range resource.Lexicon.Synsets {
		if synset.Lexfile == lexFile {
			synsets = append(synsets, synset)
		}
	}

	return
}

func synsetByLexFile(lexFile string) func(s Synset) bool {
	return func(s Synset) bool {
		return s.Lexfile == lexFile
	}
}

func filterSynsetByLexFile(lexFile string) func(s Synset, index int) bool {
	return func(s Synset, index int) bool {
		return s.Lexfile == lexFile
	}
}

func synsetsByLexFile(synsets []Synset, lexFile string) []Synset {
	return lo.Filter(synsets, filterSynsetByLexFile(lexFile))
}
