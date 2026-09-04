# Package text

Package `text` provides core natural language preprocessing utilities, including sentence segmentation, whitespace and punctuation tokenization, English stopword removal, and lexical overlap calculations.

## Overview

Text preprocessing is a foundational step in natural language processing and computational linguistics. Package `text` provides lightweight, pure Go text manipulation tools:

1. **Sentence Segmentation**: Splitting continuous prose into discrete sentences, accounting for punctuation, quotes, abbreviations, and honorifics (*Mr.*, *Dr.*, *U.S.*).
2. **Tokenization**: Breaking sentences into words, punctuation marks, and multi-word units with configurable lowercasing and punctuation stripping.
3. **Stopword Filtering**: Identifying and removing high-frequency grammatical closed-class words (articles, conjunctions, prepositions) that carry minimal lexical semantic content.
4. **Lexical Overlap (Jaccard Index)**: Computing token intersection-over-union between two texts, useful for gloss matching and sentence similarity.

## Key Functions

- `SentenceSegment(rawText string) []string`: Segments multi-sentence prose into discrete sentences.
- `Tokenize(text string) []string`: Splits a string into words, handling whitespace and punctuation.
- `IsStopword(word string) bool`: Returns `true` if the word belongs to the standard English stopword lexicon.
- `FilterStopwords(tokens []string) []string`: Returns a slice containing only content words.
- `JaccardSimilarity(text1, text2 string) float64`: Computes the Jaccard similarity coefficient (0.0 to 1.0) between the token sets of two strings.

## Example

```go
package main

import (
    "fmt"
    "github.com/raitucarp/gown/text"
)

func main() {
    doc := "Dr. Watson visited 221B Baker St. yesterday. He met Sherlock Holmes there."

    // 1. Sentence segmentation
    sentences := text.SentenceSegment(doc)
    fmt.Printf("Sentences (%d):\n", len(sentences))
    for i, s := range sentences {
        fmt.Printf("  [%d] %s\n", i+1, s)
    }

    // 2. Tokenization & stopword removal
    tokens := text.Tokenize(sentences[0])
    contentWords := text.FilterStopwords(tokens)
    fmt.Println("Content words:", contentWords)

    // 3. Jaccard lexical overlap
    sim := text.JaccardSimilarity("the fast brown dog", "the quick brown canine")
    fmt.Printf("Token overlap: %.2f\n", sim)
}
```
