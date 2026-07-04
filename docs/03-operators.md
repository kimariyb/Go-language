# 3. 运算符

## 3.1 算术运算符

```go
// 1、+ 运算符
fmt.Println("10+3=", 10+3) // =13

// 2、- 运算符
fmt.Println("10-3=", 10-3) // =7

// 3、* 运算符
fmt.Println("10*3=", 10*3) // =30

// 4、/ 运算符
//除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
fmt.Println("10/3=", 10/3)     //3
fmt.Println("10.0/3=", 10.0/3) //3.3333333333333335

// 5、% 运算符
// 取余注意 余数 = 被除数 -（被除数/除数）* 除数
fmt.Println("10%3=", 10%3)     // =1
fmt.Println("-10%3=", -10%3)   // -1
fmt.Println("10%-3=", 10%-3)   // =1
fmt.Println("-10%-3=", -10%-3) // =-1
```

> [!caution]
>
> ++（自增）和 --（自减）在 Go 语言中是单独的语句，并不是运算符。  

```go
var i int = 8
var a int
a = i++ //错误， i++只能独立使用
a = i-- //错误, i--只能独立使用
```

```go
var i int = 1
++i // 错误， 在 golang 没有 前++
--i // 错误， 在 golang 没有 前--
fmt.Println("i=", i)
```

> [!caution]
>
> 在 golang 中，++ 和 -- 只能独立使用。在 golang 中没有前++   

```go
var i int = 1
i++ // 正确写法
fmt.Println("i=", i)
```

## 3.2 关系运算符

```go
var n1 int = 9
var n2 int = 8

// 1、== 运算符
fmt.Println(n1 == n2) // false
// 2、!= 运算符
fmt.Println(n1 != n2) // true
// 3、> 运算符
fmt.Println(n1 > n2) // true
// 4、>= 运算符
fmt.Println(n1 >= n2) // true
// 5、< 运算符
fmt.Println(n1 < n2) // flase
// 6、<= 运算符
fmt.Println(n1 <= n2) // flase
```

## 3.3 逻辑运算符

```go
// 1、&& 逻辑运算符 
var age int = 40
if age > 30 && age < 50 {
    fmt.Println("ok1")
}
if age > 30 && age < 40 {
    fmt.Println("ok2")
}

// 2、|| 逻辑运算符
if age > 30 || age < 50 {
    fmt.Println("ok3")
}
if age > 30 || age < 40 {
    fmt.Println("ok4")
}
// 3、! 逻辑运算符
if age > 30 {
    fmt.Println("ok5")
}
if !(age > 30) {
    fmt.Println("ok6")
}
```

逻辑运算符短路：在逻辑运算中，**当能够确定整个表达式的结果时，就不再计算剩余部分**的特性。

```go
// 示例1：&& 短路
fmt.Println("=== && 短路演示 ===")

a, b := 10, 20

// 第一个条件为 false，第二个条件不会执行
if a > 100 && b > 0 {
    fmt.Println("条件成立")
} else {
    fmt.Println("条件不成立，b>0不会被执行")
}

// 示例2：|| 短路
fmt.Println("\n=== || 短路演示 ===")

// 第一个条件为true，第二个条件不会执行
if a < 100 || b < 0 {
    fmt.Println("条件成立，b<0不会被执行")
}
```

