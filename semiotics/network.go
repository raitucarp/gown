package semiotics

import (
	"github.com/raitucarp/gown"
)

// SemioticNode represents a sign vertex in a semiotic network.
type SemioticNode struct {
	Signifier string   `json:"signifier"`
	Signified string   `json:"signified"`
	Mode      SignMode `json:"mode"`
}

// SemioticLink represents an interpretative or signifying edge between two signs.
type SemioticLink struct {
	Source      string `json:"source"`
	Target      string `json:"target"`
	Translation string `json:"translation"` // "synonymy", "hypernymy", "metaphor", "symbolic"
}

// SemioticNetwork models Umberto Eco's and C. S. Peirce's concept of unlimited semiosis:
// a sign's interpretant becomes the representamen of a new sign.
type SemioticNetwork struct {
	Nodes map[string]SemioticNode `json:"nodes"`
	Links []SemioticLink          `json:"links"`
}

// NewSemioticNetwork constructs an empty semiotic network.
func NewSemioticNetwork() *SemioticNetwork {
	return &SemioticNetwork{
		Nodes: make(map[string]SemioticNode),
	}
}

// BuildSemioticNetwork constructs a semiotic network starting from a seed word,
// tracing signs through WordNet semantic relations.
func BuildSemioticNetwork(res *gown.LexicalResource, seedWord string, maxHops int) *SemioticNetwork {
	net := NewSemioticNetwork()
	visited := make(map[string]bool)
	queue := []string{seedWord}
	visited[seedWord] = true

	for hop := 0; hop < maxHops && len(queue) > 0; hop++ {
		nextQueue := []string{}

		for _, currWord := range queue {
			sign := CreateSaussureanSign(res, currWord)
			mode := ClassifySignMode(currWord)

			net.Nodes[currWord] = SemioticNode{
				Signifier: sign.Signifier,
				Signified: sign.Signified,
				Mode:      mode,
			}

			entries := res.Lookup(currWord)
			for _, e := range entries {
				// Synonyms
				for _, syn := range e.Relation().Synonyms() {
					targetWord := syn.Lemma.WrittenForm
					net.Links = append(net.Links, SemioticLink{
						Source:      currWord,
						Target:      targetWord,
						Translation: "synonymy",
					})
					if !visited[targetWord] && len(visited) < 30 {
						visited[targetWord] = true
						nextQueue = append(nextQueue, targetWord)
					}
				}

				// Hypernyms
				for _, hyp := range e.Relation().Hypernyms() {
					targetWord := hyp.Lemma.WrittenForm
					net.Links = append(net.Links, SemioticLink{
						Source:      currWord,
						Target:      targetWord,
						Translation: "hypernymy",
					})
					if !visited[targetWord] && len(visited) < 30 {
						visited[targetWord] = true
						nextQueue = append(nextQueue, targetWord)
					}
				}
			}
		}

		queue = nextQueue
	}

	return net
}
