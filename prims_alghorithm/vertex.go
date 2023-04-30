package primsalghorithm

import "fmt"

type Vertex struct {
	name          string
	adjacencyEdge []*Edge
}

func NewVertex(name string) *Vertex {
	return &Vertex{name: name, adjacencyEdge: make([]*Edge, 0)}
}

func (v *Vertex) AddNeighbor(vertex *Edge) {
	v.adjacencyEdge = append(v.adjacencyEdge, vertex)
}

func (v *Vertex) GetAdjacency() []*Edge {
	return v.adjacencyEdge
}

func (v Vertex) String() string {
	return fmt.Sprintf("%s", v.name)
}
