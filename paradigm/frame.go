package paradigm

import (
	"strings"

	"github.com/raitucarp/gown"
)

// FrameElement defines an argument role within a semantic frame.
type FrameElement struct {
	Name       string `json:"name"`
	CoreType   string `json:"core_type"` // "Core", "Peripheral", "Extra-Thematic"
	Definition string `json:"definition"`
}

// SemanticFrame represents a conceptual scene or construction (e.g. "Ingestion", "Commerce_buy").
type SemanticFrame struct {
	Name          string         `json:"name"`
	Definition    string         `json:"definition"`
	Elements      []FrameElement `json:"elements"`
	LexicalUnits  []string       `json:"lexical_units"`
	SynsetMatches []*gown.Synset `json:"synset_matches,omitempty"`
}

// IngestionFrame provides a standard semantic frame for ingestion events (eat, drink, devour, swallow).
func IngestionFrame() SemanticFrame {
	return SemanticFrame{
		Name:       "Ingestion",
		Definition: "An Ingestor consumes Food or Drink.",
		Elements: []FrameElement{
			{Name: "Ingestor", CoreType: "Core", Definition: "The entity consuming food."},
			{Name: "Ingestibles", CoreType: "Core", Definition: "The food or drink consumed."},
			{Name: "Manner", CoreType: "Peripheral", Definition: "The manner of consumption."},
			{Name: "Means", CoreType: "Peripheral", Definition: "The instrument used."},
		},
		LexicalUnits: []string{"eat", "consume", "devour", "swallow", "drink", "ingest", "sip", "gulp", "chew"},
	}
}

// MotionFrame provides a standard semantic frame for motion.
func MotionFrame() SemanticFrame {
	return SemanticFrame{
		Name:       "Motion",
		Definition: "A Theme moves along a Path or toward a Goal.",
		Elements: []FrameElement{
			{Name: "Theme", CoreType: "Core", Definition: "The entity in motion."},
			{Name: "Source", CoreType: "Core", Definition: "The starting location."},
			{Name: "Goal", CoreType: "Core", Definition: "The end destination."},
			{Name: "Path", CoreType: "Peripheral", Definition: "The trajectory traversed."},
		},
		LexicalUnits: []string{"run", "walk", "fly", "swim", "crawl", "travel", "move"},
	}
}

// MatchWithWordNet resolves the lexical units of the semantic frame against WordNet synsets.
func (f *SemanticFrame) MatchWithWordNet(res *gown.LexicalResource) []*gown.Synset {
	var synsets []*gown.Synset
	seen := make(map[string]bool)

	for _, lu := range f.LexicalUnits {
		entries := res.Lookup(lu)
		for _, e := range entries {
			for _, s := range e.Synsets() {
				if s != nil && !seen[s.ID] {
					seen[s.ID] = true
					synsets = append(synsets, s)
				}
			}
		}
	}

	f.SynsetMatches = synsets
	return synsets
}

// EvokesFrame returns true if a word is one of the lexical units that evokes this frame.
func (f *SemanticFrame) EvokesFrame(word string) bool {
	w := strings.ToLower(strings.TrimSpace(word))
	for _, lu := range f.LexicalUnits {
		if strings.EqualFold(lu, w) {
			return true
		}
	}
	return false
}
