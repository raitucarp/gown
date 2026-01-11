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

type Noun struct {
	*LexicalEntry
	synsets []Synset
}

type Nouns []Noun

func (resource *LexicalResource) Nouns() (entries Nouns) {
	synsetsById := resource.SynsetsById()
	lexEntries := resource.groupEntryByPos(NounPos)
	for _, lexicalEntry := range lexEntries {
		noun := Noun{LexicalEntry: &lexicalEntry}
		for _, sense := range lexicalEntry.Senses {
			synset, ok := synsetsById[sense.Synset]
			if ok {
				noun.synsets = append(noun.synsets, synset)
			}
		}

		entries = append(entries, noun)
	}

	return
}

func (nouns Nouns) filteredByLexFile(kind NounKind) (
	filteredNouns Nouns,
) {

	nounMap := map[string]*Noun{}

	for _, noun := range nouns {
		entryId := noun.ID
		for _, synset := range noun.synsets {
			if synset.Lexfile == string(kind) {
				if _, ok := nounMap[entryId]; !ok {
					newNoun := noun
					newNoun.synsets = []Synset{}
					nounMap[entryId] = &newNoun
				}

				nounMap[entryId].synsets =
					append(nounMap[entryId].synsets, synset)
			}
		}
	}

	for _, noun := range nounMap {
		filteredNouns = append(filteredNouns, *noun)
	}

	return filteredNouns
}

func (nouns Nouns) Tops() Nouns {
	return nouns.filteredByLexFile(NounTops)
}

func (nouns Nouns) Act() Nouns {
	return nouns.filteredByLexFile(NounAct)
}

func (nouns Nouns) Animal() Nouns {
	return nouns.filteredByLexFile(NounAnimal)
}

func (nouns Nouns) Artifact() Nouns {
	return nouns.filteredByLexFile(NounArtifact)
}

func (nouns *Nouns) Attribute() Nouns {
	return nouns.filteredByLexFile(NounAttribute)
}

func (nouns Nouns) Body() Nouns {
	return nouns.filteredByLexFile(NounBody)
}

func (nouns Nouns) Cognition() Nouns {
	return nouns.filteredByLexFile(NounCognition)
}

func (nouns Nouns) Communication() Nouns {
	return nouns.filteredByLexFile(NounCommunication)
}

func (nouns Nouns) Event() Nouns {
	return nouns.filteredByLexFile(NounEvent)
}

func (nouns Nouns) Feeling() Nouns {
	return nouns.filteredByLexFile(NounFeeling)
}

func (nouns Nouns) Food() Nouns {
	return nouns.filteredByLexFile(NounFood)
}

func (nouns Nouns) Group() Nouns {
	return nouns.filteredByLexFile(NounGroup)
}

func (nouns Nouns) Location() Nouns {
	return nouns.filteredByLexFile(NounLocation)
}

func (nouns *Nouns) Motive() Nouns {
	return nouns.filteredByLexFile(NounMotive)
}

func (nouns Nouns) Object() Nouns {
	return nouns.filteredByLexFile(NounObject)
}

func (nouns Nouns) Person() Nouns {
	return nouns.filteredByLexFile(NounPerson)
}

func (nouns Nouns) Phenomenon() Nouns {
	return nouns.filteredByLexFile(NounPhenomenon)
}

func (nouns Nouns) Plant() Nouns {
	return nouns.filteredByLexFile(NounPlant)
}

func (nouns Nouns) Possession() Nouns {
	return nouns.filteredByLexFile(NounPossession)
}

func (nouns Nouns) Process() Nouns {
	return nouns.filteredByLexFile(NounProcess)
}

func (nouns Nouns) Quantity() Nouns {
	return nouns.filteredByLexFile(NounQuantity)
}

func (nouns Nouns) Relation() Nouns {
	return nouns.filteredByLexFile(NounRelation)
}

func (nouns Nouns) Shape() Nouns {
	return nouns.filteredByLexFile(NounShape)
}

func (nouns Nouns) State() Nouns {
	return nouns.filteredByLexFile(NounState)
}

func (nouns Nouns) Substance() Nouns {
	return nouns.filteredByLexFile(NounSubstance)
}

func (nouns Nouns) Time() Nouns {
	return nouns.filteredByLexFile(NounTime)
}
