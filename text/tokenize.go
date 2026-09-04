package text

import (
	"strings"
	"unicode"
)

// DefaultStopwords contains standard English stopwords.
var DefaultStopwords = map[string]struct{}{
	"a": {}, "about": {}, "above": {}, "after": {}, "again": {}, "against": {}, "all": {}, "am": {}, "an": {},
	"and": {}, "any": {}, "are": {}, "aren't": {}, "as": {}, "at": {}, "be": {}, "because": {}, "been": {},
	"before": {}, "being": {}, "below": {}, "between": {}, "both": {}, "but": {}, "by": {}, "can't": {},
	"cannot": {}, "could": {}, "couldn't": {}, "did": {}, "didn't": {}, "do": {}, "does": {}, "doesn't": {},
	"doing": {}, "don't": {}, "down": {}, "during": {}, "each": {}, "few": {}, "for": {}, "from": {},
	"further": {}, "had": {}, "hadn't": {}, "has": {}, "hasn't": {}, "have": {}, "haven't": {}, "having": {},
	"he": {}, "he'd": {}, "he'll": {}, "he's": {}, "her": {}, "here": {}, "here's": {}, "hers": {},
	"herself": {}, "him": {}, "himself": {}, "his": {}, "how": {}, "how's": {}, "i": {}, "i'd": {},
	"i'll": {}, "i'm": {}, "i've": {}, "if": {}, "in": {}, "into": {}, "is": {}, "isn't": {}, "it": {},
	"it's": {}, "its": {}, "itself": {}, "let's": {}, "me": {}, "more": {}, "most": {}, "mustn't": {},
	"my": {}, "myself": {}, "no": {}, "nor": {}, "not": {}, "of": {}, "off": {}, "on": {}, "once": {},
	"only": {}, "or": {}, "other": {}, "ought": {}, "our": {}, "ours": {}, "ourselves": {}, "out": {},
	"over": {}, "own": {}, "same": {}, "shan't": {}, "she": {}, "she'd": {}, "she'll": {}, "she's": {},
	"should": {}, "shouldn't": {}, "so": {}, "some": {}, "such": {}, "than": {}, "that": {}, "that's": {},
	"the": {}, "their": {}, "theirs": {}, "them": {}, "themselves": {}, "then": {}, "there": {},
	"there's": {}, "these": {}, "they": {}, "they'd": {}, "they'll": {}, "they're": {}, "they've": {},
	"this": {}, "those": {}, "through": {}, "to": {}, "too": {}, "under": {}, "until": {}, "up": {},
	"very": {}, "was": {}, "wasn't": {}, "we": {}, "we'd": {}, "we'll": {}, "we're": {}, "we've": {},
	"were": {}, "weren't": {}, "what": {}, "what's": {}, "when": {}, "when's": {}, "where": {},
	"where's": {}, "which": {}, "while": {}, "who": {}, "who's": {}, "whom": {}, "why": {}, "why's": {},
	"with": {}, "won't": {}, "would": {}, "wouldn't": {}, "you": {}, "you'd": {}, "you'll": {},
	"you're": {}, "you've": {}, "your": {}, "yours": {}, "yourself": {}, "yourselves": {},
	"one": {}, "two": {}, "also": {}, "often": {}, "used": {}, "especially": {}, "usually": {},
}

// IsStopword checks if a lowercased word is a stopword.
func IsStopword(word string) bool {
	_, ok := DefaultStopwords[strings.ToLower(word)]
	return ok
}

// Tokenize splits a sentence into clean, lowercased word tokens.
func Tokenize(text string) []string {
	var tokens []string
	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r) && r != '-' && r != '\''
	})
	for _, w := range words {
		w = strings.Trim(w, "'-")
		if w != "" {
			tokens = append(tokens, strings.ToLower(w))
		}
	}
	return tokens
}

// ExtractContentWords tokenizes text and removes punctuation, numbers, and stopwords.
func ExtractContentWords(text string) []string {
	tokens := Tokenize(text)
	var content []string
	for _, t := range tokens {
		if !IsStopword(t) && len(t) > 1 {
			content = append(content, t)
		}
	}
	return content
}

// SentenceSegment splits text into sentences.
func SentenceSegment(text string) []string {
	var sentences []string
	var cur strings.Builder

	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		cur.WriteRune(r)
		if r == '.' || r == '!' || r == '?' {
			// Check if end of sentence or abbreviation
			if i+1 == len(runes) || unicode.IsSpace(runes[i+1]) {
				s := strings.TrimSpace(cur.String())
				if s != "" {
					sentences = append(sentences, s)
				}
				cur.Reset()
			}
		}
	}
	if cur.Len() > 0 {
		s := strings.TrimSpace(cur.String())
		if s != "" {
			sentences = append(sentences, s)
		}
	}
	return sentences
}

// JaccardSimilarity calculates the Jaccard similarity coefficient between two token sets.
func JaccardSimilarity(tokens1, tokens2 []string) float64 {
	set1 := make(map[string]struct{}, len(tokens1))
	for _, t := range tokens1 {
		set1[t] = struct{}{}
	}
	set2 := make(map[string]struct{}, len(tokens2))
	for _, t := range tokens2 {
		set2[t] = struct{}{}
	}

	if len(set1) == 0 && len(set2) == 0 {
		return 1.0
	}

	intersection := 0
	for t := range set1 {
		if _, ok := set2[t]; ok {
			intersection++
		}
	}

	union := len(set1) + len(set2) - intersection
	if union == 0 {
		return 0.0
	}
	return float64(intersection) / float64(union)
}
