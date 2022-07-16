package main

import "fmt"

//Go 并不提供典型的，类型驱动的子类化概念，但通过将类型内嵌到结构体或接口中， 它就能 “借鉴” 部分实现。(类似继承)

type A struct {
	ax, ay int
}
type B struct {
	A
	bx, by float32
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}
