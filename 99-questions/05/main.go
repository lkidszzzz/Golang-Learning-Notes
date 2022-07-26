//1.题目：输入三个整数x,y,z，请把这三个数由小到大输出。

package main

import "fmt"

var a, b, c int

func main() {
	fmt.Println("请输入要比较的三个数：")
	fmt.Scan(&a, &b, &c)
	if a > b {
		a, b = b, a
	}
	if a > c {
		a, c = c, a
	}
	if b > c {
		b, c = c, b
	}
	fmt.Println("这三个数的从大到小的顺序是：", a, b, c)
}
