package semiotics_test

import (
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/semiotics"
)

func TestSaussureanSign(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	sign := semiotics.CreateSaussureanSign(res, "good")
	if sign.Signifier != "good" {
		t.Errorf("Expected signifier 'good', got '%s'", sign.Signifier)
	}
	if len(sign.Signified) == 0 {
		t.Errorf("Expected non-empty signified concept")
	}

	t.Logf("Saussurean Sign: %s", sign.String())
}

func TestPeirceanTriad(t *testing.T) {
	// Icon (onomatopoeia)
	icon := semiotics.ClassifySignMode("buzz")
	if icon != semiotics.ModeIcon {
		t.Errorf("Expected 'buzz' to be Icon, got %s", icon)
	}

	// Index (causal connection)
	index := semiotics.ClassifySignMode("smoke")
	if index != semiotics.ModeIndex {
		t.Errorf("Expected 'smoke' to be Index, got %s", index)
	}

	// Symbol (arbitrary linguistic sign)
	symbol := semiotics.ClassifySignMode("justice")
	if symbol != semiotics.ModeSymbol {
		t.Errorf("Expected 'justice' to be Symbol, got %s", symbol)
	}

	triad := semiotics.CreatePeirceanTriad("buzz", "sound of bee", "concept of insect noise")
	if triad.Mode != semiotics.ModeIcon {
		t.Errorf("Expected Triad mode to be Icon, got %s", triad.Mode)
	}
}

func TestConnotation(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	nobleAnalysis := semiotics.AnalyzeConnotation(res, "noble")
	if nobleAnalysis.Valence != semiotics.ValencePositive {
		t.Errorf("Expected 'noble' to have positive valence, got %s", nobleAnalysis.Valence)
	}

	terribleAnalysis := semiotics.AnalyzeConnotation(res, "terrible")
	if terribleAnalysis.Valence != semiotics.ValenceNegative {
		t.Errorf("Expected 'terrible' to have negative valence, got %s", terribleAnalysis.Valence)
	}
}

func TestSemioticSquare(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	square := semiotics.GenerateSemioticSquare(res, "good")
	if square.S1 != "good" {
		t.Errorf("Expected S1 to be 'good', got '%s'", square.S1)
	}
	if square.S2 != "bad" && square.S2 != "evil" {
		t.Logf("Note: S2 is '%s'", square.S2)
	}

	rendered := square.Render()
	t.Logf("Greimas Semiotic Square for 'good':\n%s", rendered)
}

func TestSemioticNetwork(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	net := semiotics.BuildSemioticNetwork(res, "dog", 2)
	if len(net.Nodes) == 0 {
		t.Errorf("Expected semiotic network nodes")
	}

	t.Logf("Semiotic network for 'dog': %d nodes, %d links", len(net.Nodes), len(net.Links))
}
