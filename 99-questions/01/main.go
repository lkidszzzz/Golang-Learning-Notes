//1.题目：有1、2、3、4个数字，能组成多少个互不相同且无重复数字的三位数？都是多少？

package main

import "fmt"

func main() {
	count := 0
	fmt.Println("组成的互不相同且无重复数字的三位数如下：")
	for a := 1; a < 5; a++ {
		for b := 1; b < 5; b++ {
			for c := 1; c < 5; c++ {
				if (a != b) && (a != c) && (b != c) {
					count++
					fmt.Println(a, b, c)
				}
			}
		}
	}
	fmt.Println("总共能组成", count, "个互不相同且无重复数字的三位数。")
}
