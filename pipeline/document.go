package pipeline

import (
	"github.com/raitucarp/gown/discourse"
	"github.com/raitucarp/gown/pragmatics"
	"github.com/raitucarp/gown/semantics"
	"github.com/raitucarp/gown/semiotics"
	"github.com/raitucarp/gown/syntax"
)

// WordAnalysis aggregates linguistic information for a single token.
type WordAnalysis struct {
	Surface     string                    `json:"surface"`
	Lemma       string                    `json:"lemma"`
	POS         string                    `json:"pos"`
	CVPattern   string                    `json:"cv_pattern"`
	Syllables   int                       `json:"syllables"`
	SignMode    semiotics.SignMode        `json:"sign_mode"`
	Connotation semiotics.ConnotationAnalysis `json:"connotation"`
}

// SentenceAnalysis aggregates all linguistic layers for a single sentence.
type SentenceAnalysis struct {
	ID             int                               `json:"id"`
	Raw            string                            `json:"raw"`
	Words          []WordAnalysis                    `json:"words"`
	SyntaxTree     *syntax.SyntaxNode                `json:"syntax_tree,omitempty"`
	Relations      []syntax.GrammaticalRelation      `json:"relations,omitempty"`
	Roles          semantics.PredicateArgumentStructure `json:"roles,omitempty"`
	SpeechAct      pragmatics.IllocutionaryForce     `json:"speech_act"`
	Deixis         []pragmatics.DeicticExpression    `json:"deixis,omitempty"`
	Presupposition []pragmatics.Presupposition       `json:"presuppositions,omitempty"`
	Implicatures   []pragmatics.Implicature          `json:"implicatures,omitempty"`
	Politeness     pragmatics.PolitenessAnalysis     `json:"politeness"`
}

// LinguisticDocument represents a multi-layer computational linguistic analysis of an entire text.
type LinguisticDocument struct {
	RawText           string                            `json:"raw_text"`
	Sentences         []SentenceAnalysis                `json:"sentences"`
	EDUs              []discourse.EDU                   `json:"edus"`
	RSTTree           *discourse.RSTNode                `json:"rst_tree,omitempty"`
	CoreferenceChains []discourse.CoreferenceChain      `json:"coreference_chains,omitempty"`
	TopicTracking     []discourse.SentenceTopic         `json:"topic_tracking,omitempty"`
	ThemeProgression  []discourse.ThematicProgressionStep `json:"theme_progression,omitempty"`
	SemioticSquares   map[string]semiotics.SemioticSquare `json:"semiotic_squares,omitempty"`
}
