package main

import (
	"encoding/json"
	"fmt"
)

type TreeNode struct {
	Value int       `json:"value"`
	Left  *TreeNode `json:"left"`
	Right *TreeNode `json:"right"`
}

func createNode(value int) *TreeNode {
	return &TreeNode{Value: value}
}

func main() {
	var root TreeNode
	root = TreeNode{Value: 3}
	root.Left = &TreeNode{}
	root.Right = &TreeNode{5, nil, nil}
	root.Right.Left = new(TreeNode)

	root.Right.Right = createNode(10)

	nodes := []TreeNode{
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
}
