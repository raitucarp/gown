package pragmatics

import (
	"strings"
)

// ScalarScale represents an ordered Horn scale where weaker terms implicate negation of stronger terms.
type ScalarScale struct {
	Weaker   string `json:"weaker"`
	Stronger string `json:"stronger"`
	Implicate string `json:"implicate"`
}

var standardHornScales = []ScalarScale{
	{Weaker: "some", Stronger: "all", Implicate: "not all"},
	{Weaker: "sometimes", Stronger: "always", Implicate: "not always"},
	{Weaker: "possible", Stronger: "certain", Implicate: "not certain"},
	{Weaker: "may", Stronger: "must", Implicate: "not required"},
	{Weaker: "warm", Stronger: "hot", Implicate: "not hot"},
	{Weaker: "cool", Stronger: "cold", Implicate: "not cold"},
}

// Implicature represents a pragmatically inferred conversational implicature.
type Implicature struct {
	SourceTerm string `json:"source_term"`
	Inference  string `json:"inference"`
	Maxim      string `json:"maxim"` // "Quantity", "Relation", etc.
}

// DetectScalarImplicatures identifies scalar implicatures based on Horn scales.
func DetectScalarImplicatures(utterance string) []Implicature {
	words := strings.Fields(strings.ToLower(utterance))
	var implicatures []Implicature

	wordMap := make(map[string]bool)
	for _, w := range words {
		wClean := strings.Trim(w, ".,!?;:\"'")
		wordMap[wClean] = true
	}

	for _, scale := range standardHornScales {
		if wordMap[scale.Weaker] && !wordMap[scale.Stronger] {
			implicatures = append(implicatures, Implicature{
				SourceTerm: scale.Weaker,
				Inference:  scale.Implicate,
				Maxim:      "Quantity",
			})
		}
	}

	return implicatures
}
