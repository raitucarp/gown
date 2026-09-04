package similarity

import (
	"math"
	"sync"

	"github.com/raitucarp/gown"
)

// InformationContentCalculator computes intrinsic Information Content (IC) for WordNet synsets
// using the algorithm of Seco, Gomez, and Patricio (2004):
// IC(s) = 1 - (log(hypo(s) + 1) / log(max_nodes))
type InformationContentCalculator struct {
	res      *gown.LexicalResource
	icMap    map[string]float64
	maxNodes float64
	mu       sync.RWMutex
}

// NewICCalculator constructs an Information Content calculator.
func NewICCalculator(res *gown.LexicalResource) *InformationContentCalculator {
	return &InformationContentCalculator{
		res:   res,
		icMap: make(map[string]float64),
	}
}

// IC returns the Information Content of a synset, computing and caching it on demand.
func (calc *InformationContentCalculator) IC(synset *gown.Synset) float64 {
	if synset == nil {
		return 0.0
	}

	calc.mu.RLock()
	if val, ok := calc.icMap[synset.ID]; ok {
		calc.mu.RUnlock()
		return val
	}
	calc.mu.RUnlock()

	calc.mu.Lock()
	defer calc.mu.Unlock()

	if val, ok := calc.icMap[synset.ID]; ok {
		return val
	}

	if calc.maxNodes == 0 {
		calc.maxNodes = float64(len(calc.res.Lexicon.Synsets))
		if calc.maxNodes <= 1 {
			calc.maxNodes = 100000
		}
	}

	hypos := calc.countHyponyms(synset)
	score := 1.0 - (math.Log(float64(hypos)+1.0) / math.Log(calc.maxNodes))
	if score < 0 {
		score = 0
	} else if score > 1.0 {
		score = 1.0
	}

	calc.icMap[synset.ID] = score
	return score
}

func (calc *InformationContentCalculator) countHyponyms(synset *gown.Synset) int {
	visited := make(map[string]bool)
	queue := []string{synset.ID}
	visited[synset.ID] = true
	synsetsById := calc.res.SynsetsById()

	count := 0
	for len(queue) > 0 {
		currID := queue[0]
		queue = queue[1:]

		s := synsetsById[currID]
		if s == nil {
			continue
		}

		for _, rel := range s.SynsetRelations {
			if rel.RelType == string(gown.SynsetRelationTypeHyponym) ||
				rel.RelType == string(gown.SynsetRelationTypeInstanceHyponym) {
				if !visited[rel.Target] {
					visited[rel.Target] = true
					count++
					queue = append(queue, rel.Target)
				}
			}
		}
	}

	return count
}
