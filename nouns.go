package gown

type NounKind string

const (
	NounTops          NounKind = "noun.Tops"
	NounAct           NounKind = "noun.act"
	NounAnimal        NounKind = "noun.animal"
	NounArtifact      NounKind = "noun.artifact"
	NounAttribute     NounKind = "noun.attribute"
	NounBody          NounKind = "noun.body"
	NounCognition     NounKind = "noun.cognition"
	NounCommunication NounKind = "noun.communication"
	NounEvent         NounKind = "noun.event"
	NounFeeling       NounKind = "noun.feeling"
	NounFood          NounKind = "noun.food"
	NounGroup         NounKind = "noun.group"
	NounLocation      NounKind = "noun.location"
	NounMotive        NounKind = "noun.motive"
	NounObject        NounKind = "noun.object"
	NounPerson        NounKind = "noun.person"
	NounPhenomenon    NounKind = "noun.phenomenon"
	NounPlant         NounKind = "noun.plant"
	NounPossession    NounKind = "noun.possession"
	NounProcess       NounKind = "noun.process"
	NounQuantity      NounKind = "noun.quantity"
	NounRelation      NounKind = "noun.relation"
	NounShape         NounKind = "noun.shape"
	NounState         NounKind = "noun.state"
	NounSubstance     NounKind = "noun.substance"
	NounTime          NounKind = "noun.time"
)

type Noun LexicalEntry
type Nouns LexicalEntries

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

func (nouns Nouns) LexicalEntries() LexicalEntries {
	return LexicalEntries(nouns)
}

func (noun Noun) LexicalEntry() LexicalEntry {
	return LexicalEntry(noun)
}

func (noun Noun) String() string {
	return noun.Lemma.WrittenForm
}

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
