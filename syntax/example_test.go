package syntax_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown/syntax"
)

func ExampleActiveToPassive() {
	// Active clause: "The dog chased the cat"
	// (S (NP (Det The) (N dog)) (VP (V chased) (NP (Det the) (N cat))))
	active := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("Det", "The"), syntax.NewLeaf("N", "dog")),
		syntax.NewNode("VP", syntax.NewLeaf("V", "chased"), syntax.NewNode("NP", syntax.NewLeaf("Det", "the"), syntax.NewLeaf("N", "cat"))),
	)

	passive, err := syntax.ActiveToPassive(active)
	if err != nil {
		log.Fatalf("ActiveToPassive failed: %v", err)
	}

	fmt.Println(passive.Yield())
	// Output:
	// the cat was chased by The dog
}

func ExampleExtractRelations() {
	root := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("Det", "The"), syntax.NewLeaf("N", "chef")),
		syntax.NewNode("VP",
			syntax.NewLeaf("V", "cooked"),
			syntax.NewNode("NP", syntax.NewLeaf("Det", "the"), syntax.NewLeaf("N", "meal")),
		),
	)

	rels := syntax.ExtractRelations(root)
	for _, r := range rels {
		fmt.Printf("%s: %s\n", r.Role, r.Text)
	}
	// Output:
	// subject: The chef
	// predicate_verb: cooked
	// direct_object: the meal
}

func ExampleFindHead() {
	np := syntax.NewNode("NP",
		syntax.NewLeaf("Det", "the"),
		syntax.NewLeaf("Adj", "playful"),
		syntax.NewLeaf("N", "dolphin"),
	)

	head := syntax.FindHead(np)
	fmt.Printf("Head noun: %s\n", head.Terminal)
	// Output:
	// Head noun: dolphin
}
