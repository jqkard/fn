package ds

import (
	"fmt"
	"strings"
)

type (
	Edge      [2]Vertex
	Vertex    = string
	VertexSet = *Set[Vertex]
	EdgeSet   = *Set[Edge]
)

type Graph struct {
	Vertices    []Vertex
	Edges       []Edge
	IndexOf     map[Vertex]int
	EdgesOf     map[Vertex][]Edge
	NeighborsOf map[Vertex]VertexSet
}

func NewEdge(edge string) Edge {
	parts := strings.Split(edge, "-")
	return Edge{parts[0], parts[1]}
}

func (e Edge) String() string {
	return fmt.Sprintf("%s-%s", e[0], e[1])
}

func (e Edge) Tuple() (Vertex, Vertex) {
	return e[0], e[1]
}

func NewGraph(vertices string, edgePairs string) *Graph {
	g := &Graph{}
	g.Vertices = strings.Fields(vertices)
	g.IndexOf = make(map[Vertex]int)
	for i, vertex := range g.Vertices {
		g.IndexOf[vertex] = i
	}
	g.Edges = make([]Edge, 0)
	g.EdgesOf = make(map[Vertex][]Edge)
	g.NeighborsOf = make(map[Vertex]VertexSet)
	for _, edgePair := range strings.Fields(edgePairs) {
		edge := NewEdge(edgePair)
		v1, v2 := edge.Tuple()
		g.Edges = append(g.Edges, edge)
		g.EdgesOf[v1] = append(g.EdgesOf[v1], edge)
		g.EdgesOf[v2] = append(g.EdgesOf[v2], edge)
		if _, ok := g.NeighborsOf[v1]; !ok {
			g.NeighborsOf[v1] = NewSet[Vertex]()
		}
		if _, ok := g.NeighborsOf[v2]; !ok {
			g.NeighborsOf[v2] = NewSet[Vertex]()
		}
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

func (g Graph) ActiveNeighbors(vertex Vertex, activeEdges EdgeSet) []Vertex {
	neighbors := NewSet[Vertex]()
	for _, edge := range g.EdgesOf[vertex] {
		if activeEdges != nil && !activeEdges.Contains(edge) {
			continue
		}
		neighbors.Add(edge[0])
		neighbors.Add(edge[1])
	}
	neighbors.Delete(vertex)
	return neighbors.Items()
}

func (g Graph) IsClique(vertices []Vertex) bool {
	vertexSet := SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := SetFrom(g.Neighbors(vertex))
		adjacent.Add(vertex)
		if !vertexSet.Difference(adjacent).IsEmpty() {
			return false
		}
	}
	return true
}

func (g Graph) IsIndependentSet(vertices []Vertex) bool {
	vertexSet := SetFrom(vertices)
	for _, vertex := range vertices {
		adjacent := SetFrom(g.Neighbors(vertex))
		if vertexSet.Intersection(adjacent).Len() != 0 {
			return false
		}
	}
	return true
}

func (g Graph) BFSTraversal(start Vertex, activeEdges EdgeSet) []Vertex {
	q := NewQueue[Vertex](0)
	q.Enqueue(start)
	visited := NewSet[Vertex]()
	for !q.IsEmpty() {
		current := q.Dequeue()
		if visited.Contains(current) {
			continue
		}
		visited.Add(current)
		for _, neighbor := range g.ActiveNeighbors(current, activeEdges) {
			if visited.Contains(neighbor) {
				continue
			}
			q.Enqueue(neighbor)
		}
	}
	return visited.Items()
}
