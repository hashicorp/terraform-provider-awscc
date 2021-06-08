package depgraph

import (
	"fmt"
)

// Graph implements a simple dependency graph.
type Graph struct {
	nodes         map[string]struct{}
	outgoingEdges map[string]map[string]struct{}
	incomingEdges map[string]map[string]struct{}
}

// New returns a new empty dependency graph.
func New() *Graph {
	return &Graph{
		nodes:         make(map[string]struct{}),
		outgoingEdges: make(map[string]map[string]struct{}),
		incomingEdges: make(map[string]map[string]struct{}),
	}
}

// Len returns the number of nodes in the graph.
func (g *Graph) Len() int {
	return len(g.nodes)
}

// AddNode adds the specified string to the graph.
func (g *Graph) AddNode(s string) {
	if _, ok := g.nodes[s]; !ok {
		g.nodes[s] = struct{}{}
		g.outgoingEdges[s] = make(map[string]struct{})
		g.incomingEdges[s] = make(map[string]struct{})
	}
}

// RemoveNode removes the specified string from the graph if it is present.
func (g *Graph) RemoveNode(s string) {
	if _, ok := g.nodes[s]; ok {
		for _, edges := range g.outgoingEdges {
			delete(edges, s)
		}

		for _, edges := range g.incomingEdges {
			delete(edges, s)
		}

		delete(g.nodes, s)
		delete(g.outgoingEdges, s)
		delete(g.incomingEdges, s)
	}
}

// HasNode returns whether the specified string is in the graph.
func (g *Graph) HasNode(s string) bool {
	_, ok := g.nodes[s]

	return ok
}

// AddDependency adds a dependency between two nodes.
// If either node doesn't exist an error is returned.
func (g *Graph) AddDependency(from, to string) error {
	if !g.HasNode(from) {
		return nonExistentNodeError(from)
	}
	if !g.HasNode(to) {
		return nonExistentNodeError(to)
	}

	if _, ok := g.outgoingEdges[from][to]; !ok {
		g.outgoingEdges[from][to] = struct{}{}
	}

	if _, ok := g.incomingEdges[to][from]; !ok {
		g.incomingEdges[to][from] = struct{}{}
	}

	return nil
}

// RemoveDependency removes a dependency between two nodes.
// If either node doesn't exist no error is returned.
func (g *Graph) RemoveDependency(from, to string) {
	if g.HasNode(from) {
		delete(g.outgoingEdges[from], to)
	}

	if g.HasNode(to) {
		delete(g.incomingEdges[to], from)
	}
}

// DirectDependenciesOf returns the nodes that are the direct dependencies of the specified node.
// Returns an error if the specified node doesn't exist.
func (g *Graph) DirectDependenciesOf(s string) ([]string, error) {
	if !g.HasNode(s) {
		return nil, nonExistentNodeError(s)
	}

	deps := make([]string, 0)

	for k := range g.outgoingEdges[s] {
		deps = append(deps, k)
	}

	return deps, nil
}

// DirectDependentsOf returns the nodes that directly depend on the specified node.
// Returns an error if the specified node doesn't exist.
func (g *Graph) DirectDependentsOf(s string) ([]string, error) {
	if !g.HasNode(s) {
		return nil, nonExistentNodeError(s)
	}

	deps := make([]string, 0)

	for k := range g.incomingEdges[s] {
		deps = append(deps, k)
	}

	return deps, nil
}

func nonExistentNodeError(s string) error {
	return fmt.Errorf("node does not exist: %s", s)
}
