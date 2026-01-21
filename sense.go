package gown

// Sense represents a specific meaning of a word, linking a lexical entry to a synset.
type Sense struct {
	// Adjposition specifies the position of the adjective (postnominal, prenominal, predicative, etc.).
	Adjposition string `xml:"adjposition,attr"`
	// ID uniquely identifies this sense.
	ID string `xml:"id,attr"`
	// Subcat specifies subcategorization information for verbs.
	Subcat string `xml:"subcat,attr"`
	// Synset is the ID of the synset that this sense belongs to.
	Synset string `xml:"synset,attr"`
	// SenseRelations lists semantic relationships to other senses.
	SenseRelations []SenseRelation `xml:"SenseRelation"`

	// resource is an internal reference to the containing LexicalResource.
	resource *LexicalResource
	// lexicalEntry is an internal reference to the containing LexicalEntry.
	lexicalEntry *LexicalEntry
}

// GetSynset returns the synset that this sense belongs to.
func (sense *Sense) GetSynset() *Synset {
	synsetsById := sense.resource.SynsetsById()
	if synset, ok := synsetsById[sense.Synset]; ok {
		return synset
	}
	return nil
}

// Definitions returns all definitions from the synset associated with this sense.
func (sense *Sense) Definitions() (definitions []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		definitions = append(definitions, synsetsById[sense.Synset].Definitions...)
	}
	return
}

// Examples returns all usage examples from the synset associated with this sense.
func (sense *Sense) Examples() (examples []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		for _, example := range synsetsById[sense.Synset].Examples {
			examples = append(examples, example.Text)
		}
	}
	return
}
