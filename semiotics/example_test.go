package semiotics_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/semiotics"
)

func ExampleGenerateSemioticSquare() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	square := semiotics.GenerateSemioticSquare(res, "good")

	fmt.Printf("S1: %s, S2: %s\n", square.S1, square.S2)
	// Output:
	// S1: good, S2: bad
}

func ExampleAnalyzeConnotation() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	con := semiotics.AnalyzeConnotation(res, "noble")
	fmt.Printf("Word: %s, Valence: %s\n", con.Word, con.Valence)
	// Output:
	// Word: noble, Valence: positive
}
