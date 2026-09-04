package syntax_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown/syntax"
)

func buildTestTree() *syntax.SyntaxNode {
	// (S (NP (Det The) (N dog)) (VP (V chased) (NP (Det the) (N cat))))
	np1 := syntax.NewNode("NP",
		syntax.NewLeaf("Det", "The"),
		syntax.NewLeaf("N", "dog"),
	)
	np2 := syntax.NewNode("NP",
		syntax.NewLeaf("Det", "the"),
		syntax.NewLeaf("N", "cat"),
	)
	vp := syntax.NewNode("VP",
		syntax.NewLeaf("V", "chased"),
		np2,
	)
	return syntax.NewNode("S", np1, vp)
}

func TestTreeTraversalAndProperties(t *testing.T) {
	root := buildTestTree()

	if root.Yield() != "The dog chased the cat" {
		t.Errorf("Expected yield 'The dog chased the cat', got '%s'", root.Yield())
	}
	if root.Depth() != 4 {
		t.Errorf("Expected tree depth 4, got %d", root.Depth())
	}
	if root.Size() != 9 {
		t.Errorf("Expected tree size 9, got %d", root.Size())
	}

	leaves := root.Leaves()
	if len(leaves) != 5 {
		t.Errorf("Expected 5 leaves, got %d", len(leaves))
	}

	// Leaf leaves call
	leaf := syntax.NewLeaf("N", "test")
	if len(leaf.Leaves()) != 1 || leaf.Leaves()[0] != leaf {
		t.Errorf("Leaf.Leaves() failed")
	}

	// Siblings & Ancestors
	np1 := root.Children[0]
	sibs := np1.Siblings()
	if len(sibs) != 1 || sibs[0].Label != "VP" {
		t.Errorf("Expected sibling of NP1 to be VP, got %+v", sibs)
	}
	if root.Siblings() != nil {
		t.Errorf("Root should have no siblings")
	}

	ancs := leaves[0].Ancestors()
	if len(ancs) != 2 || ancs[0].Label != "NP" || ancs[1].Label != "S" {
		t.Errorf("Expected ancestors [NP, S], got %+v", ancs)
	}

	// Descendants
	descs := root.Descendants()
	if len(descs) != 8 {
		t.Errorf("Expected 8 descendants for root, got %d", len(descs))
	}

	// Clone with features
	var nilNode *syntax.SyntaxNode
	if nilNode.Clone() != nil {
		t.Errorf("nilNode.Clone() should return nil")
	}
	root.Features["mood"] = "declarative"
	cloned := root.Clone()
	if cloned.Features["mood"] != "declarative" {
		t.Errorf("Cloned features not preserved")
	}
	if cloned.Yield() != root.Yield() || cloned.Size() != root.Size() {
		t.Errorf("Cloned tree differs from original")
	}

	// Multi-level render hitting all branch connectors (isLast and !isLast with prefix)
	multiLevel := syntax.NewNode("S",
		syntax.NewNode("NP",
			syntax.NewNode("Nom", syntax.NewLeaf("Adj", "big"), syntax.NewLeaf("N", "dog")),
			syntax.NewNode("PP", syntax.NewLeaf("P", "with"), syntax.NewLeaf("N", "spots")),
		),
		syntax.NewNode("VP", syntax.NewLeaf("V", "barks")),
	)
	mRender := multiLevel.Render()
	if !strings.Contains(mRender, "│   ") || !strings.Contains(mRender, "├── ") {
		t.Errorf("Render failed to format interior branch connectors:\n%s", mRender)
	}

	// AddChild with nil
	cloned.AddChild(nil)

	// ReplaceChild
	oldNP := cloned.Children[0]
	replacementNP := syntax.NewNode("NP", syntax.NewLeaf("Det", "A"), syntax.NewLeaf("N", "wolf"))
	if !cloned.ReplaceChild(oldNP, replacementNP) {
		t.Errorf("ReplaceChild failed")
	}
	if !strings.HasPrefix(cloned.Yield(), "A wolf") {
		t.Errorf("Expected yield to begin with 'A wolf', got '%s'", cloned.Yield())
	}
	dummy := syntax.NewLeaf("N", "dummy")
	if cloned.ReplaceChild(dummy, replacementNP) {
		t.Errorf("ReplaceChild with non-existent child should return false")
	}

	// BracketedString
	bStr := root.BracketedString()
	if !strings.Contains(bStr, "(S (NP (Det The)") {
		t.Errorf("Unexpected bracketed string: %s", bStr)
	}

	// Render ASCII
	rendered := root.Render()
	if !strings.Contains(rendered, "S") || !strings.Contains(rendered, "chased") {
		t.Errorf("Unexpected rendered tree: %s", rendered)
	}
}

