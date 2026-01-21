package gown

// NounKind represents a semantic category of nouns based on WordNet's lexfile classification.
type NounKind string

const (
	// NounTops is the semantic category for top-level nouns.
	NounTops NounKind = "noun.Tops"
	// NounAct is the semantic category for nouns representing acts and actions.
	NounAct NounKind = "noun.act"
	// NounAnimal is the semantic category for animal nouns.
	NounAnimal NounKind = "noun.animal"
	// NounArtifact is the semantic category for man-made objects.
	NounArtifact NounKind = "noun.artifact"
	// NounAttribute is the semantic category for attribute nouns.
	NounAttribute NounKind = "noun.attribute"
	// NounBody is the semantic category for body parts.
	NounBody NounKind = "noun.body"
	// NounCognition is the semantic category for cognition-related nouns.
	NounCognition NounKind = "noun.cognition"
	// NounCommunication is the semantic category for communication nouns.
	NounCommunication NounKind = "noun.communication"
	// NounEvent is the semantic category for event nouns.
	NounEvent NounKind = "noun.event"
	// NounFeeling is the semantic category for emotion and feeling nouns.
	NounFeeling NounKind = "noun.feeling"
	// NounFood is the semantic category for food nouns.
	NounFood NounKind = "noun.food"
	// NounGroup is the semantic category for group and collective nouns.
	NounGroup NounKind = "noun.group"
	// NounLocation is the semantic category for location nouns.
	NounLocation NounKind = "noun.location"
	// NounMotive is the semantic category for motive nouns.
	NounMotive NounKind = "noun.motive"
	// NounObject is the semantic category for physical object nouns.
	NounObject NounKind = "noun.object"
	// NounPerson is the semantic category for person nouns.
	NounPerson NounKind = "noun.person"
	// NounPhenomenon is the semantic category for phenomenon nouns.
	NounPhenomenon NounKind = "noun.phenomenon"
	// NounPlant is the semantic category for plant nouns.
	NounPlant NounKind = "noun.plant"
	// NounPossession is the semantic category for possession nouns.
	NounPossession NounKind = "noun.possession"
	// NounProcess is the semantic category for process nouns.
	NounProcess NounKind = "noun.process"
	// NounQuantity is the semantic category for quantity nouns.
	NounQuantity NounKind = "noun.quantity"
	// NounRelation is the semantic category for relation nouns.
	NounRelation NounKind = "noun.relation"
	// NounShape is the semantic category for shape nouns.
	NounShape NounKind = "noun.shape"
	// NounState is the semantic category for state nouns.
	NounState NounKind = "noun.state"
	// NounSubstance is the semantic category for substance nouns.
	NounSubstance NounKind = "noun.substance"
	// NounTime is the semantic category for time nouns.
	NounTime NounKind = "noun.time"
)

// Noun represents a single noun lexical entry.
type Noun LexicalEntry

// Nouns represents a collection of nouns.
type Nouns LexicalEntries

// Nouns returns all nouns from this lexical resource.
func (resource *LexicalResource) Nouns() (nouns Nouns) {
	entries := Nouns(
		LexicalEntries(resource.Lexicon.LexicalEntries).
			filterByPos(NounPos),
	)

	return entries
}

func (nouns Nouns) filterByLexFile(kind NounKind) (
	filteredNouns Nouns,
) {
	lexicalEntries := nouns.LexicalEntries().
		filterByLexFile(string(kind))
	filteredNouns = Nouns(lexicalEntries)
	return
}

// LexicalEntries converts this Nouns collection to a LexicalEntries collection.
func (nouns Nouns) LexicalEntries() LexicalEntries {
	return LexicalEntries(nouns)
}

// LexicalEntry converts this Noun to a LexicalEntry.
func (noun Noun) LexicalEntry() LexicalEntry {
	return LexicalEntry(noun)
}

// String returns the written form of this noun.
func (noun Noun) String() string {
	return noun.Lemma.WrittenForm
}

