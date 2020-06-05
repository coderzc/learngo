package main

import (
	"fmt"
	"learngo/src/model/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	nodeCunt := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCunt++
		fmt.Println(nodeCunt)
	})

}
