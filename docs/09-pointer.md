# 9. 指针

指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址。

Go 语言中的指针操作非常简单，我们只需要记住两个符号：`&`（取地址）和 `*`（根据地址取值）

## 9.1 指针地址和指针类型  

每个变量在运行时都拥有一个地址，这个地址代表变量在内存中的位置。Go 语言中使用 `&` 字符放在变量前面对变量进行取地址操作。Go 语言中的值类型（`int`、`float`、`bool`、`string`、`array`、`struct`）都有对应的指针类型，如：`*int`、`*int64`、`*string` 等。

```go
ptr := &v // 比如 v 的类型为 T
```

其中：  

- `v`：代表被取地址的变量，类型为 T。
- `ptr`：用于接收地址的变量，`ptr` 的类型就为 `*T`，称做 `T` 的指针类型。`*` 代表指针。

```go
func main() {
	var a = 10
	var b = &a
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
}

// 输出：
// 10, int
// 0xc00000a0b8, *int
```

## 9.2 指针取值

在对普通变量使用 `&` 操作符取地址后会获得这个变量的指针，然后可以对指针使用 `*` 操作，也就是指针取值，代码如下。

```go
func main() {
	var a = 10
	var b = &a

	c := *b // 指针取值（根据指针的值去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)
}

// 输出：
// type of c:int
// value of c:10
```

> [!Note]
>
> 取地址操作符 `&` 和取值操作符 `*` 是一对互补操作符，`&` 取出地址，`*` 根据地址取出地址指向的值。

```go
func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	a := 10
	modify1(a)
	fmt.Printf("%d\n", a) // 10
	modify2(&a)
	fmt.Printf("%d\n", a) // 100
}
```

## 9.3 new()/make()

在 Go 语言中对于引用类型的变量，在使用时不仅要声明它，还要为它分配内存空间。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。 Go 语言中 `new()` 和 `make()` 是内建的两个函数，主要用来分配内存。

### 9.3.1 new() 函数分配内存

`new()` 是一个内置的函数，它的函数签名如下：

```go
func new(Type) *Type
```

其中：  

- `Type` 表示类型，`new()` 函数只接受一个参数，这个参数是一个类型
- `*Type` 表示类型指针，`new()` 函数返回一个指向该类型内存地址的指针。

**实际开发中 `new()` 函数不太常用**，使用 `new()` 函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。 

```go
a := new(int)
b := new(bool)

fmt.Printf("%v, %T\n", a, a) // 0xc00009e068, *int
fmt.Printf("%v, %T\n", b, b) // 0xc00009e080, *bool
fmt.Printf("%v, %T\n", *a, *a) // 0, int
fmt.Printf("%v, %T\n", *b, *b) // false, bool
```

```go
var a *int
a = new(int)
*a = 10
fmt.Println(*a)
```

### 9.3.2 make() 函数分配内存  

`make()` 也是用于内存分配的，区别于 `new()`，它只用于 `slice`、 `map` 以及 `channel` 的内存创建，而且它返回的类型就是这三个类型本身，而不是他们的指针类型，因为这三种类型就是引用类型，所以就没有必要返回他们的指针了。`make()` 函数的函数签名如下：

```go
func make(t Type, size ...IntegerType) Type
```

```go
var userinfo map[string]string
userinfo = make(map[string]string)
userinfo["username"] = "张三"
fmt.Println(userinfo)
```

### 9.3.3 new() 与 make() 的区别 

1. 二者都是用来做内存分配的。
2. `make()` 只用于 `slice`、`map` 以及 `channel` 的初始化，返回的还是这三个引用类型本身。
3. 而 `new()` 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

