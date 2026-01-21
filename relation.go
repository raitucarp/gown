package gown

// SenseRelationKind represents a type of semantic relationship between senses.
type SenseRelationKind string

const (
	// SenseRelationAlso indicates an "also" relationship between senses.
	SenseRelationAlso SenseRelationKind = "also"
	// SenseRelationAntonym indicates an antonymous relationship between senses.
	SenseRelationAntonym SenseRelationKind = "antonym"
	// SenseRelationDerivation indicates a derivation relationship between senses.
	SenseRelationDerivation SenseRelationKind = "derivation"
	// SenseRelationExemplifies indicates an exemplification relationship.
	SenseRelationExemplifies SenseRelationKind = "exemplifies"
	// SenseRelationIsExemplifiedBy indicates being exemplified by another sense.
	SenseRelationIsExemplifiedBy SenseRelationKind = "is_exemplified_by"
	// SenseRelationOther indicates an unspecified relationship.
	SenseRelationOther SenseRelationKind = "other"
	// SenseRelationParticiple indicates a participle relationship.
	SenseRelationParticiple SenseRelationKind = "participle"
	// SenseRelationPertainym indicates a pertainym relationship.
	SenseRelationPertainym SenseRelationKind = "pertainym"
	// SenseRelationSimilar indicates a similarity relationship.
	SenseRelationSimilar SenseRelationKind = "similar"
)

// DublinCoreRelType represents a Dublin Core type of semantic relationship.
type DublinCoreRelType string

const (
	// DublinCoreRelTypeAgent indicates an agent relationship.
	DublinCoreRelTypeAgent DublinCoreRelType = "agent"
	// DublinCoreRelTypeBodyPart indicates a body part relationship.
	DublinCoreRelTypeBodyPart DublinCoreRelType = "body_part"
	// DublinCoreRelTypeByMeansOf indicates a "by means of" relationship.
	DublinCoreRelTypeByMeansOf DublinCoreRelType = "by_means_of"
	// DublinCoreRelTypeDestination indicates a destination relationship.
	DublinCoreRelTypeDestination DublinCoreRelType = "destination"
	// DublinCoreRelTypeEvent indicates an event relationship.
	DublinCoreRelTypeEvent DublinCoreRelType = "event"
	// DublinCoreRelTypeInstrument indicates an instrument relationship.
	DublinCoreRelTypeInstrument DublinCoreRelType = "instrument"
	// DublinCoreRelTypeLocation indicates a location relationship.
	DublinCoreRelTypeLocation DublinCoreRelType = "location"
	// DublinCoreRelTypeMaterial indicates a material relationship.
	DublinCoreRelTypeMaterial DublinCoreRelType = "material"
	// DublinCoreRelTypeProperty indicates a property relationship.
	DublinCoreRelTypeProperty DublinCoreRelType = "property"
	// DublinCoreRelTypeResult indicates a result relationship.
	DublinCoreRelTypeResult DublinCoreRelType = "result"
	// DublinCoreRelTypeState indicates a state relationship.
	DublinCoreRelTypeState DublinCoreRelType = "state"
	// DublinCoreRelTypeUndergoer indicates an undergoer relationship.
	DublinCoreRelTypeUndergoer DublinCoreRelType = "undergoer"
	// DublinCoreRelTypeUses indicates a "uses" relationship.
	DublinCoreRelTypeUses DublinCoreRelType = "uses"
	// DublinCoreRelTypeVehicle indicates a vehicle relationship.
	DublinCoreRelTypeVehicle DublinCoreRelType = "vehicle"
)

// SynsetRelationType represents a type of semantic relationship between synsets.
type SynsetRelationType string

