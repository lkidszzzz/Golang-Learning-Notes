package main

import "fmt"

//func adder() func(int) int {
//	sum := 0
//	return func(v int) int {
//		sum += v
//		return sum
//	}
//}

//正统函数式编程：

type iAdder func(int) (int, iAdder)

func adder0(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder0(base + v)
	}
}

func main() {
	a := adder0(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a = a(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n",
			i, s)
	}
}
