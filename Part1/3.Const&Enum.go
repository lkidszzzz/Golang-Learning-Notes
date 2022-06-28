package main

import (
	"fmt"
	"math"
)

func consts() {
	//可以选择规定或不规定类型
	const (
		filename string = "test.txt"
		a, b            = 3, 4
	)
	var c int = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename)
	fmt.Println(c)
}

func enums() {
	const (
		golang = 0
		java   = 1
		cpp    = 2
	)
	fmt.Println(golang, java, cpp)
	const (
		//iota常用于常量声明中；
		//本质为一个从0开始的递增计数器，在每个const定义行后加1；
		//不管const定义行中有没有iota，仍然会递增；
		x = iota
		y
		z
	)
	fmt.Println(x, y, z)
	const (
		xx = iota
		_
		yy
		zz
	)
	fmt.Println(xx, yy, zz)
	const (
		//运用iota设计复杂表达式控制枚举方式
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}
