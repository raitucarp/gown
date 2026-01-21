# gown (Go WordNet)

Go package for working with WordNet lexical databases. **gown** provides easy-to-use APIs for searching, filtering, and analyzing semantic relationships in the Open English WordNet (OEWN) database.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Core Concepts](#core-concepts)
- [API Documentation](#api-documentation)
  - [Resource Management](#resource-management)
  - [Searching](#searching)
  - [Filtering by Part of Speech](#filtering-by-part-of-speech)
  - [Word vs Collocation](#word-vs-collocation)
  - [Relations](#relations)
  - [Random Selection](#random-selection)
- [Data Types](#data-types)
- [Project Structure](#project-structure)
- [License](#license)
- [Contributing](#contributing)

## Overview

**gown** is a Go binding for the Open English WordNet that embeds the lexical database and provides a fluent API for semantic analysis. Whether you're building NLP applications, conducting linguistic research, or working with word semantics, gown makes it straightforward to access and manipulate WordNet data.

## Features

- **Embedded WordNet Data**: Ships with the Open English WordNet (OEWN) in a binary-optimized GOB format
- **Part of Speech Support**: Full support for nouns, verbs, adjectives, and adverbs
- **Semantic Relations**: Access to antonyms, synonyms, derivations, and other semantic relationships
- **Flexible Search**: Search by lemma (word form) or by definition
- **Sense Management**: Work with word senses, definitions, and examples
- **Collocation Support**: Distinguish between single words and multi-word expressions
- **Type-Safe API**: Strongly-typed methods for each part of speech
- **Zero Dependencies Runtime**: Minimal external dependencies (only uses samber/lo for utility functions)

## Installation

```bash
go get github.com/raitucarp/gown
```

## Quick Start

```go
package main

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
)

func main() {
	// Load the lexical resource
	resource, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatal(err)
	}

	// Search for a word
	nouns := resource.SearchLemma("dog").Nouns()
	for _, noun := range nouns {
		fmt.Println(noun.String())
		// Print definitions
		for _, def := range noun.Definitions() {
			fmt.Printf("  - %s\n", def)
		}
	}

	// Get random words
	randomNouns := resource.Nouns().Random(5)
	fmt.Println("Random nouns:", randomNouns)
}
```

## Core Concepts

### LexicalResource

The main entry point representing the entire WordNet database. Contains lexicons with lexical entries and synsets.

### LexicalEntry

Represents a word in the database, associated with a lemma (base form) and multiple senses pointing to synsets.

### Synset

A set of synonymous words representing a single concept. Contains definitions, examples, and relationships to other synsets.

### Sense

The connection between a lexical entry and a synset, representing a specific meaning of a word.

### Part of Speech (POS)

Supported categories:

- **Noun** (NounPos)
- **Verb** (VerbPos)
- **Adjective** (AdjectivePos, AdjectiveSatellitePos)
- **Adverb** (AdverbPos)

## API Documentation

### Resource Management

```go
// Load the embedded WordNet database
resource, err := gown.ReadLexicalResource()

// Get all lexical entries
entries := resource.Lexicon.LexicalEntries

// Get all synsets
synsets := resource.Lexicon.Synsets
```

### Searching

```go
// Search by lemma (word form) - case-sensitive prefix/substring match
results := resource.SearchLemma("run")

// Search by definition
definitionMatches := resource.SearchLemmaByDefinition("moving fast")

// Search on specific POS collections
nouns := resource.Nouns().SearchLemma("cat")
verbs := resource.Verbs().SearchLemma("jump")
```

### Filtering by Part of Speech

```go
// Get all nouns
nouns := resource.Nouns()

// Get all verbs
verbs := resource.Verbs()

// Get all adjectives (includes satellite adjectives)
adjectives := resource.Adjectives()

// Get all adverbs
adverbs := resource.Adverbs()
```

### Word vs Collocation

```go
// Get only single words (not multi-word expressions)
singleWords := resource.Words()

// Get only multi-word expressions
collocations := resource.Collocations()

// Available on POS types too
nounWords := resource.Nouns().Words()
nounCollocations := resource.Nouns().Collocations()
```

### Relations

gown provides access to various semantic relationships:

#### Sense Relations

- `also` - Also related
- `antonym` - Opposite meaning
- `derivation` - Derived from
- `exemplifies` - Provides examples
- `is_exemplified_by` - Is exemplified by
- `participle` - Participle form
- `pertainym` - Related attribute
- `similar` - Similar meaning

#### Synset Relations

- `also`, `attribute`, `causes`, `entails`
- `domain_region`, `domain_topic`
- `holo_member`, `holo_part`, `holo_substance` (holonym relations)
- `mero_member`, `mero_part`, `mero_substance` (meronym relations)
- `hypernym` / `hyponym` (is-a relationships)

#### Dublin Core Relations

Thematic role relations including: agent, body_part, by_means_of, destination, event, instrument, location, material, property, result, state, undergoer, uses, vehicle

### Random Selection

```go
// Get n random lemmas
randomEntries := resource.RandomLemma(10)

// Get random from specific collections
randomNouns := resource.Nouns().Random(5)
randomVerbs := resource.Verbs().Random(5)
randomAdjectives := resource.Adjectives().Random(3)
```

## Data Types

### Main Structures

- **LexicalResource**: Root container for all WordNet data
- **Lexicon**: Collection of lexical entries and synsets
- **LexicalEntry**: Individual word entry with lemma and senses
- **Synset**: Concept/meaning containing related words
- **Sense**: Link between a lexical entry and a synset
- **Lemma**: Base form of a word with POS and pronunciations
- **Form**: Alternative written forms
- **Pronunciation**: Phonetic variants
- **Example**: Example usage of a sense
- **SenseRelation**: Semantic relationship between senses
- **SynsetRelation**: Semantic relationship between synsets
- **SyntacticBehaviour**: Subcategorization frames for verbs

### Type Aliases

- **Nouns**: Collection of noun lexical entries
- **Verbs**: Collection of verb lexical entries
- **Adjectives**: Collection of adjective lexical entries
- **Adverbs**: Collection of adverb lexical entries
- **LexicalEntries**: Collection of lexical entries

## License

gown is distributed under the Creative Commons Attribution 4.0 International License. The underlying WordNet data is derived from:

1. **Princeton WordNet** - Licensed under the WordNet License
2. **Open English WordNet (OEWN)** - Licensed under the WordNet License

Attribution must be given to both Princeton WordNet and the Open English WordNet team when using this library. See [LICENSE](LICENSE) for full details.

## Contributing

Contributions are welcome! Whether it's bug reports, feature requests, or pull requests, please feel free to engage with the project.

For issues or suggestions, please open an issue on the repository.

---

**Built with ❤️ for the Go and NLP communities**
