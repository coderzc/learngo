package base

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/coderzc/learngo/model/tree"
)

func TestStruct(t *testing.T) {
	var root tree.Node
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)

	root.Right.Right = tree.CreateNode(10)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	//Marshal失败时err!=nil
	jsonNodes, err := json.Marshal(root)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonNodes))

	root.SetValue(10)

	root.Print()

	pRoot := &root
	pRoot.Print()

	pRoot = nil
	//pRoot.setValue(10)

	root.SetValue2(99)

	fmt.Println(root)
	root.ResetNode()
	fmt.Println(root)

}

type MyTreeNode struct {
	node *tree.Node
}

func (myTreeNode *MyTreeNode) postOrder() {
	if myTreeNode == nil {
		return
	}
	(&MyTreeNode{myTreeNode.node.Left}).postOrder()
	(&MyTreeNode{myTreeNode.node.Right}).postOrder()
	myTreeNode.node.Print()
}
