package main

import "fmt"

func adder1() func(int) int {
	sum := 0
	//闭包{函数体[局部变量+自由变量(...树)]}
	return func(x int) int {
		sum += x
		return sum
	}
}

//"正统函数式编程"(不能有变量)
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(x int) (int, iAdder) {
		return base + x, adder2(base + x)
	}
}

func main() {
	a1 := adder1()
	for i := 0; i < 10; i++ {
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, a1(i))
	}

	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Printf("0 + 1 + ... + %d = %d\n", i, s)
	}
}
