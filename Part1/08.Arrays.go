package main

import "fmt"

//Go一般不直接使用数组，多用切片！

//数组是值类型，[5]是数组，[]是切片。
func printArray1(arr [3]int) {
	arr[0] = 233
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func printArray2(arr *[3]int) {
	arr[0] = 233
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [6]int
	//使用:=需要赋初值
	arr2 := [3]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr1, arr2, arr3)
	//遍历arr3
	//for i := 0; i < len(arr3); i++ {
	//	fmt.Println(arr3[i])
	//}
	for i, v := range arr3 {
		fmt.Println(i, v)
	}
	//只要数值
	for _, x := range arr3 {
		fmt.Println(x)
	}
	//printArray(arr2)报错，因为Go认为[3]和[5]的数组是不同的数据类型。
	printArray1(arr2)  //输出arr2[0]被改变的结果。
	fmt.Println(arr2)  //arr2[0]不被改变，因为Go只有值传递，调用函数会拷贝数组，而不会修改原数组。
	printArray2(&arr2) //输出arr2[0]被改变的结果。
	fmt.Println(arr2)  //使用指针改变arr2[0]。

	//二维数组
	var grid [4][5]int
	fmt.Println(grid)
}
