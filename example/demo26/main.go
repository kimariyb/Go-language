package main
// Go语言教程: 第14章 协程 - 并发素数计算

import (
	"fmt"
	"sync"
	"time"
)

const maxNum = 120000
const workers = 8

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}

	// 检查到 sqrt(num)
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func main() {
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)
	done := make(chan bool)

	var wg sync.WaitGroup

	// 启动 worker
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for num := range jobs {
				if isPrime(num) {
					results <- num
				}
			}
		}(i)
	}

	// 收集结果
	go func() {
		wg.Wait()
		close(results)
		done <- true
	}()

	// 发送任务
	go func() {
		for i := 0; i <= maxNum; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// 统计
	count := 0
	primes := make([]int, 0, 10000)

	collecting := true
	for collecting {
		select {
		case prime, ok := <-results:
			if !ok {
				collecting = false
			} else {
				count++
				primes = append(primes, prime)
			}
		case <-time.After(5 * time.Second):
			fmt.Println("处理超时")
			return
		}
	}

	fmt.Printf("素数数量: %d\n", count)
}
