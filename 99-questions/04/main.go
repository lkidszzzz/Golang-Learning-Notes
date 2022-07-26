//1.题目：输入某年某月某日，判断这一天是这一年的第几天？
//特殊情况，闰年且输入月份大于3时需考虑多加一天。
//if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
//	feb = 29
//} else {
//	feb = 28
//}
//然后switch case计算

package main

import (
	"fmt"
	"strconv"
	"time"
)

var year, month, day int

func main() {
	fmt.Println("请分别输入年、月、日：")
	fmt.Scan(&year, &month, &day)
	a, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-01-01")
	if month > 9 {
		if day > 9 {
			b, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-"+strconv.Itoa(day))
			d := b.Sub(a)
			fmt.Println(d.Hours()/24 + 1)
		} else {
			b, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-"+strconv.Itoa(month)+"-0"+strconv.Itoa(day))
			d := b.Sub(a)
			fmt.Println(d.Hours()/24 + 1)
		}
	} else {
		if day > 9 {
			b, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-0"+strconv.Itoa(month)+"-"+strconv.Itoa(day))
			d := b.Sub(a)
			fmt.Println(d.Hours()/24 + 1)
		} else {
			b, _ := time.Parse("2006-01-02", strconv.Itoa(year)+"-0"+strconv.Itoa(month)+"-0"+strconv.Itoa(day))
			d := b.Sub(a)
			fmt.Println(d.Hours()/24 + 1)
		}
	}
}
