package expansion_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/expansion"
)

func ExampleExpand() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	tree, err := expansion.Expand(res, "wolf",
		expansion.WithMaxDepth(2),
		expansion.WithMaxNodes(10),
		expansion.WithPOS(gown.NounPos),
	)
	if err != nil {
		log.Fatalf("expand failed: %v", err)
	}

	fmt.Printf("Tree expanded successfully: %t\n", tree.TotalNodes > 1)
	// Output:
	// Tree expanded successfully: true
}

func ExampleExpandDefinition() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	tree, err := expansion.ExpandDefinition(res, "dog",
		expansion.WithMaxDepth(2),
		expansion.WithMaxNodes(15),
		expansion.WithPOS(gown.NounPos),
	)
	if err != nil {
		log.Fatalf("expand definition failed: %v", err)
	}

	fmt.Printf("Definition tokens found: %t\n", len(tree.Root.Children) > 0)
	// Output:
	// Definition tokens found: true
}
