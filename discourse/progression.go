package discourse

import (
	"strings"

	"github.com/raitucarp/gown/functional"
	"github.com/raitucarp/gown/text"
)

// ProgressionType categorizes Daneš's thematic progression patterns.
type ProgressionType string

const (
	ProgressionConstant ProgressionType = "constant_theme" // T1 -> T1 (same theme maintained)
	ProgressionLinear   ProgressionType = "linear_progression" // R1 -> T2 (rheme becomes next theme)
	ProgressionNew      ProgressionType = "new_theme" // Theme changes completely
)

// ThematicProgressionStep represents the transition between two consecutive clause themes.
type ThematicProgressionStep struct {
	SentenceID int             `json:"sentence_id"`
	Theme      string          `json:"theme"`
	Rheme      string          `json:"rheme"`
	Type       ProgressionType `json:"type"`
}

// AnalyzeThemeProgression analyzes a text across sentences according to Daneš's theme progression theory.
func AnalyzeThemeProgression(documentText string) []ThematicProgressionStep {
	sentences := text.SentenceSegment(documentText)
	if len(sentences) == 0 {
		return nil
	}

	var steps []ThematicProgressionStep
	var prevTheme string
	var prevRheme string

	for i, s := range sentences {
		tr := functional.SplitThemeRheme(s)
		currTheme := strings.ToLower(strings.TrimSpace(tr.Theme))
		currRheme := strings.ToLower(strings.TrimSpace(tr.Rheme))

		progType := ProgressionNew
		if i > 0 {
			if strings.Contains(currTheme, prevTheme) || strings.Contains(prevTheme, currTheme) {
				progType = ProgressionConstant
			} else if strings.Contains(prevRheme, currTheme) {
				progType = ProgressionLinear
			}
		}

		steps = append(steps, ThematicProgressionStep{
			SentenceID: i + 1,
			Theme:      tr.Theme,
			Rheme:      tr.Rheme,
			Type:       progType,
		})

		prevTheme = currTheme
		prevRheme = currRheme
	}

	return steps
}
