package similarity_test

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/similarity"
)

func ExampleCompare() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	// Compare semantic similarity between "dog" and "wolf" using Wu-Palmer
	score, err := similarity.Compare(res, "dog", "wolf",
		similarity.WithMetric(similarity.MetricWuPalmer),
		similarity.WithPOS(gown.NounPos),
	)
	if err != nil {
		log.Fatalf("comparison failed: %v", err)
	}

	fmt.Printf("Dog-Wolf Wu-Palmer similarity >= 0.8: %t\n", score >= 0.8)
	// Output:
	// Dog-Wolf Wu-Palmer similarity >= 0.8: true
}

func ExampleWuPalmer() {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to read lexical resource: %v", err)
	}

	dogEntries := res.Lookup("dog", gown.WithPOS(gown.NounPos))
	if len(dogEntries) == 0 || len(dogEntries[0].Synsets()) == 0 {
		return
	}

	s1 := dogEntries[0].Synsets()[0]
	// Identity score for a synset with itself is 1.0
	score := similarity.WuPalmer(res, s1, s1)
	fmt.Printf("Self Wu-Palmer similarity: %.1f\n", score)
	// Output:
	// Self Wu-Palmer similarity: 1.0
}
