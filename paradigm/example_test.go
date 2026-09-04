package paradigm_test

import (
	"fmt"

	"github.com/raitucarp/gown/paradigm"
)

func ExampleDependencyTree() {
	// Represent "The dog barks" in Universal Dependencies
	tree := paradigm.DependencyTree{
		Nodes: []paradigm.DependencyNode{
			{ID: 1, Form: "The", Lemma: "the", UPOS: "DET", Head: 2, DepRel: "det"},
			{ID: 2, Form: "dog", Lemma: "dog", UPOS: "NOUN", Head: 3, DepRel: "nsubj"},
			{ID: 3, Form: "barks", Lemma: "bark", UPOS: "VERB", Head: 0, DepRel: "root"},
		},
	}

	root := tree.Root()
	fmt.Printf("Root predicate: %s\n", root.Form)
	// Output:
	// Root predicate: barks
}

func ExampleForwardApply() {
	// Determiner: NP / N
	det := paradigm.RightFunctor(paradigm.CatNP, paradigm.CatN)
	// Noun: N
	n := paradigm.CatN

	// Forward application: (NP / N) + N => NP
	result, ok := paradigm.ForwardApply(det, n)
	fmt.Printf("Applied: %t, Result: %s\n", ok, result.String())
	// Output:
	// Applied: true, Result: NP
}
