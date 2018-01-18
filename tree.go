package treelib

type Tree struct {
	root  *Node
	nodes map[string](*Node)
}

type Node struct {
	value     int
	nid       string
	parentNid string
	children  [](*Node)
}

func (tree *Tree) CreateNode(value int, nid string, parentNid string) *Node {
	node := new(Node)
	if tree.root == nil {
		tree.root = node
	} else if parent, ok := tree.nodes[parentNid]; ok {
		parent.children = append(parent.children, node)
	} else {
		return nil
	}
	node.nid = nid
	node.value = value
	node.parentNid = parentNid
	tree.nodes[nid] = node
	return node
}

func (tree *Tree) GetNode(nid string) *Node {
	return tree.nodes[nid]
}

func NewTree() *Tree {
	tree := new(Tree)
	tree.nodes = make(map[string]*Node)
	return tree
}
