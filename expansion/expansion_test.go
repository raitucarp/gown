package expansion_test

import (
	"context"
	"strings"
	"testing"

	"github.com/raitucarp/gown"
	"github.com/raitucarp/gown/expansion"
)

func TestExpansionComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	// 1. BFS expansion with relations filter
	treeBFS, err := expansion.Expand(res, "eat",
		expansion.WithMaxDepth(2),
		expansion.WithMaxNodes(20),
		expansion.WithRelations(string(gown.SynsetRelationTypeHypernym)),
		expansion.WithStrategy(expansion.StrategyBFS),
		expansion.WithPOS(gown.VerbPos))
	if err != nil {
		t.Fatalf("Expand BFS error: %v", err)
	}
	if treeBFS.TotalNodes <= 1 {
		t.Errorf("Expected tree to expand beyond root, got %d nodes", treeBFS.TotalNodes)
	}

	// 2. DFS expansion
	treeDFS, err := expansion.Expand(res, "eat",
		expansion.WithMaxDepth(2),
		expansion.WithMaxNodes(15),
		expansion.WithStrategy(expansion.StrategyDFS),
		expansion.WithPOS(gown.VerbPos))
	if err != nil {
		t.Fatalf("Expand DFS error: %v", err)
	}
	if treeDFS.TotalNodes <= 1 {
		t.Errorf("Expected DFS tree to expand beyond root, got %d nodes", treeDFS.TotalNodes)
	}

	// 3. Missing word expansion (returns single root node)
	missingTree, err := expansion.Expand(res, "nonexistentxyz123")
	if err != nil {
		t.Fatalf("Expand missing word returned error: %v", err)
	}
	if missingTree.TotalNodes != 1 || missingTree.Root.Word != "nonexistentxyz123" {
		t.Errorf("Unexpected missing word tree: %+v", missingTree)
	}

	// 4. Context cancellation
	ctxCanceled, cancel := context.WithCancel(context.Background())
	cancel()
	treeCanceled, errCancel := expansion.Expand(res, "dog",
		expansion.WithContext(ctxCanceled),
		expansion.WithMaxDepth(3))
	if errCancel == nil {
		t.Errorf("Expected context cancellation error")
	}
	_ = treeCanceled

	// 5. Render methods on empty/nil trees
	var nilTree *expansion.Tree
	if nilTree.Render() != "<empty tree>" {
		t.Errorf("nilTree.Render() expected '<empty tree>', got '%s'", nilTree.Render())
	}
	emptyTree := &expansion.Tree{}
	if emptyTree.Render() != "<empty tree>" {
		t.Errorf("emptyTree.Render() expected '<empty tree>', got '%s'", emptyTree.Render())
	}

	// 6. Valid rendered tree check
	rendered := treeBFS.Render()
	if !strings.Contains(rendered, "eat") {
		t.Errorf("Rendered tree missing root word 'eat':\n%s", rendered)
	}
}

func TestDefinitionExpansionComprehensive(t *testing.T) {
	res, err := gown.ReadLexicalResource()
	if err != nil {
		t.Fatalf("Failed to read lexical resource: %v", err)
	}

	tree, err := expansion.ExpandDefinition(res, "dog",
		expansion.WithMaxDepth(2),
		expansion.WithMaxNodes(20),
		expansion.WithPOS(gown.NounPos))
	if err != nil {
		t.Fatalf("ExpandDefinition error: %v", err)
	}

	foundToken := false
	for _, child := range tree.Root.Children {
		for _, grandChild := range child.Children {
			if grandChild.Type == expansion.NodeToken {
				foundToken = true
				break
			}
		}
	}
	if !foundToken {
		t.Errorf("Expected definition token node in expanded definition tree")
	}
}
