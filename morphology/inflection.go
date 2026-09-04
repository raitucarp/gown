package morphology

import (
	"strings"

	"github.com/raitucarp/gown"
)

// InflectionKind identifies the grammatical inflection type.
type InflectionKind string

const (
	InflectionPlural      InflectionKind = "plural"
	InflectionPastTense   InflectionKind = "past_tense"
	InflectionPastPart    InflectionKind = "past_participle"
	InflectionProgressive InflectionKind = "progressive"
	Inflection3rdPersonSg InflectionKind = "3rd_person_singular"
	InflectionComparative InflectionKind = "comparative"
	InflectionSuperlative InflectionKind = "superlative"
	InflectionBase        InflectionKind = "base"
)

// InflectionInfo describes the detected inflectional status of a word form.
type InflectionInfo struct {
	Word      string
	BaseLemma string
	Kind      InflectionKind
	POS       gown.POS
}

// DetectInflections analyzes the surface word form and returns possible inflectional interpretations.
func DetectInflections(res *gown.LexicalResource, word string) []InflectionInfo {
	word = strings.TrimSpace(strings.ToLower(word))
	if word == "" {
		return nil
	}

	var infos []InflectionInfo
	seen := make(map[string]bool)

	addInfo := func(base string, kind InflectionKind, pos gown.POS) {
		key := base + ":" + string(kind) + ":" + string(pos)
		if !seen[key] && base != "" {
			seen[key] = true
			infos = append(infos, InflectionInfo{
				Word:      word,
				BaseLemma: base,
				Kind:      kind,
				POS:       pos,
			})
		}
	}

	// Check if base form
	exact := res.LookupExact(word)
	if len(exact) > 0 {
		for _, e := range exact {
			addInfo(word, InflectionBase, gown.POS(e.Lemma.PartOfSpeech))
		}
	}

	// Plural noun detection
	nounLemmas := res.Morphy(word, gown.NounPos)
	for _, lemma := range nounLemmas {
		if lemma != word {
			addInfo(lemma, InflectionPlural, gown.NounPos)
		}
	}

	// Verb inflections
	verbLemmas := res.Morphy(word, gown.VerbPos)
	for _, lemma := range verbLemmas {
		if lemma != word {
			if strings.HasSuffix(word, "ing") {
				addInfo(lemma, InflectionProgressive, gown.VerbPos)
			} else if strings.HasSuffix(word, "ed") || strings.HasSuffix(word, "en") {
				addInfo(lemma, InflectionPastTense, gown.VerbPos)
			} else if strings.HasSuffix(word, "s") || strings.HasSuffix(word, "es") {
				addInfo(lemma, Inflection3rdPersonSg, gown.VerbPos)
			} else {
				addInfo(lemma, InflectionPastTense, gown.VerbPos)
			}
		}
	}

	// Adjective comparative/superlative
	adjLemmas := res.Morphy(word, gown.AdjectivePos)
	for _, lemma := range adjLemmas {
		if lemma != word {
			if strings.HasSuffix(word, "est") || strings.HasSuffix(word, "iest") {
				addInfo(lemma, InflectionSuperlative, gown.AdjectivePos)
			} else if strings.HasSuffix(word, "er") || strings.HasSuffix(word, "ier") {
				addInfo(lemma, InflectionComparative, gown.AdjectivePos)
			}
		}
	}

	return infos
}
