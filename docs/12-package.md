# 12. 包

包（package）是多个 Go 源码的集合，是一种高级的代码复用方案。Golang 中的包可以分为三种：1、系统内置包；2、自定义包；3、第三方包。

## 12.1 包管理工具 go mod

Go1.11 版本之后无需手动配置环境变量，使用 `go mod` 管理项目，也不需要非得把项目放到 `$GOPATH` 指定目录下，你可以在你磁盘的任何位置新建一个项目, Go1.13 以后可以彻底不要 `$GOPATH` 了。

### 12.1.1 go mod init 初始化项目  

实际项目开发中我们首先要在我们项目目录中用 `go mod` 命令生成一个 go.mod 文件管理我们项目的依赖。

```sh
go mod init demo21
```

### 12.1.2 其他命令

| 命令              | 用途             | 常用参数     |
| ----------------- | ---------------- | ------------ |
| `go mod init`     | 初始化模块       |              |
| `go mod tidy`     | 整理依赖         |              |
| `go mod download` | 下载依赖         |              |
| `go mod graph`    | 查看依赖图       |              |
| `go mod why`      | 查看依赖原因     | `-m`模块级别 |
| `go mod vendor`   | 创建 vendor 目录 |              |

## 12.2 自定义包

包（package）是多个 Go 源码的集合，一个包可以简单理解为一个存放多个 `.go` 文件的文件夹。该文件夹下面的所有 go 文件都要在代码的第一行添加如下代码，声明该文件归属的包。

```go
package Name
```

> [!caution]
>
> - 一个文件夹下面直接包含的文件只能归属一个 `package`， 同样一个 `package` 的文件不能在多个文件夹下。
> - 包名可以不和文件夹的名字一样， 包名不能包含 `-` 符号。
> - 包名为 `main` 的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含 `main` 包的源代码则不会得到可执行文件。

### 12.2.1 定义包

在当前目录下 `mkdir` 一个 test 文件夹，同时在该文件夹下 `touch` 一个 calc.go 文件。这样表示一个名为 test 的包被创建了。

```sh
mkdir test 
cd test && touch calc.go
```

```go
package test

var Age = 18 // 首字母大写表示公有属性
var name = "小王" // 首字母小写表示私有属性

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}
```

### 12.2.2 导入包

单行导入的格式如下：

```go
import "package 1"
import "package 2"
```

多行导入的格式如下：

```go
import (
    "package 1"
    "package 2"
)
```

如果只希望导入包，而不使用包内部的数据时，可以使用匿名导入包。具体的格式如下：

```go
import _ "package path"
```

在导入包名的时候，我们还可以为导入的包设置别名。通常用于导入的包名太长或者导入的包名冲突的情况。具体语法格式如下：

```go
import alias "package path"
```

具体案例：

```go
package main

import (
	"demo21/test"
	"fmt"
)

func main() {
	add := test.Add(10, 20)
	fmt.Println(add)

	sub := test.Sub(10, 20)
	fmt.Println(sub)
}
```

## 12.3 init() 初始化函数  

在 Go 语言程序执行时导入包语句会自动触发包内部 `init()` 函数的调用。需要注意的是：`init()` 函数没有参数也没有返回值。`init()` 函数在程序运行时自动被调用执行，不能在代码中主动调用它。

```go
package main

import "fmt"

var x int = 10

const pi = 3.14

func init() {
	fmt.Printf("x is %d\n", x)
}

func main() {
	fmt.Printf("pi is %.2f\n", pi)
}

// 输出：
// x is 10
// pi is 3.14
```

## 12.4 第三方包

我们可以在 https://pkg.go.dev/ 查找看常见的 Golang 第三方包。

### 12.4.1 初始化项目

```sh
go mod init project
```

### 12.4.2 下载安装这个包（非必须）  

例如要下载解决 `float` 精度损失的包 `decimal`。

```sh
go get github.com/shopspring/decimal
```

### 12.4.3 go mod tidy 下载丢失的包  

`go mod tidy` 命令可以增加丢失的 `module`，同时去掉未用的 `module`（推荐）

```sh
go mod tidy
```

