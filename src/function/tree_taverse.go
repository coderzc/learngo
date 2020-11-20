package main

import (
	"fmt"
	"learngo/src/model/tree"
)

func main() {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{Value: 5}
	root.Right.Left = new(tree.Node)

	nodeCunt := 0
	root.TraverseFunc(func(node *tree.Node) {
		nodeCunt++
		fmt.Println(node.Value)
	})

	channel := root.TraverseWithChannel()
	maxNode := 0
	for node := range channel {
		if node.Value > maxNode {
			maxNode = node.Value
		}
	}
	fmt.Printf("Max node value:%v", maxNode)

}
