package semantics

import (
	"fmt"
	"strings"

	"github.com/raitucarp/gown"
)

// ThematicRole represents a semantic case role in semantic role labeling (SRL).
type ThematicRole string

const (
	RoleAgent       ThematicRole = "agent"       // The initiator/doer of the action (animate)
	RolePatient     ThematicRole = "patient"     // The entity affected/undergoing change of state
	RoleTheme       ThematicRole = "theme"       // The entity moved or described
	RoleExperiencer ThematicRole = "experiencer" // The entity experiencing an emotion or sensation
	RoleStimulus    ThematicRole = "stimulus"    // The entity causing the sensation/emotion
	RoleBeneficiary ThematicRole = "beneficiary" // The entity that benefits from the action
	RoleInstrument  ThematicRole = "instrument"  // The tool/means used
	RoleLocation    ThematicRole = "location"    // Place of occurrence
	RoleSource      ThematicRole = "source"      // Origin point
	RoleGoal        ThematicRole = "goal"        // Destination point
	RoleTime        ThematicRole = "time"        // Temporal specification
)

// SemanticArgument associates a surface argument with its assigned thematic role.
type SemanticArgument struct {
	Role ThematicRole `json:"role"`
	Text string       `json:"text"`
	Head string       `json:"head"`
}

// PredicateArgumentStructure represents a semantic proposition: Predicate(Arg1, Arg2, ...).
type PredicateArgumentStructure struct {
	Predicate string             `json:"predicate"`
	Arguments []SemanticArgument `json:"arguments"`
}

// AssignRoles assigns thematic roles to syntactic arguments (subject, direct object, etc.)
// based on the semantic process type of the verb.
func AssignRoles(verb string, subject, directObject, prepPhrase string, prepType ...string) PredicateArgumentStructure {
	pas := PredicateArgumentStructure{
		Predicate: verb,
	}

	vLower := strings.ToLower(verb)

	// Mental verbs (know, think, see, hear, fear, love, like)
	isMental := false
	mentalVerbs := []string{"see", "hear", "feel", "smell", "taste", "know", "think", "believe", "love", "hate", "like", "fear"}
	for _, mv := range mentalVerbs {
		if vLower == mv {
			isMental = true
			break
		}
	}

	if isMental {
		if subject != "" {
			pas.Arguments = append(pas.Arguments, SemanticArgument{Role: RoleExperiencer, Text: subject, Head: subject})
		}
		if directObject != "" {
			pas.Arguments = append(pas.Arguments, SemanticArgument{Role: RoleStimulus, Text: directObject, Head: directObject})
		}
	} else {
		// Material / Action verbs (eat, hit, break, run, build)
		if subject != "" {
			pas.Arguments = append(pas.Arguments, SemanticArgument{Role: RoleAgent, Text: subject, Head: subject})
		}
		if directObject != "" {
			pas.Arguments = append(pas.Arguments, SemanticArgument{Role: RolePatient, Text: directObject, Head: directObject})
		}
	}

	if prepPhrase != "" {
		role := RoleLocation
		if len(prepType) > 0 {
			switch prepType[0] {
			case "with":
				role = RoleInstrument
			case "for":
				role = RoleBeneficiary
			case "from":
				role = RoleSource
			case "to", "into":
				role = RoleGoal
			}
		}
		pas.Arguments = append(pas.Arguments, SemanticArgument{Role: role, Text: prepPhrase, Head: prepPhrase})
	}

	return pas
}

// String formats the semantic proposition as Predicate(Role: Arg, ...).
func (pas PredicateArgumentStructure) String() string {
	var args []string
	for _, arg := range pas.Arguments {
		args = append(args, fmt.Sprintf("%s: %s", arg.Role, arg.Text))
	}
	return fmt.Sprintf("%s(%s)", pas.Predicate, strings.Join(args, ", "))
}

// WordNetThematicRoleCheck checks if an entity word satisfies an expected thematic role's semantic category.
func WordNetThematicRoleCheck(res *gown.LexicalResource, word string, expectedRole ThematicRole) bool {
	entries := res.Lookup(word)
	for _, e := range entries {
		for _, s := range e.Synsets() {
			switch expectedRole {
			case RoleAgent, RoleExperiencer:
				if s.Lexfile == "noun.person" || s.Lexfile == "noun.animal" {
					return true
				}
			case RoleInstrument:
				if s.Lexfile == "noun.artifact" {
					return true
				}
			case RoleLocation:
				if s.Lexfile == "noun.location" {
					return true
				}
			case RoleTime:
				if s.Lexfile == "noun.time" {
					return true
				}
			}
		}
	}
	return false
}
