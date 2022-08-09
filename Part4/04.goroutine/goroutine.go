package main

import (
	"fmt"
	"runtime"
	"time"
)

//协程Coroutine
//轻量级“线程”
//非抢占式多任务处理，由协程主动交出控制权
//编译器/解释器/虚拟机层面的多任务

//任何函数只需加上go就能送给调度器运行
//不需要在定义时区分是否为异步函数
//调度器在合适的点进行切换(I/O,select,channel, 等待锁，函数调用（有时），runtime.Gosched())

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
