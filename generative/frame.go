package generative

import (
	"strings"

	"github.com/raitucarp/gown"
)

// Valency indicates the number and nature of core verb arguments.
type Valency string

const (
	ValencyIntransitive      Valency = "intransitive"
	ValencyTransitive        Valency = "transitive"
	ValencyDitransitive      Valency = "ditransitive"
	ValencyComplexTransitive Valency = "complex_transitive"
	ValencyCopular           Valency = "copular"
)

// ArgumentRole represents a syntactic argument position.
type ArgumentRole string

const (
	ArgSubject       ArgumentRole = "subject"
	ArgDirectObject  ArgumentRole = "direct_object"
	ArgIndirectObject ArgumentRole = "indirect_object"
	ArgPrepPhrase    ArgumentRole = "prepositional_phrase"
	ArgClause        ArgumentRole = "clause"
)

// SubcatFrame represents an analyzed syntactic subcategorization frame.
type SubcatFrame struct {
	RawFrame  string         `json:"raw_frame"`
	Valency   Valency        `json:"valency"`
	Arguments []ArgumentRole `json:"arguments"`
}

// ParseSubcatFrame decomposes a WordNet syntactic behavior string into structured arguments.
// Examples:
// "Somebody ----s" -> Intransitive [Subject]
// "Somebody ----s somebody" -> Transitive [Subject, DirectObject]
// "Somebody ----s somebody something" -> Ditransitive [Subject, IndirectObject, DirectObject]
// "Somebody ----s that CLAUSE" -> ComplexTransitive [Subject, Clause]
func ParseSubcatFrame(frame string) SubcatFrame {
	f := strings.TrimSpace(frame)
	sf := SubcatFrame{RawFrame: f}

	parts := strings.Split(f, "----s")
	if len(parts) < 2 {
		parts = strings.Split(f, "----")
	}

	sf.Arguments = append(sf.Arguments, ArgSubject)

	if len(parts) < 2 {
		sf.Valency = ValencyIntransitive
		return sf
	}

	postVerb := strings.TrimSpace(parts[1])
	if postVerb == "" {
		sf.Valency = ValencyIntransitive
		return sf
	}

	tokens := strings.Fields(postVerb)
	if strings.Contains(postVerb, "CLAUSE") || strings.Contains(postVerb, "whether") {
		sf.Valency = ValencyComplexTransitive
		sf.Arguments = append(sf.Arguments, ArgClause)
	} else if len(tokens) >= 2 && (tokens[0] == "somebody" || tokens[0] == "something") && (tokens[1] == "somebody" || tokens[1] == "something") {
		sf.Valency = ValencyDitransitive
		sf.Arguments = append(sf.Arguments, ArgIndirectObject, ArgDirectObject)
	} else if strings.HasPrefix(postVerb, "to ") || strings.HasPrefix(postVerb, "from ") || strings.HasPrefix(postVerb, "with ") || strings.HasPrefix(postVerb, "on ") {
		sf.Valency = ValencyIntransitive
		sf.Arguments = append(sf.Arguments, ArgPrepPhrase)
	} else if len(tokens) > 0 {
		sf.Valency = ValencyTransitive
		sf.Arguments = append(sf.Arguments, ArgDirectObject)
		if strings.Contains(postVerb, "PP") || strings.Contains(postVerb, "to ") {
			sf.Arguments = append(sf.Arguments, ArgPrepPhrase)
		}
	} else {
		sf.Valency = ValencyIntransitive
	}

	return sf
}

// VerbSubcatFrames extracts and parses all syntactic frames associated with a verb lexical entry.
func VerbSubcatFrames(verb gown.LexicalEntry) []SubcatFrame {
	var frames []SubcatFrame
	seen := make(map[string]bool)

	for _, sense := range verb.Senses {
		if sense.Subcat != "" && !seen[sense.Subcat] {
			seen[sense.Subcat] = true
			frames = append(frames, ParseSubcatFrame(sense.Subcat))
		}
	}
	return frames
}
