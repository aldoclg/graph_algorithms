package disjointset

import "fmt"

type Edge struct {
	startVertex  *Vertex
	targetVertex *Vertex
	weight       float64
}

type SortableEdges []*Edge

func NewEdge(startVertex, targetVertex *Vertex, weight float64) *Edge {
	return &Edge{startVertex, targetVertex, weight}
}

func (this *Edge) GetStartEdge() *Vertex {
	return this.startVertex
}

func (this *Edge) GetTargetEdge() *Vertex {
	return this.targetVertex
}

func (this *Edge) GetWeight() float64 {
	return this.weight
}

func (e Edge) String() string {
	return fmt.Sprintf("%s ==%.1f==>> %s", e.GetStartEdge(), e.weight, e.GetTargetEdge())
}

func (s SortableEdges) Len() int {
	return len(s)
}

func (s SortableEdges) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortableEdges) Less(i, j int) bool {
	return s[i].weight < s[j].weight
}
