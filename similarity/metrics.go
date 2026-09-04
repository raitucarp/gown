package similarity

import (
	"math"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/graph"
)

// Metric identifies the semantic similarity measure to apply.
type Metric string

const (
	MetricWuPalmer        Metric = "wu_palmer"
	MetricPath            Metric = "path"
	MetricLeacockChodorow Metric = "leacock_chodorow"
	MetricResnik          Metric = "resnik"
	MetricLin             Metric = "lin"
	MetricJiangConrath    Metric = "jiang_conrath"
	MetricLevenshtein     Metric = "levenshtein"
)

// WuPalmer computes the Wu-Palmer similarity score between two synsets (0.0 to 1.0).
// Sim_wup = 2 * depth(LCS) / (depth(s1) + depth(s2))
func WuPalmer(res *gown.LexicalResource, s1, s2 *gown.Synset) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}
	if s1.ID == s2.ID {
		return 1.0
	}

	lcs, lcsDepth := graph.LowestCommonHypernym(res, s1, s2)
	if lcs == nil || lcsDepth <= 0 {
		return 0.0
	}

	d1 := graph.SynsetDepth(res, s1)
	d2 := graph.SynsetDepth(res, s2)

	if d1+d2 == 0 {
		return 0.0
	}

	score := (2.0 * float64(lcsDepth)) / float64(d1+d2)
	if score > 1.0 {
		score = 1.0
	}
	return score
}

// PathSimilarity computes path-based distance similarity: 1 / (1 + shortest_path_distance).
func PathSimilarity(res *gown.LexicalResource, s1, s2 *gown.Synset) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}
	if s1.ID == s2.ID {
		return 1.0
	}

	anc1 := graph.HypernymAncestors(res, s1)
	anc2 := graph.HypernymAncestors(res, s2)

	minDist := math.MaxInt32
	for id, dist1 := range anc1 {
		if dist2, ok := anc2[id]; ok {
			total := dist1 + dist2
			if total < minDist {
				minDist = total
			}
		}
	}

	if minDist == math.MaxInt32 {
		return 0.0
	}

	return 1.0 / (1.0 + float64(minDist))
}

// LeacockChodorow computes the LCH similarity measure:
// Sim_lch = -ln(length / (2 * max_depth))
func LeacockChodorow(res *gown.LexicalResource, s1, s2 *gown.Synset, maxDepth ...int) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}

	md := 20
	if len(maxDepth) > 0 && maxDepth[0] > 0 {
		md = maxDepth[0]
	}

	if s1.ID == s2.ID {
		return -math.Log(1.0 / (2.0 * float64(md)))
	}

	pathSim := PathSimilarity(res, s1, s2)
	if pathSim <= 0.0 {
		return 0.0
	}

	// length is shortest path + 1
	dist := (1.0 / pathSim) - 1.0
	length := dist + 1.0

	score := -math.Log(length / (2.0 * float64(md)))
	if score < 0 {
		return 0
	}
	return score
}

// Resnik computes Resnik similarity: IC(LCS).
func Resnik(res *gown.LexicalResource, s1, s2 *gown.Synset, calc *InformationContentCalculator) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}
	lcs, _ := graph.LowestCommonHypernym(res, s1, s2)
	if lcs == nil {
		return 0.0
	}
	return calc.IC(lcs)
}

// Lin computes Lin similarity: 2 * IC(LCS) / (IC(s1) + IC(s2)).
func Lin(res *gown.LexicalResource, s1, s2 *gown.Synset, calc *InformationContentCalculator) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}
	if s1.ID == s2.ID {
		return 1.0
	}

	lcs, _ := graph.LowestCommonHypernym(res, s1, s2)
	if lcs == nil {
		return 0.0
	}

	icLCS := calc.IC(lcs)
	ic1 := calc.IC(s1)
	ic2 := calc.IC(s2)

	denom := ic1 + ic2
	if denom <= 0.0 {
		return 0.0
	}

	score := (2.0 * icLCS) / denom
	if score > 1.0 {
		score = 1.0
	}
	return score
}

// JiangConrath computes Jiang-Conrath distance and converts to similarity:
// Dist_jcn = IC(s1) + IC(s2) - 2 * IC(LCS)
// Sim_jcn = 1 / (1 + Dist_jcn)
func JiangConrath(res *gown.LexicalResource, s1, s2 *gown.Synset, calc *InformationContentCalculator) float64 {
	if s1 == nil || s2 == nil {
		return 0.0
	}
	if s1.ID == s2.ID {
		return 1.0
	}

	lcs, _ := graph.LowestCommonHypernym(res, s1, s2)
	if lcs == nil {
		return 0.0
	}

	icLCS := calc.IC(lcs)
	ic1 := calc.IC(s1)
	ic2 := calc.IC(s2)

	dist := ic1 + ic2 - (2.0 * icLCS)
	if dist < 0 {
		dist = 0
	}
	return 1.0 / (1.0 + dist)
}

// Levenshtein computes the normalized edit similarity between two strings (0.0 to 1.0).
func Levenshtein(a, b string) float64 {
	ra := []rune(a)
	rb := []rune(b)
	la, lb := len(ra), len(rb)

	if la == 0 && lb == 0 {
		return 1.0
	}
	if la == 0 || lb == 0 {
		return 0.0
	}

	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 1
			if ra[i-1] == rb[j-1] {
				cost = 0
			}
			dp[i][j] = min(
				dp[i-1][j]+1,
				dp[i][j-1]+1,
				dp[i-1][j-1]+cost,
			)
		}
	}

	maxLen := max(la, lb)
	distance := dp[la][lb]
	return 1.0 - (float64(distance) / float64(maxLen))
}
