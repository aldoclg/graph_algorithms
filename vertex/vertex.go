package vertex

import "fmt"

type Vertex struct {
	name         string
	Visited      bool
	BeingVisited bool
	adjacency    []*Vertex
}

func New(name string) Vertex {
	return Vertex{name: name}
}

func (v *Vertex) AddNeighbor(vertex *Vertex) {
	fmt.Printf("append %s value %v \n", v.name, vertex.name)
	v.adjacency = append(v.adjacency, vertex)
}

func (v *Vertex) GetAdjacency() []*Vertex {
	return v.adjacency
}

func (v *Vertex) ShowNeighbors() {
	for _, e := range v.adjacency {
		fmt.Println(e.name)
	}
}

func (v Vertex) String() string {
	o := fmt.Sprintf("%s %v %v", v.name, v.Visited, v.BeingVisited)
	for _, e := range v.adjacency {
		o = o + fmt.Sprintf(" %v", e.name)
	}
	return o
}
