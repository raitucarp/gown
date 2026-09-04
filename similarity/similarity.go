package similarity

import (
	"fmt"

	"github.com/raitucarp/gown"
)

// CompareConfig configures semantic similarity comparisons.
type CompareConfig struct {
	Metric   Metric
	POS      gown.POS
	ICCalc   *InformationContentCalculator
	MaxDepth int
}

// CompareOption is a functional option for similarity scoring.
type CompareOption func(*CompareConfig)

// WithMetric specifies the similarity metric.
func WithMetric(m Metric) CompareOption {
	return func(c *CompareConfig) {
		c.Metric = m
	}
}

// WithPOS restricts comparison to synsets of a specific part of speech.
func WithPOS(pos gown.POS) CompareOption {
	return func(c *CompareConfig) {
		c.POS = pos
	}
}

// WithICCalculator provides a pre-initialized Information Content calculator.
func WithICCalculator(calc *InformationContentCalculator) CompareOption {
	return func(c *CompareConfig) {
		c.ICCalc = calc
	}
}

// WithMaxDepth configures maximum tree depth for Leacock-Chodorow.
func WithMaxDepth(depth int) CompareOption {
	return func(c *CompareConfig) {
		c.MaxDepth = depth
	}
}

// Compare computes the semantic similarity score between two words by finding the maximum similarity
// across their respective synset senses.
func Compare(res *gown.LexicalResource, word1, word2 string, opts ...CompareOption) (float64, error) {
	cfg := CompareConfig{
		Metric:   MetricWuPalmer,
		MaxDepth: 20,
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	if cfg.Metric == MetricLevenshtein {
		return Levenshtein(word1, word2), nil
	}

	entries1 := res.Lookup(word1)
	entries2 := res.Lookup(word2)
	if len(entries1) == 0 {
		return 0.0, fmt.Errorf("word not found: %s", word1)
	}
	if len(entries2) == 0 {
		return 0.0, fmt.Errorf("word not found: %s", word2)
	}

	var synsets1, synsets2 []*gown.Synset
	for _, e := range entries1 {
		if cfg.POS != "" && gown.POS(e.Lemma.PartOfSpeech) != cfg.POS {
			continue
		}
		for _, s := range e.Synsets() {
			if s != nil {
				synsets1 = append(synsets1, s)
			}
		}
	}
	for _, e := range entries2 {
		if cfg.POS != "" && gown.POS(e.Lemma.PartOfSpeech) != cfg.POS {
			continue
		}
		for _, s := range e.Synsets() {
			if s != nil {
				synsets2 = append(synsets2, s)
			}
		}
	}

	if len(synsets1) == 0 || len(synsets2) == 0 {
		return 0.0, fmt.Errorf("no matching synsets for %s and %s with POS %s", word1, word2, cfg.POS)
	}

	if (cfg.Metric == MetricResnik || cfg.Metric == MetricLin || cfg.Metric == MetricJiangConrath) && cfg.ICCalc == nil {
		cfg.ICCalc = NewICCalculator(res)
	}

	maxScore := -1.0
	for _, s1 := range synsets1 {
		for _, s2 := range synsets2 {
			score := 0.0
			switch cfg.Metric {
			case MetricWuPalmer:
				score = WuPalmer(res, s1, s2)
			case MetricPath:
				score = PathSimilarity(res, s1, s2)
			case MetricLeacockChodorow:
				score = LeacockChodorow(res, s1, s2, cfg.MaxDepth)
			case MetricResnik:
				score = Resnik(res, s1, s2, cfg.ICCalc)
			case MetricLin:
				score = Lin(res, s1, s2, cfg.ICCalc)
			case MetricJiangConrath:
				score = JiangConrath(res, s1, s2, cfg.ICCalc)
			}
			if score > maxScore {
				maxScore = score
			}
		}
	}

	if maxScore < 0 {
		return 0.0, nil
	}
	return maxScore, nil
}
