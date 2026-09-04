package functional

import (
	"strings"
)

// MoodType classifies interpersonal grammatical mood.
type MoodType string

const (
	MoodDeclarative   MoodType = "declarative"
	MoodInterrogative MoodType = "interrogative"
	MoodImperative    MoodType = "imperative"
	MoodExclamative   MoodType = "exclamative"
)

// Polarity indicates positive or negative grammatical polarity.
type Polarity string

const (
	PolarityPositive Polarity = "positive"
	PolarityNegative Polarity = "negative"
)

// ModalityKind classifies modal stance.
type ModalityKind string

const (
	ModalityProbability ModalityKind = "probability" // e.g. maybe, probably, must
	ModalityUsuality    ModalityKind = "usuality"    // e.g. usually, sometimes, always
	ModalityObligation  ModalityKind = "obligation"  // e.g. should, must, ought to
	ModalityInclination ModalityKind = "inclination" // e.g. willing, want, intend
	ModalityNone        ModalityKind = "none"
)

// SpeechFunction classifies the communicative act.
type SpeechFunction string

const (
	SpeechStatement SpeechFunction = "statement"
	SpeechQuestion  SpeechFunction = "question"
	SpeechCommand   SpeechFunction = "command"
	SpeechOffer     SpeechFunction = "offer"
)

// InterpersonalAnalysis holds interpersonal meaning attributes.
type InterpersonalAnalysis struct {
	Mood     MoodType       `json:"mood"`
	Polarity Polarity       `json:"polarity"`
	Modality ModalityKind   `json:"modality"`
	Speech   SpeechFunction `json:"speech_function"`
}

// AnalyzeInterpersonal parses a simple clause for interpersonal meaning indicators.
func AnalyzeInterpersonal(clause string) InterpersonalAnalysis {
	text := strings.TrimSpace(clause)
	analysis := InterpersonalAnalysis{
		Mood:     MoodDeclarative,
		Polarity: PolarityPositive,
		Modality: ModalityNone,
		Speech:   SpeechStatement,
	}

	lower := strings.ToLower(text)

	// Polarity
	negativeMarkers := []string{"not", "n't", "never", "no", "nobody", "nothing", "nowhere"}
	words := strings.Fields(lower)
	for _, w := range words {
		w = strings.Trim(w, "!?,.")
		for _, neg := range negativeMarkers {
			if w == neg || strings.HasSuffix(w, "n't") {
				analysis.Polarity = PolarityNegative
				break
			}
		}
	}

	// Modality
	if strings.Contains(lower, "must") || strings.Contains(lower, "should") || strings.Contains(lower, "ought to") {
		analysis.Modality = ModalityObligation
	} else if strings.Contains(lower, "probably") || strings.Contains(lower, "possibly") || strings.Contains(lower, "might") || strings.Contains(lower, "may") {
		analysis.Modality = ModalityProbability
	} else if strings.Contains(lower, "always") || strings.Contains(lower, "usually") || strings.Contains(lower, "often") || strings.Contains(lower, "rarely") {
		analysis.Modality = ModalityUsuality
	} else if strings.Contains(lower, "willing") || strings.Contains(lower, "want") || strings.Contains(lower, "will") {
		analysis.Modality = ModalityInclination
	}

	// Mood & Speech Function
	if strings.HasSuffix(text, "?") || strings.HasPrefix(lower, "is ") || strings.HasPrefix(lower, "are ") ||
		strings.HasPrefix(lower, "do ") || strings.HasPrefix(lower, "does ") || strings.HasPrefix(lower, "did ") ||
		strings.HasPrefix(lower, "can ") || strings.HasPrefix(lower, "could ") || strings.HasPrefix(lower, "would ") ||
		strings.HasPrefix(lower, "what ") || strings.HasPrefix(lower, "who ") || strings.HasPrefix(lower, "where ") ||
		strings.HasPrefix(lower, "when ") || strings.HasPrefix(lower, "why ") || strings.HasPrefix(lower, "how ") {
		analysis.Mood = MoodInterrogative
		analysis.Speech = SpeechQuestion
	} else if strings.HasSuffix(text, "!") && !strings.Contains(lower, "you ") && !strings.Contains(lower, "i ") {
		analysis.Mood = MoodImperative
		analysis.Speech = SpeechCommand
	}

	return analysis
}
