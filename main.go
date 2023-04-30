package main

import (
	"fmt"
	"sync"

	cycledetection "github.com/aldoclg/graph_algorithms/cycle_detection"
	disjointset "github.com/aldoclg/graph_algorithms/disjoint_set"
	"github.com/aldoclg/graph_algorithms/dsf"
	kruskalalgorithm "github.com/aldoclg/graph_algorithms/kruskal_algorithm"
	primsalghorithm "github.com/aldoclg/graph_algorithms/prims_alghorithm"
	"github.com/aldoclg/graph_algorithms/topological_ordering"
	"github.com/aldoclg/graph_algorithms/vertex"
)

var wg sync.WaitGroup

func main() {

	v1 := createGraph()
	v2 := createGraph()

	wg.Add(2)

	defer fmt.Println("Exit")

	go func() {
		dsf.Traverse(v2)
		defer fmt.Println("Exit dsf")
		defer wg.Done()
	}()

	go func() {
		dsf.TraverseRecursive(v1)
		defer fmt.Println("Exit dsf recursive")
		defer wg.Done()
	}()

	wg.Wait()

	var matrix = [][]int{
		{1, 1, 1, 1, 1, 1},
		{2, 1, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 1},
		{0, 1, 0, 1, 0, 0},
		{0, 1, 0, 1, 1, 0},
		{0, 0, 0, 1, 1, 0},
	}

	mazeSolver := dsf.NewMazeSolver(matrix, 1, 0)

	mazeSolver.FindDestination()

	ordering := topological_ordering.NewTopologicalOrder()
	list := createGraphTopologicalOrder()

	for _, v := range list {
		if !v.Visited {
			ordering.DFS(v)
		}
	}

	stack := ordering.GetStack()

	fmt.Println("Graph size", len(list))
	fmt.Println("Stack size", stack.Size())

	for stack.IsNotEmpty() {
		v := stack.Pop()
		fmt.Printf("%v\n", v)
	}

	cycleDetection := new(cycledetection.CycleDetection)
	vertexes := createCycledGraph()

	cycleDetection.Run(vertexes)

	v, e := createGraphList3()

	mst := kruskalalgorithm.Compute(v, e)

	for _, edge := range mst {
		fmt.Println(edge)
	}

	graphs := createGraphList4()
	lazyPrimsAlghorithm := primsalghorithm.NewLazyAlghorithm(graphs)

	lazyPrimsAlghorithm.Run(graphs[0])
	lazyPrimsAlghorithm.Print()

}

func createGraph() *vertex.Vertex {
	a := vertex.New("a")
	b := vertex.New("b")
	c := vertex.New("c")
	d := vertex.New("d")
	e := vertex.New("e")
	f := vertex.New("f")
	g := vertex.New("g")
	h := vertex.New("h")

	a.AddNeighbor(&b)
	a.AddNeighbor(&f)
	a.AddNeighbor(&g)

	b.AddNeighbor(&a)
	b.AddNeighbor(&c)
	b.AddNeighbor(&d)

	c.AddNeighbor(&b)

	d.AddNeighbor(&b)
	d.AddNeighbor(&e)

	g.AddNeighbor(&a)
	g.AddNeighbor(&h)

	h.AddNeighbor(&g)

	a.ShowNeighbors()

	return &a
}

func createGraphTopologicalOrder() []*vertex.Vertex {
	list := make([]*vertex.Vertex, 6)

	for i := 0; i <= 5; i++ {
		v := vertex.New(fmt.Sprintf("%d", i))
		list[i] = &v
	}

	list[2].AddNeighbor(list[3])
	list[3].AddNeighbor(list[1])
	list[4].AddNeighbor(list[0])
	list[4].AddNeighbor(list[1])
	list[5].AddNeighbor(list[0])
	list[5].AddNeighbor(list[2])

	return list
}

func createCycledGraph() []*vertex.Vertex {

	list := make([]*vertex.Vertex, 7)

	a := vertex.New("a")
	b := vertex.New("b")
	c := vertex.New("c")
	d := vertex.New("d")
	e := vertex.New("e")
	f := vertex.New("f")
	g := vertex.New("g")

	f.AddNeighbor(&a)
	a.AddNeighbor(&c)
	a.AddNeighbor(&e)
	c.AddNeighbor(&b)
	c.AddNeighbor(&d)
	e.AddNeighbor(&f)
	e.AddNeighbor(&g)

	list[0] = &f
	list[1] = &a
	list[2] = &c
	list[3] = &e
	list[4] = &b
	list[5] = &d
	list[6] = &g

	return list

}

