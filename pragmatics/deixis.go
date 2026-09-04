package pragmatics

import (
	"strings"
)

// DeixisType categorizes the deictic reference dimension.
type DeixisType string

const (
	DeixisPerson   DeixisType = "person"
	DeixisSpatial  DeixisType = "spatial"
	DeixisTemporal DeixisType = "temporal"
	DeixisDiscourse DeixisType = "discourse"
)

// DeicticExpression represents an identified deictic item in text.
type DeicticExpression struct {
	Word       string     `json:"word"`
	Type       DeixisType `json:"type"`
	Proximity  string     `json:"proximity,omitempty"` // "proximal", "distal", "speaker", "addressee"
	ResolvedTo string     `json:"resolved_to,omitempty"`
}

var personDeictics = map[string]string{
	"i": "speaker", "me": "speaker", "my": "speaker", "mine": "speaker", "myself": "speaker",
	"we": "speaker_group", "us": "speaker_group", "our": "speaker_group", "ours": "speaker_group",
	"you": "addressee", "your": "addressee", "yours": "addressee", "yourself": "addressee",
}

var spatialDeictics = map[string]string{
	"here": "proximal", "this": "proximal", "these": "proximal",
	"there": "distal", "that": "distal", "those": "distal",
	"yonder": "distal",
}

var temporalDeictics = map[string]string{
	"now": "present", "today": "present", "tonight": "present",
	"yesterday": "past", "recently": "past", "then": "remote",
	"tomorrow": "future", "soon": "future",
}

var discourseDeictics = map[string]string{
	"former": "anaphoric", "latter": "anaphoric", "above": "anaphoric",
	"below": "cataphoric", "following": "cataphoric",
}

// IdentifyDeixis scans an utterance and tags all deictic expressions.
func IdentifyDeixis(utterance string) []DeicticExpression {
	words := strings.Fields(utterance)
	var expressions []DeicticExpression

	for _, raw := range words {
		w := strings.ToLower(strings.Trim(raw, ".,!?\"'()"))
		if prox, ok := personDeictics[w]; ok {
			expressions = append(expressions, DeicticExpression{
				Word:      raw,
				Type:      DeixisPerson,
				Proximity: prox,
			})
		} else if prox, ok := spatialDeictics[w]; ok {
			expressions = append(expressions, DeicticExpression{
				Word:      raw,
				Type:      DeixisSpatial,
				Proximity: prox,
			})
		} else if prox, ok := temporalDeictics[w]; ok {
			expressions = append(expressions, DeicticExpression{
				Word:      raw,
				Type:      DeixisTemporal,
				Proximity: prox,
			})
		} else if prox, ok := discourseDeictics[w]; ok {
			expressions = append(expressions, DeicticExpression{
				Word:      raw,
				Type:      DeixisDiscourse,
				Proximity: prox,
			})
		}
	}

	return expressions
}

// ResolveDeixis resolves person, spatial, and temporal deictic expressions using the provided context.
func ResolveDeixis(utterance string, ctx *PragmaticContext) []DeicticExpression {
	deictics := IdentifyDeixis(utterance)
	if ctx == nil {
		return deictics
	}

	for i := range deictics {
		d := &deictics[i]
		switch d.Type {
		case DeixisPerson:
			if d.Proximity == "speaker" {
				d.ResolvedTo = ctx.Speaker
			} else if d.Proximity == "addressee" {
				d.ResolvedTo = ctx.Addressee
			}
		case DeixisSpatial:
			if d.Proximity == "proximal" {
				d.ResolvedTo = ctx.Location
			}
		case DeixisTemporal:
			if d.Proximity == "present" {
				d.ResolvedTo = ctx.Time
			}
		}
	}

	return deictics
}
