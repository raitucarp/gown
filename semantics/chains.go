package semantics

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/text"
)

// LexicalChain represents a sequence of related words linked across a text.
type LexicalChain struct {
	Words []string `json:"words"`
	Score float64  `json:"score"`
}

// BuildLexicalChains constructs cohesive lexical chains from a text document using WordNet relations.
func BuildLexicalChains(res *gown.LexicalResource, textContent string) []LexicalChain {
	tokens := text.ExtractContentWords(textContent)
	if len(tokens) == 0 {
		return nil
	}

	var chains []LexicalChain

	for _, token := range tokens {
		matchedChainIndex := -1

		for i, chain := range chains {
			for _, w := range chain.Words {
				if isSemanticallyRelated(res, token, w) {
					matchedChainIndex = i
					break
				}
			}
			if matchedChainIndex != -1 {
				break
			}
		}

		if matchedChainIndex != -1 {
			chains[matchedChainIndex].Words = append(chains[matchedChainIndex].Words, token)
			chains[matchedChainIndex].Score += 1.0
		} else {
			chains = append(chains, LexicalChain{
				Words: []string{token},
				Score: 1.0,
			})
		}
	}

	return chains
}

func isSemanticallyRelated(res *gown.LexicalResource, w1, w2 string) bool {
	if strings.EqualFold(w1, w2) {
		return true
	}

	entries1 := res.Lookup(w1)
	for _, e1 := range entries1 {
		// Check synonyms
		for _, syn := range e1.Relation().Synonyms() {
			if strings.EqualFold(syn.Lemma.WrittenForm, w2) {
				return true
			}
		}
		// Check hypernyms
		for _, hyp := range e1.Relation().Hypernyms() {
			if strings.EqualFold(hyp.Lemma.WrittenForm, w2) {
				return true
			}
		}
		// Check hyponyms
		for _, hyp := range e1.Relation().Hyponyms() {
			if strings.EqualFold(hyp.Lemma.WrittenForm, w2) {
				return true
			}
		}
	}

	return false
}
