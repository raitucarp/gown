package text_test

import (
	"fmt"

	"github.com/raitucarp/gown/text"
)

func ExampleTokenize() {
	tokens := text.Tokenize("Natural language processing is fascinating!")
	fmt.Printf("Token count: %d, First: %s\n", len(tokens), tokens[0])
	// Output:
	// Token count: 5, First: natural
}

func ExampleExtractContentWords() {
	words := text.ExtractContentWords("The dog chased the cat through the garden")
	fmt.Printf("Content words: %v\n", words)
	// Output:
	// Content words: [dog chased cat garden]
}

func ExampleJaccardSimilarity() {
	sim := text.JaccardSimilarity([]string{"apple", "banana"}, []string{"banana", "cherry"})
	fmt.Printf("Jaccard similarity: %.2f\n", sim)
	// Output:
	// Jaccard similarity: 0.33
}
