package gown

import "github.com/samber/lo"

func (resource *LexicalResource) FilterByPos(pos POS) (entries []LexicalEntry, synsets []Synset) {
	entries = resource.GroupEntryByPos(pos)
	synsets = resource.GroupSynsetsByPos(pos)

	return
}

func (resource *LexicalResource) FilterSynsetsByLexFile(lexFile string) (synsets []Synset) {
	for _, synset := range resource.Lexicon.Synsets {
		if synset.Lexfile == lexFile {
			synsets = append(synsets, synset)
		}
	}

	return
}

func SynsetByLexFile(lexFile string) func(s Synset) bool {
	return func(s Synset) bool {
		return s.Lexfile == lexFile
	}
}

func FilterSynsetByLexFile(lexFile string) func(s Synset, index int) bool {
	return func(s Synset, index int) bool {
		return s.Lexfile == lexFile
	}
}

func SynsetsByLexFile(synsets []Synset, lexFile string) []Synset {
	return lo.Filter(synsets, FilterSynsetByLexFile(lexFile))
}
