package main

import "fmt"

func main() {
	// 创建一个带缓冲的 channel，容量为 1
	ch := make(chan int, 10)

	// 1. 发送: 把 10 发送到 ch 中
	ch <- 10
	ch <- 20
	ch <- 30

	// 2. 接收: 从 ch 中接收数据
	i := <-ch
	fmt.Println(i)

	// 3. 关闭: 关闭 channel
	close(ch)

	var ch1 = make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch1 <- i + 1
	}

	close(ch1) // 关闭 channel

	for val := range ch1 {
		fmt.Println(val)
	}

	var ch2 = make(chan int, 5)

	for i := 0; i < 5; i++ {
		ch2 <- i + 1
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-ch2)
	}

}
