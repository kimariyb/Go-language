# 8. 函数

函数是组织好的、可重复使用的、用于执行指定任务的代码块。Go 语言中支持：函数、匿名函数和闭包。

## 8.1 函数的声明

Go 语言中定义函数使用 `func` 关键字，具体格式如下：

```go
func funcName (para Type) (return Type) {
    funcBody
}
```

其中：

- `funcName`：函数名，由字母、数字、下划线组成。但函数名的第一个字母不能是数字。在同一个包内，函数名也称不能重名。
- `para Type`：参数，由参数变量和参数变量的类型组成，多个参数之间使用 `,` 分隔。
- `return Type`：返回值，由返回值变量和其变量类型组成，也可以只写返回值的类型，多个返回值必须用 `()` 包裹，并用 `,` 分隔。
- `funcBody`：函数体，实现指定功能的代码块。

```go
func intSum(x int, y int) int {
	return x + y
}
```

## 8.2 函数的调用

定义了函数之后， 我们可以通过 `funcName()` 的方式调用函数。 例如我们调用上面定义函数，代码如下：

```go
func main() {
    res := intSum(10, 20)
    fmt.Println(res)
}
```

> [!tip]
>
> 调用有返回值的函数时，可以不接收其返回值。

## 8.3 函数参数

### 8.3.1 类型简写 

函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：

```go
func intSum(x, y int) int {
	return x + y
}
```

### 8.3.2 可变参数

可变参数是指函数的参数数量不固定。 Go 语言中的可变参数通过在参数名后加 `...` 来标识。

> [!tip]
>
> 可变参数通常要作为函数的最后一个参数。

```go
func intSum2(x ...int) int {
	fmt.Printf("%T\n", x) // []int
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}
```

## 8.4 函数返回值

Go 语言中通过 `return` 关键字向外输出返回值。

### 8.4.1 函数多返回值  

Go 语言中函数支持多返回值，函数如果有多个返回值时必须用 `()` 将所有返回值包裹起来。

```go
func calc(x, y int) (int, int) {
    sum := x + y
    sub := x - y
    return sum, sub
}
```

### 8.4.2 返回值命名

函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 `return` 关键字返回。

```go
func calc(x, y int) (sum, sub int) {
    sum = x + y
    sub = x - y
    return
}
```

## 8.5 函数类型与变量

我们可以使用 `type` 关键字来定义一个函数类型，具体格式如下：

```go
type Calculation func(int, int) int
```

凡是满足这个条件的函数都是 `Calculation` 类型的函数， 例如下面的 `add()` 和 `sub()` 是 `Calculation` 类型。

```go
var c Calculation = add
fmt.Printf("%T\n", c) // main.Calculation
fmt.Println(c(1, 2)) // 3
```

## 8.6 高级函数

### 8.6.1 函数作为参数  

```go
func add(x, y int) int {
	return x + y
} 

func calc(x, y int, operator func(int, int) int) int {
	return operator(x, y)
} 

func main() {
    ret2 := calc(10, 20, add)
    fmt.Println(ret2) // 30
}
```

### 8.6.2 函数作为返回值

```go
func doCalc(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	var a = doCalc("+")
	fmt.Println(a(5, 2)) // 7
}
```

## 8.7 匿名函数和闭包

### 8.7.1 匿名函数

在 Go 语言中函数内部不能再像之前那样定义函数了，只能定义匿名函数。匿名函数就是没有函数名的函数，匿名函数的定义格式如下：

```go
func(para Type) (return Type) {
    funcBody
}
```

匿名函数因为没有函数名，所以没办法像普通函数那样调用，所以匿名函数需要保存到某个变量或者作为立即执行函数：

```go
fmt.Println(
    func(x, y int) int {
        return x + y
    }(10, 20)) // 30
```

> [!tip]
>
> 匿名函数多用于实现回调函数和闭包

### 8.7.2 闭包

