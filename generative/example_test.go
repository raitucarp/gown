package generative_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/generative"
)

func ExampleGrammar_Generate() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	grammar := generative.StandardEnglishGrammar()
	tree, err := grammar.Generate(generative.GeneratorConfig{
		MaxDepth: 4,
		Lexicon:  res,
	})
	if err != nil {
		log.Fatalf("generation failed: %v", err)
	}

	fmt.Printf("Generated sentence has words: %t\n", len(tree.Sentence()) > 0)
	// Output:
	// Generated sentence has words: true
}

func ExampleUnify() {
	fs1 := generative.NewFeatureStructure().
		Set("cat", "NP").
		Set("num", "sg")

	fs2 := generative.NewFeatureStructure().
		Set("num", "sg").
		Set("pers", 3)

	unified, ok := generative.Unify(fs1, fs2)
	fmt.Printf("Unification ok: %t, num: %v\n", ok, unified.Get("num"))
	// Output:
	// Unification ok: true, num: sg
}

func ExampleParseSubcatFrame() {
	frame := generative.ParseSubcatFrame("Somebody ----s something to somebody")
	fmt.Printf("Valency: %s\n", frame.Valency)
	// Output:
	// Valency: transitive
}
