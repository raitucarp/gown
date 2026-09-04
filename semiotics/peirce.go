package semiotics

import (
	"strings"
)

// SignMode represents Charles Sanders Peirce's trichotomy of signs.
type SignMode string

const (
	ModeIcon   SignMode = "icon"   // Signifies by virtue of resemblance/imitation
	ModeIndex  SignMode = "index"  // Signifies by virtue of physical/causal/deictic contiguity
	ModeSymbol SignMode = "symbol" // Signifies by virtue of arbitrary social convention
)

// PeirceanTriad represents Peirce's triadic model of semiosis:
// Representamen (sign vehicle), Object (referent in reality), Interpretant (resulting mental effect).
type PeirceanTriad struct {
	Representamen string   `json:"representamen"`
	Object        string   `json:"object"`
	Interpretant  string   `json:"interpretant"`
	Mode          SignMode `json:"mode"`
}

var onomatopoeiaWords = map[string]bool{
	"buzz": true, "hiss": true, "meow": true, "bark": true, "splash": true, "pop": true,
	"bang": true, "boom": true, "click": true, "chirp": true, "whisper": true, "tick": true,
	"clatter": true, "crash": true, "rustle": true, "snort": true, "purr": true, "moo": true,
}

var indexicalWords = map[string]bool{
	"smoke": true, "footprint": true, "shadow": true, "symptom": true, "echo": true,
	"knock": true, "here": true, "there": true, "this": true, "that": true, "now": true,
}

// ClassifySignMode determines whether a word acts primarily as an Icon, Index, or Symbol.
func ClassifySignMode(word string) SignMode {
	w := strings.ToLower(strings.TrimSpace(word))
	if onomatopoeiaWords[w] {
		return ModeIcon
	}
	if indexicalWords[w] {
		return ModeIndex
	}
	return ModeSymbol
}

// CreatePeirceanTriad builds a triadic sign representation.
func CreatePeirceanTriad(representamen, object, interpretant string) PeirceanTriad {
	mode := ClassifySignMode(representamen)
	return PeirceanTriad{
		Representamen: representamen,
		Object:        object,
		Interpretant:  interpretant,
		Mode:          mode,
	}
}
