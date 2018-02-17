package treelib

import (
	"github.com/oleiade/lane"
	// "github.com/rogpeppe/godef/go/ast"
	"fmt"
)

type Tree struct {
	root  *Node
	nodes map[string](*Node)
}

type Node struct {
	value     interface{}
	nid       string
	parentNid string
	children  [](*Node)
}

func (tree *Tree) CreateNode(value interface{}, nid string, parentNid string) *Node {
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

func (tree *Tree) String() string {
	if tree.root == nil {
		return ""
	}

	nodesToPrint := lane.NewQueue()
	nodesToPrint.Enqueue(tree.root)
	var result string = ""
	for !nodesToPrint.Empty() {
		var currentNode *Node = nodesToPrint.Dequeue().(*Node)
		result += fmt.Sprintf("%v", currentNode.value)
	}
	return result
}

func NewTree() *Tree {
	tree := new(Tree)
	tree.nodes = make(map[string]*Node)
	return tree
}
