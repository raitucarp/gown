package gown

type SenseRelationKind string

const (
	SenseRelationAlso            SenseRelationKind = "also"
	SenseRelationAntonym         SenseRelationKind = "antonym"
	SenseRelationDerivation      SenseRelationKind = "derivation"
	SenseRelationExemplifies     SenseRelationKind = "exemplifies"
	SenseRelationIsExemplifiedBy SenseRelationKind = "is_exemplified_by"
	SenseRelationOther           SenseRelationKind = "other"
	SenseRelationParticiple      SenseRelationKind = "participle"
	SenseRelationPertainym       SenseRelationKind = "pertainym"
	SenseRelationSimilar         SenseRelationKind = "similar"
)

type DublinCoreRelType string

const (
	DublinCoreRelTypeAgent       DublinCoreRelType = "agent"
	DublinCoreRelTypeBodyPart    DublinCoreRelType = "body_part"
	DublinCoreRelTypeByMeansOf   DublinCoreRelType = "by_means_of"
	DublinCoreRelTypeDestination DublinCoreRelType = "destination"
	DublinCoreRelTypeEvent       DublinCoreRelType = "event"
	DublinCoreRelTypeInstrument  DublinCoreRelType = "instrument"
	DublinCoreRelTypeLocation    DublinCoreRelType = "location"
	DublinCoreRelTypeMaterial    DublinCoreRelType = "material"
	DublinCoreRelTypeProperty    DublinCoreRelType = "property"
	DublinCoreRelTypeResult      DublinCoreRelType = "result"
	DublinCoreRelTypeState       DublinCoreRelType = "state"
	DublinCoreRelTypeUndergoer   DublinCoreRelType = "undergoer"
	DublinCoreRelTypeUses        DublinCoreRelType = "uses"
	DublinCoreRelTypeVehicle     DublinCoreRelType = "vehicle"
)

type SynsetRelationType string

const (
	SynsetRelationTypeAlso             SynsetRelationType = "also"
	SynsetRelationTypeAttribute        SynsetRelationType = "attribute"
	SynsetRelationTypeCauses           SynsetRelationType = "causes"
	SynsetRelationTypeDomainRegion     SynsetRelationType = "domain_region"
	SynsetRelationTypeDomainTopic      SynsetRelationType = "domain_topic"
	SynsetRelationTypeEntails          SynsetRelationType = "entails"
	SynsetRelationTypeExemplifies      SynsetRelationType = "exemplifies"
	SynsetRelationTypeHasDomainRegion  SynsetRelationType = "has_domain_region"
	SynsetRelationTypeHasDomainTopic   SynsetRelationType = "has_domain_topic"
	SynsetRelationTypeHoloMember       SynsetRelationType = "holo_member"
	SynsetRelationTypeHoloPart         SynsetRelationType = "holo_part"
	SynsetRelationTypeHoloSubstance    SynsetRelationType = "holo_substance"
	SynsetRelationTypeHypernym         SynsetRelationType = "hypernym"
	SynsetRelationTypeHyponym          SynsetRelationType = "hyponym"
	SynsetRelationTypeInstanceHypernym SynsetRelationType = "instance_hypernym"
	SynsetRelationTypeInstanceHyponym  SynsetRelationType = "instance_hyponym"
	SynsetRelationTypeIsCausedBy       SynsetRelationType = "is_caused_by"
	SynsetRelationTypeIsEntailedBy     SynsetRelationType = "is_entailed_by"
	SynsetRelationTypeIsExemplifiedBy  SynsetRelationType = "is_exemplified_by"
	SynsetRelationTypeMeroMember       SynsetRelationType = "mero_member"
	SynsetRelationTypeMeroPart         SynsetRelationType = "mero_part"
	SynsetRelationTypeMeroSubstance    SynsetRelationType = "mero_substance"
	SynsetRelationTypeSimilar          SynsetRelationType = "similar"
)

type LexicalRelation struct {
	entry *LexicalEntry
}

