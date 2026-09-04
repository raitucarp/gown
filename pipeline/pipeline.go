package pipeline

import (
	"strings"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/discourse"
	"github.com/raitucarp/gown/phonology"
	"github.com/raitucarp/gown/pragmatics"
	"github.com/raitucarp/gown/semantics"
	"github.com/raitucarp/gown/semiotics"
	"github.com/raitucarp/gown/syntax"
	"github.com/raitucarp/gown/text"
)

// Pipeline orchestrates analysis across all linguistic disciplines:
// Orthography -> Morphology -> Syntax -> Semantics -> Pragmatics -> Discourse -> Semiotics.
type Pipeline struct {
	res *gown.LexicalResource
}

// NewPipeline constructs a linguistic processing pipeline backed by WordNet.
func NewPipeline(res *gown.LexicalResource) *Pipeline {
	return &Pipeline{res: res}
}

// Process runs full multi-layer computational linguistic analysis on raw text.
func (p *Pipeline) Process(documentText string) *LinguisticDocument {
	rawSentences := text.SentenceSegment(documentText)
	doc := &LinguisticDocument{
		RawText:         documentText,
		SemioticSquares: make(map[string]semiotics.SemioticSquare),
	}

	for sIdx, sText := range rawSentences {
		sentAnalysis := SentenceAnalysis{
			ID:             sIdx + 1,
			Raw:            sText,
			SpeechAct:      pragmatics.ClassifySpeechAct(sText),
			Deixis:         pragmatics.IdentifyDeixis(sText),
			Presupposition: pragmatics.ExtractPresuppositions(sText),
			Implicatures:   pragmatics.DetectScalarImplicatures(sText),
			Politeness:     pragmatics.AnalyzePoliteness(sText),
		}

		tokens := text.Tokenize(sText)
		for _, tok := range tokens {
			lemmas := p.res.MorphyAll(tok)
			lemma := tok
			pos := "unknown"
			if len(lemmas) > 0 {
				lemma = lemmas[0]
				entries := p.res.Lookup(lemma)
				if len(entries) > 0 {
					pos = string(entries[0].Lemma.PartOfSpeech)
				}
			}

			wAnalysis := WordAnalysis{
				Surface:     tok,
				Lemma:       lemma,
				POS:         pos,
				CVPattern:   gown.OrthographicCV(tok),
				Syllables:   phonology.CountSyllables(tok),
				SignMode:    semiotics.ClassifySignMode(tok),
				Connotation: semiotics.AnalyzeConnotation(p.res, tok),
			}
			sentAnalysis.Words = append(sentAnalysis.Words, wAnalysis)

			// Generate semiotic square for key content words
			if !text.IsStopword(tok) && len(tok) > 3 && len(doc.SemioticSquares) < 5 {
				if _, exists := doc.SemioticSquares[lemma]; !exists {
					doc.SemioticSquares[lemma] = semiotics.GenerateSemioticSquare(p.res, lemma)
				}
			}
		}

		// Syntactic parsing heuristic: S -> NP VP
		if len(tokens) >= 3 {
			// e.g. "The dog chased the cat"
			// Subject: tokens[0..1], Verb: tokens[2], Object: tokens[3..]
			splitIdx := 1
			if len(tokens) > 3 && (tokens[0] == "the" || tokens[0] == "a") {
				splitIdx = 2
			}

			npSubj := syntax.NewNode("NP")
			for i := 0; i < splitIdx; i++ {
				npSubj.AddChild(syntax.NewLeaf("W", tokens[i]))
			}

			vp := syntax.NewNode("VP")
			vp.AddChild(syntax.NewLeaf("V", tokens[splitIdx]))
			if splitIdx+1 < len(tokens) {
				npObj := syntax.NewNode("NP")
				for i := splitIdx + 1; i < len(tokens); i++ {
					npObj.AddChild(syntax.NewLeaf("W", tokens[i]))
				}
				vp.AddChild(npObj)
			}

			sTree := syntax.NewNode("S", npSubj, vp)
			sentAnalysis.SyntaxTree = sTree
			sentAnalysis.Relations = syntax.ExtractRelations(sTree)

			// Semantic thematic roles
			subjText := npSubj.Yield()
			verbText := tokens[splitIdx]
			objText := ""
			if len(tokens) > splitIdx+1 {
				objText = strings.Join(tokens[splitIdx+1:], " ")
			}
			sentAnalysis.Roles = semantics.AssignRoles(verbText, subjText, objText, "")
		}

		doc.Sentences = append(doc.Sentences, sentAnalysis)
	}

	// Discourse layer
	doc.EDUs = discourse.SegmentEDUs(documentText)
	doc.RSTTree = discourse.BuildRSTTree(doc.EDUs)
	doc.CoreferenceChains = discourse.TrackCoreference(documentText)
	doc.TopicTracking, _ = discourse.TrackTopics(documentText)
	doc.ThemeProgression = discourse.AnalyzeThemeProgression(documentText)

	return doc
}
