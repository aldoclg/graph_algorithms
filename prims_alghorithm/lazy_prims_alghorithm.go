package primsalghorithm

import "fmt"

type void struct{}

type LazyPrimsAlghorithm struct {
	unvisited           map[*Vertex]*void
	minimalSpanningTree []*Edge
	heap                PriorityQueue[*Edge]
	fullCost            float64
}

func NewLazyAlghorithm(vertexes []*Vertex) *LazyPrimsAlghorithm {
	unvisited := make(map[*Vertex]*void)
	empty := new(void)

	for _, v := range vertexes {
		unvisited[v] = empty
	}
	return &LazyPrimsAlghorithm{unvisited: unvisited,
		minimalSpanningTree: make([]*Edge, 0),
		heap:                NewPriorityQueue[*Edge](),
	}
}

func (this *LazyPrimsAlghorithm) Run(vertex *Vertex) {

	delete(this.unvisited, vertex)
	for _, e := range vertex.adjacencyEdge {
		if _, ok := this.unvisited[e.targetVertex]; ok {
			this.heap.Offer(e)
		}
	}

	minimalEdge := this.heap.Poll()
	target := minimalEdge.GetTargetVertex()

	if _, ok := this.unvisited[target]; !ok {
		return
	}
	this.minimalSpanningTree = append(this.minimalSpanningTree, minimalEdge)
	this.fullCost += minimalEdge.weight
	this.Run(target)
}

func (this *LazyPrimsAlghorithm) Print() {
	fmt.Println("Cost", this.fullCost)
	for _, e := range this.minimalSpanningTree {
		fmt.Printf("%s - %s\n", e.startVertex, e.targetVertex)
	}
}
