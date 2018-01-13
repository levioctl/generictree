package treelib

type Tree struct {
	root    *Node
	rootNid string
}

type Node struct {
	value     int
	parentNid string
}

func (tree *Tree) AddNode(value int, parent string) {
	node := new(Node)
	node.value = value
	tree.root = node
}

func (tree *Tree) GetNode(nid string) *Node {
	return tree.root
}

func NewTree() *Tree {
	tree := new(Tree)
	return tree
}
