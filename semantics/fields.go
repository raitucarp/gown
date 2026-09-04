package semantics

import (
	"strings"

	"github.com/raitucarp/gown"
)

// SemanticField represents a conceptual domain grouping related lexical items.
type SemanticField struct {
	Domain string   `json:"domain"`
	Words  []string `json:"words"`
}

// ClusterBySemanticField partitions a list of words into clusters according to WordNet lexfiles.
func ClusterBySemanticField(res *gown.LexicalResource, words []string) []SemanticField {
	clusters := make(map[string][]string)

	for _, w := range words {
		wClean := strings.ToLower(strings.TrimSpace(w))
		entries := res.Lookup(wClean)
		if len(entries) == 0 {
			clusters["other"] = append(clusters["other"], wClean)
			continue
		}

		field := "other"
		for _, e := range entries {
			for _, s := range e.Synsets() {
				if s.Lexfile != "" {
					field = s.Lexfile
					break
				}
			}
			if field != "other" {
				break
			}
		}

		clusters[field] = append(clusters[field], wClean)
	}

	var result []SemanticField
	for domain, list := range clusters {
		result = append(result, SemanticField{
			Domain: domain,
			Words:  list,
		})
	}

	return result
}
