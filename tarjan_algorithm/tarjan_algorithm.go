package tarjanalgorithm

import (
	"fmt"

	"github.com/aldoclg/graph_algorithms/stack"
)

type TarjanAlgorithm struct {
	graph      []*Vertex
	stack      stack.Stack[*Vertex]
	index      int
	sccCounter int //strongly connect component counter
}

func NewTarjaAlgorithm(vetexes []*Vertex) *TarjanAlgorithm {
	return &TarjanAlgorithm{vetexes, *stack.NewStack[*Vertex](), 0, 0}
}

func (this *TarjanAlgorithm) Run() {
	for _, v := range this.graph {
		if !v.Visited {
			this.dfs(v)
		}
	}
}

func (this *TarjanAlgorithm) dfs(vertex *Vertex) {
	vertex.SetIndex(this.index)
	vertex.SetLowLink(this.index)
	vertex.Visited = true
	this.index++
	this.stack.Push(vertex)
	vertex.OnStack = true

	for _, v := range vertex.GetAdjacency() {
		if !v.Visited {
			this.dfs(v)
			vertex.SetLowLink(min(vertex.lowLink, v.lowLink))
		} else if v.OnStack {
			vertex.SetLowLink(min(vertex.lowLink, v.index))
		}
	}

	if vertex.GetLowLink() == vertex.GetIndex() {
		for true {

			if this.stack.IsEmpty() {
				break
			}

			w := this.stack.Pop()
			w.OnStack = false
			w.SetComponentId(this.sccCounter)

			if w.GetName() == vertex.GetName() {
				break
			}
		}
		this.sccCounter++
	}
}

func min(number1, number2 int) int {
	if number1 < number2 {
		return number1
	}
	return number2
}

func (this *TarjanAlgorithm) Print() {
	for _, v := range this.graph {
		fmt.Printf("%v\n", v)
	}
}
