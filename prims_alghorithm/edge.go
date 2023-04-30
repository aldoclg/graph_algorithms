package primsalghorithm

type Edge struct {
	startVertex  *Vertex
	targetVertex *Vertex
	weight       float64
}

func NewEdge(startVertex, targetVertex *Vertex, weight float64) *Edge {
	return &Edge{startVertex, targetVertex, weight}
}

func (this *Edge) GetStartVertex() *Vertex {
	return this.startVertex
}

func (this *Edge) GetTargetVertex() *Vertex {
	return this.targetVertex
}

func (this *Edge) GetWeight() float64 {
	return this.weight
}

func (this *Edge) SetTargetVertex(vertex *Vertex) {
	this.targetVertex = vertex
}

func (this *Edge) Compare(that *Edge) bool {
	return this.weight < that.weight
}
