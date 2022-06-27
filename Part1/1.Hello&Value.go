package main

import "fmt"

//在函数外面定义变量不能用:= 因为需要关键字开头
//函数外部定义的变量不是全局变量，范围为package
var (
	aa = 0
	bb = false
	ss = "hello,hello."
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b, c int = 1, 2, 3
	var s string = "hello"
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 0, true, "def", "hello"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 0, true, "def", "hello"
	//:=是对var...=的缩写，仅第一次加:
	a = 1
	fmt.Println(a, b, c, s)
}

func main() {
	fmt.Println("Hello world!")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	fmt.Println(aa, bb, ss)
}
