package main

import (
	"Part4/02.functional/fib"
	"bufio"
	"fmt"
	"os"
)

//常用:
//Open/Close
//Lock/Unlock
//PrintHeader/PrintFooter

//defer确保调用在函数结束前发生
//相当于栈，先进后出
func trydefer1() {
	//3,2,1 (相当于栈，先进后出)
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//return
	//panic("err")
	//fmt.Println(4)
}

//参数在defer语句时计算
func trydefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println(i)
		if i == 30 {
			panic("printed too many")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	trydefer1()
	trydefer2()
	writeFile("fib.txt")
}
