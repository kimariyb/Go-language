package main
// Go语言教程: 第14章 协程 - GOMAXPROCS/并发性能

import (
	"fmt"
	"runtime"
	"strconv"
	"sync"
	"time"
)

func CustomPrime() float32 {
	start := time.Now().UnixNano() // 记录开始时间

	for i := 2; i <= 10000000; i++ {
		flag := true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i, "是素数")
		}
	}

	end := time.Now().UnixNano()         // 记录结束时间
	duration := float32(end-start) / 1e6 // 转换为毫秒
	fmt.Printf("计算耗时: %.2f 毫秒\n", duration)
	return duration
}

var wg sync.WaitGroup // 1. 定义全局的 WaitGroup

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test () 你好 golang " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 4. goroutine 结束就登记-1
}

func main() {
	// 获取当前计算机上面的 CPU 个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum) // cpuNum= 20
	// 可以自己设置使用多个 cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")

	wg.Add(1)       // 2. 启动一个 goroutine 就登记 +1
	defer wg.Wait() // 3. 等待所有 goroutine 结束
	go test()       // 开启了一个协程
	for i := 1; i <= 2; i++ {
		fmt.Println(" main() 你好 golang" + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 50)
	}

	//CustomPrime() // 1011.39 ms
}
