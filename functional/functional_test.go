package functional_test

import (
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/functional"
)

func TestTransitivityClassificationComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// 1. Material process (run)
	verbsRun := res.LookupVerb("run")
	if len(verbsRun) > 0 {
		p := functional.ClassifyVerbProcess(verbsRun[0])
		if p.Process != functional.ProcessMaterial {
			t.Errorf("Expected 'run' to be Material, got %s", p.Process)
		}
	}

	// 2. Mental process (think)
	verbsThink := res.LookupVerb("think")
	if len(verbsThink) > 0 {
		p := functional.ClassifyVerbProcess(verbsThink[0])
		if p.Process != functional.ProcessMental {
			t.Errorf("Expected 'think' to be Mental, got %s", p.Process)
		}
	}

	// 3. Relational process (have)
	verbsHave := res.LookupVerb("have")
	if len(verbsHave) > 0 {
		p := functional.ClassifyVerbProcess(verbsHave[0])
		if p.Process != functional.ProcessRelational && p.Process != functional.ProcessMaterial {
			t.Errorf("Expected 'have' to be Relational or Material, got %s", p.Process)
		}
	}

	// 4. Verbal process (say)
	verbsSay := res.LookupVerb("say")
	if len(verbsSay) > 0 {
		p := functional.ClassifyVerbProcess(verbsSay[0])
		if p.Process != functional.ProcessVerbal {
			t.Errorf("Expected 'say' to be Verbal, got %s", p.Process)
		}
	}

	// 5. Behavioral process (cough)
	verbsCough := res.LookupVerb("cough")
	if len(verbsCough) > 0 {
		p := functional.ClassifyVerbProcess(verbsCough[0])
		if p.Process != functional.ProcessBehavioral {
			t.Errorf("Expected 'cough' to be Behavioral, got %s", p.Process)
		}
	}

	// 6. Existential process (exist)
	verbsExist := res.LookupVerb("exist")
	if len(verbsExist) > 0 {
		p := functional.ClassifyVerbProcess(verbsExist[0])
		if p.Process != functional.ProcessExistential {
			t.Errorf("Expected 'exist' to be Existential, got %s", p.Process)
		}
	}
}

func TestInterpersonalAnalysisComprehensive(t *testing.T) {
	// Declarative positive statement
	statement := functional.AnalyzeInterpersonal("The dog chased the cat.")
	if statement.Mood != functional.MoodDeclarative || statement.Speech != functional.SpeechStatement || statement.Polarity != functional.PolarityPositive {
		t.Errorf("Expected declarative positive statement, got %+v", statement)
	}

	// Interrogative with question mark
	q1 := functional.AnalyzeInterpersonal("Did the dog chase the cat?")
	if q1.Mood != functional.MoodInterrogative || q1.Speech != functional.SpeechQuestion {
		t.Errorf("Expected interrogative question, got %+v", q1)
	}

	// Interrogative with WH-question word
	q2 := functional.AnalyzeInterpersonal("what are you doing")
	if q2.Mood != functional.MoodInterrogative || q2.Speech != functional.SpeechQuestion {
		t.Errorf("Expected WH-interrogative question, got %+v", q2)
	}

	// Negative polarity (never, nothing, didn't)
	neg1 := functional.AnalyzeInterpersonal("He never arrived.")
	if neg1.Polarity != functional.PolarityNegative {
		t.Errorf("Expected negative polarity for 'never', got %+v", neg1)
	}
	neg2 := functional.AnalyzeInterpersonal("They didn't see anything.")
	if neg2.Polarity != functional.PolarityNegative {
		t.Errorf("Expected negative polarity for 'didn't', got %+v", neg2)
	}

	// Modality types:
	// 1. Obligation
	mObl := functional.AnalyzeInterpersonal("You should study harder.")
	if mObl.Modality != functional.ModalityObligation {
		t.Errorf("Expected obligation modality, got %+v", mObl)
	}
	// 2. Probability
	mProb := functional.AnalyzeInterpersonal("They might arrive tonight.")
	if mProb.Modality != functional.ModalityProbability {
		t.Errorf("Expected probability modality, got %+v", mProb)
	}
	// 3. Usuality
	mUsual := functional.AnalyzeInterpersonal("He usually eats lunch here.")
	if mUsual.Modality != functional.ModalityUsuality {
		t.Errorf("Expected usuality modality, got %+v", mUsual)
	}
	// 4. Inclination
	mInc := functional.AnalyzeInterpersonal("I am willing to help.")
	if mInc.Modality != functional.ModalityInclination {
		t.Errorf("Expected inclination modality, got %+v", mInc)
	}

	// Imperative command
	cmd := functional.AnalyzeInterpersonal("Go away!")
	if cmd.Mood != functional.MoodImperative || cmd.Speech != functional.SpeechCommand {
		t.Errorf("Expected imperative command, got %+v", cmd)
	}
}

func TestThemeRhemeComprehensive(t *testing.T) {
	// Standard two-word theme (article + noun)
	tr1 := functional.SplitThemeRheme("The researcher analyzed the data")
	if tr1.Theme != "The researcher" || tr1.Rheme != "analyzed the data" {
		t.Errorf("Unexpected Theme/Rheme: %+v", tr1)
	}

	// Single word theme
	tr2 := functional.SplitThemeRheme("Researchers analyze data")
	if tr2.Theme != "Researchers" || tr2.Rheme != "analyze data" {
		t.Errorf("Unexpected Theme/Rheme: %+v", tr2)
	}

	// One-word sentence
	tr3 := functional.SplitThemeRheme("Run!")
	if tr3.Theme != "Run!" || tr3.Rheme != "" {
		t.Errorf("Unexpected Theme/Rheme for single word: %+v", tr3)
	}

	// Empty string
	tr4 := functional.SplitThemeRheme("")
	if tr4.Theme != "" || tr4.Rheme != "" {
		t.Errorf("Unexpected Theme/Rheme for empty: %+v", tr4)
	}
}

func TestAnalyzeCohesionComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// Short text (< 2 content words)
	if ties := functional.AnalyzeCohesion(res, "hello"); ties != nil {
		t.Errorf("Expected nil for single-word text, got %+v", ties)
	}

	// Text with repetition, synonymy, and hyponymy
	passage := "The canine barked loudly. A hound ran quickly. The fast dog was a good canine."
	ties := functional.AnalyzeCohesion(res, passage)

	if len(ties) == 0 {
		t.Fatalf("Expected cohesion ties in passage")
	}

	hasRepetition := false
	hasSynonym := false
	for _, tie := range ties {
		if tie.Type == functional.CohesionRepetition {
			hasRepetition = true
		}
		if tie.Type == functional.CohesionSynonymy || tie.Type == functional.CohesionHypernymy || tie.Type == functional.CohesionHyponymy {
			hasSynonym = true
		}
	}

	if !hasRepetition {
		t.Errorf("Expected repetition tie for 'canine'")
	}
	if !hasSynonym {
		t.Errorf("Expected semantic relation tie")
	}
}
