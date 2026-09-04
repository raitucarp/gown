package expansion

import (
	"context"
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/text"
)

// Config specifies controls for lexical expansion.
type Config struct {
	MaxDepth           int
	MaxNodes           int
	RelationTypes      []string
	POS                gown.POS
	Strategy           TraversalStrategy
	ExpandDefinitions  bool
	IncludeSynonyms    bool
	Context            context.Context
}

// Option is a functional configuration option for expansion.
type Option func(*Config)

// WithMaxDepth sets the maximum search depth.
func WithMaxDepth(depth int) Option {
	return func(c *Config) {
		c.MaxDepth = depth
	}
}

// WithMaxNodes sets the maximum total nodes in the expansion tree.
func WithMaxNodes(max int) Option {
	return func(c *Config) {
		c.MaxNodes = max
	}
}

// WithRelations filters expansion by synset relation types (e.g. "hypernym", "hyponym").
func WithRelations(relations ...string) Option {
	return func(c *Config) {
		c.RelationTypes = relations
	}
}

// WithPOS restricts expansion to a specific part of speech.
func WithPOS(pos gown.POS) Option {
	return func(c *Config) {
		c.POS = pos
	}
}

// WithStrategy chooses between BFS and DFS.
func WithStrategy(strategy TraversalStrategy) Option {
	return func(c *Config) {
		c.Strategy = strategy
	}
}

// WithDefinitionExpansion toggles recursive token extraction from glosses.
func WithDefinitionExpansion(expand bool) Option {
	return func(c *Config) {
		c.ExpandDefinitions = expand
	}
}

// WithContext supplies a cancellation context.
func WithContext(ctx context.Context) Option {
	return func(c *Config) {
		c.Context = ctx
	}
}

// Expand builds a semantic expansion tree starting from a given word or lemma.
func Expand(res *gown.LexicalResource, word string, opts ...Option) (*Tree, error) {
	cfg := Config{
		MaxDepth:        2,
		MaxNodes:        50,
		IncludeSynonyms: true,
		Strategy:        StrategyBFS,
		Context:         context.Background(),
	}
	for _, opt := range opts {
		opt(&cfg)
	}

	entries := res.Lookup(word)
	if len(entries) == 0 {
		return &Tree{
			Root: &Node{
				ID:   word,
				Word: word,
				Type: NodeWord,
			},
			TotalNodes: 1,
		}, nil
	}

	root := &Node{
		ID:    word,
		Word:  word,
		Type:  NodeWord,
		Depth: 0,
	}

	visitedSynsets := make(map[string]bool)
	totalNodes := 1

	synsetsById := res.SynsetsById()
	lexicalsById := res.LexicalsById()

	// Gather initial synsets for the word
	var initialSynsets []*gown.Synset
	for _, entry := range entries {
		if cfg.POS != "" && gown.POS(entry.Lemma.PartOfSpeech) != cfg.POS {
			continue
		}
		for _, s := range entry.Synsets() {
			if s != nil && !visitedSynsets[s.ID] {
				visitedSynsets[s.ID] = true
				initialSynsets = append(initialSynsets, s)
			}
		}
	}

	type queueItem struct {
		synset *gown.Synset
		parent *Node
		depth  int
	}

	var queue []queueItem
	for _, s := range initialSynsets {
		sNode := &Node{
			ID:         s.ID,
			Word:       formatSynsetMembers(s, lexicalsById),
			Type:       NodeSynset,
			POS:        s.PartOfSpeech,
			Definition: s.PrimaryDefinition(),
			Relation:   "sense",
			Depth:      1,
		}
		root.Children = append(root.Children, sNode)
		totalNodes++

		if cfg.ExpandDefinitions && s.PrimaryDefinition() != "" {
			defWords := text.ExtractContentWords(s.PrimaryDefinition())
			for _, dw := range defWords {
				if totalNodes >= cfg.MaxNodes {
					break
				}
				dwNode := &Node{
					ID:       dw,
					Word:     dw,
					Type:     NodeToken,
					Relation: "definition_token",
					Depth:    2,
				}
				sNode.Children = append(sNode.Children, dwNode)
				totalNodes++
			}
		}

		if cfg.MaxDepth > 1 {
			queue = append(queue, queueItem{synset: s, parent: sNode, depth: 1})
		}
	}

	for len(queue) > 0 && totalNodes < cfg.MaxNodes {
		select {
		case <-cfg.Context.Done():
			return &Tree{Root: root, TotalNodes: totalNodes}, cfg.Context.Err()
		default:
		}

		var current queueItem
		if cfg.Strategy == StrategyDFS {
			current = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		} else {
			current = queue[0]
			queue = queue[1:]
		}

		if current.depth >= cfg.MaxDepth {
			continue
		}

		for _, rel := range current.synset.SynsetRelations {
			if totalNodes >= cfg.MaxNodes {
				break
			}
			if len(cfg.RelationTypes) > 0 && !containsString(cfg.RelationTypes, rel.RelType) {
				continue
			}

			target := synsetsById[rel.Target]
			if target == nil || visitedSynsets[target.ID] {
				continue
			}
			visitedSynsets[target.ID] = true

			relNode := &Node{
				ID:         target.ID,
				Word:       formatSynsetMembers(target, lexicalsById),
				Type:       NodeSynset,
				POS:        target.PartOfSpeech,
				Definition: target.PrimaryDefinition(),
				Relation:   rel.RelType,
				Depth:      current.depth + 1,
			}
			current.parent.Children = append(current.parent.Children, relNode)
			totalNodes++

			if current.depth+1 < cfg.MaxDepth {
				queue = append(queue, queueItem{
					synset: target,
					parent: relNode,
					depth:  current.depth + 1,
				})
			}
		}
	}

	return &Tree{
		Root:       root,
		TotalNodes: totalNodes,
	}, nil
}

// ExpandDefinition performs recursive definition-driven exploration.
func ExpandDefinition(res *gown.LexicalResource, word string, opts ...Option) (*Tree, error) {
	opts = append(opts, WithDefinitionExpansion(true))
	return Expand(res, word, opts...)
}

func formatSynsetMembers(s *gown.Synset, lexicalsById map[string]*gown.LexicalEntry) string {
	var members []string
	for _, m := range s.Members {
		if entry, ok := lexicalsById[m]; ok {
			members = append(members, entry.Lemma.WrittenForm)
		}
	}
	if len(members) == 0 {
		return s.ID
	}
	return strings.Join(members, ", ")
}

func containsString(slice []string, val string) bool {
	for _, s := range slice {
		if s == val {
			return true
		}
	}
	return false
}
