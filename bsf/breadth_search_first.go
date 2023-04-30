package bsf

import (
	"fmt"

	"github.com/aldoclg/graph_algorithms/queue"
	"github.com/aldoclg/graph_algorithms/vertex"
)

func Traverse(root vertex.Vertex) {
	queue := queue.NewQueue[vertex.Vertex]()
	root.Visited = true
	queue.Enqueue(root)

	for queue.IsNotEmpty() {
		actual := queue.Dequeue()
		fmt.Printf("[BSF] Actual vertex %v\n", actual)
		for _, v := range actual.GetAdjacency() {
			if v.Visited == false {
				v.Visited = true
				queue.Enqueue(*v)
			}
		}
	}
}
