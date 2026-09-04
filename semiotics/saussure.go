package semiotics

import (
	"fmt"
	"strings"

	"github.com/raitucarp/gown"
)

// SaussureanSign represents Ferdinand de Saussure's dyadic model of the linguistic sign:
// Sign = Signifier (sound-image / written form) + Signified (concept / mental representation).
type SaussureanSign struct {
	Signifier string   `json:"signifier"` // The sensory form (sound / word)
	Signified string   `json:"signified"` // The concept / mental meaning
	SynsetID  string   `json:"synset_id,omitempty"`
	Oppositions []string `json:"oppositions,omitempty"` // Words contrasted to define structural value
}

// CreateSaussureanSign constructs a sign from a WordNet lexical entry and its primary synset.
func CreateSaussureanSign(res *gown.LexicalResource, word string) *SaussureanSign {
	entries := res.Lookup(word)
	if len(entries) == 0 {
		return &SaussureanSign{
			Signifier: word,
			Signified: "unknown concept",
		}
	}

	entry := entries[0]
	signified := "concept of " + word
	synID := ""

	if len(entry.Synsets()) > 0 && entry.Synsets()[0] != nil {
		syn := entry.Synsets()[0]
		synID = syn.ID
		signified = syn.PrimaryDefinition()
	}

	// Structural value (valeur): oppositions via antonyms and synonyms
	var oppositions []string
	for _, ant := range entry.Relation().Antonyms() {
		oppositions = append(oppositions, ant.Lemma.WrittenForm)
	}

	return &SaussureanSign{
		Signifier:   entry.Lemma.WrittenForm,
		Signified:   signified,
		SynsetID:    synID,
		Oppositions: oppositions,
	}
}

// String formats the Saussurean sign.
func (s SaussureanSign) String() string {
	oppStr := ""
	if len(s.Oppositions) > 0 {
		oppStr = fmt.Sprintf(" [opposed to: %s]", strings.Join(s.Oppositions, ", "))
	}
	return fmt.Sprintf("Sign{%s / \"%s\"}%s", s.Signifier, s.Signified, oppStr)
}
