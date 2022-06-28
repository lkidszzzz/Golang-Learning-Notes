package main

import (
	"fmt"
	"io/ioutil"
)

func filenames() {
	const filename1 = "test.txt"
	contents1, err1 := ioutil.ReadFile(filename1)
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Printf("%s\n", contents1)
	}
	//也可以缩写成如下形式：
	//这种写法在出了if后无法再访问contents和err
	const filename2 = "non-existent.txt"
	if contents2, err2 := ioutil.ReadFile(filename2); err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%s\n", contents2)
	}
}

func eval(a, b int, op string) int {
	var result int
	//switch会自动break，不需要break时需使用fallthrough。
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	return result
}

func grade(score int) string {
	g := ""
	//switch后面可以没有表达式
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score!: %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	filenames()
	fmt.Println("10+10=", eval(10, 10, "+"))
	fmt.Println("10/10=", eval(10, 10, "/"))
	fmt.Println(grade(0), grade(60), grade(100))
	fmt.Println(grade(-1))
}
