package semiotics

import (
	"fmt"

	"github.com/raitucarp/gown"
)

// SemioticSquare represents A. J. Greimas's structural semiotic square (*carré sémiotique*).
// It maps the fundamental elementary structure of signification.
type SemioticSquare struct {
	S1      string `json:"s1"`       // Term 1 (e.g. "good", "life")
	S2      string `json:"s2"`       // Contrary (e.g. "evil", "death")
	NotS1   string `json:"not_s1"`   // Contradictory to S1 (e.g. "not good", "non-life")
	NotS2   string `json:"not_s2"`   // Contradictory to S2 (e.g. "not evil", "non-death")
}

// GenerateSemioticSquare constructs a semiotic square starting from a seed term by querying
// WordNet antonyms for S2 and calculating contradictory negations.
func GenerateSemioticSquare(res *gown.LexicalResource, term string) SemioticSquare {
	s1 := term
	s2 := "non-" + term

	entries := res.Lookup(term)
	for _, e := range entries {
		antonyms := e.Relation().Antonyms()
		if len(antonyms) > 0 {
			s2 = antonyms[0].Lemma.WrittenForm
			break
		}
	}

	return SemioticSquare{
		S1:    s1,
		S2:    s2,
		NotS1: "not " + s1,
		NotS2: "not " + s2,
	}
}

// Render formats the semiotic square as an ASCII diagram showing contrariety, contradiction, and implication.
func (sq SemioticSquare) Render() string {
	return fmt.Sprintf(
`      %s  <==== Contrariety ====>  %s
          \                      /
           \    Contradiction   /
            \                  /
             \                /
              \              /
     %s  <--------------------  %s
         (Implication: %s => %s, %s => %s)`,
		sq.S1, sq.S2,
		sq.NotS2, sq.NotS1,
		sq.NotS2, sq.S1, sq.NotS1, sq.S2,
	)
}
