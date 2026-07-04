# 14. 协程

## 14.1 高并发基础

### 14.1.1 线程与进程

- **进程（Process）**就是程序在操作系统中的一次执行过程，**是系统进行资源分配和调度的基本单位**，进程是一个动态概念，是程序在执行过程中分配和管理资源的基本单位，每一个进程都有一个自己的地址空间。 **一个进程至少有 5 种基本状态，它们是：初始态、执行态、等待状态、就绪状态以及终止状态。**
- **线程**是进程的一个执行实例，**是程序执行的最小单元**，它是比进程更小的能独立运行的基本单位。一个进程可以创建多个线程，同一个进程中的多个线程可以并发执行，一个程序要运行的话至少有一个进程。

### 14.1.2 并行与并发

- **并发：**多个线程同时竞争一个位置，竞争到的才可以执行，每一个时间段只有一个线程在执行。
- **并行：**多个线程可以同时执行，每一个时间段，可以有多个线程同时执行。  

通俗的讲**多线程程序在单核 CPU 上面运行就是并发**，**多线程程序在多核 CPU 上运行就是并行**，如果线程数大于 CPU 核数，则多线程程序在多个 CPU 上面运行既有并行又有并发。

## 14.2 Goroutine 的定义

**Golang 中的主线程：**（可以理解为线程/也可以理解为进程），在一个 Golang 程序的主线程上可以起多个协程。Golang 中多协程可以实现并行或者并发。

**协程：**可以理解为用户级线程，这是对内核透明的，也就是系统并不知道有协程的存在，是完全由用户自己的程序进行调度的。

Golang 的一大特色就是从语言层面原生支持协程，在函数或者方法前面加 `go` 关键字就可创建一个协程。可以说 Golang 中的协程就是goroutine 。

> [!tip]
>
> Golang 中的多协程有点类似其他语言中的多线程。

**多协程和多线程：**Golang 中每个 goroutine (协程) 默认占用内存远比 Java、 C 的线程少。OS 线程（操作系统线程）一般都有固定的栈内存（通常为 2MB 左右），一个 goroutine (协程) 占用内存非常小，只有 2KB 左右，多协程 goroutine 切换调度开销方面远比线程要少。

## 14.3 Goroutine 的使用

```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test () hello, world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {
	go test() // 开启了一个协程
	for i := 1; i <= 10; i++ {
		fmt.Println(" main() hello, golang" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
```

上面代码看上去没有问题，但是要注意主线程执行完毕后即使协程没有执行完毕，程序会退出，所以我们需要对上面代码进行改造。`sync.WaitGroup` 可以实现主线程等待协程执行完毕。

```go
package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup // 1. 定义全局的 WaitGroup

func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test () 你好 golang " + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() // 4. goroutine 结束就登记-1
}

func main() {
	wg.Add(1)       // 2. 启动一个 goroutine 就登记 +1
	defer wg.Wait() // 3. 等待所有 goroutine 结束
	go test()       // 开启了一个协程
	for i := 1; i <= 2; i++ {
		fmt.Println(" main() 你好 golang" + strconv.Itoa(i))
		time.Sleep(time.Millisecond * 50)
	}
}
```

多次执行上面的代码，会发现每次打印的数字的顺序都不一致。这是因为 10 个 goroutine 是并发执行的，而 goroutine 的调度是随机的。

## 14.4 设置 CPU 数量

Go 运行时的调度器使用 `GOMAXPROCS` 参数来确定需要使用多少个 OS 线程来同时执行 Go 代码。默认值是机器上的 CPU 核心数。例如在一个 8 核心的机器上，调度器会把 Go 代码同时调度到 8 个 OS 线程上。

