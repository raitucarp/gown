package discourse_test

import (
	"fmt"

	"github.com/raitucarp/gown/discourse"
)

func ExampleSegmentEDUs() {
	text := "The sun was shining, although dark clouds hovered on the horizon."
	edus := discourse.SegmentEDUs(text)

	fmt.Printf("Total EDUs: %d\n", len(edus))
	fmt.Printf("EDU 1: %s\n", edus[0].Text)
	// Output:
	// Total EDUs: 2
	// EDU 1: The sun was shining
}

func ExampleBuildRSTTree() {
	edus := []discourse.EDU{
		{ID: 1, Text: "The dog barked loudly"},
		{ID: 2, Text: "because it saw a cat"},
	}

	tree := discourse.BuildRSTTree(edus)
	fmt.Printf("Tree relation: %s, Children: %d\n", tree.Relation, len(tree.Children))
	// Output:
	// Tree relation: Sequence, Children: 2
}

func ExampleTrackTopics() {
	text := "The dog barked. The dog ran fast."
	topics, transitions := discourse.TrackTopics(text)

	fmt.Printf("Sentences: %d, Transitions: %d\n", len(topics), len(transitions))
	// Output:
	// Sentences: 2, Transitions: 1
}
