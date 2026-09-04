package functional

import (
	"strings"
)

// ThemeRheme represents the information structure of a clause in SFL.
type ThemeRheme struct {
	Clause string `json:"clause"`
	Theme  string `json:"theme"`
	Rheme  string `json:"rheme"`
}

// SplitThemeRheme decomposes a clause into its Theme (initial topical constituent)
// and Rheme (the remainder of the message).
func SplitThemeRheme(clause string) ThemeRheme {
	words := strings.Fields(strings.TrimSpace(clause))
	if len(words) <= 1 {
		return ThemeRheme{
			Clause: clause,
			Theme:  clause,
			Rheme:  "",
		}
	}

	// Simple heuristic: subject NP / first constituent (1 to 2 words if article, else 1)
	themeLen := 1
	firstLower := strings.ToLower(words[0])
	if firstLower == "the" || firstLower == "a" || firstLower == "an" || firstLower == "this" || firstLower == "that" {
		if len(words) > 2 {
			themeLen = 2
		}
	}

	theme := strings.Join(words[:themeLen], " ")
	rheme := strings.Join(words[themeLen:], " ")

	return ThemeRheme{
		Clause: clause,
		Theme:  theme,
		Rheme:  rheme,
	}
}
