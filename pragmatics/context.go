package pragmatics

// PragmaticContext models the communicative situation in which an utterance occurs.
type PragmaticContext struct {
	Speaker      string            `json:"speaker"`
	Addressee    string            `json:"addressee"`
	Location     string            `json:"location"` // e.g. "here", "library", "London"
	Time         string            `json:"time"`     // e.g. "now", "2026-09-04"
	CommonGround map[string]bool   `json:"common_ground"` // propositions mutually accepted by participants
	Attributes   map[string]string `json:"attributes"`    // arbitrary extra context parameters
}

// NewContext creates a default communicative context.
func NewContext(speaker, addressee string) *PragmaticContext {
	return &PragmaticContext{
		Speaker:      speaker,
		Addressee:    addressee,
		Location:     "here",
		Time:         "now",
		CommonGround: make(map[string]bool),
		Attributes:   make(map[string]string),
	}
}

// AddPresupposition registers a proposition into the shared common ground.
func (c *PragmaticContext) AddPresupposition(proposition string) {
	c.CommonGround[proposition] = true
}

// Presupposed returns true if the proposition is part of the common ground.
func (c *PragmaticContext) Presupposed(proposition string) bool {
	return c.CommonGround[proposition]
}
