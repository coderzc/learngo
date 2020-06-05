package tree

import "fmt"

func (node *Node)Traverse()  {
	if node == nil{
		return
	}
	node.Left.Traverse()
	fmt.Println(node.Value)
	node.Right.Traverse()
}

func (node *Node)TraverseFunc(f func(*Node))  {
	if node == nil{
		return
	}
	node.Left.Traverse()
	f(node)
	node.Right.Traverse()
}
