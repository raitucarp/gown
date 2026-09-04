package discourse

import (
	"strings"

	"github.com/raitucarp/gown/text"
)

// Mention represents an occurrence of an entity in the text.
type Mention struct {
	ID         int    `json:"id"`
	SentenceID int    `json:"sentence_id"`
	Text       string `json:"text"`
	IsPronoun  bool   `json:"is_pronoun"`
}

// CoreferenceChain groups all mentions that refer to the same real-world entity.
type CoreferenceChain struct {
	EntityName string    `json:"entity_name"`
	Mentions   []Mention `json:"mentions"`
}

// TrackCoreference identifies nominal mentions and resolves simple anaphora across sentences.
func TrackCoreference(documentText string) []CoreferenceChain {
	sentences := text.SentenceSegment(documentText)
	var chains []CoreferenceChain

	pronouns := map[string]bool{
		"he": true, "him": true, "his": true,
		"she": true, "her": true, "hers": true,
		"it": true, "its": true,
		"they": true, "them": true, "their": true,
	}

	mentionID := 1
	var lastNominalEntity string

	for sID, sent := range sentences {
		words := strings.Fields(sent)
		for _, w := range words {
			wClean := strings.ToLower(strings.Trim(w, ".,!?;:\"'()"))
			if wClean == "" {
				continue
			}

			if pronouns[wClean] {
				// Pronoun mention resolved to last observed nominal entity
				resolved := lastNominalEntity
				if resolved == "" {
					resolved = wClean
				}

				m := Mention{
					ID:         mentionID,
					SentenceID: sID + 1,
					Text:       w,
					IsPronoun:  true,
				}
				mentionID++

				found := false
				for i := range chains {
					if chains[i].EntityName == resolved {
						chains[i].Mentions = append(chains[i].Mentions, m)
						found = true
						break
					}
				}
				if !found {
					chains = append(chains, CoreferenceChain{
						EntityName: resolved,
						Mentions:   []Mention{m},
					})
				}
			} else if !text.IsStopword(wClean) && len(wClean) > 2 {
				// Content nominal candidate
				lastNominalEntity = wClean
				m := Mention{
					ID:         mentionID,
					SentenceID: sID + 1,
					Text:       w,
					IsPronoun:  false,
				}
				mentionID++

				found := false
				for i := range chains {
					if strings.EqualFold(chains[i].EntityName, wClean) {
						chains[i].Mentions = append(chains[i].Mentions, m)
						found = true
						break
					}
				}
				if !found {
					chains = append(chains, CoreferenceChain{
						EntityName: wClean,
						Mentions:   []Mention{m},
					})
				}
			}
		}
	}

	return chains
}