func TestHeadDetectionComprehensive(t *testing.T) {
	// Nil node
	if head := syntax.FindHead(nil); head != nil {
		t.Errorf("FindHead(nil) should be nil")
	}

	// Single leaf
	singleLeaf := syntax.NewLeaf("N", "cat")
	if head := syntax.FindHead(singleLeaf); head != singleLeaf {
		t.Errorf("FindHead(leaf) should return itself")
	}

	// NP with multiple nouns / fallback
	np := syntax.NewNode("NP",
		syntax.NewLeaf("Det", "the"),
		syntax.NewLeaf("Adj", "big"),
		syntax.NewLeaf("N", "house"),
	)
	if head := syntax.FindHead(np); head == nil || head.Terminal != "house" {
		t.Errorf("Expected NP head 'house', got %v", head)
	}

	// NP fallback to rightmost
	npFallback := syntax.NewNode("NP", syntax.NewLeaf("X", "foo"), syntax.NewLeaf("Y", "bar"))
	if head := syntax.FindHead(npFallback); head == nil || head.Terminal != "bar" {
		t.Errorf("Expected NP fallback head 'bar', got %v", head)
	}

	// VP with modal / verb
	vp := syntax.NewNode("VP",
		syntax.NewLeaf("MD", "will"),
		syntax.NewLeaf("V", "sing"),
	)
	if head := syntax.FindHead(vp); head == nil || head.Terminal != "will" {
		t.Errorf("Expected VP head 'will', got %v", head)
	}

	// VP fallback
	vpFallback := syntax.NewNode("VP", syntax.NewLeaf("X", "first"), syntax.NewLeaf("Y", "second"))
	if head := syntax.FindHead(vpFallback); head == nil || head.Terminal != "first" {
		t.Errorf("Expected VP fallback head 'first', got %v", head)
	}

	// PP with preposition
	pp := syntax.NewNode("PP",
		syntax.NewLeaf("IN", "in"),
		syntax.NewNode("NP", syntax.NewLeaf("N", "room")),
	)
	if head := syntax.FindHead(pp); head == nil || head.Terminal != "in" {
		t.Errorf("Expected PP head 'in', got %v", head)
	}

	// PP fallback
	ppFallback := syntax.NewNode("PP", syntax.NewLeaf("X", "prep"))
	if head := syntax.FindHead(ppFallback); head == nil || head.Terminal != "prep" {
		t.Errorf("Expected PP fallback head 'prep', got %v", head)
	}

	// S with VP
	sNode := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "birds")),
		syntax.NewNode("VP", syntax.NewLeaf("V", "fly")),
	)
	if head := syntax.FindHead(sNode); head == nil || head.Terminal != "fly" {
		t.Errorf("Expected S head 'fly', got %v", head)
	}

	// AP / ADJ
	ap := syntax.NewNode("AP",
		syntax.NewLeaf("ADV", "very"),
		syntax.NewLeaf("JJ", "bright"),
	)
	if head := syntax.FindHead(ap); head == nil || head.Terminal != "bright" {
		t.Errorf("Expected AP head 'bright', got %v", head)
	}

	// ADV
	advP := syntax.NewNode("ADVP",
		syntax.NewLeaf("RB", "quickly"),
	)
	if head := syntax.FindHead(advP); head == nil || head.Terminal != "quickly" {
		t.Errorf("Expected ADVP head 'quickly', got %v", head)
	}

	// Other node category default fallback
	otherNode := syntax.NewNode("FRAG",
		syntax.NewLeaf("X", "hello"),
	)
	if head := syntax.FindHead(otherNode); head == nil || head.Terminal != "hello" {
		t.Errorf("Expected default fallback head 'hello', got %v", head)
	}
}