func createGraphList3() ([]*disjointset.Vertex, []*disjointset.Edge) {
	graphList := make([]*disjointset.Vertex, 8)

	edgeList := make([]*disjointset.Edge, 0)

	v0 := disjointset.NewVertex("A")
	v1 := disjointset.NewVertex("B")
	v2 := disjointset.NewVertex("C")
	v3 := disjointset.NewVertex("D")
	v4 := disjointset.NewVertex("E")
	v5 := disjointset.NewVertex("F")
	v6 := disjointset.NewVertex("G")
	v7 := disjointset.NewVertex("H")

	v0.AddNeighbor(disjointset.NewEdge(v0, v1, 5))
	v0.AddNeighbor(disjointset.NewEdge(v0, v4, 9))
	v0.AddNeighbor(disjointset.NewEdge(v0, v7, 8))

	v1.AddNeighbor(disjointset.NewEdge(v1, v2, 12))
	v1.AddNeighbor(disjointset.NewEdge(v1, v3, 15))
	v1.AddNeighbor(disjointset.NewEdge(v1, v7, 6))

	v2.AddNeighbor(disjointset.NewEdge(v2, v3, 3))
	v2.AddNeighbor(disjointset.NewEdge(v2, v6, 11))

	v3.AddNeighbor(disjointset.NewEdge(v3, v6, 9))

	v4.AddNeighbor(disjointset.NewEdge(v4, v7, 5))
	v4.AddNeighbor(disjointset.NewEdge(v4, v5, 4)) //13
	v4.AddNeighbor(disjointset.NewEdge(v4, v6, 20))

	v5.AddNeighbor(disjointset.NewEdge(v5, v6, 13))
	v5.AddNeighbor(disjointset.NewEdge(v5, v2, 1)) //14

	v7.AddNeighbor(disjointset.NewEdge(v7, v2, 7))
	v7.AddNeighbor(disjointset.NewEdge(v7, v5, 6))

	graphList[0] = v0
	graphList[1] = v1
	graphList[2] = v2
	graphList[3] = v3
	graphList[4] = v4
	graphList[5] = v5
	graphList[6] = v6
	graphList[7] = v7

	for _, v := range graphList {
		for _, e := range v.GetAdjacency() {
			edgeList = append(edgeList, e)
		}
	}

	return graphList, edgeList
}

func createGraphList4() []*primsalghorithm.Vertex {

	graphList := make([]*primsalghorithm.Vertex, 7)

	v0 := primsalghorithm.NewVertex("A")
	v1 := primsalghorithm.NewVertex("B")
	v2 := primsalghorithm.NewVertex("C")
	v3 := primsalghorithm.NewVertex("D")
	v4 := primsalghorithm.NewVertex("E")
	v5 := primsalghorithm.NewVertex("F")
	v6 := primsalghorithm.NewVertex("G")

	graphList[0] = v0
	graphList[1] = v1
	graphList[2] = v2
	graphList[3] = v3
	graphList[4] = v4
	graphList[5] = v5
	graphList[6] = v6

	v0.AddNeighbor(primsalghorithm.NewEdge(v0, v1, 1))
	v0.AddNeighbor(primsalghorithm.NewEdge(v0, v2, 2))
	v0.AddNeighbor(primsalghorithm.NewEdge(v0, v3, 12))

	v1.AddNeighbor(primsalghorithm.NewEdge(v1, v0, 1))
	v1.AddNeighbor(primsalghorithm.NewEdge(v1, v3, 4))
	v1.AddNeighbor(primsalghorithm.NewEdge(v1, v4, 7))
	v1.AddNeighbor(primsalghorithm.NewEdge(v1, v6, 8))

	v2.AddNeighbor(primsalghorithm.NewEdge(v2, v0, 2))
	v2.AddNeighbor(primsalghorithm.NewEdge(v2, v3, 6))
	v2.AddNeighbor(primsalghorithm.NewEdge(v2, v5, 3))

	v3.AddNeighbor(primsalghorithm.NewEdge(v3, v0, 12))
	v3.AddNeighbor(primsalghorithm.NewEdge(v3, v1, 4))
	v3.AddNeighbor(primsalghorithm.NewEdge(v3, v2, 6))
	v3.AddNeighbor(primsalghorithm.NewEdge(v3, v4, 2))
	v3.AddNeighbor(primsalghorithm.NewEdge(v3, v5, 5))

	v4.AddNeighbor(primsalghorithm.NewEdge(v4, v1, 7))
	v4.AddNeighbor(primsalghorithm.NewEdge(v4, v3, 2))
	v4.AddNeighbor(primsalghorithm.NewEdge(v4, v5, 4))
	v4.AddNeighbor(primsalghorithm.NewEdge(v4, v6, 9))

	v5.AddNeighbor(primsalghorithm.NewEdge(v5, v2, 3))
	v5.AddNeighbor(primsalghorithm.NewEdge(v5, v3, 5))
	v5.AddNeighbor(primsalghorithm.NewEdge(v5, v4, 4))
	v5.AddNeighbor(primsalghorithm.NewEdge(v5, v6, 2))

	v6.AddNeighbor(primsalghorithm.NewEdge(v6, v1, 8))
	v6.AddNeighbor(primsalghorithm.NewEdge(v6, v4, 9))
	v6.AddNeighbor(primsalghorithm.NewEdge(v6, v5, 2))

	return graphList

}
