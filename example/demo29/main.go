package main
// Go语言教程: 第14章 协程 - 读写锁 sync.RWMutex

import (
	"fmt"
	"sync"
	"time"
)

var count int
var mutex sync.RWMutex
var wg sync.WaitGroup

func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 3)
	mutex.Unlock()
	wg.Done()
}

func read() {
	mutex.RLock()
	fmt.Println("执行读操作")
	time.Sleep(time.Second * 3)
	mutex.RUnlock()
	wg.Done()
}

func main() {
	defer wg.Wait()

	// 开启 10 个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 开启 10 个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
}
