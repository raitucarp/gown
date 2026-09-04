package functional_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/functional"
)

func ExampleClassifyVerbProcess() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	verbs := res.LookupVerb("run")
	if len(verbs) > 0 {
		profile := functional.ClassifyVerbProcess(verbs[0])
		fmt.Printf("Verb: run, Process: %s\n", profile.Process)
	}
	// Output:
	// Verb: run, Process: material
}

func ExampleAnalyzeInterpersonal() {
	clause := "Did the students finish their assignment?"
	analysis := functional.AnalyzeInterpersonal(clause)

	fmt.Printf("Mood: %s, Speech: %s, Polarity: %s\n", analysis.Mood, analysis.Speech, analysis.Polarity)
	// Output:
	// Mood: interrogative, Speech: question, Polarity: positive
}

func ExampleSplitThemeRheme() {
	tr := functional.SplitThemeRheme("The researcher analyzed the experimental dataset")
	fmt.Printf("Theme: %s | Rheme: %s\n", tr.Theme, tr.Rheme)
	// Output:
	// Theme: The researcher | Rheme: analyzed the experimental dataset
}
