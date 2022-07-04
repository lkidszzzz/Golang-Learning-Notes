package main

import "fmt"

//Golang有指针，但没有指针运算。Golang只有值传递。

func swap1(a, b *int) {
	*b, *a = *a, *b
}

func swap2(a, b int) (int, int) {
	return b, a
}

func main() {
	a1, b1 := 0, 1
	swap1(&a1, &b1)
	a2, b2 := 0, 1
	swap2(a2, b2)

	var p *int
	fmt.Printf("%v\n", p) //打印 nil
	var i int             //定义一个整形变量 i
	p = &i                //使得p指向i，获取i的地址
	fmt.Printf("%v\n", p) //打印内存地址
	fmt.Println(*p)
	*p = 233
	fmt.Printf("%v\n", *p)
	fmt.Printf("%v\n", i)
}
