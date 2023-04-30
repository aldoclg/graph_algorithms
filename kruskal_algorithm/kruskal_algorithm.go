package kruskalalgorithm

import (
	"sort"

	disjointset "github.com/aldoclg/graph_algorithms/disjoint_set"
)

func Compute(vertexes []*disjointset.Vertex, edges []*disjointset.Edge) []*disjointset.Edge {
	disjointSet := disjointset.NewDisjointSet(vertexes)

	minimumSpanningTree := make([]*disjointset.Edge, 0)

	sort.Sort(disjointset.SortableEdges(edges))

	for _, edge := range edges {
		u := edge.GetStartEdge()
		v := edge.GetTargetEdge()
		if disjointSet.Find(u.GetNode()) != disjointSet.Find(v.GetNode()) {
			minimumSpanningTree = append(minimumSpanningTree, edge)
		}
		disjointSet.Union(u.GetNode(), v.GetNode())
	}

	return minimumSpanningTree
}
