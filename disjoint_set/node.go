package disjointset

type Node struct {
	id         int
	height     int
	parentNode *Node
}

func NewNode(id, height int, parentNode *Node) *Node {
	return &Node{id, height, parentNode}
}

func (this *Node) GetParent() *Node {
	return this.parentNode
}

func (this *Node) SetParent(node *Node) {
	this.parentNode = node
}

func (this *Node) GetHeight() int {
	return this.height
}

func (this *Node) SetHeight(height int) {
	this.height = height
}
