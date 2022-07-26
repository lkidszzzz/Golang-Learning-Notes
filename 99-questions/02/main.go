//1.题目：企业发放的奖金根据利润提成。
//利润(I)低于或等于10万元时，奖金可提10%；
//利润高于10万元，低于20万元时，低于10万元的部分按10%提成，高于10万元的部分，可可提成7.5%；
//20万到40万之间时，高于20万元的部分，可提成5%；40万到60万之间时高于40万元的部分，可提成3%；
//60万到100万之间时，高于60万元的部分，可提成1.5%；
//高于100万元时，超过100万元的部分按1%提成从键盘输入当月利润，求应发放奖金总数？

package main

import "fmt"

var interest, bonus float32

func main() {
	fmt.Println("请输入利润金额：")
	fmt.Scan(&interest)
	switch {
	case interest <= 100000:
		bonus = 0.1 * interest
		fmt.Println("奖金金额：", bonus)
	case interest > 100000 && interest <= 200000:
		bonus = 0.1*100000 + (interest-100000)*0.075
		fmt.Println("奖金金额：", bonus)
	case interest > 200000 && interest <= 400000:
		bonus = 0.1*100000 + 0.075*100000 + (interest-200000)*0.05
		fmt.Println("奖金金额：", bonus)
	case interest > 400000 && interest <= 600000:
		bonus = 0.1*100000 + 0.075*100000 + 0.05*200000 + (interest-400000)*0.03
		fmt.Println("奖金金额：", bonus)
	case interest > 600000 && interest <= 1000000:
		bonus = 0.1*100000 + 0.075*100000 + 0.05*200000 + 0.03*200000 + (interest-600000)*0.015
		fmt.Println("奖金金额：", bonus)
	case interest > 1000000:
		bonus = 0.1*100000 + 0.075*100000 + 0.05*200000 + 0.03*200000 + 0.15*400000 + (interest-1000000)*0.01
		fmt.Println("奖金金额：", bonus)
	}
}