Go 语言中可以通过 `runtime.GOMAXPROCS()` 函数设置当前程序并发时占用的 CPU 逻辑核心数。

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机上面的 CPU 个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum) // cpuNum= 20
    
	// 可以自己设置使用多个 cpu
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
```

## 14.5 管道

管道（Channel）是 Golang 在语言级别上提供的 goroutine 间的通讯方式，我们可以使用 `channel` 在多个 goroutine 之间传递消息。如果说 goroutine 是 Go 程序并发的执行体，`channel` 就是它们之间的连接。`channel` 是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

> [!tip]
>
> Golang 的并发模型是 CSP（Communicating Sequential Processes），提倡**通过通信共享内存**而不是**通过共享内存而实现通信**。

Go 语言中的 `channel` 是一种特殊的类型，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。 

### 14.5.1 管道的声明

`channel` 是一种类型，一种引用类型。声明管道类型的格式如下：

```go
var ch1 chan Type
```

声明的管道后需要使用 `make()` 函数初始化之后才能使用。  

```go
make(chan Type, cap)
```

### 14.5.2 管道的操作

管道有发送（send）、接收(receive）和关闭（close）三种操作。发送和接收都使用 `<-` 符号。

```go
package main

import "fmt"

func main() {
	// 创建一个带缓冲的 channel，容量为 1
	ch := make(chan int, 10)

	// 1. 发送: 把 10 发送到 ch 中
	ch <- 10

	// 2. 接收: 从 ch 中接收数据
	i := <-ch
	fmt.Println(i)

	// 3. 关闭: 关闭 channel
	close(ch)
}
```

关于关闭管道需要注意的事情是，只有在通知接收方 goroutine 所有的数据都发送完毕的时候才需要关闭管道。管道是可以被垃圾回收机制回收的，它和关闭文件是不一样的，**在结束操作之后关闭文件是必须要做的，但关闭管道不是必须的**。

> [!caution]
>
> 1. 对一个关闭的管道再发送值就会导致 `panic`。
>
> 2. 对一个关闭的管道进行接收会一直获取值直到管道为空。  
> 3. 对一个关闭的并且没有值的管道执行接收操作会得到对应类型的零值。
> 4. 关闭一个已经关闭的管道会导致 `panic`。

### 14.5.3 管道阻塞

如果创建管道的时候没有指定容量，那么我们可以叫这个管道为**无缓冲的管道**。无缓冲的管道又称为阻塞的管道。我们来看一下下面的代码：

```go
func main() {
    ch := make(chan int)
    ch <- 10
    fmt.Println("发送成功")
}
```

面这段代码能够通过编译，但是执行的时候会出现以下错误：

```go
fatal error: all goroutines are asleep - deadlock!
```

解决上面问题的方法还有一种就是使用有缓冲区的管道。我们可以在使用 `make()` 函数初始化管道的时候为其指定管道的容量，例如：

```go
func main() {
    ch := make(chan int, 1) // 创建一个容量为 1 的有缓冲区管道
    ch <- 10
    fmt.Println("发送成功")
}
```

只要管道的容量大于零，那么该管道就是有缓冲的管道，管道的容量表示管道中能存放元素的数量。 

管道阻塞具体代码如下：  

```go
func main() {
    ch := make(chan int, 1)
    ch <- 10
    ch <- 12
    fmt.Println("发送成功")
}

// 输出：
// fatal error: all goroutines are asleep - deadlock!
```

解决办法：

```go
func main() {
    ch := make(chan int, 1)
    ch <- 10 // 放进去
    <-ch // 取走
    ch <- 12 // 放进去
    <-ch // 取走
    ch <- 17 // 还可以放进去
    fmt.Println("发送成功")
}

// 输出：
// 发送成功
```

### 14.5.4 遍历管道

当向管道中发送完数据时，我们可以通过 `close()` 函数来关闭管道。当管道被关闭时，再往该管道发送值会引发 `panic`，从该管道取值的操作会先取完管道中的值，再然后取到的值一直都是对应类型的零值。可以通过 `for` 或 `for range` 语句遍历管道。

```go
var ch1 = make(chan int, 5)

for i := 0; i < 5; i++ {
    ch1 <- i + 1
}

close(ch1) // 关闭 channel

for val := range ch1 {
    fmt.Println(val)
}
```

> [!Caution]
>
> 使用 `for range` 遍历管道，当管道被关闭的时候就会退出 `for range`，如果没有关闭管道就会报错误。

```go
var ch2 = make(chan int, 5)

for i := 0; i < 5; i++ {
    ch2 <- i + 1
}

