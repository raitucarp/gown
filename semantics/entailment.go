package semantics

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/graph"
)

// EntailmentType indicates how an entailment relationship is derived.
type EntailmentType string

const (
	EntailmentHyponymy EntailmentType = "hyponymy" // dog entails animal
	EntailmentWordNet  EntailmentType = "wordnet"  // snore entails sleep
	EntailmentIdentity EntailmentType = "identity" // dog entails dog
	EntailmentNone     EntailmentType = "none"
)

// CheckEntailment checks if premise entails hypothesis through hypernymy or WordNet entailment.
// e.g. "dog" entails "animal", "snore" entails "sleep".
func CheckEntailment(res *gown.LexicalResource, premise, hypothesis string) (bool, EntailmentType) {
	p := strings.ToLower(strings.TrimSpace(premise))
	h := strings.ToLower(strings.TrimSpace(hypothesis))

	if p == h {
		return true, EntailmentIdentity
	}

	pEntries := res.Lookup(p)
	hEntries := res.Lookup(h)
	if len(pEntries) == 0 || len(hEntries) == 0 {
		return false, EntailmentNone
	}

	// 1. Check direct WordNet verb entailment relation (premise entails hypothesis)
	for _, pe := range pEntries {
		for _, entailed := range pe.Relation().Entails() {
			if strings.EqualFold(entailed.Lemma.WrittenForm, h) {
				return true, EntailmentWordNet
			}
		}
	}

	// 2. Check taxonomic hyponymy (premise is a hyponym of hypothesis => premise is a specific kind of hypothesis)
	for _, pe := range pEntries {
		for _, ps := range pe.Synsets() {
			anc := graph.HypernymAncestors(res, ps)
			for _, he := range hEntries {
				for _, hs := range he.Synsets() {
					if _, ok := anc[hs.ID]; ok {
						return true, EntailmentHyponymy
					}
				}
			}
		}
	}

	return false, EntailmentNone
}

// CheckContradiction tests whether two words are incompatible opposites via antonymy.
// e.g. "hot" vs "cold", "alive" vs "dead", "happy" vs "sad".
func CheckContradiction(res *gown.LexicalResource, word1, word2 string) bool {
	w1 := strings.ToLower(strings.TrimSpace(word1))
	w2 := strings.ToLower(strings.TrimSpace(word2))

	if w1 == w2 {
		return false
	}

	entries1 := res.Lookup(w1)
	for _, e1 := range entries1 {
		for _, ant := range e1.Relation().Antonyms() {
			if strings.EqualFold(ant.Lemma.WrittenForm, w2) {
				return true
			}
		}
	}

	return false
}
