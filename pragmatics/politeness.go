package pragmatics

import (
	"strings"
)

// PolitenessStrategy indicates the face-management strategy employed.
type PolitenessStrategy string

const (
	StrategyPositivePoliteness PolitenessStrategy = "positive_politeness" // claiming common ground, compliments
	StrategyNegativePoliteness PolitenessStrategy = "negative_politeness" // hedging, minimizing imposition, indirectness
	StrategyBaldOnRecord       PolitenessStrategy = "bald_on_record"       // direct, unmitigated imperative
	StrategyOffRecord          PolitenessStrategy = "off_record"          // indirect hints, ambiguous
)

// PolitenessAnalysis evaluates an utterance for interpersonal politeness features.
type PolitenessAnalysis struct {
	Strategy       PolitenessStrategy `json:"strategy"`
	HedgeScore     float64            `json:"hedge_score"`
	MitigationTags []string           `json:"mitigation_tags"`
}

var hedges = []string{
	"please", "could you", "would you", "perhaps", "maybe", "kindly",
	"if possible", "i wonder if", "sorry to bother", "just", "a little",
}

// AnalyzePoliteness inspects an utterance for mitigation and politeness markers.
func AnalyzePoliteness(utterance string) PolitenessAnalysis {
	lower := strings.ToLower(utterance)
	var tags []string

	for _, h := range hedges {
		if strings.Contains(lower, h) {
			tags = append(tags, h)
		}
	}

	analysis := PolitenessAnalysis{
		MitigationTags: tags,
		HedgeScore:     float64(len(tags)) / 3.0,
	}
	if analysis.HedgeScore > 1.0 {
		analysis.HedgeScore = 1.0
	}

	if len(tags) >= 2 {
		analysis.Strategy = StrategyNegativePoliteness
	} else if len(tags) == 1 {
		if strings.Contains(lower, "please") {
			analysis.Strategy = StrategyNegativePoliteness
		} else {
			analysis.Strategy = StrategyPositivePoliteness
		}
	} else {
		if strings.HasSuffix(lower, "!") {
			analysis.Strategy = StrategyBaldOnRecord
		} else {
			analysis.Strategy = StrategyPositivePoliteness
		}
	}

	return analysis
}