闭包可以理解成**定义在一个函数内部的函数**。在本质上，闭包是将函数内部和函数外部连接起来的桥梁。或者说是函数和其引用环境的组合体。首先我们来看一个例子：

```go
package main

import "fmt"

func adder() func(int) int {
	var x int

	return func(y int) int {
		x += y
		return x
	}
}

func main() {
	var f = adder()
	fmt.Println(f(5)) // 5
	fmt.Println(f(5)) // 10
	fmt.Println(f(5)) // 15
}

```

变量 `f` 是一个函数并且它引用了其外部作用域中的 `x` 变量，此时 `f` 就是一个闭包。在 `f` 的生命周期内，变量 `x` 也一直有效。

```go
func Calc(base int) (func(int) int, func(int) int) {
	add := func(y int) int {
		base += y
		return base
	}
	sub := func(y int) int {
		base -= y
		return base
	}
	return add, sub
}

func main() {
	f1, f2 := Calc(10)
	fmt.Println(f1(5), f2(5)) // 15 10
	fmt.Println(f1(5), f2(5)) // 15 10
	fmt.Println(f1(5), f2(5)) // 15 10
}
```

## 8.8 defer

Go 语言中的 `defer` 语句会将其后面跟随的语句进行延迟处理。在 `defer` 归属的函数即将返回时，将延迟处理的语句按 `defer` 定义的逆序进行执行，也就是说，**先被 `defer` 的语句最后被执行，最后被 `defer` 的语句，最先被执行**。

```go
func main() {
	fmt.Println("Start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("End")
}

// 输出：
// Start
// End
// 3
// 2
// 1
```

> [!tip]
>
> 由于 `defer` 语句延迟调用的特性，所以 `defer` 语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等

### 8.8.1 defer 执行时机  

在 Go 语言的函数中 `return` 语句在底层并不是原子操作，它分为**给返回值赋值**和 `RET` 指令两步。而 `defer` 语句执行的时机就在返回值赋值操作后，`RET` 指令执行前。

```go
func f1() int {
    x := 5
    defer func() { x = 10 }()
    return x
    // 返回 5，不是 10
}
```

**底层执行顺序**：

1. 创建匿名返回值 `r` (在栈上)
2. 将 `x` 的值 5 赋值给 `r`
3. 执行 `defer：x = 10` (不影响 `r`)
4. 返回 `r` 的值 5

### 8.8.2 defer 示例

```go
func calc(index string, a, b int) int {
    res := a + b
    fmt.Println(index, a, b, res)
    return res
} 

func main() {
    x := 1
    y := 2
    defer calc("AA", x, calc("A", x, y))
    x = 10
    defer calc("BB", x, calc("B", x, y))
    y = 20
}

// 输出：
// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4
```

> [!tip]
>
> `defer` 注册要延迟执行的函数时该函数所有的参数都需要确定其值。

## 8.9 panic()/recover()

Go 语言使用 `panic()/recover()` 模式来处理错误。`panic()` 可以在任何地方引发，但 `recover()` 只有在 `defer` 调用的函数中有效。 

### 8.9.1 panic()/recover() 的基本使用

```go
package main

import "fmt"

func f1() {
	fmt.Println("hello")
}

func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	panic("panic")
}

func main() {
	f1()
	f2()
	fmt.Println("world")
}

// 输出：
// hello
// recover: panic
// world
```

> [!caution]
>
> 1. `recover()` 必须搭配 `defer` 使用。
> 1. `defer` 一定要在可能引发 `panic()` 的语句之前定义。

### 8.9.2 defer/recover() 实现异常处理

```go
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err) // 抛出除 0 异常
		}
	}()
    
    num1 := 10
    num2 := 0
    res := num1 / num2
    fmt.Println("res=", res)
}
```

### 8.9.3 defer/panic()/recover() 抛出异常

```go
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件错误")
}

func f3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	f3()
	fmt.Println("continue...")
}

// 输出：
// 读取文件错误
// continue...
```

