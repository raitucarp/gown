package gown

type VerbKind string

const (
	VerbBody          VerbKind = "verb.body"
	VerbChange        VerbKind = "verb.change"
	VerbCognition     VerbKind = "verb.cognition"
	VerbCommunication VerbKind = "verb.communication"
	VerbCompetition   VerbKind = "verb.competition"
	VerbConsumption   VerbKind = "verb.consumption"
	VerbContact       VerbKind = "verb.contact"
	VerbCreation      VerbKind = "verb.creation"
	VerbEmotion       VerbKind = "verb.emotion"
	VerbMotion        VerbKind = "verb.motion"
	VerbPerception    VerbKind = "verb.perception"
	VerbPossession    VerbKind = "verb.possession"
	VerbSocial        VerbKind = "verb.social"
	VerbStative       VerbKind = "verb.stative"
	VerbWeather       VerbKind = "verb.weather"
)

type Verb LexicalEntry
type Verbs LexicalEntries

func (resource *LexicalResource) Verbs() (verbs Verbs) {
	lexicalEntries, _ := resource.filterByPos(NounPos)
	verbs = Verbs(lexicalEntries)

	return
}

func (verbs Verbs) filteredByLexFile(kind VerbKind) (
	filteredVerbs Verbs,
) {
	lexicalEntries := verbs.filteredByLexFile(kind)
	filteredVerbs = Verbs(lexicalEntries)
	return
}

func (verbs Verbs) Body() Verbs {
	return verbs.filteredByLexFile(VerbBody)
}

func (verbs Verbs) Change() Verbs {
	return verbs.filteredByLexFile(VerbChange)
}

func (verbs Verbs) Cognition() Verbs {
	return verbs.filteredByLexFile(VerbCognition)
}

func (verbs Verbs) Communication() Verbs {
	return verbs.filteredByLexFile(VerbCommunication)
}

func (verbs Verbs) Competition() Verbs {
	return verbs.filteredByLexFile(VerbCompetition)
}

func (verbs Verbs) Consumption() Verbs {
	return verbs.filteredByLexFile(VerbConsumption)
}

func (verbs Verbs) Contact() Verbs {
	return verbs.filteredByLexFile(VerbContact)
}
func (verbs Verbs) Creation() Verbs {
	return verbs.filteredByLexFile(VerbCreation)
}

func (verbs Verbs) Emotion() Verbs {
	return verbs.filteredByLexFile(VerbEmotion)
}

func (verbs Verbs) Motion() Verbs {
	return verbs.filteredByLexFile(VerbMotion)
}

func (verbs Verbs) Perception() Verbs {
	return verbs.filteredByLexFile(VerbPerception)
}

func (verbs Verbs) Possession() Verbs {
	return verbs.filteredByLexFile(VerbPossession)
}
func (verbs Verbs) Social() Verbs {
	return verbs.filteredByLexFile(VerbSocial)
}

func (verbs Verbs) Stative() Verbs {
	return verbs.filteredByLexFile(VerbStative)
}

func (verbs Verbs) Weather() Verbs {
	return verbs.filteredByLexFile(VerbWeather)
}
