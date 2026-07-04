package main
// Go语言教程: 第14章 协程 - 互斥锁 sync.Mutex

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
var mutex sync.Mutex

func add() {
	mutex.Lock()
	count++
	fmt.Println("Count: ", count)
	mutex.Unlock()
	time.Sleep(time.Millisecond)
	wg.Done()
}

func main() {
	// 设置多个CPU核心来增加并发竞争的可能性
	runtime.GOMAXPROCS(runtime.NumCPU())

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go add()
	}
	wg.Wait()
	fmt.Printf("Final count value: %d\n", count)
}
