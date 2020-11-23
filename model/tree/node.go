package tree

import "fmt"

type Node struct {
	Value int   `json:"value"`
	Left  *Node `json:"left"`
	Right *Node `json:"right"`
}

func CreateNode(value int) *Node {
	return &Node{Value: value}
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.Value = value
}

func (node Node) SetValue2(value int) {
	node.Value = value
}

func (node *Node) Print() {
	fmt.Println(node.Value)
}

func (node *Node) ResetNode() {
	*node = Node{}
}
