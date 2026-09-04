package semiotics

import (
	"strings"

	"github.com/raitucarp/gown"
)

// Valence represents emotional or evaluative orientation.
type Valence string

const (
	ValencePositive Valence = "positive"
	ValenceNeutral  Valence = "neutral"
	ValenceNegative Valence = "negative"
)

// ConnotationAnalysis separates literal denotation from cultural/evaluative connotation.
type ConnotationAnalysis struct {
	Word        string   `json:"word"`
	Denotation  string   `json:"denotation"`
	Valence     Valence  `json:"valence"`
	Register    string   `json:"register"` // "formal", "informal", "slang", "neutral"
	Associations []string `json:"associations"`
}

var positiveMarkers = []string{"good", "great", "excellent", "beautiful", "noble", "honor", "love", "pleasant", "virtue", "joy"}
var negativeMarkers = []string{"bad", "evil", "terrible", "ugly", "corrupt", "hate", "pain", "unpleasant", "vicious", "decay"}

// AnalyzeConnotation inspects the literal WordNet definition and evaluative markers of a word.
func AnalyzeConnotation(res *gown.LexicalResource, word string) ConnotationAnalysis {
	entries := res.Lookup(word)
	denotation := ""
	if len(entries) > 0 && len(entries[0].Synsets()) > 0 && entries[0].Synsets()[0] != nil {
		denotation = entries[0].Synsets()[0].PrimaryDefinition()
	}

	analysis := ConnotationAnalysis{
		Word:       word,
		Denotation: denotation,
		Valence:    ValenceNeutral,
		Register:   "neutral",
	}

	lower := strings.ToLower(word + " " + denotation)

	// Check register from WordNet glosses
	if strings.Contains(lower, "informal") || strings.Contains(lower, "slang") {
		analysis.Register = "informal"
	} else if strings.Contains(lower, "formal") || strings.Contains(lower, "archaic") || strings.Contains(lower, "literary") {
		analysis.Register = "formal"
	}

	posScore := 0
	negScore := 0
	for _, m := range positiveMarkers {
		if strings.Contains(lower, m) {
			posScore++
			analysis.Associations = append(analysis.Associations, "+"+m)
		}
	}
	for _, m := range negativeMarkers {
		if strings.Contains(lower, m) {
			negScore++
			analysis.Associations = append(analysis.Associations, "-"+m)
		}
	}

	if posScore > negScore {
		analysis.Valence = ValencePositive
	} else if negScore > posScore {
		analysis.Valence = ValenceNegative
	}

	return analysis
}
