package semantics

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/graph"
)

// SemanticRestriction defines a conceptual constraint required for an argument.
type SemanticRestriction string

const (
	RestrictAnimate   SemanticRestriction = "animate"   // person or animal
	RestrictHuman     SemanticRestriction = "human"     // person
	RestrictFood      SemanticRestriction = "food"      // food or edible substance
	RestrictLiquid    SemanticRestriction = "liquid"    // drinkable or liquid substance
	RestrictPhysical  SemanticRestriction = "physical"  // physical object or entity
	RestrictLocation  SemanticRestriction = "location"  // location or place
	RestrictArtifact  SemanticRestriction = "artifact"  // human-made tool or object
	RestrictDocument  SemanticRestriction = "document"  // written material or communication
)

// SelectionalProfile specifies the selectional constraints of a verb.
type SelectionalProfile struct {
	Verb           string               `json:"verb"`
	SubjectRestriction SemanticRestriction  `json:"subject_restriction,omitempty"`
	ObjectRestriction  SemanticRestriction  `json:"object_restriction,omitempty"`
}

// GetVerbSelectionalProfile returns standard selectional preferences for common English verbs.
func GetVerbSelectionalProfile(verb string) SelectionalProfile {
	vLower := strings.ToLower(verb)
	switch vLower {
	case "eat", "devour", "chew", "swallow", "ingest":
		return SelectionalProfile{
			Verb:               verb,
			SubjectRestriction: RestrictAnimate,
			ObjectRestriction:  RestrictFood,
		}
	case "drink", "sip", "gulp":
		return SelectionalProfile{
			Verb:               verb,
			SubjectRestriction: RestrictAnimate,
			ObjectRestriction:  RestrictLiquid,
		}
	case "read", "browse", "peruse":
		return SelectionalProfile{
			Verb:               verb,
			SubjectRestriction: RestrictHuman,
			ObjectRestriction:  RestrictDocument,
		}
	case "drive", "pilot", "steer":
		return SelectionalProfile{
			Verb:               verb,
			SubjectRestriction: RestrictHuman,
			ObjectRestriction:  RestrictArtifact,
		}
	default:
		return SelectionalProfile{
			Verb:               verb,
			SubjectRestriction: RestrictPhysical,
			ObjectRestriction:  RestrictPhysical,
		}
	}
}

// SatisfiesRestriction checks if a noun word satisfies a semantic restriction using WordNet taxonomy.
// It prioritizes exact-case lemma matches (e.g. "rock" vs capitalized proper noun "Rock") and evaluates
// primary senses to prevent metaphorical or proper-noun senses from causing false positives in literal selectional checking.
func SatisfiesRestriction(res *gown.LexicalResource, noun string, restriction SemanticRestriction) bool {
	entries := res.Lookup(noun, gown.WithPOS(gown.NounPos))
	if len(entries) == 0 {
		return false
	}

	// Prefer exact case match if available (e.g. common noun "rock" vs proper noun "Rock")
	var targetEntries gown.LexicalEntries
	for _, e := range entries {
		if e.Lemma.WrittenForm == noun {
			targetEntries = append(targetEntries, e)
		}
	}
	if len(targetEntries) == 0 {
		targetEntries = entries
	}

	for _, e := range targetEntries {
		synsets := e.Synsets()
		if len(synsets) > 0 {
			if matchesRestriction(res, synsets[0], restriction) {
				return true
			}
		}
	}
	return false
}

func matchesRestriction(res *gown.LexicalResource, s *gown.Synset, restriction SemanticRestriction) bool {
	if s == nil {
		return false
	}

	switch restriction {
	case RestrictAnimate:
		return s.Lexfile == "noun.person" || s.Lexfile == "noun.animal"
	case RestrictHuman:
		return s.Lexfile == "noun.person"
	case RestrictFood:
		if s.Lexfile == "noun.food" {
			return true
		}
		// Check hypernyms for "food"
		anc := graph.HypernymAncestors(res, s)
		for id := range anc {
			ancSyn := res.SynsetByID(id)
			if ancSyn != nil && (ancSyn.Lexfile == "noun.food" || strings.Contains(ancSyn.PrimaryDefinition(), "food")) {
				return true
			}
		}
	case RestrictLiquid:
		if s.Lexfile == "noun.substance" || s.Lexfile == "noun.food" {
			def := strings.ToLower(s.PrimaryDefinition())
			if strings.Contains(def, "liquid") || strings.Contains(def, "beverage") || strings.Contains(def, "drink") {
				return true
			}
		}
	case RestrictArtifact:
		return s.Lexfile == "noun.artifact"
	case RestrictLocation:
		return s.Lexfile == "noun.location"
	case RestrictDocument:
		return s.Lexfile == "noun.communication"
	case RestrictPhysical:
		return s.Lexfile == "noun.artifact" || s.Lexfile == "noun.object" || s.Lexfile == "noun.animal" || s.Lexfile == "noun.plant"
	}

	return false
}

// CheckSelectionalViolation tests whether subject and direct object satisfy the verb's selectional restrictions.
func CheckSelectionalViolation(res *gown.LexicalResource, verb, subject, object string) (bool, string) {
	profile := GetVerbSelectionalProfile(verb)

	if subject != "" && profile.SubjectRestriction != "" {
		if !SatisfiesRestriction(res, subject, profile.SubjectRestriction) {
			return true, "Subject '" + subject + "' violates selectional restriction [" + string(profile.SubjectRestriction) + "] for verb '" + verb + "'"
		}
	}

	if object != "" && profile.ObjectRestriction != "" {
		if !SatisfiesRestriction(res, object, profile.ObjectRestriction) {
			return true, "Object '" + object + "' violates selectional restriction [" + string(profile.ObjectRestriction) + "] for verb '" + verb + "'"
		}
	}

	return false, ""
}
