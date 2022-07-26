//1.打印出所有的“水仙花数”，所谓“水仙花数”是指一个三位数，其各位数字立方和等于该数本身。
//例如：153是一个“水仙花数”，因为153=1的三次方＋5的三次方＋3的三次方。

package main

import "fmt"

var a, b, c int

func main() {
	fmt.Println("水仙花数：")
	for a = 1; a < 10; a++ {
		for b = 0; b < 10; b++ {
			for c = 0; c < 10; c++ {
				if (a*a*a + b*b*b + c*c*c) == (100*a + 10*b + c) {
					fmt.Printf("%d%d%d\n", a, b, c)
				}
			}
		}
	}
}
