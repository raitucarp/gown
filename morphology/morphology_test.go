package morphology_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/morphology"
)

func TestDetectInflectionsComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// 1. Empty word
	if infos := morphology.DetectInflections(res, ""); infos != nil {
		t.Errorf("DetectInflections('') expected nil, got %+v", infos)
	}

	// 2. Base form
	baseInfos := morphology.DetectInflections(res, "cat")
	foundBase := false
	for _, inf := range baseInfos {
		if inf.Kind == morphology.InflectionBase && inf.BaseLemma == "cat" {
			foundBase = true
			break
		}
	}
	if !foundBase {
		t.Errorf("Expected base inflection for 'cat', got %+v", baseInfos)
	}

	// 3. Plural noun: irregular "children" -> "child"
	pluralInfos := morphology.DetectInflections(res, "children")
	foundPlural := false
	for _, inf := range pluralInfos {
		if inf.Kind == morphology.InflectionPlural && inf.BaseLemma == "child" {
			foundPlural = true
			break
		}
	}
	if !foundPlural {
		t.Errorf("Expected plural inflection 'child' for 'children', got: %+v", pluralInfos)
	}

	// 4. Regular plural noun: "dogs" -> "dog"
	dogsInfos := morphology.DetectInflections(res, "dogs")
	foundDogs := false
	for _, inf := range dogsInfos {
		if inf.Kind == morphology.InflectionPlural && inf.BaseLemma == "dog" {
			foundDogs = true
			break
		}
	}
	if !foundDogs {
		t.Errorf("Expected plural 'dog' for 'dogs', got: %+v", dogsInfos)
	}

	// 5. Progressive verb: "running" -> "run"
	runningInfos := morphology.DetectInflections(res, "running")
	foundProg := false
	for _, inf := range runningInfos {
		if inf.Kind == morphology.InflectionProgressive && inf.BaseLemma == "run" {
			foundProg = true
			break
		}
	}
	if !foundProg {
		t.Errorf("Expected progressive inflection 'run' for 'running', got: %+v", runningInfos)
	}

	// 6. Past tense verb: "walked" -> "walk"
	walkedInfos := morphology.DetectInflections(res, "walked")
	foundPast := false
	for _, inf := range walkedInfos {
		if inf.Kind == morphology.InflectionPastTense && inf.BaseLemma == "walk" {
			foundPast = true
			break
		}
	}
	if !foundPast {
		t.Errorf("Expected past tense 'walk' for 'walked', got: %+v", walkedInfos)
	}

	// 7. 3rd person singular verb: "runs" -> "run"
	runsInfos := morphology.DetectInflections(res, "runs")
	found3rd := false
	for _, inf := range runsInfos {
		if inf.Kind == morphology.Inflection3rdPersonSg && inf.BaseLemma == "run" {
			found3rd = true
			break
		}
	}
	if !found3rd {
		t.Errorf("Expected 3rd person singular 'run' for 'runs', got: %+v", runsInfos)
	}

	// 8. Comparative adjective: "faster" -> "fast"
	fasterInfos := morphology.DetectInflections(res, "faster")
	foundComp := false
	for _, inf := range fasterInfos {
		if inf.Kind == morphology.InflectionComparative && inf.BaseLemma == "fast" {
			foundComp = true
			break
		}
	}
	if !foundComp {
		t.Errorf("Expected comparative 'fast' for 'faster', got: %+v", fasterInfos)
	}

	// 9. Superlative adjective: "fastest" -> "fast"
	fastestInfos := morphology.DetectInflections(res, "fastest")
	foundSuper := false
	for _, inf := range fastestInfos {
		if inf.Kind == morphology.InflectionSuperlative && inf.BaseLemma == "fast" {
			foundSuper = true
			break
		}
	}
	if !foundSuper {
		t.Errorf("Expected superlative 'fast' for 'fastest', got: %+v", fastestInfos)
	}
}

func TestGenerateLexicalFamilyComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	fam := morphology.GenerateLexicalFamily(res, "act")
	if len(fam.Members) == 0 {
		t.Fatalf("Expected lexical family members for 'act'")
	}
	if fam.Root != "act" {
		t.Errorf("Expected fam.Root 'act', got '%s'", fam.Root)
	}

	hasDerivation := false
	hasSuffix := false
	hasPrefix := false
	for _, m := range fam.Members {
		if m.Relation == "derivation" {
			hasDerivation = true
		}
		if strings.HasPrefix(m.Relation, "suffix:") {
			hasSuffix = true
		}
		if strings.HasPrefix(m.Relation, "prefix:") {
			hasPrefix = true
		}
	}

	if !hasDerivation {
		t.Errorf("Expected derivational relations in lexical family")
	}
	if !hasSuffix {
		t.Errorf("Expected suffix relations in lexical family")
	}
	if !hasPrefix {
		t.Errorf("Expected prefix relations in lexical family")
	}

	// Nonexistent word family
	famEmpty := morphology.GenerateLexicalFamily(res, "xyznonexistentword")
	if len(famEmpty.Members) != 0 {
		t.Errorf("Expected 0 members for nonexistent word family, got %d", len(famEmpty.Members))
	}
}

func TestSplitCompoundComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// 1. Short word (< 4 runes) returns nil
	if splits := morphology.SplitCompound(res, "cat"); splits != nil {
		t.Errorf("Expected nil for short word 'cat', got %+v", splits)
	}

	// 2. Closed compound "sunflower" -> "sun" + "flower"
	splits := morphology.SplitCompound(res, "sunflower")
	found := false
	for _, s := range splits {
		if len(s.Parts) == 2 && s.Parts[0] == "sun" && s.Parts[1] == "flower" {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Expected compound split ['sun', 'flower'] for 'sunflower', got %+v", splits)
	}

	// 3. Hyphenated compound: "ice-cream"
	hyphenSplits := morphology.SplitCompound(res, "ice-cream")
	if len(hyphenSplits) == 0 || hyphenSplits[0].Parts[0] != "ice" || hyphenSplits[0].Parts[1] != "cream" {
		t.Errorf("Expected hyphenated split for 'ice-cream', got %+v", hyphenSplits)
	}

	// 4. Space-separated compound: "ice cream"
	spaceSplits := morphology.SplitCompound(res, "ice cream")
	if len(spaceSplits) == 0 || spaceSplits[0].Parts[0] != "ice" || spaceSplits[0].Parts[1] != "cream" {
		t.Errorf("Expected space-separated split for 'ice cream', got %+v", spaceSplits)
	}

	// 5. Unsplittable / nonexistent word
	if unSplits := morphology.SplitCompound(res, "zzzzzzzz"); len(unSplits) != 0 {
		t.Errorf("Expected 0 splits for gibberish word, got %d", len(unSplits))
	}
}
