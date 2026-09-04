package semantics_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/semantics"
)

func ExampleDisambiguateLesk() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	// Disambiguate "bank" in a financial context
	sentence := "He deposited his savings and money at the local bank."
	result := semantics.DisambiguateLesk(res, "bank", sentence)

	fmt.Printf("Disambiguated word: %s\n", result.Word)
	fmt.Printf("Found definition: %t\n", len(result.Definition) > 0)
	// Output:
	// Disambiguated word: bank
	// Found definition: true
}

func ExampleCheckSelectionalViolation() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	// "The rock ate the apple" -> rock is not animate
	violation, msg := semantics.CheckSelectionalViolation(res, "eat", "rock", "apple")
	fmt.Printf("Violation detected: %t\n", violation)
	fmt.Printf("Has message: %t\n", len(msg) > 0)
	// Output:
	// Violation detected: true
	// Has message: true
}

func ExampleAssignRoles() {
	pas := semantics.AssignRoles("devour", "wolf", "meat", "in the forest", "in")
	fmt.Println(pas.String())
	// Output:
	// devour(agent: wolf, patient: meat, location: in the forest)
}
