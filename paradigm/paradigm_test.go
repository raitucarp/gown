package paradigm_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/paradigm"
)

func TestDependencyTreeComprehensive(t *testing.T) {
	// "The dog barks loudly"
	tree := paradigm.DependencyTree{
		Nodes: []paradigm.DependencyNode{
			{ID: 1, Form: "The", Lemma: "the", UPOS: "DET", Head: 2, DepRel: "det"},
			{ID: 2, Form: "dog", Lemma: "dog", UPOS: "NOUN", Head: 3, DepRel: "nsubj"},
			{ID: 3, Form: "barks", Lemma: "bark", UPOS: "VERB", Head: 0, DepRel: "root"},
			{ID: 4, Form: "loudly", Lemma: "loudly", UPOS: "ADV", Head: 3, DepRel: "advmod"},
		},
	}

	root := tree.Root()
	if root == nil || root.Form != "barks" {
		t.Errorf("Expected root to be 'barks', got %+v", root)
	}

	deps := tree.DependentsOf(3)
	if len(deps) != 2 {
		t.Errorf("Expected 2 dependents of 'barks', got %d", len(deps))
	}

	// No dependents
	noDeps := tree.DependentsOf(4)
	if len(noDeps) != 0 {
		t.Errorf("Expected 0 dependents for leaf, got %d", len(noDeps))
	}

	// Tree without root
	noRootTree := paradigm.DependencyTree{
		Nodes: []paradigm.DependencyNode{
			{ID: 1, Form: "a", Head: 2},
			{ID: 2, Form: "b", Head: 1},
		},
	}
	if r := noRootTree.Root(); r != nil {
		t.Errorf("Expected nil root for headless tree, got %+v", r)
	}

	conllu := tree.CoNLLU()
	if !strings.Contains(conllu, "barks\tbark\tVERB") {
		t.Errorf("Expected CoNLL-U format to contain verbs: %s", conllu)
	}
}

func TestSemanticFramesComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// 1. Ingestion Frame
	ingestion := paradigm.IngestionFrame()
	if !ingestion.EvokesFrame("eat") || !ingestion.EvokesFrame("devour") {
		t.Errorf("Expected 'eat' and 'devour' to evoke Ingestion frame")
	}
	if ingestion.EvokesFrame("airplane") {
		t.Errorf("Expected 'airplane' NOT to evoke Ingestion frame")
	}
	matchesIng := ingestion.MatchWithWordNet(res)
	if len(matchesIng) == 0 {
		t.Errorf("Expected WordNet synset matches for Ingestion frame")
	}

	// 2. Motion Frame
	motion := paradigm.MotionFrame()
	if !motion.EvokesFrame("run") || !motion.EvokesFrame("walk") || !motion.EvokesFrame("fly") {
		t.Errorf("Expected run, walk, fly to evoke Motion frame")
	}
	if motion.EvokesFrame("pizza") {
		t.Errorf("Expected 'pizza' NOT to evoke Motion frame")
	}
	matchesMotion := motion.MatchWithWordNet(res)
	if len(matchesMotion) == 0 {
		t.Errorf("Expected WordNet synset matches for Motion frame")
	}
	if len(motion.Elements) != 4 {
		t.Errorf("Expected 4 elements in motion frame, got %d", len(motion.Elements))
	}
}

func TestCategorialGrammarComprehensive(t *testing.T) {
	// AtomicCategory tests
	catS := paradigm.CatS
	catNP := paradigm.CatNP
	catN := paradigm.CatN

	if !catS.Equals(paradigm.CatS) {
		t.Errorf("CatS should equal CatS")
	}
	if catS.Equals(catNP) {
		t.Errorf("CatS should not equal CatNP")
	}
	if catS.String() != "S" {
		t.Errorf("CatS.String() expected 'S', got '%s'", catS.String())
	}

	// ComplexCategory string and equals
	left := paradigm.LeftFunctor(catS, catNP) // S \ NP
	right := paradigm.RightFunctor(catNP, catN) // NP / N

	if left.String() != `(S\NP)` {
		t.Errorf("left functor string expected '(S\\NP)', got '%s'", left.String())
	}
	if right.String() != "(NP/N)" {
		t.Errorf("right functor string expected '(NP/N)', got '%s'", right.String())
	}

	if !left.Equals(paradigm.LeftFunctor(catS, catNP)) {
		t.Errorf("left functor should equal identical LeftFunctor")
	}
	if left.Equals(right) {
		t.Errorf("left functor should not equal right functor")
	}
	if left.Equals(catS) {
		t.Errorf("ComplexCategory should not equal AtomicCategory")
	}
	if catS.Equals(left) {
		t.Errorf("AtomicCategory should not equal ComplexCategory")
	}

	// ForwardApply tests
	// 1. Success: (NP / N) + N => NP
	resFA, okFA := paradigm.ForwardApply(right, catN)
	if !okFA || !resFA.Equals(catNP) {
		t.Errorf("ForwardApply failed: got %v, ok=%t", resFA, okFA)
	}
	// 2. Mismatched argument: (NP / N) + S => failure
	_, okMismatchFA := paradigm.ForwardApply(right, catS)
	if okMismatchFA {
		t.Errorf("ForwardApply with mismatched argument should fail")
	}
	// 3. Wrong slash: LeftFunctor passed to ForwardApply => failure
	_, okWrongSlashFA := paradigm.ForwardApply(left, catNP)
	if okWrongSlashFA {
		t.Errorf("ForwardApply with LeftFunctor should fail")
	}
	// 4. Non-complex functor passed to ForwardApply => failure
	_, okAtomicFA := paradigm.ForwardApply(catS, catN)
	if okAtomicFA {
		t.Errorf("ForwardApply with AtomicCategory should fail")
	}

	// BackwardApply tests
	// 1. Success: NP + (S \ NP) => S
	resBA, okBA := paradigm.BackwardApply(catNP, left)
	if !okBA || !resBA.Equals(catS) {
		t.Errorf("BackwardApply failed: got %v, ok=%t", resBA, okBA)
	}
	// 2. Mismatched argument: N + (S \ NP) => failure
	_, okMismatchBA := paradigm.BackwardApply(catN, left)
	if okMismatchBA {
		t.Errorf("BackwardApply with mismatched argument should fail")
	}
	// 3. Wrong slash: RightFunctor passed to BackwardApply => failure
	_, okWrongSlashBA := paradigm.BackwardApply(catN, right)
	if okWrongSlashBA {
		t.Errorf("BackwardApply with RightFunctor should fail")
	}
	// 4. Non-complex functor passed to BackwardApply => failure
	_, okAtomicBA := paradigm.BackwardApply(catNP, catS)
	if okAtomicBA {
		t.Errorf("BackwardApply with AtomicCategory should fail")
	}
}
