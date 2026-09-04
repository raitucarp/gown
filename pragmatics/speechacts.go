package pragmatics

import (
	"strings"
)

// SpeechActClass represents Searle's classification of illocutionary acts.
type SpeechActClass string

const (
	ActAssertive   SpeechActClass = "assertive"   // Commits speaker to the truth (claim, state, report)
	ActDirective   SpeechActClass = "directive"   // Attempts to get addressee to act (ask, order, request)
	ActCommissive  SpeechActClass = "commissive"  // Commits speaker to future action (promise, offer, vow)
	ActExpressive  SpeechActClass = "expressive"  // Expresses psychological state (thank, apologize, welcome)
	ActDeclaration SpeechActClass = "declaration" // Alters the institutional state (declare, resign, baptize)
)

// IllocutionaryForce characterizes the pragmatic intent of an utterance.
type IllocutionaryForce struct {
	Class           SpeechActClass `json:"class"`
	PerformativeVerb string         `json:"performative_verb,omitempty"`
	Confidence      float64        `json:"confidence"`
}

var performativeVerbs = map[string]SpeechActClass{
	"state": ActAssertive, "assert": ActAssertive, "claim": ActAssertive, "report": ActAssertive, "conclude": ActAssertive,
	"ask": ActDirective, "request": ActDirective, "order": ActDirective, "command": ActDirective, "beg": ActDirective, "advise": ActDirective,
	"promise": ActCommissive, "swear": ActCommissive, "vow": ActCommissive, "offer": ActCommissive, "pledge": ActCommissive,
	"thank": ActExpressive, "apologize": ActExpressive, "congratulate": ActExpressive, "welcome": ActExpressive, "deplore": ActExpressive,
	"declare": ActDeclaration, "baptize": ActDeclaration, "pronounce": ActDeclaration, "appoint": ActDeclaration, "resign": ActDeclaration,
}

// ClassifySpeechAct identifies the primary speech act class and illocutionary force of an utterance.
func ClassifySpeechAct(utterance string) IllocutionaryForce {
	lower := strings.ToLower(strings.TrimSpace(utterance))

	// 1. Explicit performative formula: "I [performative] that/to..."
	for verb, actClass := range performativeVerbs {
		if strings.HasPrefix(lower, "i "+verb) || strings.HasPrefix(lower, "we "+verb) {
			return IllocutionaryForce{
				Class:           actClass,
				PerformativeVerb: verb,
				Confidence:      0.95,
			}
		}
	}

	// 2. Directives: questions and imperatives
	if strings.HasSuffix(lower, "?") {
		return IllocutionaryForce{Class: ActDirective, Confidence: 0.85}
	}
	if strings.HasPrefix(lower, "please ") || strings.HasPrefix(lower, "could you ") || strings.HasPrefix(lower, "can you ") {
		return IllocutionaryForce{Class: ActDirective, Confidence: 0.90}
	}

	// 3. Expressives: greetings and acknowledgments
	if strings.HasPrefix(lower, "thank") || strings.HasPrefix(lower, "sorry") || strings.HasPrefix(lower, "congratulations") {
		return IllocutionaryForce{Class: ActExpressive, Confidence: 0.90}
	}

	// 4. Commissives: first person future commitment
	if strings.HasPrefix(lower, "i will ") || strings.HasPrefix(lower, "i promise ") || strings.HasPrefix(lower, "we will ") {
		return IllocutionaryForce{Class: ActCommissive, Confidence: 0.80}
	}

	// 5. Default declarative: Assertive
	return IllocutionaryForce{Class: ActAssertive, Confidence: 0.75}
}
