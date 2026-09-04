package phonology

import (
	"strings"
	"unicode"
)

// PhonemeType classifies phonemes into phonetic categories.
type PhonemeType string

const (
	PhonemeVowel     PhonemeType = "vowel"
	PhonemeConsonant PhonemeType = "consonant"
	PhonemeDiphthong PhonemeType = "diphthong"
)

// IPAPhoneme describes an IPA symbol and its acoustic/articulatory features.
type IPAPhoneme struct {
	Symbol   string
	Type     PhonemeType
	Voiced   bool
	Manner   string // "plosive", "fricative", "nasal", "approximant", "affricate"
	Place    string // "bilabial", "labiodental", "alveolar", "palatal", "velar", "glottal"
	Sonority int    // 1 (plosive) to 5 (vowel)
}

// IPATable maps standard English IPA symbols to articulatory properties.
var IPATable = map[string]IPAPhoneme{
	// Vowels (Sonority 5)
	"i": {Symbol: "i", Type: PhonemeVowel, Sonority: 5},
	"ɪ": {Symbol: "ɪ", Type: PhonemeVowel, Sonority: 5},
	"e": {Symbol: "e", Type: PhonemeVowel, Sonority: 5},
	"ɛ": {Symbol: "ɛ", Type: PhonemeVowel, Sonority: 5},
	"æ": {Symbol: "æ", Type: PhonemeVowel, Sonority: 5},
	"ɑ": {Symbol: "ɑ", Type: PhonemeVowel, Sonority: 5},
	"ɒ": {Symbol: "ɒ", Type: PhonemeVowel, Sonority: 5},
	"ɔ": {Symbol: "ɔ", Type: PhonemeVowel, Sonority: 5},
	"ʊ": {Symbol: "ʊ", Type: PhonemeVowel, Sonority: 5},
	"u": {Symbol: "u", Type: PhonemeVowel, Sonority: 5},
	"ʌ": {Symbol: "ʌ", Type: PhonemeVowel, Sonority: 5},
	"ə": {Symbol: "ə", Type: PhonemeVowel, Sonority: 5},
	"ɜ": {Symbol: "ɜ", Type: PhonemeVowel, Sonority: 5},

	// Diphthongs (Sonority 5)
	"eɪ": {Symbol: "eɪ", Type: PhonemeDiphthong, Sonority: 5},
	"aɪ": {Symbol: "aɪ", Type: PhonemeDiphthong, Sonority: 5},
	"ɔɪ": {Symbol: "ɔɪ", Type: PhonemeDiphthong, Sonority: 5},
	"aʊ": {Symbol: "aʊ", Type: PhonemeDiphthong, Sonority: 5},
	"oʊ": {Symbol: "oʊ", Type: PhonemeDiphthong, Sonority: 5},
	"əʊ": {Symbol: "əʊ", Type: PhonemeDiphthong, Sonority: 5},

	// Approximants & Glides (Sonority 4)
	"l": {Symbol: "l", Type: PhonemeConsonant, Manner: "approximant", Place: "alveolar", Sonority: 4},
	"r": {Symbol: "r", Type: PhonemeConsonant, Manner: "approximant", Place: "alveolar", Sonority: 4},
	"ɹ": {Symbol: "ɹ", Type: PhonemeConsonant, Manner: "approximant", Place: "alveolar", Sonority: 4},
	"w": {Symbol: "w", Type: PhonemeConsonant, Manner: "approximant", Place: "bilabial", Sonority: 4},
	"j": {Symbol: "j", Type: PhonemeConsonant, Manner: "approximant", Place: "palatal", Sonority: 4},

	// Nasals (Sonority 3)
	"m": {Symbol: "m", Type: PhonemeConsonant, Manner: "nasal", Place: "bilabial", Sonority: 3},
	"n": {Symbol: "n", Type: PhonemeConsonant, Manner: "nasal", Place: "alveolar", Sonority: 3},
	"ŋ": {Symbol: "ŋ", Type: PhonemeConsonant, Manner: "nasal", Place: "velar", Sonority: 3},

	// Fricatives (Sonority 2)
	"f": {Symbol: "f", Type: PhonemeConsonant, Manner: "fricative", Place: "labiodental", Sonority: 2},
	"v": {Symbol: "v", Type: PhonemeConsonant, Manner: "fricative", Place: "labiodental", Sonority: 2},
	"θ": {Symbol: "θ", Type: PhonemeConsonant, Manner: "fricative", Place: "dental", Sonority: 2},
	"ð": {Symbol: "ð", Type: PhonemeConsonant, Manner: "fricative", Place: "dental", Sonority: 2},
	"s": {Symbol: "s", Type: PhonemeConsonant, Manner: "fricative", Place: "alveolar", Sonority: 2},
	"z": {Symbol: "z", Type: PhonemeConsonant, Manner: "fricative", Place: "alveolar", Sonority: 2},
	"ʃ": {Symbol: "ʃ", Type: PhonemeConsonant, Manner: "fricative", Place: "palato-alveolar", Sonority: 2},
	"ʒ": {Symbol: "ʒ", Type: PhonemeConsonant, Manner: "fricative", Place: "palato-alveolar", Sonority: 2},
	"h": {Symbol: "h", Type: PhonemeConsonant, Manner: "fricative", Place: "glottal", Sonority: 2},

	// Plosives & Affricates (Sonority 1)
	"p":  {Symbol: "p", Type: PhonemeConsonant, Manner: "plosive", Place: "bilabial", Sonority: 1},
	"b":  {Symbol: "b", Type: PhonemeConsonant, Manner: "plosive", Place: "bilabial", Sonority: 1},
	"t":  {Symbol: "t", Type: PhonemeConsonant, Manner: "plosive", Place: "alveolar", Sonority: 1},
	"d":  {Symbol: "d", Type: PhonemeConsonant, Manner: "plosive", Place: "alveolar", Sonority: 1},
	"k":  {Symbol: "k", Type: PhonemeConsonant, Manner: "plosive", Place: "velar", Sonority: 1},
	"g":  {Symbol: "g", Type: PhonemeConsonant, Manner: "plosive", Place: "velar", Sonority: 1},
	"tʃ": {Symbol: "tʃ", Type: PhonemeConsonant, Manner: "affricate", Place: "palato-alveolar", Sonority: 1},
	"dʒ": {Symbol: "dʒ", Type: PhonemeConsonant, Manner: "affricate", Place: "palato-alveolar", Sonority: 1},
}

