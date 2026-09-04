package generative

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/raitucarp/gown"
)

// ProductionRule represents a context-free grammar rewrite rule: LHS -> RHS[0] RHS[1] ...
type ProductionRule struct {
	LHS string   `json:"lhs"`
	RHS []string `json:"rhs"`
}

// Grammar represents a context-free phrase structure grammar.
type Grammar struct {
	StartSymbol string           `json:"start_symbol"`
	Rules       []ProductionRule `json:"rules"`
	ruleMap     map[string][]ProductionRule
}

// NewGrammar constructs a new CFG with the specified start symbol.
func NewGrammar(start string) *Grammar {
	return &Grammar{
		StartSymbol: start,
		ruleMap:     make(map[string][]ProductionRule),
	}
}

// AddRule registers a phrase structure production rule.
func (g *Grammar) AddRule(lhs string, rhs ...string) {
	rule := ProductionRule{LHS: lhs, RHS: rhs}
	g.Rules = append(g.Rules, rule)
	g.ruleMap[lhs] = append(g.ruleMap[lhs], rule)
}

// StandardEnglishGrammar returns a classic Chomskyan phrase structure grammar:
// S -> NP VP
// NP -> Det N | Det Adj N | ProperNoun
// VP -> V | V NP | V NP PP
// PP -> P NP
func StandardEnglishGrammar() *Grammar {
	g := NewGrammar("S")
	g.AddRule("S", "NP", "VP")
	g.AddRule("NP", "Det", "N")
	g.AddRule("NP", "Det", "Adj", "N")
	g.AddRule("NP", "ProperNoun")
	g.AddRule("VP", "V")
	g.AddRule("VP", "V", "NP")
	g.AddRule("VP", "V", "NP", "PP")
	g.AddRule("PP", "P", "NP")

	// Terminal productions for closed classes
	g.AddRule("Det", "the")
	g.AddRule("Det", "a")
	g.AddRule("P", "in")
	g.AddRule("P", "on")
	g.AddRule("P", "with")
	g.AddRule("P", "by")
	g.AddRule("ProperNoun", "Alice")
	g.AddRule("ProperNoun", "Bob")

	return g
}

// GeneratorConfig configures grammar-driven sentence derivation.
type GeneratorConfig struct {
	MaxDepth int
	Lexicon  *gown.LexicalResource
}

// Generate derives a parse tree from the grammar, optionally inserting open-class words
// (N, V, Adj, Adv) dynamically from WordNet.
func (g *Grammar) Generate(cfg GeneratorConfig) (*ParseNode, error) {
	if cfg.MaxDepth == 0 {
		cfg.MaxDepth = 10
	}
	return g.expandSymbol(g.StartSymbol, 0, cfg)
}

func (g *Grammar) expandSymbol(symbol string, depth int, cfg GeneratorConfig) (*ParseNode, error) {
	if depth > cfg.MaxDepth {
		return nil, fmt.Errorf("exceeded max derivation depth %d for symbol %s", cfg.MaxDepth, symbol)
	}

	// 1. Check if symbol should be populated from WordNet
	if cfg.Lexicon != nil {
		if terminal := g.getWordNetLexicalItem(symbol, cfg.Lexicon); terminal != "" {
			return &ParseNode{
				Symbol:   symbol,
				Terminal: terminal,
			}, nil
		}
	}

	rules := g.ruleMap[symbol]
	if len(rules) == 0 {
		// Terminal symbol
		return &ParseNode{
			Symbol:   symbol,
			Terminal: symbol,
		}, nil
	}

	// Pick a random production rule for this non-terminal
	rule := rules[rand.Intn(len(rules))]

	// If RHS is a single lowercase token or terminal, produce leaf
	if len(rule.RHS) == 1 && strings.ToLower(rule.RHS[0]) == rule.RHS[0] && !g.isNonTerminal(rule.RHS[0]) {
		return &ParseNode{
			Symbol:   symbol,
			Terminal: rule.RHS[0],
		}, nil
	}

	node := &ParseNode{Symbol: symbol}
	for _, rhsSym := range rule.RHS {
		child, err := g.expandSymbol(rhsSym, depth+1, cfg)
		if err != nil {
			return nil, err
		}
		node.Children = append(node.Children, child)
	}

	return node, nil
}

func (g *Grammar) isNonTerminal(s string) bool {
	return len(g.ruleMap[s]) > 0
}

func (g *Grammar) getWordNetLexicalItem(symbol string, res *gown.LexicalResource) string {
	switch symbol {
	case "N", "Noun":
		nouns := res.Nouns().Words().Random(1)
		if len(nouns) > 0 {
			return nouns[0].Lemma.WrittenForm
		}
	case "V", "Verb":
		verbs := res.Verbs().Words().Random(1)
		if len(verbs) > 0 {
			return verbs[0].Lemma.WrittenForm
		}
	case "Adj", "Adjective":
		adjs := res.Adjectives().Words().Random(1)
		if len(adjs) > 0 {
			return adjs[0].Lemma.WrittenForm
		}
	case "Adv", "Adverb":
		advs := res.Adverbs().Words().Random(1)
		if len(advs) > 0 {
			return advs[0].Lemma.WrittenForm
		}
	}
	return ""
}