const (
	// SynsetRelationTypeAlso indicates an "also" relationship.
	SynsetRelationTypeAlso SynsetRelationType = "also"
	// SynsetRelationTypeAttribute indicates an attribute relationship.
	SynsetRelationTypeAttribute SynsetRelationType = "attribute"
	// SynsetRelationTypeCauses indicates a causal relationship.
	SynsetRelationTypeCauses SynsetRelationType = "causes"
	// SynsetRelationTypeDomainRegion indicates a domain region relationship.
	SynsetRelationTypeDomainRegion SynsetRelationType = "domain_region"
	// SynsetRelationTypeDomainTopic indicates a domain topic relationship.
	SynsetRelationTypeDomainTopic SynsetRelationType = "domain_topic"
	// SynsetRelationTypeEntails indicates an entailment relationship.
	SynsetRelationTypeEntails SynsetRelationType = "entails"
	// SynsetRelationTypeExemplifies indicates an exemplification relationship.
	SynsetRelationTypeExemplifies SynsetRelationType = "exemplifies"
	// SynsetRelationTypeHasDomainRegion indicates a "has domain region" relationship.
	SynsetRelationTypeHasDomainRegion SynsetRelationType = "has_domain_region"
	// SynsetRelationTypeHasDomainTopic indicates a "has domain topic" relationship.
	SynsetRelationTypeHasDomainTopic SynsetRelationType = "has_domain_topic"
	// SynsetRelationTypeHoloMember indicates a holonym member relationship.
	SynsetRelationTypeHoloMember SynsetRelationType = "holo_member"
	// SynsetRelationTypeHoloPart indicates a holonym part relationship.
	SynsetRelationTypeHoloPart SynsetRelationType = "holo_part"
	// SynsetRelationTypeHoloSubstance indicates a holonym substance relationship.
	SynsetRelationTypeHoloSubstance SynsetRelationType = "holo_substance"
	// SynsetRelationTypeHypernym indicates a hypernym (more general) relationship.
	SynsetRelationTypeHypernym SynsetRelationType = "hypernym"
	// SynsetRelationTypeHyponym indicates a hyponym (more specific) relationship.
	SynsetRelationTypeHyponym SynsetRelationType = "hyponym"
	// SynsetRelationTypeInstanceHypernym indicates an instance hypernym relationship.
	SynsetRelationTypeInstanceHypernym SynsetRelationType = "instance_hypernym"
	// SynsetRelationTypeInstanceHyponym indicates an instance hyponym relationship.
	SynsetRelationTypeInstanceHyponym SynsetRelationType = "instance_hyponym"
	// SynsetRelationTypeIsCausedBy indicates a "is caused by" relationship.
	SynsetRelationTypeIsCausedBy SynsetRelationType = "is_caused_by"
	// SynsetRelationTypeIsEntailedBy indicates a "is entailed by" relationship.
	SynsetRelationTypeIsEntailedBy SynsetRelationType = "is_entailed_by"
	// SynsetRelationTypeIsExemplifiedBy indicates a "is exemplified by" relationship.
	SynsetRelationTypeIsExemplifiedBy SynsetRelationType = "is_exemplified_by"
	// SynsetRelationTypeMeroMember indicates a meronym member relationship.
	SynsetRelationTypeMeroMember SynsetRelationType = "mero_member"
	// SynsetRelationTypeMeroPart indicates a meronym part relationship.
	SynsetRelationTypeMeroPart SynsetRelationType = "mero_part"
	// SynsetRelationTypeMeroSubstance indicates a meronym substance relationship.
	SynsetRelationTypeMeroSubstance SynsetRelationType = "mero_substance"
	// SynsetRelationTypeSimilar indicates a similarity relationship.
	SynsetRelationTypeSimilar SynsetRelationType = "similar"
)

// LexicalRelation provides methods to find semantic relationships for a lexical entry.
type LexicalRelation struct {
	// entry is the lexical entry for which to find relations.
	entry *LexicalEntry
}

// Synonyms returns all synonyms of this entry (other members of the same synsets).
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

// Also returns entries related via "also" sense and synset relations.
func (rel *LexicalRelation) Also() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationAlso)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeAlso)...)
	return
}

// Antonyms returns antonyms of this entry.
func (rel *LexicalRelation) Antonyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationAntonym)
	return
}

// Derivations returns words derived from this entry.
func (rel *LexicalRelation) Derivations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationDerivation)
	return
}

// Exemplifies returns entries exemplified by this entry.
func (rel *LexicalRelation) Exemplifies() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationExemplifies)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeExemplifies)...)
	return
}

// IsExemplifiedBy returns entries that exemplify this entry.
func (rel *LexicalRelation) IsExemplifiedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationIsExemplifiedBy)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeIsExemplifiedBy)...)
	return
}

// Others returns entries related via other unspecified relations.
func (rel *LexicalRelation) Others() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationOther)
	return
}

// Participles returns participle forms of this entry.
func (rel *LexicalRelation) Participles() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationParticiple)
	return
}

// Pertainyms returns pertainyms of this entry.
func (rel *LexicalRelation) Pertainyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationPertainym)
	return
}

// Similars returns entries similar to this entry.
func (rel *LexicalRelation) Similars() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseRelationType(SenseRelationSimilar)
	entries = append(entries, rel.entry.filterBySynsetRelation(SynsetRelationTypeSimilar)...)
	return
}

