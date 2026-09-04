package similarity_test

import (
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/similarity"
)

func TestSimilarityCompareAllMetrics(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	calc := similarity.NewICCalculator(res)

	tests := []struct {
		name    string
		w1, w2  string
		opts    []similarity.CompareOption
		wantErr bool
		minVal  float64
		maxVal  float64
	}{
		{
			name:   "Wu-Palmer dog vs cat",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricWuPalmer), similarity.WithPOS(gown.NounPos)},
			minVal: 0.5,
			maxVal: 1.0,
		},
		{
			name:   "Wu-Palmer dog vs dog identity",
			w1:     "dog",
			w2:     "dog",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricWuPalmer)},
			minVal: 0.99,
			maxVal: 1.0,
		},
		{
			name:   "Path similarity dog vs cat",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricPath)},
			minVal: 0.05,
			maxVal: 1.0,
		},
		{
			name:   "Leacock-Chodorow dog vs cat default depth",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricLeacockChodorow)},
			minVal: 0.5,
			maxVal: 5.0,
		},
		{
			name:   "Leacock-Chodorow dog vs wolf with custom depth",
			w1:     "dog",
			w2:     "wolf",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricLeacockChodorow), similarity.WithMaxDepth(25)},
			minVal: 0.5,
			maxVal: 5.0,
		},
		{
			name:   "Resnik dog vs cat with precomputed IC",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricResnik), similarity.WithICCalculator(calc)},
			minVal: 0.1,
			maxVal: 1.0,
		},
		{
			name:   "Resnik dog vs wolf with nil IC (auto-initializes)",
			w1:     "dog",
			w2:     "wolf",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricResnik)},
			minVal: 0.1,
			maxVal: 1.0,
		},
		{
			name:   "Lin dog vs cat",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricLin), similarity.WithICCalculator(calc)},
			minVal: 0.1,
			maxVal: 1.0,
		},
		{
			name:   "Jiang-Conrath dog vs cat",
			w1:     "dog",
			w2:     "cat",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricJiangConrath), similarity.WithICCalculator(calc)},
			minVal: 0.1,
			maxVal: 1.0,
		},
		{
			name:   "Levenshtein identical",
			w1:     "kitten",
			w2:     "kitten",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricLevenshtein)},
			minVal: 0.99,
			maxVal: 1.0,
		},
		{
			name:   "Levenshtein different",
			w1:     "kitten",
			w2:     "sitting",
			opts:   []similarity.CompareOption{similarity.WithMetric(similarity.MetricLevenshtein)},
			minVal: 0.1,
			maxVal: 0.9,
		},
		{
			name:    "Missing word1 error",
			w1:      "nonexistentxyzword123",
			w2:      "dog",
			wantErr: true,
		},
		{
			name:    "Missing word2 error",
			w1:      "dog",
			w2:      "nonexistentxyzword123",
			wantErr: true,
		},
		{
			name:    "POS filter excludes all synsets",
			w1:      "dog",
			w2:      "cat",
			opts:    []similarity.CompareOption{similarity.WithPOS(gown.AdverbPos)},
			wantErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			score, err := similarity.Compare(res, tc.w1, tc.w2, tc.opts...)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error for %s, got nil", tc.name)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error for %s: %v", tc.name, err)
			}
			if score < tc.minVal || score > tc.maxVal {
				t.Errorf("%s: score %.4f out of expected range [%.4f, %.4f]", tc.name, score, tc.minVal, tc.maxVal)
			}
		})
	}
}

