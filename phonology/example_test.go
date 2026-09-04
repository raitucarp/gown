package phonology_test

import (
	"fmt"

	"github.com/raitucarp/gown/phonology"
)

func ExampleCountSyllables() {
	count := phonology.CountSyllables("linguistics")
	fmt.Printf("Syllables in 'linguistics': %d\n", count)
	// Output:
	// Syllables in 'linguistics': 3
}

func ExampleSyllabify() {
	sylls := phonology.Syllabify("cat")
	if len(sylls) > 0 {
		fmt.Printf("Onset: %s, Nucleus: %s, Coda: %s\n", sylls[0].Onset, sylls[0].Nucleus, sylls[0].Coda)
	}
	// Output:
	// Onset: c, Nucleus: a, Coda: t
}

func ExampleAreRhymes() {
	rhyme := phonology.AreRhymes("cat", "hat")
	fmt.Printf("cat and hat rhyme: %t\n", rhyme)
	// Output:
	// cat and hat rhyme: true
}
