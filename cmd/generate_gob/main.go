package main

import (
	"encoding/gob"
	"log"

	"github.com/raitucarp/gown"
)

func registerTypes() {
	gob.Register(gown.LexicalResource{})
	gob.Register(gown.Lexicon{})
	gob.Register(gown.LexicalEntry{})
	gob.Register([]gown.LexicalEntry{})
	gob.Register(gown.Synset{})
	gob.Register([]gown.Synset{})
	gob.Register(gown.SyntacticBehaviour{})
	gob.Register([]gown.SyntacticBehaviour{})
	gob.Register(gown.Form{})
	gob.Register([]gown.Form{})
	gob.Register(gown.Sense{})
	gob.Register([]gown.Sense{})
	gob.Register(gown.SenseRelation{})
	gob.Register([]gown.SenseRelation{})
	gob.Register(gown.Pronunciation{})
	gob.Register([]gown.Pronunciation{})
	gob.Register(gown.Example{})
	gob.Register([]gown.Example{})
	gob.Register(gown.Synset{})
	gob.Register([]gown.Synset{})
	gob.Register(gown.SynsetRelation{})
	gob.Register([]gown.SynsetRelation{})
	gob.Register(gown.SyntacticBehaviour{})
	gob.Register([]gown.SyntacticBehaviour{})
}

func main() {

	resource, err := readLexicalResource("./data/wn.xml")
	if err != nil {
		log.Fatal(err)
		return
	}

	registerTypes()

	gobFile, err := createOewnGobFile()
	if err != nil {
		panic(err)
	}
	defer gobFile.Close()

	gzipWriter, bufferedWriter, err := newGzipAndBufferedWriter(gobFile)
	if err != nil {
		panic(err)
	}
	defer gzipWriter.Close()
	defer bufferedWriter.Flush()

	err = encodeOewn(bufferedWriter, resource)
	if err != nil {
		panic(err)
	}

}
