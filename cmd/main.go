package main

import (
	"fmt"
	"github.com/levioctl/treelib"
)


func main() {
	tree := treelib.NewTree()
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
	fmt.Println(tree)
}
