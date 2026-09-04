package pipeline_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/pipeline"
)

func ExamplePipeline_Process() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	p := pipeline.NewPipeline(res)
	doc := p.Process("The friendly dog barked joyfully.")

	fmt.Printf("Sentences: %d, Words: %d\n", len(doc.Sentences), len(doc.Sentences[0].Words))
	// Output:
	// Sentences: 1, Words: 5
}
