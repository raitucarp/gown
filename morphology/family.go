package morphology

import (
	"sort"
	"strings"

	"github.com/raitucarp/gown"
)

// LexicalMember represents a member of a derivational word family.
type LexicalMember struct {
	Word     string   `json:"word"`
	POS      gown.POS `json:"pos"`
	Relation string   `json:"relation"` // "derivation", "affix_variant", etc.
}

// LexicalFamily represents a morphological and derivational family grouped around a root.
type LexicalFamily struct {
	Root    string          `json:"root"`
	Members []LexicalMember `json:"members"`
}

// DerivationalAffixes lists common English prefixes and suffixes for family expansion.
var DerivationalAffixes = struct {
	Prefixes []string
	Suffixes []string
}{
	Prefixes: []string{"un", "re", "in", "im", "dis", "en", "non", "pre", "pro", "sub", "inter", "trans", "over", "under", "co", "de"},
	Suffixes: []string{"tion", "ation", "sion", "ment", "ness", "ity", "ty", "er", "or", "ist", "ism", "able", "ible", "al", "ial", "ful", "less", "ive", "ous", "ious", "ic", "ical", "ly", "ize", "ise", "ify", "ate", "en"},
}

// GenerateLexicalFamily gathers all derivationally related forms, WordNet sense derivations,
// and affixal derivations for a root word.
func GenerateLexicalFamily(res *gown.LexicalResource, word string) *LexicalFamily {
	word = strings.TrimSpace(strings.ToLower(word))
	fam := &LexicalFamily{
		Root: word,
	}

	seen := make(map[string]bool)
	seen[word] = true

	addMember := func(w string, pos gown.POS, rel string) {
		wLower := strings.ToLower(w)
		if wLower == "" || seen[wLower] {
			return
		}
		seen[wLower] = true
		fam.Members = append(fam.Members, LexicalMember{
			Word:     wLower,
			POS:      pos,
			Relation: rel,
		})
	}

	// 1. Traverse WordNet derivational relations from entry.Relation().Derivations()
	entries := res.Lookup(word)
	for _, entry := range entries {
		derivs := entry.Relation().Derivations()
		for _, d := range derivs {
			addMember(d.Lemma.WrittenForm, gown.POS(d.Lemma.PartOfSpeech), "derivation")
		}
	}

	// 2. Affix variations checked against WordNet dictionary
	for _, suffix := range DerivationalAffixes.Suffixes {
		candidates := []string{
			word + suffix,
			strings.TrimSuffix(word, "e") + suffix,
			strings.TrimSuffix(word, "y") + "i" + suffix,
		}
		for _, cand := range candidates {
			exact := res.LookupExact(cand)
			for _, e := range exact {
				addMember(e.Lemma.WrittenForm, gown.POS(e.Lemma.PartOfSpeech), "suffix:"+suffix)
			}
		}
	}

	for _, prefix := range DerivationalAffixes.Prefixes {
		cand := prefix + word
		exact := res.LookupExact(cand)
		for _, e := range exact {
			addMember(e.Lemma.WrittenForm, gown.POS(e.Lemma.PartOfSpeech), "prefix:"+prefix)
		}
	}

	// Sort members alphabetically
	sort.Slice(fam.Members, func(i, j int) bool {
		return fam.Members[i].Word < fam.Members[j].Word
	})

	return fam
}
