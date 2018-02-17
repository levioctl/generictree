package treelib

import (
	"testing"

	assert "github.com/stretchr/testify/assert"
)

func Test_CreateRoot(t *testing.T) {
	tree := NewTree()
	var rootNode *Node = tree.CreateNode(30, "root", "")
	assert.Equal(t, rootNode.value, 30, "Root node should have a value of 30")
}

func Test_GetRoot(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	rootNode := tree.GetNode("root")
	assert.Equal(t, rootNode.value, 30, "Root node should have a value of 30")
}

func Test_CreateChild(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	childNode := tree.CreateNode(50, "child", tree.root.nid)
	assert.Equal(t, childNode.value, 50, "Child node should have a value of 30")
}

func Test_GetChild(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	tree.CreateNode(50, "child", tree.root.nid)
	childNode := tree.GetNode("child")
	assert.Equal(t, childNode.value, 50, "Child node should have a value of 30")
}

func Test_RootValuePersistsAfterChildIsBorn(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	tree.CreateNode(50, "child", tree.root.nid)
	rootNode := tree.GetNode("root")
	assert.Equal(t, rootNode.value, 30, "Root node should have a value of 30")
}

func Test_ParentNid(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	childNode := tree.CreateNode(50, "child", tree.root.nid)
	assert.Equal(t, childNode.parentNid, "root", "Parent NID of child is not 'root'")
}

func Test_Children(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(30, "root", "")
	tree.CreateNode(100, "child1", tree.root.nid)
	tree.CreateNode(200, "child2", tree.root.nid)
	tree.CreateNode(110, "child1-of-child1", "child1")

	childrenOfRoot := tree.GetNode("root").children
	assert.Equal(t, childrenOfRoot[0].value, 100)
	assert.Equal(t, childrenOfRoot[1].value, 200)
	childrenOfChild1 := tree.GetNode("child1").children
	assert.Equal(t, childrenOfChild1[0].value, 110)
}


// Does not work yet
func _Test_String(t *testing.T) {
	tree := NewTree()
	tree.CreateNode(100, "root", "")
	tree.CreateNode(110, "first", "root")
	tree.CreateNode(120, "second", "root")
	tree.CreateNode(130, "third", "root")
	tree.CreateNode(111, "firstfirst", "first")
	tree.CreateNode(112, "firstsecond", "first")
	tree.CreateNode(121, "secondfirst", "second")
	tree.CreateNode(131, "thirdfirst", "third")
	tree.CreateNode(131, "thirdsecond", "third")
	tree.CreateNode(131, "thirdthird", "third")

	expected := `100
├── 110
│   ├── 111
│   └── 112
├── 120
│   └── 120
└─── 130
    ├── 131
    ├── 132
    └── 133
`
	assert.Equal(t, expected, tree.String())
}
