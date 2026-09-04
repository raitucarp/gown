package semantics

import (
	"math"

	"github.com/raitucarp/gown"
)

// PolysemyReport summarizes polysemy and sense inventory for a word.
type PolysemyReport struct {
	Word         string            `json:"word"`
	TotalSenses  int               `json:"total_senses"`
	SensesByPOS  map[gown.POS]int  `json:"senses_by_pos"`
	IsPolysemous bool              `json:"is_polysemous"` // True if TotalSenses > 1
	Entropy      float64           `json:"entropy"`       // Sense distribution entropy
	Definitions  []string          `json:"definitions"`
}

// AnalyzePolysemy inspects the senses and semantic distribution of a word.
func AnalyzePolysemy(res *gown.LexicalResource, word string) PolysemyReport {
	entries := res.Lookup(word)
	report := PolysemyReport{
		Word:        word,
		SensesByPOS: make(map[gown.POS]int),
	}

	for _, e := range entries {
		pos := gown.POS(e.Lemma.PartOfSpeech)
		numSenses := len(e.Senses)
		report.TotalSenses += numSenses
		report.SensesByPOS[pos] += numSenses

		for _, s := range e.Senses {
			syn := s.GetSynset()
			if syn != nil && syn.PrimaryDefinition() != "" {
				report.Definitions = append(report.Definitions, syn.PrimaryDefinition())
			}
		}
	}

	report.IsPolysemous = report.TotalSenses > 1

	// Calculate POS distribution entropy
	if report.TotalSenses > 0 {
		var h float64
		total := float64(report.TotalSenses)
		for _, cnt := range report.SensesByPOS {
			p := float64(cnt) / total
			if p > 0 {
				h -= p * math.Log2(p)
			}
		}
		report.Entropy = h
	}

	return report
}

// IsHomonym checks if the word has entries across distinct, unrelated semantic fields or parts of speech.
func IsHomonym(res *gown.LexicalResource, word string) bool {
	entries := res.Lookup(word)
	if len(entries) > 1 {
		// Word has multiple distinct lexical entries (often indicative of distinct etymologies or POS)
		return true
	}
	return false
}
