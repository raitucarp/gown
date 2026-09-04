package paradigm

import (
	"fmt"
)

// SlashDirection indicates whether an argument is expected on the left or right.
type SlashDirection int

const (
	SlashRight SlashDirection = iota // A / B (expects B on the right)
	SlashLeft                        // A \ B (expects B on the left)
)

// CategoryType represents a semantic or syntactic type in Categorial Grammar.
type CategoryType interface {
	String() string
	Equals(other CategoryType) bool
}

// AtomicCategory represents primitive types (e.g. "S", "NP", "N").
type AtomicCategory string

func (a AtomicCategory) String() string { return string(a) }
func (a AtomicCategory) Equals(other CategoryType) bool {
	if o, ok := other.(AtomicCategory); ok {
		return a == o
	}
	return false
}

const (
	CatS  AtomicCategory = "S"
	CatNP AtomicCategory = "NP"
	CatN  AtomicCategory = "N"
)

// ComplexCategory represents directional functor types (Result / Argument or Result \ Argument).
type ComplexCategory struct {
	Result    CategoryType
	Slash     SlashDirection
	Argument  CategoryType
}

func (c ComplexCategory) String() string {
	slashStr := "/"
	if c.Slash == SlashLeft {
		slashStr = `\`
	}
	return fmt.Sprintf("(%s%s%s)", c.Result.String(), slashStr, c.Argument.String())
}

func (c ComplexCategory) Equals(other CategoryType) bool {
	o, ok := other.(ComplexCategory)
	if !ok {
		return false
	}
	return c.Slash == o.Slash && c.Result.Equals(o.Result) && c.Argument.Equals(o.Argument)
}

// RightFunctor constructs A / B (expects B on the right).
func RightFunctor(result, argument CategoryType) ComplexCategory {
	return ComplexCategory{Result: result, Slash: SlashRight, Argument: argument}
}

// LeftFunctor constructs A \ B (expects B on the left).
func LeftFunctor(result, argument CategoryType) ComplexCategory {
	return ComplexCategory{Result: result, Slash: SlashLeft, Argument: argument}
}

// ForwardApply applies functor A/B to argument B on its right: (A / B) + B => A.
func ForwardApply(functor, argument CategoryType) (CategoryType, bool) {
	comp, ok := functor.(ComplexCategory)
	if !ok || comp.Slash != SlashRight {
		return nil, false
	}
	if comp.Argument.Equals(argument) {
		return comp.Result, true
	}
	return nil, false
}

// BackwardApply applies functor A\B to argument B on its left: B + (A \ B) => A.
func BackwardApply(argument, functor CategoryType) (CategoryType, bool) {
	comp, ok := functor.(ComplexCategory)
	if !ok || comp.Slash != SlashLeft {
		return nil, false
	}
	if comp.Argument.Equals(argument) {
		return comp.Result, true
	}
	return nil, false
}
