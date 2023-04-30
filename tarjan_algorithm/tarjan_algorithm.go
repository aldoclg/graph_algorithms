package tarjanalgorithm

import (
	"fmt"
	"math"
	"vertex/stack"
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
			newLowLink := math.Min(float64(vertex.lowLink), float64(v.lowLink))
			vertex.SetLowLink(int(newLowLink))
		} else if v.OnStack {
			newLowLink := math.Min(float64(vertex.lowLink), float64(v.index))
			vertex.SetLowLink(int(newLowLink))
		}

		if vertex.lowLink == vertex.index {
			for true {
				w := this.stack.Pop()
				w.OnStack = false
				w.SetComponentId(this.sccCounter)

				if w == vertex {
					break
				}
			}

			this.sccCounter++
		}
	}
}

func (this *TarjanAlgorithm) print() {
	for _, v := range this.graph {
		fmt.Printf("%v\v", v)
	}
}
