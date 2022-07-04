package main

import "fmt"

//Slice本身没有数据，是对底层array的一个view。
//由ptr（头部指针），len（数组或slice的长度），cap（数组或slice的容量）三个变量实现。
//Go语言中切片的内部结构包含地址、大小和容量，切片一般用于快速地操作一块数据集合。

func printSlices() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	//半开半闭，输出2,3,4,5
	fmt.Println("arr[2:6]:", arr[2:6])
	fmt.Println("arr[:6]:", arr[:6])
	fmt.Println("arr[2:]:", arr[2:])
	fmt.Println("arr[:]:", arr[:])
}

func updateSlices(s []int) {
	s[0] = 100
}

func reSilces(s []int) {
	s = s[:5]
	s = s[2:]
}

func silcesOps() {
	var s []int
	for i := 0; i < 6; i++ {
		s = append(s, 2*i+1)
		//cap不够会自动*2
		fmt.Printf("len:%d,cap:%d || ", len(s), cap(s))
		fmt.Printf("i:%d,s:%d\n", i, s)
	}
	//创建提前指定len的切片
	//s := make([]int, 16)
	//创建提前指定cap的切片
	//s := make([]int, 16, 32)

	fmt.Println("Copying Slice")
	s1 := make([]int, 8, 8)
	copy(s1, s)
	fmt.Println("s:", s)
	fmt.Println("s1:", s1)

	fmt.Println("Deleting elements from Slice")
	//Go没有切片删除元素的方法
	s1 = append(s1[:3], s1[4:]...)
	fmt.Println("s1:", s1)

	fmt.Println("Popping from front")
	front := s1[0]
	s1 = s1[1:]
	fmt.Println("Popping from back")
	tail := s1[len(s1)-1]
	s1 = s1[:len(s1)-1]
	fmt.Printf("front:%d,tail:%d,s1:%d", front, tail, s1)
}

func main() {
	printSlices()
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:]
	s2 := arr[:]
	updateSlices(s1)
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2)
	reSilces(s2)
	fmt.Println("Reslice s2:", s2)

	fmt.Println("Extending slice")
	arr[2] = 2
	s1 = arr[2:6]
	s2 = s1[3:5]
	//取的是s1[3],s1[4]
	//s1[4]取不出来:index out of range
	//但由于slice是对arr的view，slice仍然能截取并找到其下标对应的值（如果不超过cap）。
	fmt.Println("s1:", s1)
	fmt.Println("s2:", s2) //s2:[5 6]
	fmt.Println("len(s1):", len(s1))
	fmt.Println("cap(s1):", cap(s1))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	fmt.Println("s5:", s5)
	//s4和s5view的是拷贝了数据、增加了长度的新array。
	//添加元素时如果超过cap，系统会重新分配更大的底层数组。
	//由于值传递的关系，必须接受append的返回值。
	fmt.Println("arr:", arr)

	silcesOps()
}
