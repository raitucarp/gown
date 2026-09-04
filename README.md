# gown: Computational Linguistics and WordNet for Go

`gown` is a pure Go computational linguistics library and lexical database interface built on the Open English WordNet (OEWN). It provides embedded, zero-network dictionary lookup, morphological lemmatization, graph-theoretical ontology traversal, semantic similarity measures, and dedicated modules for syntax, semantics, pragmatics, discourse analysis, semiotics, phonology, and generative grammar.

The library is designed for software engineers, natural language processing (NLP) practitioners, and academic researchers in linguistics and the humanities.

---

## Table of Contents

- [Introduction](#introduction)
  - [What is WordNet?](#what-is-wordnet)
  - [Design Philosophy](#design-philosophy)
- [Package Architecture](#package-architecture)
- [Installation](#installation)
- [Quick Start](#quick-start)
- [Linguistic Modules](#linguistic-modules)
  - [1. Lexical Lookup and Morphy](#1-lexical-lookup-and-morphy)
  - [2. Knowledge Graph Traversal (`graph`)](#2-knowledge-graph-traversal-graph)
  - [3. Semantic Similarity and Relatedness (`similarity`)](#3-semantic-similarity-and-relatedness-similarity)
  - [4. Lexical and Definition Expansion (`expansion`)](#4-lexical-and-definition-expansion-expansion)
  - [5. Phonology, Syllables, and Meter (`phonology`, `pattern`)](#5-phonology-syllables-and-meter-phonology-pattern)
  - [6. Formal Syntax and Generative Grammar (`syntax`, `generative`)](#6-formal-syntax-and-generative-grammar-syntax-generative)
  - [7. Lexical Semantics and Sense Disambiguation (`semantics`)](#7-lexical-semantics-and-sense-disambiguation-semantics)
  - [8. Pragmatics, Deixis, and Speech Acts (`pragmatics`)](#8-pragmatics-deixis-and-speech-acts-pragmatics)
  - [9. Discourse Analysis and Rhetorical Structure (`discourse`)](#9-discourse-analysis-and-rhetorical-structure-discourse)
  - [10. Structural Semiotics (`semiotics`)](#10-structural-semiotics-semiotics)
  - [11. Systemic Functional Linguistics (`functional`)](#11-systemic-functional-linguistics-functional)
  - [12. Unified Linguistic Pipeline (`pipeline`)](#12-unified-linguistic-pipeline-pipeline)
- [Verification and Test Coverage](#verification-and-test-coverage)
- [Attribution and License](#attribution-and-license)

---

## Introduction

### What is WordNet?

In traditional dictionaries, words are arranged alphabetically. However, human language is organized conceptually. WordNet is a lexical database developed by cognitive scientists and linguists at Princeton University. 

In WordNet:
- **Synsets (Synonym Sets)**: The fundamental building blocks. A synset represents an underlying concept defined by a set of synonymous words (for example, `{dog, domestic dog, Canis familiaris}`).
- **Lexical Entries & Lemmas**: The base dictionary forms of individual words (such as `run`, `happy`, or `astronomy`).
- **Semantic Relations**: Conceptual pointers connecting synsets, such as hypernyms ("car is a vehicle"), hyponyms ("robin is a bird"), and meronyms ("wheel is a part of a bicycle").
- **Lexical Relations**: Direct relationships between specific word forms, such as antonymy ("light" vs "dark") or derivation ("teach" -> "teacher").

`gown` bundles the Open English WordNet (OEWN) directly into compiled Go binary data, requiring no external database servers, no Python bridges, and no internet access at runtime.

### Design Philosophy

- **Zero External Runtime Dependencies**: Standard Go library with pure Go data structures.
- **Determinism and Speed**: In-memory hash indexing ensures lookups and traversals run in sub-microsecond time.
- **Interoperability**: Standardized representations across phonology, morphology, syntax, semantics, and discourse.
- **High Test Coverage**: Rigorous test suites covering 94%+ statement coverage across every package with verified runnable examples.

---

## Package Architecture

| Package | Domain | Core Capabilities |
| :--- | :--- | :--- |
| `gown` | Lexicon & Morphy | O(1) lemma lookups, POS filtering, reverse dictionary search, Princeton Morphy lemmatization |
| `graph` | Graph Theory & Ontology | BFS/DFS, shortest paths, Dijkstra, Lowest Common Hypernym (LCS), degree centrality |
| `similarity` | Lexical Semantics | Path, Wu-Palmer, Leacock-Chodorow, Information Content (Seco), Resnik, Lin, Jiang-Conrath |
| `expansion` | Tree Expansion | Recursive hypernym/hyponym trees, recursive definition gloss expansion, cycle detection |
| `phonology` | Phonology & Prosody | Syllable counting, Sonority Sequencing Principle, onset-nucleus-coda, rhyme, meter, IPA |
| `pattern` | Orthographic Templates | Consonant-Vowel (CV) template matching (`CVC`, `CCVVC`, wildcards) |
| `morphology` | Morphological Analysis | Inflection identification, compound decomposition, derivational family extraction |
| `syntax` | Grammatical Structure | Syntactic category classification, subcategorization validation, constituent tree manipulation |
| `generative` | Generative Linguistics | Context-Free Grammars (CFG), lexical insertion, Penn Treebank formatting, feature unification |
| `semantics` | Lexical Semantics | Polysemy quantification, entropy, homonym detection, Lesk Word Sense Disambiguation (WSD) |
| `pragmatics` | Pragmatics & Context | Austin/Searle speech act classification, deixis extraction, presuppositions, Gricean implicature |
| `discourse` | Discourse Analysis | Elementary Discourse Units (EDUs), Rhetorical Structure Theory (RST), coreference tracking |
| `semiotics` | Structural Semiotics | Saussurean signifier/signified, Peircean triads, Greimas semiotic squares, connotation valence |
| `functional` | Functional Linguistics | Hallidayan transitivity processes, interpersonal mood/modality, Theme-Rheme decomposition |
| `paradigm` | Grammatical Frameworks | Universal Dependencies, FrameNet semantic frames, Categorial Grammar directional types |
| `pipeline` | End-to-End Analysis | Multi-layer document processing unifying all linguistic levels into a structured model |
| `text` | Text Utilities | Tokenization, sentence segmentation, stopword filtering, Jaccard token overlap |

---

## Installation

Ensure Go 1.21 or later is installed on your system:

```bash
go get github.com/raitucarp/gown
```

---

## Quick Start

The following program demonstrates loading the lexical database, performing basic lookups, querying synonyms and definitions, and finding hypernyms:

```go
package main

import (
	"fmt"
	"log"

	"github.com/raitucarp/gown"
)

func main() {
	// 1. Initialize lexical resource
	res, err := gown.ReadLexicalResource()
	if err != nil {
		log.Fatalf("failed to load WordNet: %v", err)
	}

	// 2. Look up noun entries for "canine"
	entries := res.LookupNoun("canine")
	for _, entry := range entries {
		fmt.Printf("Lemma: %s (POS: %s)\n", entry.Lemma.WrittenForm, entry.PartOfSpeech())
		for _, def := range entry.Definitions() {
			fmt.Printf("  Definition: %s\n", def)
		}
	}

	// 3. Inspect semantic relations (hypernyms / broader concepts)
	if len(entries) > 0 {
		hypernyms := entries[0].Relation().Hypernyms()
		fmt.Println("Hypernyms:")
		for _, hyp := range hypernyms {
			fmt.Printf("  - %s\n", hyp.Lemma.WrittenForm)
		}
	}
}
```

---

## Linguistic Modules

### 1. Lexical Lookup and Morphy

The root `gown` package provides fast indexed lookups and inflection reduction via the Princeton WordNet Morphy algorithm.

```go
// Lookup with functional options
entries := res.Lookup("running", gown.WithPOS(gown.VerbPos), gown.WithMorphy())

// Irregular forms are normalized automatically
miceLemmas := res.Morphy("mice", gown.NounPos)       // ["mouse"]
betterLemmas := res.Morphy("better", gown.AdjectivePos) // ["good", "well"]

// Reverse dictionary search (query definitions and examples)
results := res.ReverseLookup("optical instrument for viewing distant objects", 
    gown.WithReversePOS(gown.NounPos), 
    gown.WithReverseLimit(5))
```

### 2. Knowledge Graph Traversal (`graph`)

The `graph` package converts WordNet synsets and relations into a directed property graph.

```go
import "github.com/raitucarp/gown/graph"

builder := graph.NewBuilder(res)
g := builder.BuildSynsetGraph(gown.SynsetRelationHypernym)

// Shortest path between two synsets
path, found := g.ShortestPath(dogSynset.ID, animalSynset.ID)

// Lowest Common Subsumer (LCS) / Lowest Common Hypernym
lcsID, depth, err := g.LowestCommonHypernym(catSynset.ID, dogSynset.ID)

// Node centrality and degree distribution
inDeg, outDeg := g.Degree(dogSynset.ID)
```

### 3. Semantic Similarity and Relatedness (`similarity`)

The `similarity` package provides standard information-theoretic and path-based metrics to measure semantic distance between concepts.

```go
import "github.com/raitucarp/gown/similarity"

// Path-based distance metrics
wup := similarity.WuPalmer(res, dogSynset, wolfSynset)       // 0.0 - 1.0 scale
lch, _ := similarity.LeacockChodorow(res, dogSynset, catSynset)

// Information Content (IC) and Corpus-independent metrics (Seco et al.)
calc := similarity.NewInformationContent(res)
linSim := similarity.Lin(res, dogSynset, wolfSynset, calc)
resnikSim := similarity.Resnik(res, dogSynset, catSynset, calc)
jcnSim := similarity.JiangConrath(res, dogSynset, wolfSynset, calc)
```

### 4. Lexical and Definition Expansion (`expansion`)

Generate hierarchical trees or recursively follow definition tokens to analyze terminological depth.

```go
import "github.com/raitucarp/gown/expansion"

// Recursive hypernym tree expansion
opts := expansion.DefaultTreeOptions()
opts.MaxDepth = 4
tree, err := expansion.ExpandRelationTree(res, dogSynset.ID, gown.SynsetRelationHypernym, opts)

// Render formatted ASCII hierarchy
fmt.Println(tree.Render())

// Recursive definition gloss expansion
glossTree, err := expansion.ExpandDefinitionTree(res, "carnivore", 3)
```

### 5. Phonology, Syllables, and Meter (`phonology`, `pattern`)

Analyze sound patterns, syllable structure, poetic meter, and phoneme sequences.

```go
import (
	"github.com/raitucarp/gown/pattern"
	"github.com/raitucarp/gown/phonology"
)

// Syllable count and Sonority Sequencing Principle syllabification
count := phonology.CountSyllables("anticipation") // 5
sylls := phonology.Syllabify("linguistics")       // Onset, Nucleus, Coda decomposition

// Phonetic rhyme and prosodic meter
isRhyme := phonology.CheckRhyme("night", "bright") // RhymePerfect
foot := phonology.DetectMeterFoot([]int{0, 1})     // "Iamb"

// Orthographic consonant-vowel templates
pattern.OrthographicCV("banana") // "CVCVCV"
matches, _ := res.FindByPattern("CVCVC", 10)
```

### 6. Formal Syntax and Generative Grammar (`syntax`, `generative`)

Parse context-free grammars, construct constituent trees, and unify syntactic feature structures.

```go
import (
	"github.com/raitucarp/gown/generative"
	"github.com/raitucarp/gown/syntax"
)

// Context-Free Grammar (CFG) definition
grammar, _ := generative.NewCFG(`
	S  -> NP VP
	NP -> Det N
	VP -> V NP
`)

// Lexical insertion from WordNet to generate valid syntactic trees
tree, _ := generative.GenerateSentence(grammar, res, "S")
fmt.Println(tree.PennString()) // (S (NP (Det The) (N dog)) (VP (V chased) (NP (Det a) (N ball))))

// Feature structure unification
f1 := generative.FeatureStructure{"POS": "N", "NUM": "sg", "PERS": "3"}
f2 := generative.FeatureStructure{"NUM": "sg", "CASE": "nom"}
unified, ok := generative.UnifyFeatures(f1, f2)
```

### 7. Lexical Semantics and Sense Disambiguation (`semantics`)

Quantify semantic ambiguity, compute polysemy entropy, and disambiguate word senses in context.

```go
import "github.com/raitucarp/gown/semantics"

// Measure lexical ambiguity
polysemy := semantics.AnalyzePolysemy(res, "bank")
fmt.Printf("Senses: %d, Shannon Entropy: %.2f\n", polysemy.SenseCount, polysemy.Entropy)

// Simplified and Extended Lesk Word Sense Disambiguation (WSD)
sentence := "He deposited money in his account at the bank."
bestSynset, confidence := semantics.SimplifiedLesk(res, "bank", sentence)
```

### 8. Pragmatics, Deixis, and Speech Acts (`pragmatics`)

Analyze contextual language usage, speaker intent, and conversational implicature.

```go
import "github.com/raitucarp/gown/pragmatics"

// Austin-Searle illocutionary force classification
force := pragmatics.ClassifySpeechAct("Could you please open the window?")
// force.Primary: Directive, Politeness: Polite

// Deictic reference extraction (person, place, time)
deictics := pragmatics.IdentifyDeixis("We will meet here tomorrow at noon.")

// Scalar implicatures (Gricean maxims)
implicatures := pragmatics.DetectScalarImplicatures("Some students passed the exam.")
// Implicature: "Not all students passed the exam"
```

### 9. Discourse Analysis and Rhetorical Structure (`discourse`)

Segment texts into discourse units, build rhetorical trees, and track entity coreference.

```go
import "github.com/raitucarp/gown/discourse"

text := "Although the storm was heavy, the ship reached the harbor safely. The crew cheered."

// Elementary Discourse Unit (EDU) segmentation
edus := discourse.SegmentEDUs(text)

// Rhetorical Structure Theory (RST) parsing
rstTree := discourse.BuildRSTTree(edus)

// Pronoun and nominal coreference resolution
chains := discourse.ResolveCoreference(text)
```

### 10. Structural Semiotics (`semiotics`)

Examine semiotic systems using Saussurean signs, Peircean triads, and Greimas squares.

```go
import "github.com/raitucarp/gown/semiotics"

// Greimas Semiotic Square (Carré Sémiotique)
square := semiotics.GenerateSemioticSquare(res, "good")
fmt.Println(square.Render())
// Maps S1 (good) <--> S2 (bad) and their contradictory negations (~S1, ~S2)

// Connotation and affective valence
connotation := semiotics.AnalyzeConnotation(res, "noble")
// Valence: positive, Register: formal
```

### 11. Systemic Functional Linguistics (`functional`)

Deconstruct text according to M. A. K. Halliday's Systemic Functional Grammar.

```go
import "github.com/raitucarp/gown/functional"

// Transitivity analysis (Material, Mental, Relational, Verbal, Behavioral, Existential)
process := functional.ClassifyProcess("think") // ProcessMental

// Clause Theme-Rheme information structure
theme, rheme := functional.AnalyzeThemeRheme("Yesterday morning, the council approved the budget.")
```

### 12. Unified Linguistic Pipeline (`pipeline`)

The `pipeline` package integrates all modules into an end-to-end analyzer that processes raw text into a structured document.

```go
import "github.com/raitucarp/gown/pipeline"

p := pipeline.NewPipeline(res)
doc := p.Process("The diligent student read the manuscript carefully.")

for _, sent := range doc.Sentences {
    fmt.Printf("Sentence: %s\n", sent.Raw)
    fmt.Printf("  Speech Act: %s\n", sent.SpeechAct.Primary)
    for _, word := range sent.Words {
        fmt.Printf("  Token: %-12s POS: %-4s Syllables: %d Pattern: %s\n", 
            word.Surface, word.POS, word.Syllables, word.CVPattern)
    }
}
```

---

## Verification and Test Coverage

The library maintains strict unit test suites with multiple test scenarios, edge-case handling, and runnable GoDoc examples. Statement test coverage across all packages exceeds 94%:

```text
github.com/raitucarp/gown               94.4%
github.com/raitucarp/gown/discourse     95.8%
github.com/raitucarp/gown/expansion     97.3%
github.com/raitucarp/gown/functional    94.3%
github.com/raitucarp/gown/generative    96.4%
github.com/raitucarp/gown/graph         97.5%
github.com/raitucarp/gown/morphology    95.6%
github.com/raitucarp/gown/paradigm     100.0%
github.com/raitucarp/gown/phonology     95.8%
github.com/raitucarp/gown/pipeline     100.0%
github.com/raitucarp/gown/pragmatics    96.6%
github.com/raitucarp/gown/semantics     94.3%
github.com/raitucarp/gown/semiotics     96.6%
github.com/raitucarp/gown/similarity    94.5%
github.com/raitucarp/gown/syntax        96.8%
github.com/raitucarp/gown/text          98.0%
```

To run all unit tests and verify coverage:

```bash
go test -v -cover ./...
```

---

## Attribution and License

`gown` is released under the [MIT License](LICENSE.md).

The lexical data embedded within this library originates from:
1. **Princeton WordNet**: Developed by Princeton University under the WordNet 3.0 License.
2. **Open English WordNet (OEWN)**: Maintained by the Global WordNet Association under the Creative Commons Attribution 4.0 International License (CC BY 4.0).

Please see [ATTRIBUTION.md](ATTRIBUTION.md) and [LICENSE.md](LICENSE.md) for full citations and licensing notices.
