package gown

import (
	"bufio"
	"bytes"
	"compress/gzip"
	_ "embed"
	"encoding/gob"
)

//go:embed data/oewn.gob
var oewn []byte

func registerTypes() {
	gob.Register(LexicalResource{})
	gob.Register(Lexicon{})
	gob.Register(LexicalEntry{})
	gob.Register([]LexicalEntry{})
	gob.Register(Synset{})
	gob.Register([]Synset{})
	gob.Register(SyntacticBehaviour{})
	gob.Register([]SyntacticBehaviour{})
	gob.Register(Form{})
	gob.Register([]Form{})
	gob.Register(Sense{})
	gob.Register([]Sense{})
	gob.Register(SenseRelation{})
	gob.Register([]SenseRelation{})
	gob.Register(Pronunciation{})
	gob.Register([]Pronunciation{})
	gob.Register(Example{})
	gob.Register([]Example{})
	gob.Register(Synset{})
	gob.Register([]Synset{})
	gob.Register(SynsetRelation{})
	gob.Register([]SynsetRelation{})
	gob.Register(SyntacticBehaviour{})
	gob.Register([]SyntacticBehaviour{})
}

// ReadLexicalResource loads the embedded Open English WordNet (OEWN) lexical database
// from the compressed GOB file and initializes all internal mappings and references.
// It returns a pointer to the LexicalResource or an error if decoding fails.
func ReadLexicalResource() (resource *LexicalResource, err error) {
	oewnReader := bytes.NewReader(oewn)

	gzipReader, err := gzip.NewReader(oewnReader)
	if err != nil {
		return
	}
	defer gzipReader.Close()

	bufferedReader := bufio.NewReader(gzipReader)

	registerTypes()
	decoder := gob.NewDecoder(bufferedReader)
	err = decoder.Decode(&resource)
	if err != nil {
		return
	}

	for index := range resource.Lexicon.LexicalEntries {
		resource.Lexicon.LexicalEntries[index].resource = resource
		for senseIndex := range resource.Lexicon.LexicalEntries[index].Senses {
			resource.Lexicon.
				LexicalEntries[index].
				Senses[senseIndex].resource = resource

			resource.Lexicon.
				LexicalEntries[index].
				Senses[senseIndex].lexicalEntry = &resource.Lexicon.
				LexicalEntries[index]

		}
	}

	return
}