// IPAToPhonemes parses an IPA string into a sequence of phonemes, ignoring stress marks.
func IPAToPhonemes(ipa string) []IPAPhoneme {
	clean := strings.Trim(ipa, "/[]")
	var result []IPAPhoneme

	runes := []rune(clean)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		if r == 'ˈ' || r == 'ˌ' || r == '.' || r == 'ː' || unicode.IsSpace(r) {
			continue
		}

		// Try two-character phonemes first (affricates and diphthongs)
		if i+1 < len(runes) {
			two := string(runes[i : i+2])
			if p, ok := IPATable[two]; ok {
				result = append(result, p)
				i++
				continue
			}
		}

		one := string(r)
		if p, ok := IPATable[one]; ok {
			result = append(result, p)
		} else {
			// Heuristic default for unlisted consonants
			result = append(result, IPAPhoneme{
				Symbol:   one,
				Type:     PhonemeConsonant,
				Sonority: 1,
			})
		}
	}

	return result
}

// IPAToCV converts an IPA transcription into its phonetic C/V pattern.
func IPAToCV(ipa string) string {
	phonemes := IPAToPhonemes(ipa)
	var sb strings.Builder
	for _, p := range phonemes {
		if p.Type == PhonemeVowel || p.Type == PhonemeDiphthong {
			sb.WriteRune('V')
		} else {
			sb.WriteRune('C')
		}
	}
	return sb.String()
}
