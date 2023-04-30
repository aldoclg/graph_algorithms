package disjointset

type DisjointSet struct {
	vertexes []*Vertex
}

func NewDisjointSet(vertexes []*Vertex) *DisjointSet {

	for _, v := range vertexes {
		v.SetNode(NewNode(0, 0, nil))
	}

	return &DisjointSet{vertexes}
}

func (this *DisjointSet) Find(node *Node) *Node {
	actual := node

	for actual.GetParent() != nil {
		actual = actual.GetParent()
	}

	root := actual
	actual = node
	var temp *Node

	for actual != root {
		temp = actual.GetParent()
		actual.SetParent(root)
		actual = temp
	}

	return actual
}

func (this *DisjointSet) Union(node1, node2 *Node) {

	root1, root2 := this.Find(node1), this.Find(node2)

	if root1 == root2 {
		return
	}

	if root1.GetHeight() > root2.GetHeight() {
		root1.SetParent(root2)
	} else if root2.GetHeight() > root1.GetHeight() {
		root2.SetParent(root1)
	} else {
		root2.SetParent(root1)
		root1.SetHeight(root1.GetHeight() + 1)
	}

}
