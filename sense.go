package gown

type Sense struct {
	Adjposition    string          `xml:"adjposition,attr"`
	ID             string          `xml:"id,attr"`
	Subcat         string          `xml:"subcat,attr"`
	Synset         string          `xml:"synset,attr"`
	SenseRelations []SenseRelation `xml:"SenseRelation"`

	resource     *LexicalResource
	lexicalEntry *LexicalEntry
}

func (sense *Sense) GetSynset() *Synset {
	synsetsById := sense.resource.SynsetsById()
	if synset, ok := synsetsById[sense.Synset]; ok {
		return synset
	}
	return nil
}

func (sense *Sense) Definitions() (definitions []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		definitions = append(definitions, synsetsById[sense.Synset].Definitions...)
	}
	return
}

func (sense *Sense) Examples() (examples []string) {
	synsetsById := sense.resource.SynsetsById()
	if _, ok := synsetsById[sense.Synset]; ok {
		for _, example := range synsetsById[sense.Synset].Examples {
			examples = append(examples, example.Text)
		}
	}
	return
}
