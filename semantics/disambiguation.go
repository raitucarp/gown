package semantics

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/text"
)

// DisambiguationResult describes the best-matching synset and overlap score.
type DisambiguationResult struct {
	Word         string       `json:"word"`
	BestSynset   *gown.Synset `json:"best_synset"`
	Definition   string       `json:"definition"`
	Score        int          `json:"score"`
	TotalSenses  int          `json:"total_senses"`
}

// LeskMode specifies standard vs extended gloss signatures.
type LeskMode int

const (
	// LeskSimplified measures overlap with the target sense definition and examples.
	LeskSimplified LeskMode = iota
	// LeskExtended includes definitions of hypernyms and hyponyms.
	LeskExtended
)

// DisambiguateLesk selects the most probable sense of targetWord within the contextSentence
// using the Lesk overlap algorithm.
func DisambiguateLesk(res *gown.LexicalResource, targetWord, contextSentence string, mode ...LeskMode) DisambiguationResult {
	m := LeskSimplified
	if len(mode) > 0 {
		m = mode[0]
	}

	entries := res.Lookup(targetWord)
	result := DisambiguationResult{
		Word: targetWord,
	}

	if len(entries) == 0 {
		return result
	}

	contextWords := text.ExtractContentWords(contextSentence)
	contextSet := make(map[string]bool)
	for _, w := range contextWords {
		// Ignore the target word itself in the context set
		if !strings.EqualFold(w, targetWord) {
			contextSet[w] = true
		}
	}

	var candidateSynsets []*gown.Synset
	for _, entry := range entries {
		for _, s := range entry.Synsets() {
			if s != nil {
				candidateSynsets = append(candidateSynsets, s)
			}
		}
	}

	result.TotalSenses = len(candidateSynsets)
	if len(candidateSynsets) == 0 {
		return result
	}

	// Fallback: first sense heuristic
	bestSynset := candidateSynsets[0]
	bestScore := -1

	for _, syn := range candidateSynsets {
		sigWords := buildGlossSignature(res, syn, m)
		score := 0
		for _, w := range sigWords {
			if contextSet[w] {
				score++
			}
		}

		if score > bestScore {
			bestScore = score
			bestSynset = syn
		}
	}

	result.BestSynset = bestSynset
	result.Score = max(bestScore, 0)
	if bestSynset != nil {
		result.Definition = bestSynset.PrimaryDefinition()
	}

	return result
}

func buildGlossSignature(res *gown.LexicalResource, syn *gown.Synset, mode LeskMode) []string {
	var textParts []string
	textParts = append(textParts, syn.Definitions...)
	for _, ex := range syn.Examples {
		textParts = append(textParts, ex.Text)
	}

	if mode == LeskExtended {
		for _, hyp := range syn.Hypernyms(res) {
			textParts = append(textParts, hyp.Definitions...)
		}
		for _, hyp := range syn.Hyponyms(res) {
			textParts = append(textParts, hyp.Definitions...)
		}
	}

	combined := strings.Join(textParts, " ")
	return text.ExtractContentWords(combined)
}
