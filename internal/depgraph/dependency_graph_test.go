package depgraph

import (
	"testing"
)

func TestDependencyGraphAddRemoveNodes(t *testing.T) {
	g := New()

	if got, expected := g.Len(), 0; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}

	g.AddNode("foo")
	if !g.HasNode("foo") {
		t.Fatalf("expected graph to contain foo")
	}
	if got, expected := g.Len(), 1; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}

	g.AddNode("bar")
	if !g.HasNode("bar") {
		t.Fatalf("expected graph to contain bar")
	}
	if got, expected := g.Len(), 2; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}

	// Add node that's already present.
	g.AddNode("bar")
	if got, expected := g.Len(), 2; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}

	if g.HasNode("baz") {
		t.Fatalf("expected graph not to contain baz")
	}

	g.RemoveNode("bar")
	if g.HasNode("bar") {
		t.Fatalf("expected graph not to contain bar")
	}
	if got, expected := g.Len(), 1; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}

	// Remove node that's not present.
	g.RemoveNode("bar")
	if got, expected := g.Len(), 1; got != expected {
		t.Fatalf("incorrect length. Expected: %d, got: %d", expected, got)
	}
}
