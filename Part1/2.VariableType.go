package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

//内建变量类型：
//bool,string...
//(u)int,(u)int8,(u)int16...(u)int64,uintptr:其中(u)代表无符号，数字代表位数，uintptr是一种指针。
//byte,rune:byte8位，rune类似于char，但有32字节，避开了许多1字节的char带来的坑。
//float32,float64,complex64,complex128:complex是复数，有实部和虚部（i=根号-1）,可以在二维坐标系中理解。
//Golang只有强制类型转换！

func complexValue() {
	//表示复数
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func euler() {
	//欧拉公式（i不被识别为复数时可写成1i），由于受到浮点数影响不能得到准确答案0。
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
	fmt.Println(cmplx.Exp(1i*math.Pi) + 1)
	fmt.Printf("%.3f\n", cmplx.Exp(1i*math.Pi)+1)
}

func triangle() {
	//Golang中均需要强制转换
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
	//强制转换过程中，浮点数直接向下取整，有时会导致较大误差，因此...
	var x, y, z float64 = 2, 5, 0
	z = math.Sqrt(float64(x*x + y*y))
	fmt.Println(math.Floor(float64(z) + 0.5))
}

func main() {
	fmt.Println()
	complexValue()
	euler()
	triangle()
}
