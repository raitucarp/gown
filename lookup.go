package gown

import (
	"strings"
)

// LookupConfig specifies options for lexical lookup.
type LookupConfig struct {
	Pos           POS
	CaseSensitive bool
	UseMorphy     bool
	ExactOnly     bool
}

// LookupOption is a functional option for configuring lexical lookups.
type LookupOption func(*LookupConfig)

// WithPOS filters lookup results by part of speech.
func WithPOS(pos POS) LookupOption {
	return func(c *LookupConfig) {
		c.Pos = pos
	}
}

// WithCaseSensitive enforces exact case matching for word lookup.
func WithCaseSensitive() LookupOption {
	return func(c *LookupConfig) {
		c.CaseSensitive = true
	}
}

// WithMorphy enables morphological normalization and inflection reduction (e.g. "running" -> "run").
func WithMorphy() LookupOption {
	return func(c *LookupConfig) {
		c.UseMorphy = true
	}
}

// WithExactOnly disables fallback substring or collocation searches.
func WithExactOnly() LookupOption {
	return func(c *LookupConfig) {
		c.ExactOnly = true
	}
}

// Lookup queries the lexical database for entries matching the word, applying functional options.
func (resource *LexicalResource) Lookup(word string, opts ...LookupOption) LexicalEntries {
	cfg := LookupConfig{
		UseMorphy: true, // Default to morphy-aware lookup
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	word = strings.TrimSpace(word)
	if word == "" {
		return nil
	}

	seen := make(map[string]bool)
	var result LexicalEntries

	addEntry := func(entry *LexicalEntry) {
		if entry == nil || seen[entry.ID] {
			return
		}
		if cfg.Pos != "" {
			entryPos := POS(entry.Lemma.PartOfSpeech)
			if entryPos != cfg.Pos {
				if !( (cfg.Pos == AdjectivePos || cfg.Pos == AdjectiveSatellitePos) &&
					(entryPos == AdjectivePos || entryPos == AdjectiveSatellitePos) ) {
					return
				}
			}
		}
		seen[entry.ID] = true
		result = append(result, *entry)
	}

	// 1. Exact or case-insensitive direct lookup
	queryForms := []string{word}
	if !cfg.CaseSensitive {
		queryForms = append(queryForms, strings.ToLower(word))
	}

	for _, q := range queryForms {
		if cfg.CaseSensitive {
			for _, entry := range resource.entriesByLemmaExact[q] {
				addEntry(entry)
			}
		} else {
			for _, entry := range resource.entriesByLemmaLower[strings.ToLower(q)] {
				addEntry(entry)
			}
		}
	}

	// If found exact matches and not requesting morphy expansion, return
	if len(result) > 0 && !cfg.UseMorphy {
		return result
	}

	// 2. Morphological reduction (Morphy) if enabled or if no exact match found
	if cfg.UseMorphy {
		var lemmas []string
		if cfg.Pos != "" {
			lemmas = resource.Morphy(word, cfg.Pos)
		} else {
			lemmas = resource.MorphyAll(word)
		}

		for _, lemma := range lemmas {
			for _, entry := range resource.entriesByLemmaLower[strings.ToLower(lemma)] {
				addEntry(entry)
			}
		}
	}

	return result
}

// LookupExact retrieves entries with exact lemma matching.
func (resource *LexicalResource) LookupExact(word string, pos ...POS) LexicalEntries {
	var opts []LookupOption
	opts = append(opts, WithExactOnly())
	if len(pos) > 0 && pos[0] != "" {
		opts = append(opts, WithPOS(pos[0]))
	}
	return resource.Lookup(word, opts...)
}

// LookupLemma finds lexical entries matching the base lemma string.
func (resource *LexicalResource) LookupLemma(lemma string) LexicalEntries {
	return resource.Lookup(lemma, WithExactOnly())
}

// LookupNoun finds noun entries for the given word, with morphological reduction.
func (resource *LexicalResource) LookupNoun(word string) Nouns {
	entries := resource.Lookup(word, WithPOS(NounPos), WithMorphy())
	return Nouns(entries)
}

// LookupVerb finds verb entries for the given word, with morphological reduction.
func (resource *LexicalResource) LookupVerb(word string) Verbs {
	entries := resource.Lookup(word, WithPOS(VerbPos), WithMorphy())
	return Verbs(entries)
}

// LookupAdjective finds adjective entries for the given word, with morphological reduction.
func (resource *LexicalResource) LookupAdjective(word string) Adjectives {
	entries := resource.Lookup(word, WithPOS(AdjectivePos), WithMorphy())
	return Adjectives(entries)
}

// LookupAdverb finds adverb entries for the given word, with morphological reduction.
func (resource *LexicalResource) LookupAdverb(word string) Adverbs {
	entries := resource.Lookup(word, WithPOS(AdverbPos), WithMorphy())
	return Adverbs(entries)
}

// SynsetByID looks up a synset by its unique ID (e.g. "oewn-02084071-n").
func (resource *LexicalResource) SynsetByID(id string) *Synset {
	return resource.SynsetsById()[id]
}

// SynsetByILI looks up a synset by its Interlingual Index (ILI, e.g. "i46714").
func (resource *LexicalResource) SynsetByILI(ili string) *Synset {
	resource.SynsetsById() // Ensure indexes are loaded
	return resource.synsetsByILI[ili]
}

// SenseByID looks up a sense by its unique ID.
func (resource *LexicalResource) SenseByID(id string) *Sense {
	return resource.SenseById()[id]
}

// SenseByKey looks up a sense by its sense key or ID.
func (resource *LexicalResource) SenseByKey(key string) *Sense {
	senseById := resource.SenseById()
	if s, ok := senseById[key]; ok {
		return s
	}
	// Also search by prefix or substring if key is formatted differently
	for _, s := range senseById {
		if s.ID == key || strings.HasSuffix(s.ID, key) {
			return s
		}
	}
	return nil
}

// ReverseLookupConfig specifies options for reverse dictionary search.
type ReverseLookupConfig struct {
	Pos       POS
	MaxResult int
}

// ReverseLookupOption configures reverse lookup.
type ReverseLookupOption func(*ReverseLookupConfig)

// WithReversePOS limits reverse search to a specific part of speech.
func WithReversePOS(pos POS) ReverseLookupOption {
	return func(c *ReverseLookupConfig) {
		c.Pos = pos
	}
}

// WithReverseLimit sets the maximum number of results for reverse lookup.
func WithReverseLimit(limit int) ReverseLookupOption {
	return func(c *ReverseLookupConfig) {
		c.MaxResult = limit
	}
}

// ReverseLookup searches for synsets whose definitions or examples contain the query string.
func (resource *LexicalResource) ReverseLookup(query string, opts ...ReverseLookupOption) []*Synset {
	cfg := ReverseLookupConfig{
		MaxResult: 50,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	qLower := strings.ToLower(strings.TrimSpace(query))
	if qLower == "" {
		return nil
	}

	var results []*Synset
	for i := range resource.Lexicon.Synsets {
		synset := &resource.Lexicon.Synsets[i]
		if cfg.Pos != "" && POS(synset.PartOfSpeech) != cfg.Pos {
			continue
		}

		matched := false
		for _, def := range synset.Definitions {
			if strings.Contains(strings.ToLower(def), qLower) {
				matched = true
				break
			}
		}
		if !matched {
			for _, ex := range synset.Examples {
				if strings.Contains(strings.ToLower(ex.Text), qLower) {
					matched = true
					break
				}
			}
		}

		if matched {
			results = append(results, synset)
			if cfg.MaxResult > 0 && len(results) >= cfg.MaxResult {
				break
			}
		}
	}

	return results
}

// PrimaryDefinition returns the first definition of the synset, or empty string.
func (s *Synset) PrimaryDefinition() string {
	if len(s.Definitions) > 0 {
		return s.Definitions[0]
	}
	return ""
}

// LexicalEntries returns the lexical entries that are members of this synset.
func (s *Synset) LexicalEntries(resource *LexicalResource) LexicalEntries {
	lexicalsById := resource.LexicalsById()
	var entries LexicalEntries
	for _, member := range s.Members {
		if entry, ok := lexicalsById[member]; ok {
			entries = append(entries, *entry)
		}
	}
	return entries
}

// RelatedSynsets returns all synsets linked by a specific relation type.
func (s *Synset) RelatedSynsets(resource *LexicalResource, relType SynsetRelationType) []*Synset {
	synsetsById := resource.SynsetsById()
	var related []*Synset
	for _, rel := range s.SynsetRelations {
		if rel.RelType == string(relType) {
			if target, ok := synsetsById[rel.Target]; ok {
				related = append(related, target)
			}
		}
	}
	return related
}

// Hypernyms returns direct hypernym synsets.
func (s *Synset) Hypernyms(resource *LexicalResource) []*Synset {
	return s.RelatedSynsets(resource, SynsetRelationTypeHypernym)
}

// Hyponyms returns direct hyponym synsets.
func (s *Synset) Hyponyms(resource *LexicalResource) []*Synset {
	return s.RelatedSynsets(resource, SynsetRelationTypeHyponym)
}
