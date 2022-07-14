package main

//封装
//名字一般使用CamelCase
//首字母大写代表public，首字母小写代表private。（针对package）
//每个目录一个包，main包包含可执行入口
//为结构定义的方法必须放在同一个包内，可以是不同文件

import (
	tree "Part2/01.tree"
	"fmt"
)

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}
	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}
	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node
	fmt.Println(root)
	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	root.Right = &tree.Node{5, nil, nil}
	root.Right.Left = new(tree.Node)
	root.Left.Right = tree.CreateNode(2)

	root.Print()
	fmt.Println()
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println()

	root.Traverse()
	fmt.Println()

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	pRoot := &root
	pRoot.Print()
	fmt.Println()
	pRoot.SetValue(200)
	pRoot.Print()

	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
