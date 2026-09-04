package pragmatics

import (
	"strings"
)

// TriggerType classifies the lexical or syntactic trigger for a presupposition.
type TriggerType string

const (
	TriggerFactiveVerb    TriggerType = "factive_verb"     // know, regret, realize
	TriggerChangeOfState  TriggerType = "change_of_state"  // stop, start, continue
	TriggerIterative      TriggerType = "iterative"        // again, another, return
	TriggerDefiniteDesc   TriggerType = "definite_article" // the X exists
)

// Presupposition represents a background assumption taken for granted by an utterance.
type Presupposition struct {
	Trigger       string      `json:"trigger"`
	Type          TriggerType `json:"type"`
	Presupposition string     `json:"presupposition"`
	SurvivesNegation bool     `json:"survives_negation"`
}

// ExtractPresuppositions detects pragmatic presupposition triggers in a clause.
func ExtractPresuppositions(utterance string) []Presupposition {
	lower := strings.ToLower(strings.TrimSpace(utterance))
	var results []Presupposition

	// 1. Factive verbs
	factiveVerbs := []string{"realize", "realized", "know", "knew", "regret", "regretted", "discovered"}
	for _, fv := range factiveVerbs {
		idx := strings.Index(lower, fv+" that ")
		if idx != -1 {
			clause := strings.TrimSpace(utterance[idx+len(fv)+6:])
			results = append(results, Presupposition{
				Trigger:          fv,
				Type:             TriggerFactiveVerb,
				Presupposition:   clause,
				SurvivesNegation: true,
			})
		}
	}

	// 2. Change of state verbs
	changeVerbs := []string{"stopped", "stops", "quit", "ceased"}
	for _, cv := range changeVerbs {
		idx := strings.Index(lower, cv+" ")
		if idx != -1 {
			activity := strings.TrimSpace(utterance[idx+len(cv)+1:])
			results = append(results, Presupposition{
				Trigger:          cv,
				Type:             TriggerChangeOfState,
				Presupposition:   "previously engaged in: " + activity,
				SurvivesNegation: true,
			})
		}
	}

	// 3. Iterative adverbs
	if strings.HasSuffix(lower, " again") || strings.Contains(lower, " again ") {
		results = append(results, Presupposition{
			Trigger:          "again",
			Type:             TriggerIterative,
			Presupposition:   "event occurred previously",
			SurvivesNegation: true,
		})
	}

	return results
}
