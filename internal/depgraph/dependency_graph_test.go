package depgraph

import (
	"reflect"
	"testing"
)

func TestDependencyGraphAddAndRemoveNodes(t *testing.T) {
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

func TestDependencyGraphDirectDependenciesAndDepdendents(t *testing.T) {
	g := New()

	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d")

	if err := g.AddDependency("a", "d"); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := g.AddDependency("a", "b"); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := g.AddDependency("b", "c"); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := g.AddDependency("d", "b"); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	got, err := g.DirectDependenciesOf("a")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"d", "b"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependencies. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependenciesOf("b")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"c"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependencies. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependenciesOf("c")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependencies. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependenciesOf("d")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"b"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependencies. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependentsOf("a")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependents. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependentsOf("b")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"a", "d"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependents. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependentsOf("c")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"b"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependents. Expected: %v, got: %v", expected, got)
	}

	got, err = g.DirectDependentsOf("d")
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if expected := []string{"a"}; !reflect.DeepEqual(got, expected) {
		t.Fatalf("incorrect direct dependents. Expected: %v, got: %v", expected, got)
	}
}
