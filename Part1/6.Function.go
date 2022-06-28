package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

//名在前，类型在后。

//这是我们在学习条件语句时写过的代码:
//func eval(a, b int, op string) int {
//	var result int
//	switch op {
//	case "+":
//		result = a + b
//	case "-":
//		result = a - b
//	case "*":
//		result = a * b
//	case "/":
//		result = a / b
//	default:
//		panic("unsupported operator:" + op)
//	}
//	return result
//}
//用golang常用的返回形式（一个值+一个err）重写它
func evalReforged(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operator: %s", op)
	}
}

//用函数的思想再次重写它
func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args "+"(%d, %d)\n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

//带余除法
func div1(a, b int) (int, int) {
	return a / b, a % b
}

//换种写法,但大部分情况还是建议上面的写法。
func div2(a, b int) (q, r int) {
	q = a / b
	r = a % b
	return
}

//可变参数列表
func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	if result, err := evalReforged(23, 3, "*"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error:", err)
	}
	//fmt.Println(evalReforged(23, 3, "*"))
	if result, err := evalReforged(23, 3, "x"); err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error:", err)
	}

	fmt.Println(apply(pow, 23, 3))
	//也可以直接写成匿名函数，省去重新定义函数pow：
	fmt.Println(apply(func(a, b int) int {
		return int(math.Pow(float64(a), float64(b)))
	}, 23, 3))

	fmt.Println(div1(23, 3))
	q, r := div2(23, 3)
	fmt.Println(q, r)
	//用下划线代替不想要接收使用的返回值
	x, _ := div1(233, 3)
	fmt.Println(x)

	fmt.Println(sum(2, 3, 3, 3, 3, 3, 3, 3))
}
