package gown

// VerbKind represents a semantic category of verbs based on WordNet's lexfile classification.
type VerbKind string

const (
	// VerbBody is the semantic category for body-related verbs.
	VerbBody VerbKind = "verb.body"
	// VerbChange is the semantic category for change verbs.
	VerbChange VerbKind = "verb.change"
	// VerbCognition is the semantic category for cognition verbs.
	VerbCognition VerbKind = "verb.cognition"
	// VerbCommunication is the semantic category for communication verbs.
	VerbCommunication VerbKind = "verb.communication"
	// VerbCompetition is the semantic category for competition verbs.
	VerbCompetition VerbKind = "verb.competition"
	// VerbConsumption is the semantic category for consumption verbs.
	VerbConsumption VerbKind = "verb.consumption"
	// VerbContact is the semantic category for contact verbs.
	VerbContact VerbKind = "verb.contact"
	// VerbCreation is the semantic category for creation verbs.
	VerbCreation VerbKind = "verb.creation"
	// VerbEmotion is the semantic category for emotion verbs.
	VerbEmotion VerbKind = "verb.emotion"
	// VerbMotion is the semantic category for motion verbs.
	VerbMotion VerbKind = "verb.motion"
	// VerbPerception is the semantic category for perception verbs.
	VerbPerception VerbKind = "verb.perception"
	// VerbPossession is the semantic category for possession verbs.
	VerbPossession VerbKind = "verb.possession"
	// VerbSocial is the semantic category for social verbs.
	VerbSocial VerbKind = "verb.social"
	// VerbStative is the semantic category for stative verbs.
	VerbStative VerbKind = "verb.stative"
	// VerbWeather is the semantic category for weather verbs.
	VerbWeather VerbKind = "verb.weather"
)

// Verb represents a single verb lexical entry.
type Verb LexicalEntry

// Verbs represents a collection of verbs.
type Verbs LexicalEntries

// Verbs returns all verbs from this lexical resource.
func (resource *LexicalResource) Verbs() (verbs Verbs) {
	entries := Verbs(
		LexicalEntries(resource.Lexicon.LexicalEntries).
			filterByPos(VerbPos),
	)

	return entries
}

func (verbs Verbs) filteredByLexFile(kind VerbKind) (
	filteredVerbs Verbs,
) {
	lexicalEntries := verbs.LexicalEntries().
		filterByLexFile(string(kind))
	filteredVerbs = Verbs(lexicalEntries)
	return
}

// LexicalEntries converts this Verbs collection to a LexicalEntries collection.
func (verbs Verbs) LexicalEntries() LexicalEntries {
	return LexicalEntries(verbs)
}

// LexicalEntry converts this Verb to a LexicalEntry.
func (verb Verb) LexicalEntry() LexicalEntry {
	return LexicalEntry(verb)
}

// String returns the written form of this verb.
func (verb Verb) String() string {
	return verb.Lemma.WrittenForm
}

// AllKind returns a map of all verb kinds with their corresponding verb collections.
func (verbs Verbs) AllKind() map[VerbKind]Verbs {
	return map[VerbKind]Verbs{
		VerbBody:          verbs.Body(),
		VerbChange:        verbs.Change(),
		VerbCognition:     verbs.Cognition(),
		VerbCommunication: verbs.Communication(),
		VerbCompetition:   verbs.Competition(),
		VerbConsumption:   verbs.Consumption(),
		VerbContact:       verbs.Contact(),
		VerbCreation:      verbs.Creation(),
		VerbEmotion:       verbs.Emotion(),
		VerbMotion:        verbs.Motion(),
		VerbPerception:    verbs.Perception(),
		VerbPossession:    verbs.Possession(),
		VerbSocial:        verbs.Social(),
		VerbStative:       verbs.Stative(),
		VerbWeather:       verbs.Weather(),
	}
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
