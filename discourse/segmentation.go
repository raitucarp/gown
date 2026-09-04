package discourse

import (
	"strings"

	"github.com/raitucarp/gown/text"
)

// EDU represents an Elementary Discourse Unit (a clause-like discourse segment).
type EDU struct {
	ID        int    `json:"id"`
	Text      string `json:"text"`
	SentenceID int   `json:"sentence_id"`
}

// SegmentEDUs splits a text document into Elementary Discourse Units using clause boundaries
// and discourse conjunctions (because, although, but, however, while).
func SegmentEDUs(documentText string) []EDU {
	sentences := text.SentenceSegment(documentText)
	var edus []EDU
	eduID := 1

	connectors := []string{", but ", ", because ", ", although ", ", however, ", ", while ", "; however, "}

	for sID, sent := range sentences {
		rawSegments := []string{sent}

		for _, conn := range connectors {
			var nextSegments []string
			for _, seg := range rawSegments {
				if strings.Contains(seg, conn) {
					parts := strings.Split(seg, conn)
					for i, p := range parts {
						p = strings.TrimSpace(p)
						if i > 0 {
							// Re-attach connector keyword for context
							connClean := strings.Trim(conn, " ,;")
							nextSegments = append(nextSegments, connClean+": "+p)
						} else {
							nextSegments = append(nextSegments, p)
						}
					}
				} else {
					nextSegments = append(nextSegments, seg)
				}
			}
			rawSegments = nextSegments
		}

		for _, seg := range rawSegments {
			clean := strings.TrimSpace(seg)
			if clean != "" {
				edus = append(edus, EDU{
					ID:         eduID,
					Text:       clean,
					SentenceID: sID + 1,
				})
				eduID++
			}
		}
	}

	return edus
}
