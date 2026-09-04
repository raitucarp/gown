package gown_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
)

func ExampleReadLexicalResource() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	entries := res.LookupExact("dog", gown.NounPos)
	if len(entries) > 0 {
		fmt.Printf("Lemma: %s, Synsets: %d\n", entries[0].Lemma.WrittenForm, len(entries[0].Synsets()))
	}
	// Output:
	// Lemma: dog, Synsets: 7
}

func ExampleLexicalResource_Lookup() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	// Lookup with Morphological normalization (running -> run)
	entries := res.Lookup("running", gown.WithPOS(gown.VerbPos), gown.WithMorphy())
	if len(entries) > 0 {
		fmt.Printf("Base form: %s\n", entries[0].Lemma.WrittenForm)
	}
	// Output:
	// Base form: run
}

func ExampleLexicalResource_Morphy() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	// Reduce irregular plural to singular
	lemmas := res.Morphy("children", gown.NounPos)
	fmt.Printf("Singular of children: %v\n", lemmas)
	// Output:
	// Singular of children: [child]
}
