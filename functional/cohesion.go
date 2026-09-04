package functional

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/text"
)

// CohesionType classifies the lexical relation connecting two text words.
type CohesionType string

const (
	CohesionRepetition CohesionType = "repetition"
	CohesionSynonymy   CohesionType = "synonymy"
	CohesionHypernymy  CohesionType = "hypernymy"
	CohesionHyponymy   CohesionType = "hyponymy"
	CohesionMeronymy   CohesionType = "meronymy"
	CohesionAntonymy   CohesionType = "antonymy"
)

// CohesiveTie represents a cohesive link between two words in a discourse.
type CohesiveTie struct {
	Word1 string       `json:"word1"`
	Word2 string       `json:"word2"`
	Type  CohesionType `json:"type"`
}

// AnalyzeCohesion discovers lexical cohesion ties between content words across a text using WordNet.
func AnalyzeCohesion(res *gown.LexicalResource, textContent string) []CohesiveTie {
	tokens := text.ExtractContentWords(textContent)
	if len(tokens) < 2 {
		return nil
	}

	var ties []CohesiveTie
	seen := make(map[string]bool)

	for i := 0; i < len(tokens); i++ {
		w1 := tokens[i]
		for j := i + 1; j < len(tokens); j++ {
			w2 := tokens[j]
			pairKey := w1 + ":" + w2
			if seen[pairKey] {
				continue
			}

			// 1. Repetition
			if w1 == w2 {
				seen[pairKey] = true
				ties = append(ties, CohesiveTie{Word1: w1, Word2: w2, Type: CohesionRepetition})
				continue
			}

			// 2. WordNet relation check
			entries1 := res.Lookup(w1)
			entries2 := res.Lookup(w2)
			if len(entries1) == 0 || len(entries2) == 0 {
				continue
			}

			found := false

			// Check synonyms
			for _, e1 := range entries1 {
				for _, syn := range e1.Relation().Synonyms() {
					if strings.EqualFold(syn.Lemma.WrittenForm, w2) {
						ties = append(ties, CohesiveTie{Word1: w1, Word2: w2, Type: CohesionSynonymy})
						seen[pairKey] = true
						found = true
						break
					}
				}
				if found {
					break
				}

				// Check hypernyms
				for _, hyp := range e1.Relation().Hypernyms() {
					if strings.EqualFold(hyp.Lemma.WrittenForm, w2) {
						ties = append(ties, CohesiveTie{Word1: w1, Word2: w2, Type: CohesionHypernymy})
						seen[pairKey] = true
						found = true
						break
					}
				}
				if found {
					break
				}

				// Check hyponyms
				for _, hyp := range e1.Relation().Hyponyms() {
					if strings.EqualFold(hyp.Lemma.WrittenForm, w2) {
						ties = append(ties, CohesiveTie{Word1: w1, Word2: w2, Type: CohesionHyponymy})
						seen[pairKey] = true
						found = true
						break
					}
				}
				if found {
					break
				}

				// Check antonyms
				for _, ant := range e1.Relation().Antonyms() {
					if strings.EqualFold(ant.Lemma.WrittenForm, w2) {
						ties = append(ties, CohesiveTie{Word1: w1, Word2: w2, Type: CohesionAntonymy})
						seen[pairKey] = true
						found = true
						break
					}
				}
				if found {
					break
				}
			}
		}
	}

	return ties
}
