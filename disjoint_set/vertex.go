package disjointset

import "fmt"

type Vertex struct {
	name      string
	adjacency []*Edge
	node      *Node //disjoint set
}

func NewVertex(name string) *Vertex {
	return &Vertex{name: name}
}

func (v *Vertex) AddNeighbor(edge *Edge) {
	v.adjacency = append(v.adjacency, edge)
}

func (v *Vertex) GetAdjacency() []*Edge {
	return v.adjacency
}

func (v *Vertex) ShowNeighbors() {
	for _, e := range v.adjacency {
		fmt.Println(e)
	}
}

func (v Vertex) String() string {
	return fmt.Sprintf("%s", v.name)
}

func (v *Vertex) SetNode(node *Node) {
	v.node = node
}

func (v *Vertex) GetNode() *Node {
	return v.node
}
