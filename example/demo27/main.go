package main
// Go语言教程: 第14章 协程 - 单向管道

import (
	"fmt"
	"time"
)

func main() {
	//// 1. 在默认情况下下，管道是双向
	//var ch1 chan int // 可读可写
	//ch1 = make(chan int, 3)
	//fmt.Printf("%v\n", ch1)
	//
	//// 2. 声明为只写
	//var ch2 chan<- int
	//ch2 = make(chan int, 3)
	//ch2 <- 20
	////fmt.Println(<-ch2)
	////invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)
	//
	//// 3. 声明为只读
	//var ch3 <-chan int
	//ch3 = make(chan int, 3)
	//// ch3 <- 20
	//// invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)
	//num := <-ch3
	//fmt.Println(num)

	//ch := make(chan int, 1)
	//ch <- 10 // 放进去
	//<-ch     // 取走
	//ch <- 12 // 放进去
	//<-ch     // 取走
	//ch <- 17 // 还可以放进去
	//fmt.Println("发送成功")

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "string_" + fmt.Sprintf("%d", i)
	}

	for {
		select {
		case v := <-intChan:
			fmt.Println("int:", v)
		case v := <-stringChan:
			fmt.Println("string:", v)
		default:
			fmt.Printf("没有数据了")
			time.Sleep(time.Second)
			return
		}
	}
}