func TestGrammaticalRelationsComprehensive(t *testing.T) {
	// Nil input
	if rels := syntax.ExtractRelations(nil); rels != nil {
		t.Errorf("ExtractRelations(nil) should be nil")
	}

	// Complex ditransitive clause with PP adjunct:
	// "The teacher gave the student a book in the classroom"
	npSubj := syntax.NewNode("NP", syntax.NewLeaf("Det", "The"), syntax.NewLeaf("N", "teacher"))
	npIO := syntax.NewNode("NP", syntax.NewLeaf("Det", "the"), syntax.NewLeaf("N", "student"))
	npDO := syntax.NewNode("NP", syntax.NewLeaf("Det", "a"), syntax.NewLeaf("N", "book"))
	ppAdjunct := syntax.NewNode("PP",
		syntax.NewLeaf("IN", "in"),
		syntax.NewNode("NP", syntax.NewLeaf("Det", "the"), syntax.NewLeaf("N", "classroom")),
	)

	vp := syntax.NewNode("VP",
		syntax.NewLeaf("V", "gave"),
		npIO,
		npDO,
		ppAdjunct,
	)
	sClause := syntax.NewNode("S", npSubj, vp)

	rels := syntax.ExtractRelations(sClause)
	roleMap := make(map[syntax.GrammaticalRole]string)
	for _, r := range rels {
		roleMap[r.Role] = r.Text
	}

	if roleMap[syntax.RoleSubject] != "The teacher" {
		t.Errorf("Expected subject 'The teacher', got '%s'", roleMap[syntax.RoleSubject])
	}
	if roleMap[syntax.RolePredicateVerb] != "gave" {
		t.Errorf("Expected verb 'gave', got '%s'", roleMap[syntax.RolePredicateVerb])
	}
	if roleMap[syntax.RoleIndirectObject] != "the student" {
		t.Errorf("Expected indirect object 'the student', got '%s'", roleMap[syntax.RoleIndirectObject])
	}
	if roleMap[syntax.RoleDirectObject] != "a book" {
		t.Errorf("Expected direct object 'a book', got '%s'", roleMap[syntax.RoleDirectObject])
	}
	if roleMap[syntax.RoleAdjunct] != "in the classroom" {
		t.Errorf("Expected adjunct 'in the classroom', got '%s'", roleMap[syntax.RoleAdjunct])
	}
	if roleMap[syntax.RolePrepObject] != "the classroom" {
		t.Errorf("Expected prep object 'the classroom', got '%s'", roleMap[syntax.RolePrepObject])
	}
}