// Agents returns entries that are agents related to this entry.
func (rel *LexicalRelation) Agents() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeAgent)
	return
}

// BodyParts returns body parts related to this entry.
func (rel *LexicalRelation) BodyParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeBodyPart)
	return
}

// ByMeansOf returns entries related by "by means of" relationship.
func (rel *LexicalRelation) ByMeansOf() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeByMeansOf)
	return
}

// Destinations returns destination entries related to this entry.
func (rel *LexicalRelation) Destinations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeDestination)
	return
}

// Events returns events related to this entry.
func (rel *LexicalRelation) Events() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeEvent)
	return
}

// Instruments returns instruments related to this entry.
func (rel *LexicalRelation) Instruments() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeInstrument)
	return
}

// Locations returns locations related to this entry.
func (rel *LexicalRelation) Locations() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeLocation)
	return
}

// Materials returns materials related to this entry.
func (rel *LexicalRelation) Materials() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeMaterial)
	return
}

// Properties returns properties related to this entry.
func (rel *LexicalRelation) Properties() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeProperty)
	return
}

// Results returns results related to this entry.
func (rel *LexicalRelation) Results() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeResult)
	return
}

// States returns states related to this entry.
func (rel *LexicalRelation) States() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeState)
	return
}

// Undergoers returns undergoers related to this entry.
func (rel *LexicalRelation) Undergoers() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeUndergoer)
	return
}

// Uses returns entries that use this entry or are used by it.
func (rel *LexicalRelation) Uses() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeUses)
	return
}

// Vehicles returns vehicles related to this entry.
func (rel *LexicalRelation) Vehicles() (entries LexicalEntries) {
	entries = rel.entry.filterBySenseDublinCoreType(DublinCoreRelTypeVehicle)
	return
}

// Attributes returns attributes related to this entry.
func (rel *LexicalRelation) Attributes() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeAttribute)
	return
}

// Causes returns entries that cause this entry.
func (rel *LexicalRelation) Causes() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeCauses)
	return
}

// DomainRegions returns domain regions related to this entry.
func (rel *LexicalRelation) DomainRegions() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeDomainRegion)

	return
}

// DomainTopics returns domain topics related to this entry.
func (rel *LexicalRelation) DomainTopics() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeDomainTopic)
	return
}

// Entails returns entries that this entry entails.
func (rel *LexicalRelation) Entails() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeEntails)
	return
}

// HasDomainRegions returns entries with domain regions related to this entry.
func (rel *LexicalRelation) HasDomainRegions() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHasDomainRegion)
	return
}

// HasDomainTopics returns entries with domain topics related to this entry.
func (rel *LexicalRelation) HasDomainTopics() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHasDomainTopic)
	return
}

// HoloMembers returns holonym members of this entry.
func (rel *LexicalRelation) HoloMembers() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloMember)
	return
}

// HoloParts returns holonym parts of this entry.
func (rel *LexicalRelation) HoloParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloPart)
	return
}

// HoloSubstances returns holonym substances of this entry.
func (rel *LexicalRelation) HoloSubstances() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHoloSubstance)
	return
}

// Hypernyms returns hypernyms (more general terms) of this entry.
func (rel *LexicalRelation) Hypernyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHypernym)
	return
}

// Hyponyms returns hyponyms (more specific terms) of this entry.
func (rel *LexicalRelation) Hyponyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeHyponym)
	return
}

// InstanceHypernyms returns instance hypernyms of this entry.
func (rel *LexicalRelation) InstanceHypernyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeInstanceHypernym)
	return
}

// InstanceHyponyms returns instance hyponyms of this entry.
func (rel *LexicalRelation) InstanceHyponyms() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeInstanceHyponym)
	return
}

// IsCausedBy returns entries that cause this entry.
func (rel *LexicalRelation) IsCausedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeIsCausedBy)
	return
}

// IsEntailedBy returns entries entailed by this entry.
func (rel *LexicalRelation) IsEntailedBy() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeIsEntailedBy)
	return
}

// MeroMembers returns meronym members of this entry.
func (rel *LexicalRelation) MeroMembers() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroMember)
	return
}

// MeroParts returns meronym parts of this entry.
func (rel *LexicalRelation) MeroParts() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroPart)
	return
}

// MeroSubstances returns meronym substances of this entry.
func (rel *LexicalRelation) MeroSubstances() (entries LexicalEntries) {
	entries = rel.entry.filterBySynsetRelation(SynsetRelationTypeMeroSubstance)
	return
}
