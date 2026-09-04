package pragmatics_test

import (
	"fmt"

	"github.com/raitucarp/gown/pragmatics"
)

func ExampleClassifySpeechAct() {
	force := pragmatics.ClassifySpeechAct("Could you please open the window?")
	fmt.Printf("Speech act: %s, Confidence: %.2f\n", force.Class, force.Confidence)
	// Output:
	// Speech act: directive, Confidence: 0.85
}

func ExampleExtractPresuppositions() {
	pres := pragmatics.ExtractPresuppositions("Mary stopped drinking coffee")
	for _, p := range pres {
		fmt.Printf("Trigger: %s, Type: %s\n", p.Trigger, p.Type)
	}
	// Output:
	// Trigger: stopped, Type: change_of_state
}

func ExampleResolveDeixis() {
	ctx := pragmatics.NewContext("Speaker", "Listener")
	ctx.Location = "Auditorium"
	ctx.Time = "10:00 AM"

	deictics := pragmatics.ResolveDeixis("I am here now", ctx)
	for _, d := range deictics {
		fmt.Printf("%s -> %s\n", d.Word, d.ResolvedTo)
	}
	// Output:
	// I -> Speaker
	// here -> Auditorium
	// now -> 10:00 AM
}