// AllKind returns a map of all noun kinds with their corresponding noun collections.
func (nouns Nouns) AllKind() map[NounKind]Nouns {
	return map[NounKind]Nouns{
		NounTops:          nouns.Tops(),
		NounAct:           nouns.Act(),
		NounAnimal:        nouns.Animal(),
		NounArtifact:      nouns.Artifact(),
		NounAttribute:     nouns.Attribute(),
		NounBody:          nouns.Body(),
		NounCognition:     nouns.Cognition(),
		NounCommunication: nouns.Communication(),
		NounEvent:         nouns.Event(),
		NounFeeling:       nouns.Feeling(),
		NounFood:          nouns.Food(),
		NounGroup:         nouns.Group(),
		NounLocation:      nouns.Location(),
		NounMotive:        nouns.Motive(),
		NounObject:        nouns.Object(),
		NounPerson:        nouns.Person(),
		NounPhenomenon:    nouns.Phenomenon(),
		NounPlant:         nouns.Plant(),
		NounPossession:    nouns.Possession(),
		NounProcess:       nouns.Process(),
		NounQuantity:      nouns.Quantity(),
		NounRelation:      nouns.Relation(),
		NounShape:         nouns.Shape(),
		NounState:         nouns.State(),
		NounSubstance:     nouns.Substance(),
		NounTime:          nouns.Time(),
	}
}

func (nouns Nouns) Tops() Nouns {
	return nouns.filterByLexFile(NounTops)
}

func (nouns Nouns) Act() Nouns {
	return nouns.filterByLexFile(NounAct)
}

func (nouns Nouns) Animal() Nouns {
	return nouns.filterByLexFile(NounAnimal)
}

func (nouns Nouns) Artifact() Nouns {
	return nouns.filterByLexFile(NounArtifact)
}

func (nouns Nouns) Attribute() Nouns {
	return nouns.filterByLexFile(NounAttribute)
}

func (nouns Nouns) Body() Nouns {
	return nouns.filterByLexFile(NounBody)
}

func (nouns Nouns) Cognition() Nouns {
	return nouns.filterByLexFile(NounCognition)
}

func (nouns Nouns) Communication() Nouns {
	return nouns.filterByLexFile(NounCommunication)
}

func (nouns Nouns) Event() Nouns {
	return nouns.filterByLexFile(NounEvent)
}

func (nouns Nouns) Feeling() Nouns {
	return nouns.filterByLexFile(NounFeeling)
}

func (nouns Nouns) Food() Nouns {
	return nouns.filterByLexFile(NounFood)
}

func (nouns Nouns) Group() Nouns {
	return nouns.filterByLexFile(NounGroup)
}

func (nouns Nouns) Location() Nouns {
	return nouns.filterByLexFile(NounLocation)
}

func (nouns *Nouns) Motive() Nouns {
	return nouns.filterByLexFile(NounMotive)
}

func (nouns Nouns) Object() Nouns {
	return nouns.filterByLexFile(NounObject)
}

func (nouns Nouns) Person() Nouns {
	return nouns.filterByLexFile(NounPerson)
}

func (nouns Nouns) Phenomenon() Nouns {
	return nouns.filterByLexFile(NounPhenomenon)
}

func (nouns Nouns) Plant() Nouns {
	return nouns.filterByLexFile(NounPlant)
}

func (nouns Nouns) Possession() Nouns {
	return nouns.filterByLexFile(NounPossession)
}

func (nouns Nouns) Process() Nouns {
	return nouns.filterByLexFile(NounProcess)
}

func (nouns Nouns) Quantity() Nouns {
	return nouns.filterByLexFile(NounQuantity)
}

func (nouns Nouns) Relation() Nouns {
	return nouns.filterByLexFile(NounRelation)
}

func (nouns Nouns) Shape() Nouns {
	return nouns.filterByLexFile(NounShape)
}

func (nouns Nouns) State() Nouns {
	return nouns.filterByLexFile(NounState)
}

func (nouns Nouns) Substance() Nouns {
	return nouns.filterByLexFile(NounSubstance)
}

func (nouns Nouns) Time() Nouns {
	return nouns.filterByLexFile(NounTime)
}
