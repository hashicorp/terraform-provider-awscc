package depgraph

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
