package morphology_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/morphology"
)

func ExampleDetectInflections() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	infos := morphology.DetectInflections(res, "children")
	for _, info := range infos {
		if info.Kind == morphology.InflectionPlural {
			fmt.Printf("Base: %s, Kind: %s\n", info.BaseLemma, info.Kind)
		}
	}
	// Output:
	// Base: child, Kind: plural
}

func ExampleSplitCompound() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	splits := morphology.SplitCompound(res, "sunflower")
	if len(splits) > 0 {
		fmt.Printf("Split: %s + %s\n", splits[0].Parts[0], splits[0].Parts[1])
	}
	// Output:
	// Split: sun + flower
}
