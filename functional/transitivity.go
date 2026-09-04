package functional

import (
	"strings"

	"github.com/raitucarp/gown"
)

// ProcessType identifies the transitivity process category in Halliday's SFL.
type ProcessType string

const (
	ProcessMaterial    ProcessType = "material"    // Doing and happening (e.g. run, hit, build)
	ProcessMental      ProcessType = "mental"      // Sensing, feeling, thinking (e.g. know, like, see)
	ProcessRelational  ProcessType = "relational"  // Being and having (e.g. be, have, belong)
	ProcessVerbal      ProcessType = "verbal"      // Saying (e.g. tell, explain, ask)
	ProcessBehavioral  ProcessType = "behavioral"  // Physiological and psychological behavior (e.g. breathe, dream, cry)
	ProcessExistential ProcessType = "existential" // Existing (e.g. exist, arise)
)

// ParticipantRole represents the experiential participant in a transitivity configuration.
type ParticipantRole string

const (
	RoleActor       ParticipantRole = "actor"
	RoleGoal        ParticipantRole = "goal"
	RoleSenser      ParticipantRole = "senser"
	RolePhenomenon  ParticipantRole = "phenomenon"
	RoleCarrier     ParticipantRole = "carrier"
	RoleAttribute   ParticipantRole = "attribute"
	RoleSayer       ParticipantRole = "sayer"
	RoleReceiver    ParticipantRole = "receiver"
	RoleTarget      ParticipantRole = "target"
	RoleBehaver     ParticipantRole = "behaver"
	RoleExistent    ParticipantRole = "existent"
)

// TransitivityProfile characterizes a verb's process type and default participant structure.
type TransitivityProfile struct {
	Process      ProcessType       `json:"process"`
	Participants []ParticipantRole `json:"participants"`
	Lexfiles     []string          `json:"lexfiles"`
}

// ClassifyVerbProcess categorizes a WordNet verb entry into an SFL Transitivity process type
// based on its semantic lexfiles.
func ClassifyVerbProcess(verb gown.LexicalEntry) TransitivityProfile {
	var lexfiles []string
	for _, syn := range verb.Synsets() {
		if syn != nil && syn.Lexfile != "" {
			lexfiles = append(lexfiles, syn.Lexfile)
		}
	}

	profile := TransitivityProfile{
		Lexfiles: lexfiles,
	}

	// Tally occurrences of lexfiles
	counts := make(map[ProcessType]int)
	for _, lf := range lexfiles {
		switch lf {
		case "verb.motion", "verb.contact", "verb.change", "verb.creation", "verb.consumption":
			counts[ProcessMaterial]++
		case "verb.cognition", "verb.perception", "verb.emotion":
			counts[ProcessMental]++
		case "verb.stative", "verb.possession":
			counts[ProcessRelational]++
		case "verb.communication":
			counts[ProcessVerbal]++
		case "verb.body":
			counts[ProcessBehavioral]++
		}
	}

	bestType := ProcessMaterial
	maxCount := 0
	for pType, cnt := range counts {
		if cnt > maxCount {
			maxCount = cnt
			bestType = pType
		}
	}

	// Handle specific existential verbs like "exist", "occur"
	lemma := strings.ToLower(verb.Lemma.WrittenForm)
	if lemma == "exist" || lemma == "occur" || lemma == "happen" {
		bestType = ProcessExistential
	}

	profile.Process = bestType
	switch bestType {
	case ProcessMaterial:
		profile.Participants = []ParticipantRole{RoleActor, RoleGoal}
	case ProcessMental:
		profile.Participants = []ParticipantRole{RoleSenser, RolePhenomenon}
	case ProcessRelational:
		profile.Participants = []ParticipantRole{RoleCarrier, RoleAttribute}
	case ProcessVerbal:
		profile.Participants = []ParticipantRole{RoleSayer, RoleReceiver, RoleTarget}
	case ProcessBehavioral:
		profile.Participants = []ParticipantRole{RoleBehaver}
	case ProcessExistential:
		profile.Participants = []ParticipantRole{RoleExistent}
	}

	return profile
}