func (rel *LexicalRelation) Synonyms() (entries LexicalEntries) {
	lexicalsById := rel.entry.resource.LexicalsById()
	for _, synset := range rel.entry.Synsets() {
		for _, member := range synset.Members {
			if lexEntry, ok := lexicalsById[member]; ok {
				if lexEntry.ID != rel.entry.ID {
					entries = append(entries, *lexEntry)
				}
			}
		}
	}

	return
}

func (rel *LexicalRelation) Also() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationAlso)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeAlso)...)
	return
}

func (rel *LexicalRelation) Antonyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationAntonym)
	return
}

func (rel *LexicalRelation) Derivations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationDerivation)
	return
}

func (rel *LexicalRelation) Exemplifies() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationExemplifies)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeExemplifies)...)
	return
}

func (rel *LexicalRelation) IsExemplifiedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationIsExemplifiedBy)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeIsExemplifiedBy)...)
	return
}

func (rel *LexicalRelation) Others() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationOther)
	return
}

func (rel *LexicalRelation) Participles() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationParticiple)
	return
}

func (rel *LexicalRelation) Pertainyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationPertainym)
	return
}

func (rel *LexicalRelation) Similars() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationSimilar)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeSimilar)...)
	return
}

func (rel *LexicalRelation) Agents() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeAgent)
	return
}

func (rel *LexicalRelation) BodyParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeBodyPart)
	return
}

func (rel *LexicalRelation) ByMeansOf() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeByMeansOf)
	return
}

func (rel *LexicalRelation) Destinations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeDestination)
	return
}

func (rel *LexicalRelation) Events() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeEvent)
	return
}

func (rel *LexicalRelation) Instruments() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeInstrument)
	return
}

func (rel *LexicalRelation) Locations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeLocation)
	return
}

func (rel *LexicalRelation) Materials() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeMaterial)
	return
}

func (rel *LexicalRelation) Properties() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeProperty)
	return
}

func (rel *LexicalRelation) Results() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeResult)
	return
}

func (rel *LexicalRelation) States() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeState)
	return
}

func (rel *LexicalRelation) Undergoers() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeUndergoer)
	return
}

func (rel *LexicalRelation) Uses() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeUses)
	return
}

func (rel *LexicalRelation) Vehicles() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeVehicle)
	return
}

func (rel *LexicalRelation) Attributes() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeAttribute)
	return
}

func (rel *LexicalRelation) Causes() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeCauses)
	return
}

func (rel *LexicalRelation) DomainRegions() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeDomainRegion)

	return
}

func (rel *LexicalRelation) DomainTopics() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeDomainTopic)
	return
}

func (rel *LexicalRelation) Entails() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeEntails)
	return
}

func (rel *LexicalRelation) HasDomainRegions() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHasDomainRegion)
	return
}

func (rel *LexicalRelation) HasDomainTopics() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHasDomainTopic)
	return
}

func (rel *LexicalRelation) HoloMembers() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloMember)
	return
}

func (rel *LexicalRelation) HoloParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloPart)
	return
}

func (rel *LexicalRelation) HoloSubstances() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloSubstance)
	return
}

func (rel *LexicalRelation) Hypernyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHypernym)
	return
}

func (rel *LexicalRelation) Hyponyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHyponym)
	return
}

func (rel *LexicalRelation) InstanceHypernyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeInstanceHypernym)
	return
}

func (rel *LexicalRelation) InstanceHyponyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeInstanceHyponym)
	return
}

func (rel *LexicalRelation) IsCausedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeIsCausedBy)
	return
}

func (rel *LexicalRelation) IsEntailedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeIsEntailedBy)
	return
}

func (rel *LexicalRelation) MeroMembers() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroMember)
	return
}

func (rel *LexicalRelation) MeroParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroPart)
	return
}

func (rel *LexicalRelation) MeroSubstances() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroSubstance)
	return
}
