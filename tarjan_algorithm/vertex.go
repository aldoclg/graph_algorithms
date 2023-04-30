package tarjanalgorithm

import "fmt"

type Vertex struct {
	name        string
	Visited     bool
	OnStack     bool
	adjacency   []*Vertex
	lowLink     int
	index       int
	componentId int
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

func (v *Vertex) GetLowLink() int {
	return v.lowLink
}

func (v *Vertex) SetLowLink(value int) {
	v.lowLink = value
}

func (v *Vertex) GetIndex() int {
	return v.index
}

func (v *Vertex) SetIndex(value int) {
	v.index = value
}

func (v *Vertex) GetComponentId() int {
	return v.componentId
}

func (v *Vertex) SetComponentId(value int) {
	v.componentId = value
}

func (v Vertex) String() string {
	o := fmt.Sprintf("%s %v %v %d", v.name, v.Visited, v.OnStack, v.componentId)
	for _, e := range v.adjacency {
		o = o + fmt.Sprintf(" %v", e.name)
	}
	return o
}
