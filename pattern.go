package gown

import (
	"regexp"
	"strings"
	"unicode"
)

// PatternMode specifies how CV patterns are calculated.
type PatternMode int

const (
	// ModeOrthographic classifies letters based on spelling rules.
	ModeOrthographic PatternMode = iota
	// ModePhonological classifies phonemes from phonetic IPA representations.
	ModePhonological
)

// PatternClassifierConfig configures vowel/consonant classification.
type PatternClassifierConfig struct {
	TreatYAsVowelWhenMedial bool
	Mode                    PatternMode
	IgnoreCase              bool
}

// DefaultPatternConfig provides standard English orthographic defaults.
var DefaultPatternConfig = PatternClassifierConfig{
	TreatYAsVowelWhenMedial: true,
	Mode:                    ModeOrthographic,
	IgnoreCase:              true,
}

// IsVowelChar returns true if a rune is an English vowel letter.
func IsVowelChar(r rune, isMedial bool, treatYAsVowel bool) bool {
	lower := unicode.ToLower(r)
	switch lower {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	case 'y':
		if treatYAsVowel && isMedial {
			return true
		}
		return false
	default:
		return false
	}
}

// OrthographicCV converts an English word to its Vowel-Consonant (CV) pattern.
// Examples:
// "cat" -> "CVC"
// "eat" -> "VVC"
// "bread" -> "CCVVC"
// "rhythm" -> "CCCCCCC" or "CCVCCC" (with Y treated as medial vowel)
func OrthographicCV(word string, cfg ...PatternClassifierConfig) string {
	c := DefaultPatternConfig
	if len(cfg) > 0 {
		c = cfg[0]
	}

	runes := []rune(word)
	if len(runes) == 0 {
		return ""
	}

	var sb strings.Builder
	for i, r := range runes {
		if !unicode.IsLetter(r) {
			continue
		}
		isMedial := i > 0 && i < len(runes)-1
		if IsVowelChar(r, isMedial, c.TreatYAsVowelWhenMedial) {
			sb.WriteRune('V')
		} else {
			sb.WriteRune('C')
		}
	}
	return sb.String()
}

// PatternQuery represents a compiled linguistic pattern.
type PatternQuery struct {
	re *regexp.Regexp
}

// CompilePattern compiles a CV pattern with wildcard and template support:
// - 'C': Consonant
// - 'V': Vowel
// - '?': Matches exactly one C or V
// - '*': Zero or more of previous character (e.g. "C*V*")
// - '+': One or more of previous character
// - '[VOWEL]': alias for 'V'
// - '[CONSONANT]': alias for 'C'
func CompilePattern(pattern string) (*PatternQuery, error) {
	pattern = strings.ToUpper(strings.TrimSpace(pattern))
	pattern = strings.ReplaceAll(pattern, "[VOWEL]", "V")
	pattern = strings.ReplaceAll(pattern, "[CONSONANT]", "C")

	// Convert wildcards:
	// If pattern contains standard regex operators, convert appropriately
	var regexStr strings.Builder
	regexStr.WriteString("^")
	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		switch ch {
		case '?':
			regexStr.WriteString("[CV]")
		case '*':
			regexStr.WriteString("*")
		case '+':
			regexStr.WriteString("+")
		case 'C', 'V':
			regexStr.WriteByte(ch)
		case '{', '}', ',', '(', ')', '|':
			regexStr.WriteByte(ch)
		default:
			// ignore or literal
		}
	}
	regexStr.WriteString("$")

	re, err := regexp.Compile(regexStr.String())
	if err != nil {
		return nil, err
	}
	return &PatternQuery{re: re}, nil
}

// Matches returns true if the candidate CV pattern matches the query.
func (pq *PatternQuery) Matches(cvPattern string) bool {
	return pq.re.MatchString(cvPattern)
}

// MatchesWord returns true if the word's CV pattern matches the query.
func (pq *PatternQuery) MatchesWord(word string) bool {
	cv := OrthographicCV(word)
	return pq.Matches(cv)
}

// FindByPattern searches the LexicalResource for all entries matching the given pattern template.
func (resource *LexicalResource) FindByPattern(pattern string, maxResults ...int) (LexicalEntries, error) {
	pq, err := CompilePattern(pattern)
	if err != nil {
		return nil, err
	}

	limit := -1
	if len(maxResults) > 0 && maxResults[0] > 0 {
		limit = maxResults[0]
	}

	var results LexicalEntries
	seen := make(map[string]bool)

	for i := range resource.Lexicon.LexicalEntries {
		entry := &resource.Lexicon.LexicalEntries[i]
		if seen[entry.Lemma.WrittenForm] {
			continue
		}
		if pq.MatchesWord(entry.Lemma.WrittenForm) {
			seen[entry.Lemma.WrittenForm] = true
			results = append(results, *entry)
			if limit > 0 && len(results) >= limit {
				break
			}
		}
	}

	return results, nil
}
