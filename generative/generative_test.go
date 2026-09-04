package generative_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/generative"
)

func TestStandardGrammarGenerationComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	grammar := generative.StandardEnglishGrammar()

	// 1. Generation with WordNet lexicon
	tree, err := grammar.Generate(generative.GeneratorConfig{
		MaxDepth: 5,
		Lexicon:  res,
	})
	if err != nil {
		t.Fatalf("Grammar generation failed: %v", err)
	}

	sentence := tree.Sentence()
	if len(sentence) == 0 {
		t.Errorf("Expected non-empty generated sentence")
	}

	bracketed := tree.BracketedString()
	if len(bracketed) == 0 {
		t.Errorf("Expected non-empty bracketed tree")
	}

	asciiTree := tree.RenderTree()
	if len(asciiTree) == 0 {
		t.Errorf("Expected non-empty ASCII tree")
	}

	// 2. Custom grammar without WordNet lexicon
	custom := generative.NewGrammar("S")
	custom.AddRule("S", "NP", "VP")
	custom.AddRule("NP", "cat")
	custom.AddRule("VP", "meows")

	customTree, err := custom.Generate(generative.GeneratorConfig{MaxDepth: 3})
	if err != nil {
		t.Fatalf("Custom grammar generation failed: %v", err)
	}
	if customTree.Sentence() != "cat meows" {
		t.Errorf("Expected 'cat meows', got '%s'", customTree.Sentence())
	}

	// 3. Exceeded max depth error
	recursive := generative.NewGrammar("S")
	recursive.AddRule("S", "S", "and", "S")
	_, errDepth := recursive.Generate(generative.GeneratorConfig{MaxDepth: 2})
	if errDepth == nil {
		t.Errorf("Expected max depth error for infinitely recursive grammar")
	}

	// 4. ParseNode methods
	leaf := &generative.ParseNode{Symbol: "N", Terminal: "dog"}
	if !leaf.IsLeaf() {
		t.Errorf("Expected leaf.IsLeaf() to be true")
	}
	if customTree.IsLeaf() {
		t.Errorf("Root tree should not be leaf")
	}
}

func TestFeatureStructureUnificationComprehensive(t *testing.T) {
	// 1. Basic unification
	fs1 := generative.NewFeatureStructure().
		Set("cat", "NP").
		Set("num", "sg")

	fs2 := generative.NewFeatureStructure().
		Set("num", "sg").
		Set("pers", 3)

	unified, ok := generative.Unify(fs1, fs2)
	if !ok {
		t.Fatalf("Expected unification to succeed")
	}
	if unified.Get("cat") != "NP" || unified.Get("num") != "sg" || unified.Get("pers") != 3 {
		t.Errorf("Unexpected unified feature structure: %v", unified)
	}

	// 2. Incompatible unification (atomic clash)
	fsIncompatible := generative.NewFeatureStructure().Set("num", "pl")
	_, okClash := generative.Unify(fs1, fsIncompatible)
	if okClash {
		t.Errorf("Expected unification with clashing 'num' values to fail")
	}

	// 3. Nil handling
	uNilNil, okNN := generative.Unify(nil, nil)
	if !okNN || uNilNil != nil {
		t.Errorf("Unify(nil, nil) failed")
	}
	uNil1, okN1 := generative.Unify(nil, fs1)
	if !okN1 || uNil1.Get("cat") != "NP" {
		t.Errorf("Unify(nil, fs1) failed")
	}
	u1Nil, ok1N := generative.Unify(fs1, nil)
	if !ok1N || u1Nil.Get("cat") != "NP" {
		t.Errorf("Unify(fs1, nil) failed")
	}

	// 4. Nested FeatureStructure unification
	inner1 := generative.NewFeatureStructure().Set("gender", "fem")
	inner2 := generative.NewFeatureStructure().Set("case", "nom")
	fsNested1 := generative.NewFeatureStructure().Set("agr", inner1)
	fsNested2 := generative.NewFeatureStructure().Set("agr", inner2)

	uNested, okNested := generative.Unify(fsNested1, fsNested2)
	if !okNested {
		t.Fatalf("Expected nested unification to succeed")
	}
	agr := uNested.Get("agr").(generative.FeatureStructure)
	if agr.Get("gender") != "fem" || agr.Get("case") != "nom" {
		t.Errorf("Nested unification failed: %v", agr)
	}

	// 5. Nested unification failure
	innerClash := generative.NewFeatureStructure().Set("gender", "masc")
	fsNestedClash := generative.NewFeatureStructure().Set("agr", innerClash)
	_, okNestedClash := generative.Unify(fsNested1, fsNestedClash)
	if okNestedClash {
		t.Errorf("Expected nested clash to fail")
	}

	// 6. Type mismatch clash (atomic vs FeatureStructure)
	fsMismatch := generative.NewFeatureStructure().Set("agr", "atomic_string")
	_, okMismatch := generative.Unify(fsNested1, fsMismatch)
	if okMismatch {
		t.Errorf("Expected type mismatch clash to fail")
	}

	// 7. String representation
	if str := fs1.String(); !strings.Contains(str, "cat") {
		t.Errorf("Unexpected string for feature structure: %s", str)
	}
}

func TestSubcatFrameParsingComprehensive(t *testing.T) {
	// 1. Intransitive
	f1 := generative.ParseSubcatFrame("Somebody ----s")
	if f1.Valency != generative.ValencyIntransitive {
		t.Errorf("Expected intransitive valency, got %s", f1.Valency)
	}

	// 2. Transitive
	f2 := generative.ParseSubcatFrame("Somebody ----s something")
	if f2.Valency != generative.ValencyTransitive {
		t.Errorf("Expected transitive valency, got %s", f2.Valency)
	}

	// 3. Ditransitive
	f3 := generative.ParseSubcatFrame("Somebody ----s somebody something")
	if f3.Valency != generative.ValencyDitransitive {
		t.Errorf("Expected ditransitive valency, got %s", f3.Valency)
	}

	// 4. ComplexTransitive with CLAUSE
	f4 := generative.ParseSubcatFrame("Somebody ----s that CLAUSE")
	if f4.Valency != generative.ValencyComplexTransitive {
		t.Errorf("Expected complex transitive valency for CLAUSE, got %s", f4.Valency)
	}

	// 5. Intransitive with Prepositional Phrase
	f5 := generative.ParseSubcatFrame("Somebody ----s to somebody")
	if f5.Valency != generative.ValencyIntransitive {
		t.Errorf("Expected intransitive with PP, got %s", f5.Valency)
	}

	// 6. Transitive with PP
	f6 := generative.ParseSubcatFrame("Somebody ----s something to somebody")
	if f6.Valency != generative.ValencyTransitive {
		t.Errorf("Expected transitive with PP, got %s", f6.Valency)
	}

	// 7. VerbSubcatFrames
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}
	verbs := res.LookupVerb("give")
	if len(verbs) > 0 {
		frames := generative.VerbSubcatFrames(verbs[0])
		if len(frames) == 0 {
			t.Logf("Note: 'give' had %d subcat frames parsed from senses", len(frames))
		}
	}
}