for i := 0; i < 5; i++ {
    fmt.Println(<-ch2)
}
```

> [!Note]
>
> 使用 `for` 遍历管道，不关闭管道，也不会报错。

## 14.6 单向管道

有的时候我们会将管道作为参数在多个任务函数间传递，很多时候我们在不同的任务函数中使用管道都会对其进行限制，比如限制管道在函数中只能发送或只能接收。

```go  
package main

import "fmt"

func main() {
	// 1. 在默认情况下下，管道是双向
	var ch1 chan int // 可读可写
	ch1 = make(chan int, 3)
	fmt.Printf("%v\n", ch1)

	// 2. 声明为只写
	var ch2 chan<- int
	ch2 = make(chan int, 3)
	ch2 <- 20
	//fmt.Println(<-ch2)
	//invalid operation: cannot receive from send-only channel ch2 (variable of type chan<- int)

	// 3. 声明为只读
	var ch3 <-chan int
	ch3 = make(chan int, 3)
	// ch3 <- 20
	// invalid operation: cannot send to receive-only channel ch3 (variable of type <-chan int)
	num := <-ch3
	fmt.Println(num)
}
```

## 14.7 select

在某些场景下我们需要同时从多个通道接收数据。这个时候就可以用到 Golang 中给我们提供的 `select` 多路复用。通常情况通道在接收数据时，如果没有数据可以接收将会发生阻塞。

比如说下面代码来实现从多个通道接受数据的时候就会发生阻塞：  

```go
for {
    // 尝试从 ch1 接收值
    data, ok := <-ch1
    // 尝试从 ch2 接收值
    data, ok := <-ch2
    ...
}
```

`select` 的使用类似于 `switch` 语句，它有一系列 `case` 分支和一个默认的分支。每个 `case` 会对应一个管道的通信（接收或发送）过程。`select` 会一直等待，直到某个 `case` 的通信操作完成时，就会执行 `case` 分支对应的语句。具体格式如下：

```go
select{
    case <-ch1:
    	...
    case data := <-ch2:
    	...
    case ch3<- data:
    	...
	default:
    	...
}
```

```go
package main

import (
	"fmt"
	"time"
)

func main() {
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
```

>[!caution]
>
>使用 `select` 获取管道中的数据不需要 `close()`。

## 14.8 并发安全

### 14.8.1 互斥锁

互斥锁是传统并发编程中对共享资源进行访问控制的主要手段，它由标准库 `sync` 中的 `Mutex` 结构体类型表示。`sync.Mutex` 类型只有两个公开的指针方法，`Lock()` 和 `Unlock()`。`Lock()` 锁定当前的共享资源，`Unlock()` 进行解锁。

```go
package main

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
	mutex.Lock() // 上锁
	count++
	fmt.Println("Count: ", count)
	mutex.Unlock() // 开锁
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
```

使用互斥锁能够保证同一时间有且只有一个 goroutine 进入临界区，其他的 goroutine 则在等待锁；当互斥锁释放后，等待的 goroutine才可以获取锁进入临界区，多个 goroutine 同时等待一个锁时，唤醒的策略是随机的。

### 14.8.2 读写互斥锁  

互斥锁的本质是当一个 goroutine 访问的时候，其他 goroutine 都不能访问。这样在资源同步，避免竞争的同时也降低了程序的并发性能。程序由原来的并行执行变成了串行执行。

> [!tip]
>
> 其实，当我们对一个不会变化的数据只做读操作的话，是不存在资源竞争的问题的。因为数据是不变的，不管怎么读取，多少 goroutine 同时读取，都是可以的。

**读写锁**可以让多个读操作并发，同时读取，但是对于写操作是完全互斥的。也就是说，当一个 goroutine 进行写操作的时候，其他 goroutine 既不能进行读操作，也不能进行写操作。

Golang 中的读写锁由结构体类型 `sync.RWMutex` 表示。此类型的方法集合中包含两对方法：  

```go
// 1. 写锁定和写解锁
func (*RWMutex) Lock()
func (*RWMutex) Unlock()

// 2. 读锁定和读解锁
func (*RWMutex) RLock()
func (*RWMutex) RUnlock()
```

```go
package main

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
```

