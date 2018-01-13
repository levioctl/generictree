package treelib

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func TestTreeRoot(t *testing.T) {
	tree := NewTree()
	tree.AddNode(30, tree.rootNid)
	rootNode := tree.GetNode(tree.rootNid)
	assert.Equal(t, rootNode.value, 30, "Root node should have a value of 30")
}