func TestTransformationsEdgeCases(t *testing.T) {
	// ActiveToPassive errors
	if _, err := syntax.ActiveToPassive(nil); err == nil {
		t.Errorf("ActiveToPassive(nil) should error")
	}
	badTree := syntax.NewNode("S", syntax.NewLeaf("N", "alone"))
	if _, err := syntax.ActiveToPassive(badTree); err == nil {
		t.Errorf("ActiveToPassive on non-transitive tree should error")
	}
	noTransitive := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "dog")),
		syntax.NewNode("VP", syntax.NewLeaf("V", "sleeps")),
	)
	if _, err := syntax.ActiveToPassive(noTransitive); err == nil {
		t.Errorf("ActiveToPassive without direct object should error")
	}

	// PassiveToActive with modifier child and full active conversion
	passiveWithMod := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "mouse")),
		syntax.NewNode("VP",
			syntax.NewLeaf("Aux", "was"),
			syntax.NewLeaf("V", "chased"),
			syntax.NewNode("PP", syntax.NewLeaf("P", "by"), syntax.NewNode("NP", syntax.NewLeaf("N", "cat"))),
			syntax.NewNode("AdvP", syntax.NewLeaf("Adv", "fast")),
		),
	)
	actFromPass, err := syntax.PassiveToActive(passiveWithMod)
	if err != nil {
		t.Fatalf("PassiveToActive with modifier failed: %v", err)
	}
	if !strings.Contains(actFromPass.Yield(), "fast") {
		t.Errorf("Expected active tree to preserve modifier 'fast', got '%s'", actFromPass.Yield())
	}

	// Passive without passive subject
	noSubjPass := syntax.NewNode("S",
		syntax.NewNode("VP", syntax.NewLeaf("V", "was"), syntax.NewLeaf("V", "eaten")),
	)
	if _, err := syntax.PassiveToActive(noSubjPass); err == nil {
		t.Errorf("PassiveToActive without passive subject should error")
	}

	// Passive without verb in VP
	noVerbPass := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "cake")),
		syntax.NewNode("VP",
			syntax.NewNode("PP", syntax.NewLeaf("P", "by"), syntax.NewNode("NP", syntax.NewLeaf("N", "baker"))),
		),
	)
	if _, err := syntax.PassiveToActive(noVerbPass); err == nil {
		t.Errorf("PassiveToActive without verb should error")
	}

	// PassiveToActive errors
	if _, err := syntax.PassiveToActive(nil); err == nil {
		t.Errorf("PassiveToActive(nil) should error")
	}
	if _, err := syntax.PassiveToActive(badTree); err == nil {
		t.Errorf("PassiveToActive on invalid tree should error")
	}
	noBy := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "cat")),
		syntax.NewNode("VP", syntax.NewLeaf("V", "chased")),
	)
	if _, err := syntax.PassiveToActive(noBy); err == nil {
		t.Errorf("PassiveToActive without by-phrase should error")
	}

	// NegateClause
	if _, err := syntax.NegateClause(nil); err == nil {
		t.Errorf("NegateClause(nil) should error")
	}
	if _, err := syntax.NegateClause(badTree); err == nil {
		t.Errorf("NegateClause without VP should error")
	}

	// NegateClause with existing auxiliary / modal
	treeWithAux := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "dog")),
		syntax.NewNode("VP", syntax.NewLeaf("MD", "can"), syntax.NewLeaf("V", "bark")),
	)
	negatedAux, err := syntax.NegateClause(treeWithAux)
	if err != nil {
		t.Fatalf("NegateClause with MD failed: %v", err)
	}
	if !strings.Contains(negatedAux.Yield(), "can not bark") {
		t.Errorf("Expected yield to contain 'can not bark', got '%s'", negatedAux.Yield())
	}

	// ToInterrogative
	if _, err := syntax.ToInterrogative(nil); err == nil {
		t.Errorf("ToInterrogative(nil) should error")
	}
	if _, err := syntax.ToInterrogative(badTree); err == nil {
		t.Errorf("ToInterrogative without NP/VP should error")
	}

	// ToInterrogative with existing Aux
	treeWithAux2 := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("N", "dog")),
		syntax.NewNode("VP", syntax.NewLeaf("Aux", "is"), syntax.NewLeaf("V", "barking")),
	)
	qTree, err := syntax.ToInterrogative(treeWithAux2)
	if err != nil {
		t.Fatalf("ToInterrogative with Aux failed: %v", err)
	}
	if !strings.HasPrefix(qTree.Yield(), "Is dog barking") {
		t.Errorf("Expected question to start with 'Is dog barking', got '%s'", qTree.Yield())
	}
}

