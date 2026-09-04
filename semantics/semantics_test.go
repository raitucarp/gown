package semantics_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/semantics"
)

func TestPolysemyReportAndHomonym(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	report := semantics.AnalyzePolysemy(res, "bank")
	if !report.IsPolysemous {
		t.Errorf("Expected 'bank' to be polysemous")
	}
	if report.TotalSenses <= 1 {
		t.Errorf("Expected multiple senses for 'bank', got %d", report.TotalSenses)
	}
	if report.Entropy <= 0 {
		t.Errorf("Expected positive entropy for 'bank', got %.4f", report.Entropy)
	}

	// Monosemous or missing word
	reportMissing := semantics.AnalyzePolysemy(res, "nonexistentwordxyz123")
	if reportMissing.IsPolysemous || reportMissing.TotalSenses != 0 || reportMissing.Entropy != 0 {
		t.Errorf("Expected empty polysemy report for missing word: %+v", reportMissing)
	}

	// IsHomonym
	if !semantics.IsHomonym(res, "bank") {
		t.Errorf("Expected 'bank' to be classified as homonym/polysemous across entries")
	}
	if semantics.IsHomonym(res, "nonexistentwordxyz123") {
		t.Errorf("Expected false for nonexistent word in IsHomonym")
	}
}

func TestSemanticFieldsClustering(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	words := []string{"dog", "cat", "apple", "banana", "walk", "run", "nonexistentwordxyz123"}
	fields := semantics.ClusterBySemanticField(res, words)

	if len(fields) < 2 {
		t.Errorf("Expected multiple semantic fields, got %d", len(fields))
	}

	hasAnimals := false
	hasFood := false
	hasOther := false
	for _, f := range fields {
		if strings.Contains(f.Domain, "animal") {
			hasAnimals = true
		}
		if strings.Contains(f.Domain, "food") || strings.Contains(f.Domain, "plant") {
			hasFood = true
		}
		if f.Domain == "other" {
			hasOther = true
		}
	}

	if !hasAnimals || !hasOther {
		t.Errorf("Expected animal and other fields in clustering: %+v", fields)
	}
	_ = hasFood
}

func TestLeskDisambiguationComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Missing target word
	missingRes := semantics.DisambiguateLesk(res, "nonexistentwordxyz123", "Some context sentence.")
	if missingRes.BestSynset != nil || missingRes.Score != 0 {
		t.Errorf("Expected nil synset for missing word")
	}

	// Context with financial clues (simplified)
	financeCtx := "He deposited cash and money in the bank vault."
	resultFinance := semantics.DisambiguateLesk(res, "bank", financeCtx, semantics.LeskSimplified)
	if resultFinance.BestSynset == nil {
		t.Fatalf("Expected disambiguated synset for 'bank'")
	}

	// Context with river clues (extended Lesk)
	riverCtx := "We sat on the muddy river bank and watched the water flow."
	resultRiver := semantics.DisambiguateLesk(res, "bank", riverCtx, semantics.LeskExtended)
	if resultRiver.BestSynset == nil {
		t.Fatalf("Expected disambiguated synset for 'bank' under Extended Lesk")
	}
}

func TestLexicalChains(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	text := "The dog barked at the hound. The cat jumped over the mouse. The canine was fast."
	chains := semantics.BuildLexicalChains(res, text)

	if len(chains) == 0 {
		t.Errorf("Expected lexical chains to be constructed")
	}
}

func TestThematicRolesComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Material action verb with instrument (with)
	pas1 := semantics.AssignRoles("cut", "John", "the bread", "a knife", "with")
	if len(pas1.Arguments) != 3 {
		t.Fatalf("Expected 3 arguments, got %d", len(pas1.Arguments))
	}
	if pas1.Arguments[2].Role != semantics.RoleInstrument {
		t.Errorf("Expected instrument role for knife, got %s", pas1.Arguments[2].Role)
	}
	str1 := pas1.String()
	if !strings.Contains(str1, "instrument: a knife") {
		t.Errorf("Unexpected PAS string: %s", str1)
	}

	// Beneficiary (for)
	pas2 := semantics.AssignRoles("bake", "Mary", "a cake", "her mother", "for")
	if pas2.Arguments[2].Role != semantics.RoleBeneficiary {
		t.Errorf("Expected beneficiary role, got %s", pas2.Arguments[2].Role)
	}

	// Source (from)
	pas3 := semantics.AssignRoles("take", "Bob", "money", "the bank", "from")
	if pas3.Arguments[2].Role != semantics.RoleSource {
		t.Errorf("Expected source role, got %s", pas3.Arguments[2].Role)
	}

	// Goal (to)
	pas4 := semantics.AssignRoles("send", "Alice", "a letter", "Paris", "to")
	if pas4.Arguments[2].Role != semantics.RoleGoal {
		t.Errorf("Expected goal role, got %s", pas4.Arguments[2].Role)
	}

	// Mental verb (fear)
	mentalPAS := semantics.AssignRoles("fear", "children", "the dark", "")
	if mentalPAS.Arguments[0].Role != semantics.RoleExperiencer {
		t.Errorf("Expected experiencer, got %s", mentalPAS.Arguments[0].Role)
	}
	if mentalPAS.Arguments[1].Role != semantics.RoleStimulus {
		t.Errorf("Expected stimulus, got %s", mentalPAS.Arguments[1].Role)
	}

	// WordNetThematicRoleCheck
	if !semantics.WordNetThematicRoleCheck(res, "doctor", semantics.RoleAgent) {
		t.Errorf("Expected 'doctor' to satisfy agent/person role")
	}
	if !semantics.WordNetThematicRoleCheck(res, "hammer", semantics.RoleInstrument) {
		t.Errorf("Expected 'hammer' to satisfy instrument/artifact role")
	}
	if !semantics.WordNetThematicRoleCheck(res, "city", semantics.RoleLocation) {
		t.Errorf("Expected 'city' to satisfy location role")
	}
	if !semantics.WordNetThematicRoleCheck(res, "morning", semantics.RoleTime) {
		t.Errorf("Expected 'morning' to satisfy time role")
	}
	if semantics.WordNetThematicRoleCheck(res, "banana", semantics.RoleAgent) {
		t.Errorf("Expected banana NOT to satisfy agent role")
	}
}

func TestSelectionalRestrictionsComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Food restriction
	if !semantics.SatisfiesRestriction(res, "apple", semantics.RestrictFood) {
		t.Errorf("Expected 'apple' to satisfy food restriction")
	}

	// Animate restriction
	if !semantics.SatisfiesRestriction(res, "dog", semantics.RestrictAnimate) {
		t.Errorf("Expected 'dog' to satisfy animate restriction")
	}

	// Human restriction
	if !semantics.SatisfiesRestriction(res, "teacher", semantics.RestrictHuman) {
		t.Errorf("Expected 'teacher' to satisfy human restriction")
	}

	// Liquid restriction
	if !semantics.SatisfiesRestriction(res, "water", semantics.RestrictLiquid) {
		t.Errorf("Expected 'water' to satisfy liquid restriction")
	}

	// Document restriction
	if !semantics.SatisfiesRestriction(res, "book", semantics.RestrictDocument) {
		t.Errorf("Expected 'book' to satisfy document restriction")
	}

	// Artifact restriction
	if !semantics.SatisfiesRestriction(res, "car", semantics.RestrictArtifact) {
		t.Errorf("Expected 'car' to satisfy artifact restriction")
	}

	// Location restriction
	if !semantics.SatisfiesRestriction(res, "city", semantics.RestrictLocation) {
		t.Errorf("Expected 'city' to satisfy location restriction")
	}

	// Physical restriction
	if !semantics.SatisfiesRestriction(res, "stone", semantics.RestrictPhysical) {
		t.Errorf("Expected 'stone' to satisfy physical restriction")
	}

	// Missing noun
	if semantics.SatisfiesRestriction(res, "nonexistentnounxyz123", semantics.RestrictAnimate) {
		t.Errorf("Expected false for missing noun")
	}

	// Selectional violation checks
	// 1. "The rock ate the apple" -> subject violation
	viol1, msg1 := semantics.CheckSelectionalViolation(res, "eat", "rock", "apple")
	if !viol1 || !strings.Contains(msg1, "Subject") {
		t.Errorf("Expected subject violation for 'rock ate apple', got %t (%s)", viol1, msg1)
	}

	// 2. "The boy drank the rock" -> object violation
	viol2, msg2 := semantics.CheckSelectionalViolation(res, "drink", "boy", "rock")
	if !viol2 || !strings.Contains(msg2, "Object") {
		t.Errorf("Expected object violation for 'boy drank rock', got %t (%s)", viol2, msg2)
	}

	// 3. Valid: "The boy drank water" -> no violation
	viol3, _ := semantics.CheckSelectionalViolation(res, "drink", "boy", "water")
	if viol3 {
		t.Errorf("Expected no violation for 'boy drank water'")
	}

	// Profiles for read and drive
	readProf := semantics.GetVerbSelectionalProfile("read")
	if readProf.ObjectRestriction != semantics.RestrictDocument {
		t.Errorf("Expected document restriction for read object")
	}
	driveProf := semantics.GetVerbSelectionalProfile("drive")
	if driveProf.ObjectRestriction != semantics.RestrictArtifact {
		t.Errorf("Expected artifact restriction for drive object")
	}
}

func TestEntailmentAndContradictionComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Identity entailment
	entIdent, tIdent := semantics.CheckEntailment(res, "dog", "dog")
	if !entIdent || tIdent != semantics.EntailmentIdentity {
		t.Errorf("Expected identity entailment for dog->dog")
	}

	// Hyponymy entailment
	entHypo, tHypo := semantics.CheckEntailment(res, "dog", "animal")
	if !entHypo || tHypo != semantics.EntailmentHyponymy {
		t.Errorf("Expected hyponymy entailment for dog->animal")
	}

	// Missing words
	entMiss, _ := semantics.CheckEntailment(res, "nonexistentwordxyz123", "animal")
	if entMiss {
		t.Errorf("Expected false entailment for missing word")
	}

	// Contradiction: hot vs cold
	if !semantics.CheckContradiction(res, "hot", "cold") {
		t.Errorf("Expected 'hot' and 'cold' to contradict each other")
	}
	if semantics.CheckContradiction(res, "hot", "warm") {
		t.Errorf("Expected 'hot' and 'warm' NOT to contradict each other")
	}
	// Same words
	if semantics.CheckContradiction(res, "hot", "hot") {
		t.Errorf("Expected false for identical words in contradiction")
	}
}

func TestPredicateLogicComprehensive(t *testing.T) {
	model := semantics.NewModel(
		semantics.PredicateLogic{Predicate: "Chased", Arguments: []string{"dog", "cat"}},
		semantics.PredicateLogic{Predicate: "Barked", Arguments: []string{"dog"}},
	)

	queryTrue := semantics.PredicateLogic{Predicate: "Chased", Arguments: []string{"dog", "cat"}}
	if !model.Evaluate(queryTrue) {
		t.Errorf("Expected Chased(dog, cat) to evaluate to true")
	}
	if queryTrue.String() != "Chased(dog, cat)" {
		t.Errorf("Unexpected string for queryTrue: %s", queryTrue.String())
	}

	queryFalse := semantics.PredicateLogic{Predicate: "Chased", Arguments: []string{"cat", "dog"}}
	if model.Evaluate(queryFalse) {
		t.Errorf("Expected Chased(cat, dog) to evaluate to false")
	}

	queryNegated := semantics.PredicateLogic{Predicate: "Chased", Arguments: []string{"cat", "dog"}, Negated: true}
	if !model.Evaluate(queryNegated) {
		t.Errorf("Expected ¬Chased(cat, dog) to evaluate to true")
	}
	if queryNegated.String() != "¬Chased(cat, dog)" {
		t.Errorf("Unexpected string for negated query: %s", queryNegated.String())
	}
}
