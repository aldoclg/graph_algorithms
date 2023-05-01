package topological_ordering

import (
	"github.com/aldoclg/graph_algorithms/stack"
	"github.com/aldoclg/graph_algorithms/vertex"
)

type TopologicalOrder struct {
	stack stack.Stack[vertex.Vertex]
}

func NewTopologicalOrder() *TopologicalOrder {
	stack := stack.NewStack[vertex.Vertex]()
	return &TopologicalOrder{stack: *stack}
}

func (to *TopologicalOrder) DFS(vertex *vertex.Vertex) {
	vertex.Visited = true

	for _, v := range vertex.GetAdjacency() {
		if !v.Visited {
			to.DFS(v)
		}
	}

	to.stack.Push(*vertex)
}

func (to *TopologicalOrder) GetStack() *stack.Stack[vertex.Vertex] {
	return &to.stack
}