func TestDirectMetricsEdgeCases(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	calc := similarity.NewICCalculator(res)

	dogEntries := res.Lookup("dog", gown.WithPOS(gown.NounPos))
	if len(dogEntries) == 0 || len(dogEntries[0].Synsets()) == 0 {
		t.Fatalf("dog synset not found")
	}
	sDog := dogEntries[0].Synsets()[0]

	catEntries := res.Lookup("cat", gown.WithPOS(gown.NounPos))
	if len(catEntries) == 0 || len(catEntries[0].Synsets()) == 0 {
		t.Fatalf("cat synset not found")
	}
	sCat := catEntries[0].Synsets()[0]

	// 1. WuPalmer edge cases
	if score := similarity.WuPalmer(res, nil, sCat); score != 0.0 {
		t.Errorf("WuPalmer(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.WuPalmer(res, sDog, nil); score != 0.0 {
		t.Errorf("WuPalmer(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.WuPalmer(res, sDog, sDog); score != 1.0 {
		t.Errorf("WuPalmer(sDog, sDog) = %.2f, expected 1.0", score)
	}

	// 2. PathSimilarity edge cases
	if score := similarity.PathSimilarity(res, nil, sCat); score != 0.0 {
		t.Errorf("PathSimilarity(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.PathSimilarity(res, sDog, nil); score != 0.0 {
		t.Errorf("PathSimilarity(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.PathSimilarity(res, sDog, sDog); score != 1.0 {
		t.Errorf("PathSimilarity(sDog, sDog) = %.2f, expected 1.0", score)
	}

	// 3. LeacockChodorow edge cases
	if score := similarity.LeacockChodorow(res, nil, sCat); score != 0.0 {
		t.Errorf("LeacockChodorow(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.LeacockChodorow(res, sDog, nil); score != 0.0 {
		t.Errorf("LeacockChodorow(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.LeacockChodorow(res, sDog, sDog); score <= 0.0 {
		t.Errorf("LeacockChodorow(sDog, sDog) = %.2f, expected > 0.0", score)
	}
	if score := similarity.LeacockChodorow(res, sDog, sDog, 30); score <= 0.0 {
		t.Errorf("LeacockChodorow(sDog, sDog, 30) = %.2f, expected > 0.0", score)
	}

	// 4. Resnik edge cases
	if score := similarity.Resnik(res, nil, sCat, calc); score != 0.0 {
		t.Errorf("Resnik(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.Resnik(res, sDog, nil, calc); score != 0.0 {
		t.Errorf("Resnik(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.Resnik(res, sDog, sDog, calc); score <= 0.0 {
		t.Errorf("Resnik(sDog, sDog) = %.2f, expected > 0.0", score)
	}

	// 5. Lin edge cases
	if score := similarity.Lin(res, nil, sCat, calc); score != 0.0 {
		t.Errorf("Lin(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.Lin(res, sDog, nil, calc); score != 0.0 {
		t.Errorf("Lin(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.Lin(res, sDog, sDog, calc); score != 1.0 {
		t.Errorf("Lin(sDog, sDog) = %.2f, expected 1.0", score)
	}

	// 6. JiangConrath edge cases
	if score := similarity.JiangConrath(res, nil, sCat, calc); score != 0.0 {
		t.Errorf("JiangConrath(nil, sCat) = %.2f, expected 0.0", score)
	}
	if score := similarity.JiangConrath(res, sDog, nil, calc); score != 0.0 {
		t.Errorf("JiangConrath(sDog, nil) = %.2f, expected 0.0", score)
	}
	if score := similarity.JiangConrath(res, sDog, sDog, calc); score != 1.0 {
		t.Errorf("JiangConrath(sDog, sDog) = %.2f, expected 1.0", score)
	}

	// 7. InformationContentCalculator edge cases
	if ic := calc.IC(nil); ic != 0.0 {
		t.Errorf("IC(nil) = %.2f, expected 0.0", ic)
	}
	ic1 := calc.IC(sDog)
	ic2 := calc.IC(sDog) // from cache
	if ic1 != ic2 {
		t.Errorf("Cached IC differs: %.4f vs %.4f", ic1, ic2)
	}
	if ic1 <= 0.0 || ic1 > 1.0 {
		t.Errorf("IC out of bounds [0, 1]: %.4f", ic1)
	}

	// 8. Levenshtein edge cases
	if sim := similarity.Levenshtein("", ""); sim != 1.0 {
		t.Errorf("Levenshtein empty-empty = %.2f, expected 1.0", sim)
	}
	if sim := similarity.Levenshtein("abc", ""); sim != 0.0 {
		t.Errorf("Levenshtein non-empty vs empty = %.2f, expected 0.0", sim)
	}
	if sim := similarity.Levenshtein("", "xyz"); sim != 0.0 {
		t.Errorf("Levenshtein empty vs non-empty = %.2f, expected 0.0", sim)
	}
	if sim := similarity.Levenshtein("linguistics", "linguistics"); sim != 1.0 {
		t.Errorf("Levenshtein identical = %.2f, expected 1.0", sim)
	}
	if sim := similarity.Levenshtein("semantics", "pragmatics"); sim <= 0.0 || sim >= 1.0 {
		t.Errorf("Levenshtein partial similarity out of range: %.4f", sim)
	}
}
