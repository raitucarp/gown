package generative

import (
	"fmt"
	"reflect"
)

// FeatureStructure represents an attribute-value matrix (AVM) for formal syntax.
type FeatureStructure map[string]any

// NewFeatureStructure constructs an empty feature structure.
func NewFeatureStructure() FeatureStructure {
	return make(FeatureStructure)
}

// Set stores a feature attribute-value pair.
func (fs FeatureStructure) Set(key string, val any) FeatureStructure {
	fs[key] = val
	return fs
}

// Get retrieves a feature value.
func (fs FeatureStructure) Get(key string) any {
	return fs[key]
}

// Unify computes the greatest lower bound (unification) of two feature structures.
// Returns the unified structure and a boolean indicating whether unification succeeded.
func Unify(fs1, fs2 FeatureStructure) (FeatureStructure, bool) {
	if fs1 == nil && fs2 == nil {
		return nil, true
	}
	if fs1 == nil {
		return cloneFS(fs2), true
	}
	if fs2 == nil {
		return cloneFS(fs1), true
	}

	result := cloneFS(fs1)

	for k, v2 := range fs2 {
		v1, exists := result[k]
		if !exists {
			result[k] = cloneVal(v2)
			continue
		}

		// Both have the feature 'k', check compatibility
		sub1, isFS1 := v1.(FeatureStructure)
		sub2, isFS2 := v2.(FeatureStructure)

		if isFS1 && isFS2 {
			unifiedSub, ok := Unify(sub1, sub2)
			if !ok {
				return nil, false
			}
			result[k] = unifiedSub
		} else if isFS1 != isFS2 {
			// Type mismatch
			return nil, false
		} else {
			// Atomic values
			if !reflect.DeepEqual(v1, v2) {
				return nil, false
			}
		}
	}

	return result, true
}

func cloneFS(fs FeatureStructure) FeatureStructure {
	if fs == nil {
		return nil
	}
	res := make(FeatureStructure, len(fs))
	for k, v := range fs {
		res[k] = cloneVal(v)
	}
	return res
}

func cloneVal(v any) any {
	if sub, ok := v.(FeatureStructure); ok {
		return cloneFS(sub)
	}
	return v
}

// String formats the feature structure as an attribute-value matrix.
func (fs FeatureStructure) String() string {
	return fmt.Sprintf("%v", map[string]any(fs))
}