func TestMatchingAndSimilarityComprehensive(t *testing.T) {
	root := buildTestTree()

	// FindSubtrees
	if res := syntax.FindSubtrees(nil, "NP"); res != nil {
		t.Errorf("FindSubtrees(nil) should be nil")
	}
	if len(syntax.FindSubtrees(root, "NONEXISTENT")) != 0 {
		t.Errorf("FindSubtrees for missing label should be empty")
	}

	// MatchesRule edge cases
	if syntax.MatchesRule(nil, "S -> NP VP") {
		t.Errorf("MatchesRule(nil) should be false")
	}
	leaf := syntax.NewLeaf("N", "dog")
	if syntax.MatchesRule(leaf, "N -> dog") {
		t.Errorf("MatchesRule on leaf should be false")
	}
	if syntax.MatchesRule(root, "invalid_rule_no_arrow") {
		t.Errorf("MatchesRule with invalid format should be false")
	}
	if syntax.MatchesRule(root, "NP -> Det N") {
		t.Errorf("MatchesRule with mismatched LHS should be false")
	}
	if syntax.MatchesRule(root, "S -> NP VP PP") {
		t.Errorf("MatchesRule with mismatched RHS length should be false")
	}
	// Wildcard matching
	if !syntax.MatchesRule(root, "S -> * VP") {
		t.Errorf("MatchesRule with wildcard * failed")
	}
	if !syntax.MatchesRule(root, "S -> _ _") {
		t.Errorf("MatchesRule with wildcard _ failed")
	}

	// FindByRule
	if res := syntax.FindByRule(nil, "S -> NP VP"); res != nil {
		t.Errorf("FindByRule(nil) should be nil")
	}
	npMatches := syntax.FindByRule(root, "NP -> Det N")
	if len(npMatches) != 2 {
		t.Errorf("Expected 2 matches for 'NP -> Det N', got %d", len(npMatches))
	}

	// ExtractRules inspection
	rules := syntax.ExtractRules(root)
	if len(rules) == 0 {
		t.Errorf("ExtractRules(root) returned empty rules")
	}
	hasSPair := false
	for _, r := range rules {
		if strings.HasPrefix(r, "S ->") {
			hasSPair = true
			break
		}
	}
	if !hasSPair {
		t.Errorf("ExtractRules missing S rule: %v", rules)
	}

	// TreeSimilarity partial overlap
	otherTree := syntax.NewNode("S",
		syntax.NewNode("NP", syntax.NewLeaf("Det", "A"), syntax.NewLeaf("N", "fox")),
		syntax.NewNode("VP", syntax.NewLeaf("V", "jumps")),
	)
	partSim := syntax.TreeSimilarity(root, otherTree)
	if partSim <= 0.0 || partSim >= 1.0 {
		t.Errorf("Expected partial TreeSimilarity in (0, 1), got %.4f", partSim)
	}

	// StructuralDistance with smaller tree first (tests negative abs branch)
	smallerTree := syntax.NewNode("S", syntax.NewLeaf("V", "goes"))
	distSmallLarge := syntax.StructuralDistance(smallerTree, root)
	if distSmallLarge <= 0.0 || distSmallLarge > 1.0 {
		t.Errorf("StructuralDistance(small, large) out of range: %.4f", distSmallLarge)
	}

	// FindSubtrees search on non-matching root
	vSubtrees := syntax.FindSubtrees(root, "V")
	if len(vSubtrees) != 1 {
		t.Errorf("Expected 1 V subtree, got %d", len(vSubtrees))
	}

	// TreeSimilarity edge cases
	if sim := syntax.TreeSimilarity(nil, nil); sim != 1.0 {
		t.Errorf("TreeSimilarity(nil, nil) expected 1.0, got %.2f", sim)
	}
	if sim := syntax.TreeSimilarity(root, nil); sim != 0.0 {
		t.Errorf("TreeSimilarity(root, nil) expected 0.0, got %.2f", sim)
	}
	if sim := syntax.TreeSimilarity(nil, root); sim != 0.0 {
		t.Errorf("TreeSimilarity(nil, root) expected 0.0, got %.2f", sim)
	}
	l1 := syntax.NewLeaf("N", "dog")
	l2 := syntax.NewLeaf("N", "dog")
	l3 := syntax.NewLeaf("N", "cat")
	if sim := syntax.TreeSimilarity(l1, l2); sim != 1.0 {
		t.Errorf("TreeSimilarity identical leaves expected 1.0, got %.2f", sim)
	}
	if sim := syntax.TreeSimilarity(l1, l3); sim != 0.0 {
		t.Errorf("TreeSimilarity different leaves expected 0.0, got %.2f", sim)
	}

	// StructuralDistance edge cases
	if dist := syntax.StructuralDistance(nil, nil); dist != 0.0 {
		t.Errorf("StructuralDistance(nil, nil) expected 0.0, got %.2f", dist)
	}
	if dist := syntax.StructuralDistance(root, nil); dist != 1.0 {
		t.Errorf("StructuralDistance(root, nil) expected 1.0, got %.2f", dist)
	}
	if dist := syntax.StructuralDistance(nil, root); dist != 1.0 {
		t.Errorf("StructuralDistance(nil, root) expected 1.0, got %.2f", dist)
	}
	if dist := syntax.StructuralDistance(root, root); dist != 0.0 {
		t.Errorf("StructuralDistance(root, root) expected 0.0, got %.2f", dist)
	}
	diffTree := syntax.NewNode("S", syntax.NewLeaf("V", "ran"))
	distDiff := syntax.StructuralDistance(root, diffTree)
	if distDiff <= 0.0 || distDiff > 1.0 {
		t.Errorf("StructuralDistance out of range: %.4f", distDiff)
	}
}
