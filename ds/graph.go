package ds

import (
	"fmt"
	"strings"

	"github.com/jqkard/fn/dict"
)

type (
	Edge      [2]Vertex
	Vertex    = string
	VertexSet = *Set[Vertex]
	EdgeSet   = *Set[Edge]
)

const edgeGlue string = "-"

type Graph struct {
	Vertices    []Vertex
	Edges       []Edge
	EdgesOf     map[Vertex][]Edge
	NeighborsOf map[Vertex]VertexSet
}

func NewEdge(edge string) Edge {
	parts := strings.Split(edge, edgeGlue)
	return Edge{parts[0], parts[1]}
}

func (e Edge) String() string {
	return fmt.Sprintf("%s%s%s", e[0], edgeGlue, e[1])
}

func (e Edge) Tuple() (Vertex, Vertex) {
	return e[0], e[1]
}

func NewGraph(vertices string, edgePairs string) *Graph {
	g := &Graph{}
	g.Vertices = strings.Fields(vertices)
	g.Edges = make([]Edge, 0)
	g.EdgesOf = make(map[Vertex][]Edge)
	g.NeighborsOf = make(map[Vertex]VertexSet)
	for _, edgePair := range strings.Fields(edgePairs) {
		edge := NewEdge(edgePair)
		v1, v2 := edge.Tuple()
		g.Edges = append(g.Edges, edge)
		g.EdgesOf[v1] = append(g.EdgesOf[v1], edge)
		g.EdgesOf[v2] = append(g.EdgesOf[v2], edge)
		dict.SetDefault(g.NeighborsOf, v1, NewSet[Vertex]())
		dict.SetDefault(g.NeighborsOf, v2, NewSet[Vertex]())
		g.NeighborsOf[v1].Add(v2)
		g.NeighborsOf[v2].Add(v1)
	}
	return g
}

func (g Graph) Neighbors(vertex Vertex) []Vertex {
	neighbors, ok := g.NeighborsOf[vertex]
	if !ok {
		return []Vertex{}
	}
	return neighbors.Items()
}
