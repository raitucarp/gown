package morphology

import (
	"strings"

	"github.com/raitucarp/gown"
)

// CompoundSplit represents a decomposition of a compound word into sub-words.
type CompoundSplit struct {
	Compound string   `json:"compound"`
	Parts    []string `json:"parts"`
}

// SplitCompound attempts to decompose an English compound word into constituent words found in WordNet.
func SplitCompound(res *gown.LexicalResource, word string) []CompoundSplit {
	word = strings.TrimSpace(strings.ToLower(word))
	if len(word) < 4 {
		return nil
	}

	var splits []CompoundSplit

	// Check if already separated by hyphen or space
	if strings.ContainsAny(word, "- ") {
		parts := strings.FieldsFunc(word, func(r rune) bool {
			return r == '-' || r == ' '
		})
		allValid := true
		for _, p := range parts {
			if len(res.LookupExact(p)) == 0 {
				allValid = false
				break
			}
		}
		if allValid {
			splits = append(splits, CompoundSplit{
				Compound: word,
				Parts:    parts,
			})
			return splits
		}
	}

	// Two-word compound splitting: word = left + right
	for i := 3; i <= len(word)-3; i++ {
		left := word[:i]
		right := word[i:]

		leftEntries := res.LookupExact(left)
		if len(leftEntries) == 0 {
			continue
		}

		rightEntries := res.LookupExact(right)
		if len(rightEntries) == 0 {
			continue
		}

		splits = append(splits, CompoundSplit{
			Compound: word,
			Parts:    []string{left, right},
		})
	}

	return splits
}
