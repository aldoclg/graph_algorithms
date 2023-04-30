package cycledetection

import (
	"fmt"

	"github.com/aldoclg/graph_algorithms/vertex"
)

type CycleDetection struct {
}

func (cd *CycleDetection) Run(vertexes []*vertex.Vertex) {
	for _, v := range vertexes {
		if !v.Visited {
			cd.dfs(v)
		}
	}
}

func (cd *CycleDetection) dfs(vertex *vertex.Vertex) {

	vertex.BeingVisited = true
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in", r)
		}
	}()
	for _, u := range vertex.GetAdjacency() {
		fmt.Println("u", u)
		if u.BeingVisited {
			panic(fmt.Sprintf("[DAG Exception] Cycle detected at %v", *u))
		}

		if !u.Visited {
			u.Visited = true
			cd.dfs(u)
		}
	}

	vertex.BeingVisited = false
	vertex.Visited = true
}
