package pragmatics_test

import (
	"testing"

	"github.com/raitucarp/gown/pragmatics"
)

func TestContextAndDeixisComprehensive(t *testing.T) {
	ctx := pragmatics.NewContext("Alice", "Bob")
	ctx.Location = "Room 101"
	ctx.Time = "noon"

	// Presupposition storage in common ground
	ctx.AddPresupposition("Earth is round")
	if !ctx.Presupposed("Earth is round") {
		t.Errorf("Expected 'Earth is round' to be in common ground")
	}
	if ctx.Presupposed("Earth is flat") {
		t.Errorf("Expected 'Earth is flat' NOT to be in common ground")
	}

	utterance := "I will meet you here now."
	deictics := pragmatics.ResolveDeixis(utterance, ctx)

	if len(deictics) != 4 {
		t.Fatalf("Expected 4 deictic expressions (I, you, here, now), got %d", len(deictics))
	}

	for _, d := range deictics {
		switch d.Word {
		case "I":
			if d.ResolvedTo != "Alice" {
				t.Errorf("Expected 'I' to resolve to 'Alice', got '%s'", d.ResolvedTo)
			}
		case "you":
			if d.ResolvedTo != "Bob" {
				t.Errorf("Expected 'you' to resolve to 'Bob', got '%s'", d.ResolvedTo)
			}
		case "here":
			if d.ResolvedTo != "Room 101" {
				t.Errorf("Expected 'here' to resolve to 'Room 101', got '%s'", d.ResolvedTo)
			}
		case "now":
			if d.ResolvedTo != "noon" {
				t.Errorf("Expected 'now' to resolve to 'noon', got '%s'", d.ResolvedTo)
			}
		}
	}
}

func TestSpeechActClassificationComprehensive(t *testing.T) {
	tests := []struct {
		utterance string
		expected  pragmatics.SpeechActClass
	}{
		{"I state that the Earth revolves around the Sun.", pragmatics.ActAssertive},
		{"We claim that the hypothesis is valid.", pragmatics.ActAssertive},
		{"The sky is blue.", pragmatics.ActAssertive},
		{"Please close the window.", pragmatics.ActDirective},
		{"Where is the nearest station?", pragmatics.ActDirective},
		{"Could you help me with this?", pragmatics.ActDirective},
		{"I promise that I will finish the task on time.", pragmatics.ActCommissive},
		{"I will call you later tonight.", pragmatics.ActCommissive},
		{"Thank you very much for your help.", pragmatics.ActExpressive},
		{"Sorry for the delay.", pragmatics.ActExpressive},
		{"I declare the meeting adjourned.", pragmatics.ActDeclaration},
	}

	for _, tt := range tests {
		act := pragmatics.ClassifySpeechAct(tt.utterance)
		if act.Class != tt.expected {
			t.Errorf("ClassifySpeechAct(%q) = %s; expected %s", tt.utterance, act.Class, tt.expected)
		}
	}
}

func TestPresuppositionsComprehensive(t *testing.T) {
	// 1. Factive verb trigger
	pres := pragmatics.ExtractPresuppositions("John realized that the keys were in the car")
	if len(pres) == 0 {
		t.Fatalf("Expected presupposition from 'realize'")
	}
	if pres[0].Type != pragmatics.TriggerFactiveVerb || pres[0].Presupposition != "the keys were in the car" {
		t.Errorf("Unexpected presupposition: %+v", pres[0])
	}

	// 2. Change of state verb trigger
	stopPres := pragmatics.ExtractPresuppositions("Alice stopped smoking")
	if len(stopPres) == 0 {
		t.Fatalf("Expected presupposition from 'stopped'")
	}
	if stopPres[0].Type != pragmatics.TriggerChangeOfState {
		t.Errorf("Expected change of state trigger, got: %+v", stopPres[0])
	}

	// 3. Iterative adverb trigger ("again")
	iterPres := pragmatics.ExtractPresuppositions("The bell rang again")
	if len(iterPres) == 0 {
		t.Fatalf("Expected presupposition from 'again'")
	}
	if iterPres[0].Type != pragmatics.TriggerIterative {
		t.Errorf("Expected iterative trigger, got %+v", iterPres[0])
	}
}

func TestScalarImplicatureComprehensive(t *testing.T) {
	implicatures := pragmatics.DetectScalarImplicatures("Some students passed the examination.")
	if len(implicatures) == 0 {
		t.Fatalf("Expected scalar implicature for 'some'")
	}
	if implicatures[0].Inference != "not all" {
		t.Errorf("Expected 'not all' implicature, got '%s'", implicatures[0].Inference)
	}

	// No implicature sentence
	noImp := pragmatics.DetectScalarImplicatures("The dog barked.")
	if len(noImp) != 0 {
		t.Errorf("Expected 0 implicatures, got %d", len(noImp))
	}
}

func TestPolitenessComprehensive(t *testing.T) {
	// 1. Negative politeness with multiple hedges
	polite := pragmatics.AnalyzePoliteness("Could you please open the door perhaps?")
	if polite.Strategy != pragmatics.StrategyNegativePoliteness {
		t.Errorf("Expected negative politeness strategy, got %s", polite.Strategy)
	}
	if len(polite.MitigationTags) < 2 {
		t.Errorf("Expected >= 2 mitigation tags, got %d", len(polite.MitigationTags))
	}

	// 2. Single hedge with please
	pPlease := pragmatics.AnalyzePoliteness("Please help me.")
	if pPlease.Strategy != pragmatics.StrategyNegativePoliteness {
		t.Errorf("Expected negative politeness for please, got %s", pPlease.Strategy)
	}

	// 3. Single hedge without please
	pMaybe := pragmatics.AnalyzePoliteness("Maybe we can leave now.")
	if pMaybe.Strategy != pragmatics.StrategyPositivePoliteness {
		t.Errorf("Expected positive politeness for maybe, got %s", pMaybe.Strategy)
	}

	// 4. Bald on record with exclamation
	direct := pragmatics.AnalyzePoliteness("Open the door!")
	if direct.Strategy != pragmatics.StrategyBaldOnRecord {
		t.Errorf("Expected bald on-record strategy, got %s", direct.Strategy)
	}

	// 5. Neutral positive politeness
	neutral := pragmatics.AnalyzePoliteness("We are going to the library.")
	if neutral.Strategy != pragmatics.StrategyPositivePoliteness {
		t.Errorf("Expected positive politeness for neutral sentence, got %s", neutral.Strategy)
	}
}
