# 1. 变量和常量

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。Go 语言中关键字和保留字都不能用作变量名。

Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且 Go 语言的变量声明后必须使用。

## 1.1 变量

### 1.1.1 var 声明变量

```go
var identifier type
```

```go
var name string
var age int
var isOk bool
```

```go
package main

import "fmt"

func main() {
    var username = "张三"
    var age int = 20
    fmt.Println(username, age)
}
```

### 1.1.2 一次定义多个变量

```go
var identifier1, identifier2 type
```

```go
package main

import "fmt"

func main() {
	var username, sex string
    username = "张三"
    sex = "男"
    fmt.Println(username, sex)
}
```

### 1.1.3 批量声明变量的时候指定类型

```go
var (
    a string
    b int
    c bool
) 
a = "张三"
b = 10
c = true
fmt.Println(a,b,c)
```

### 1.1.4 变量的初始化

Go 语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如：整型和浮点型变量的默认值为 `0`。字符串变量的默认值为空字符串。布尔型变量默认为 `false`。切片、函数、指针变量的默认为 `nil`。

当然我们也可在声明变量的时候为其指定初始值。 变量初始化的标准格式如下：

```go
var name string = "zhangsan"
var age int = 18
```

### 1.1.5 类型推导

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "zhangsan"
var age = 18
```

### 1.1.6 短变量声明法

在函数内部，可以使用更简略的 `:=` 方式声明并初始化变量。  

>[!caution]
>
>**短变量只能用于声明局部变量， 不能用于全局变量的声明**

```go
package main

import "fmt" 

// 全局变量 m
var m = 100

func main() {
    n := 10
    m := 200 // 此处声明局部变量 m
    fmt.Println(m, n)
}
```

### 1.1.7 匿名变量

在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。匿名变量用一个下划线 `_` 表示，例如：

```go
func getInfo() (int, string) {
	return 10, "张三"
}

func main() {
    _, username := getInfo()
    fmt.Println(username)
}
```

匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明。

>[!caution]
>
>1. 函数外的每个语句都必须以关键字开始（`var`、 `const`、 `func`等）
>2. `:=` 不能使用在函数外。
>3. `_` 多用于占位，表示忽略值。

## 1.2 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。常量的声明和变量声明非常类似，只是把 `var` 换成了 `const`，**常量在定义的时候必须赋值**。

### 1.2.1 使用 const 定义常量

```go
const pi = 3.1415
const e = 2.7182
```

多个常量也可以一起声明

```go
const (
    pi = 3.1415
    e = 2.7182
)
```

`const` 同时声明多个常量时， 如果省略了值则表示和上面一行的值相同。 例如：

```go
const (
    n1 = 100
    n2  // 100
    n3	// 100
)
```

## 1.3 变量、常量命名规则

1. 变量名称必须由数字、字母、下划线组成。
2. 标识符开头不能是数字。
3. 标识符不能是保留字和关键字。
4. 变量的名字是区分大小写的如: `age` 和 `Age` 是不同的变量。在实际的运用中，也建议，不要用一个单词大小写区分两个变量。
5. 变量名称一定要见名思意：变量名称建议用名词，方法名称建议用动词。
6. 变量命名一般采用驼峰式，当遇到特有名词（缩写或简称，如 DNS）的时候，特有名词根据是否私有全部大写或小写。

