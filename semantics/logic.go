package semantics

import (
	"fmt"
	"strings"
)

// PredicateLogic represents a logical formula P(arg1, arg2, ...).
type PredicateLogic struct {
	Predicate string   `json:"predicate"`
	Arguments []string `json:"arguments"`
	Negated   bool     `json:"negated"`
}

// String formats the predicate formula, e.g. "Chased(dog, cat)" or "¬Sleep(john)".
func (pl PredicateLogic) String() string {
	prefix := ""
	if pl.Negated {
		prefix = "¬"
	}
	return fmt.Sprintf("%s%s(%s)", prefix, pl.Predicate, strings.Join(pl.Arguments, ", "))
}

// SimpleModel represents a first-order logic relational model / world state.
type SimpleModel struct {
	TruePropositions map[string]bool
}

// NewModel constructs a logic model with initial ground facts.
func NewModel(facts ...PredicateLogic) *SimpleModel {
	m := &SimpleModel{TruePropositions: make(map[string]bool)}
	for _, f := range facts {
		m.Assert(f)
	}
	return m
}

// Assert adds a fact to the model.
func (m *SimpleModel) Assert(f PredicateLogic) {
	key := fmt.Sprintf("%s(%s)", strings.ToLower(f.Predicate), strings.ToLower(strings.Join(f.Arguments, ",")))
	m.TruePropositions[key] = !f.Negated
}

// Evaluate checks whether a formula evaluates to true in the model.
func (m *SimpleModel) Evaluate(f PredicateLogic) bool {
	key := fmt.Sprintf("%s(%s)", strings.ToLower(f.Predicate), strings.ToLower(strings.Join(f.Arguments, ",")))
	isTrue := m.TruePropositions[key]
	if f.Negated {
		return !isTrue
	}
	return isTrue
}
