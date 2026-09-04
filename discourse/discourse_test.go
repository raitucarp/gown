package discourse_test

import (
	"strings"
	"testing"

	"github.com/raitucarp/gown/discourse"
)

func TestEDUSegmentationAndRSTComprehensive(t *testing.T) {
	// 1. Multi-EDU text with various rhetorical markers
	text := "The dog barked loudly, because it saw a stranger. However, the stranger was friendly. If you approach gently, then he will wag his tail."
	edus := discourse.SegmentEDUs(text)

	if len(edus) < 4 {
		t.Fatalf("Expected >= 4 EDUs, got %d", len(edus))
	}

	rstTree := discourse.BuildRSTTree(edus)
	if rstTree == nil {
		t.Fatalf("Expected non-nil RST tree")
	}

	rendered := rstTree.Render()
	if !strings.Contains(rendered, "barked") {
		t.Errorf("Rendered RST tree missing text: %s", rendered)
	}

	// 2. BuildRSTTree with empty EDUs
	if nilRST := discourse.BuildRSTTree(nil); nilRST != nil {
		t.Errorf("BuildRSTTree(nil) expected nil, got %+v", nilRST)
	}

	// 3. BuildRSTTree with 1 EDU
	singleEDU := []discourse.EDU{{ID: 1, Text: "Just one clause"}}
	singleRST := discourse.BuildRSTTree(singleEDU)
	if singleRST == nil || singleRST.Text != "Just one clause" || len(singleRST.Children) != 0 {
		t.Errorf("Unexpected single RST node: %+v", singleRST)
	}

	// 4. Render on nil RSTNode
	var nilNode *discourse.RSTNode
	if nilNode.Render() != "<empty rst>" {
		t.Errorf("nilNode.Render() expected '<empty rst>', got '%s'", nilNode.Render())
	}

	// 5. Empty segmentation
	if emptyEDUs := discourse.SegmentEDUs(""); len(emptyEDUs) != 0 {
		t.Errorf("Expected 0 EDUs for empty text, got %d", len(emptyEDUs))
	}
}

func TestCoreferenceTrackingComprehensive(t *testing.T) {
	text := "John bought a new car. He drove it to work. John loved the vehicle."
	chains := discourse.TrackCoreference(text)

	if len(chains) == 0 {
		t.Fatalf("Expected coreference chains to be extracted")
	}

	// Empty text
	emptyChains := discourse.TrackCoreference("")
	if len(emptyChains) != 0 {
		t.Errorf("Expected 0 chains for empty text, got %d", len(emptyChains))
	}
}

func TestTopicTrackingAndThemeProgressionComprehensive(t *testing.T) {
	text := "The dog chased the cat. The cat climbed a high tree. A bird was singing on the branch."

	// 1. Topic tracking
	topics, transitions := discourse.TrackTopics(text)
	if len(topics) != 3 {
		t.Errorf("Expected 3 sentence topics, got %d", len(topics))
	}
	if len(transitions) != 2 {
		t.Errorf("Expected 2 topic transitions, got %d", len(transitions))
	}

	// Empty topic tracking
	emptyTopics, emptyTransitions := discourse.TrackTopics("")
	if len(emptyTopics) != 0 || len(emptyTransitions) != 0 {
		t.Errorf("Expected 0 topics/transitions for empty text")
	}

	// 2. Daneš Theme progression
	prog := discourse.AnalyzeThemeProgression(text)
	if len(prog) != 3 {
		t.Errorf("Expected 3 theme progression steps, got %d", len(prog))
	}

	// Empty theme progression
	emptyProg := discourse.AnalyzeThemeProgression("")
	if len(emptyProg) != 0 {
		t.Errorf("Expected 0 progression steps for empty text")
	}
}

func TestDiscourseGraph(t *testing.T) {
	dg := discourse.NewDiscourseGraph()
	dg.AddNode("s1", discourse.NodeSentence, "Sentence 1")
	dg.AddNode("e1", discourse.NodeEntity, "Entity Dog")
	dg.AddEdge("s1", "e1", "mentions_entity")

	if len(dg.Nodes) != 2 {
		t.Errorf("Expected 2 discourse nodes, got %d", len(dg.Nodes))
	}
	if len(dg.Edges) != 1 {
		t.Errorf("Expected 1 discourse edge, got %d", len(dg.Edges))
	}
	if dg.Edges[0].Relation != "mentions_entity" {
		t.Errorf("Unexpected relation: %s", dg.Edges[0].Relation)
	}
}
