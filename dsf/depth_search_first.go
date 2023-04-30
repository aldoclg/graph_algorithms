package dsf

import (
	"fmt"

	"github.com/aldoclg/graphalgorithm/stack"
	"github.com/aldoclg/graphalgorithm/vertex"
)

func Traverse(root *vertex.Vertex) {
	queue := stack.NewStack[vertex.Vertex]()
	root.Visited = true
	queue.Push(*root)

	for queue.IsNotEmpty() {
		actual := queue.Pop()
		fmt.Printf("[DSF] Actual vertex %v\n", actual)
		for _, v := range actual.GetAdjacency() {
			if v.Visited == false {
				v.Visited = true
				queue.Push(*v)
			}
		}
	}
}

func TraverseRecursive(root *vertex.Vertex) {
	fmt.Printf("[DSF RER] Actual vertex %v\n", *root)
	root.Visited = true
	recursiveDFS(root.GetAdjacency())
}

func recursiveDFS(list []*vertex.Vertex) {
	for _, v := range list {
		if !v.Visited {
			v.Visited = true
			fmt.Printf("[DSF RER] Actual vertex %v\n", *v)
			recursiveDFS(v.GetAdjacency())
		}
	}
}
