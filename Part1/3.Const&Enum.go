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
	const (
		x = iota
		y
		z
	)
	fmt.Println(golang, java, cpp)
	fmt.Println(x, y, z)
	const (
		xx = iota
		_
		yy
		zz
	)
	fmt.Println(xx, yy, zz)
	const (
		//通过设计表达式控制枚举方式
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
