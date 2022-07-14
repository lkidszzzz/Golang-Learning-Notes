package tree

import "fmt"

//Go语言仅支持封装，不支持继承和多态
//Go语言没有class，只有struct

type Node struct {
	Value       int
	Left, Right *Node
}

//工厂函数，返回了局部变量的地址。
func CreateNode(value int) *Node {
	return &Node{Value: value}
}

//接收者：treenode，传值。
//要改变内容必须使用指针接收者
//结构过大也考虑使用指针接收者
//一致性：如有指针接收者，最好都是指针接收者
//值接收者是Go语言特有
//值/指针接收者均可接收值/指针
func (node Node) Print() {
	fmt.Print(node.Value)
}

func (node *Node) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node.Ignored.")
		return
	}
	node.Value = value
}
