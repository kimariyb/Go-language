# 4. 流程控制

## 4.1 if else

```go
var score int = 80

if score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好")
} else if score >= 60 {
    fmt.Println("及格")
} else {
    fmt.Println("不及格")
}
```

> [!caution]
>
> Go 语言规定与 if 匹配的左括号 { 必须与 if 和表达式放在同一行，{ 放在其他位置会触发编译错误。同理，与 else 匹配的 { 也必须与 else 写在同一行， else 也必须与上一个 if 或 else if 右边的大括号在同一行。

> [!caution]
>
> Go 语言中没有三目运算符

## 4.2 for

Go 语言中的所有循环类型均可以使用 `for` 关键字来完成。

```go
for i := 0; i < 10; i++ {
    fmt.Println(i)
}
```

for 循环的初始语句可以被忽略，但是初始语句后的分号必须要写，例如：

```go
i := 0
for ; i < 10; i++ {
    fmt.Println(i)
}
```

for 循环的初始语句和结束语句都可以省略，例如：

```go
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}
```

> [!caution]
>
> Go 语言中是没有 while 语句的，我们可以通过 for 代替  

for 循环可以通过 `break`、 `goto`、 `return`、 `panic` 语句强制退出循环。

```go
k := 1
for {
    if k <= 10 {
        fmt.Println(k)
    } else {
        break // break 就是跳出这个 for 循环
    }
    k++
}
```

## 4.3 for range

Go 语言中可以使用 `for range` 遍历数组、切片、字符串、map 及通道（channel）。通过 `for range` 遍历的返回值有以下规律：

1. 数组、切片、字符串返回索引和值。  
2. map 返回键和值  
3. 通道（channel）只返回通道内的值。  

```go
s := "hello world"
for index, value := range s {
    fmt.Printf("%d:%c\n", index, value)
}

slice := []int{1, 2, 3, 4, 5}
for index, value := range slice {
    fmt.Printf("%d:%d\n", index, value)
}
```

## 4.4 switch case

使用 `switch` 语句可方便地对大量的值进行条件判断。Go 语言规定每个 `switch` 只能有一个 `default` 分支。  

```go
suffix := ".a"
switch suffix {
    case ".html":
    fmt.Println("text/html")
    case ".css":
    fmt.Println("text/css")
    case ".js":
    fmt.Println("text/javascript")
    default:
    fmt.Println("格式错误")
    
```

Go 语言中每个 case 语句中可以不写 `break`，不加 `break` 也不会出现穿透的现象。一个分支可以有多个值，多个 case 值中间使用英文逗号分隔。

```go
n := 2
switch n {
    case 1, 3, 5, 7, 9:
    fmt.Println("奇数")
    case 2, 4, 6, 8:
    fmt.Println("偶数")
    default:
    fmt.Println(n)
}
```

`fallthrough` 是 Go 语言中 `switch` 语句的一个特殊关键字，用于强制执行下一个 case 分支，而不管下一个 case 的条件是否满足。

```go
num := 2

switch num {
    case 1:
    fmt.Println("数字是 1")
    case 2:
    fmt.Println("数字是 2")
    case 3:
    fmt.Println("数字是 3")
    default:
    fmt.Println("其他数字")
} 

// 输出: 
// 数字是 2
```

```go
num := 2

switch num {
    case 1:
    fmt.Println("数字是 1")
    case 2:
    fmt.Println("数字是 2")
    fallthrough  // 强制执行下一个 case
    case 3:
    fmt.Println("数字是 3")
    case 4:
    fmt.Println("数字是 4")
    default:
    fmt.Println("其他数字")
}

// 输出:
// 数字是 2
// 数字是 3
```

## 4.5 break

Go 语言中 `break` 语句用于以下几个方面：  

1. 用于循环语句中跳出循环，并开始执行循环之后的语句。  
2. `break` 在 `switch`（开关语句）中在执行一条 case 后跳出语句的作用。  
3. 在多重循环中，可以用标号 label 标出想 `break` 的循环。  

```go
lable:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable
			}
			fmt.Println("i, j", i, j)
		}
	}
```

## 4.6 continue

`continue` 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 `for` 循环内使用。  

> [!tip]
>
> 在 continue 语句后添加标签时，表示开始标签对应的循环。  

```go
here:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i, j", i, j)
		}
	}
```

## 4.7 goto

`goto` 语句通过标签进行代码间的无条件跳转。`goto` 语句可以在快速跳出循环、避免重复退出上有一定的帮助。 

```go
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Println("i, j", i, j)
		}
	}
	return
breakTag:
	fmt.Println("breakTag")
```

