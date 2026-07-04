# 2. 基本数据类型

Go 语言中数据类型分为：基本数据类型和复合数据类型  

- 基本数据类型有：整型、浮点型、布尔型、字符串
- 复合数据类型有：数组、切片、结构体、函数、map、通道（channel）、接口等。

## 2.1 整型

|  类型  |                             范围                             | 占用空间 | 有无符号 |
| :----: | :----------------------------------------------------------: | :------: | :------: |
|  int8  |                (-128 到 127) -2^7^ 到 2^7^-1                 | 1 个字节 |    有    |
| int16  |             (-32768 到 32767) -2^15^ 到 2^15^-1              | 2 个字节 |    有    |
| int32  |        (-2147483648 到 2147483647) -2^31^ 到 2^31^-1         | 4 个字节 |    有    |
| int64  | (-9223372036854775808 到 9223372036854775807) -2^63^ 到 2^63^-1 | 8 个字节 |    有    |
| uint8  |                    (0 到 255) 0 到 2^8^-1                    | 1 个字节 |    无    |
| uint16 |                  (0 到 65535) 0 到 2^16^-1                   | 2 个字节 |    无    |
| uint32 |                (0 到 4294967295) 0 到 2^32^-1                | 4 个字节 |    无    |
| uint64 |           (0 到 18446744073709551615) 0 到 2^64^-1           | 8 个字节 |    无    |

> [!note]
>
> 实际项目中整数类型、切片、map 的元素数量等都可以用 int 来表示。在涉及到二进制传输、为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使
>
> 用 int 和 uint。

```go
var num int64
num = 123
fmt.Printf("值: %d, 类型: %T", num, num)
```

`unsafe.Sizeof(n1)` 是 `unsafe` 包的一个函数，可以返回变量占用的字节数  

```go
package main

import (
    "fmt"
    "unsafe"
) 

func main() {
    var a int8 = 120
    fmt.Printf("%T\n", a) // int8
    fmt.Println(unsafe.Sizeof(a)) // 1
}
```

## 2.2 浮点型

Go 语言支持两种浮点型数：`float32` 和 `float64`。打印浮点数时，可以使用 `fmt` 包配合动词 `%f`，代码如下：   

```go
package main

import (
    "fmt"
    "math"
) 

func main() {
    fmt.Printf("%f\n", math.Pi) //默认保留 6 位小数
    fmt.Printf("%.2f\n", math.Pi) //保留 2 位小数
}
```

> [!tip]
>
> Go 语言中浮点数默认是 `float64`

### 2.2.1 精度丢失问题

几乎所有的编程语言都有精度丢失这个问题，这是典型的二进制浮点数精度损失问题，在定长条件下，二进制小数和十进制小数互转可能有精度丢失。

```go
m1 := 8.2
m2 := 3.8
fmt.Println(m1 - m2) // 期望是 4.4， 结果打印出了 4.3999999999999995
```

> [!TIP]
>
> 可以使用第三方包来解决精度损失问题 https://github.com/shopspring/decimal

### 2.2.2 科学计数法表示浮点类型

```go
num1 := 1.234e3
num2 := 1.234e-3

fmt.Println(num1, num2)
```

## 2.3 布尔型

Go 语言中以 `bool` 类型进行声明布尔型数据， 布尔型数据只有 `true` 和 `false` 两个值。

> [!caution]
>
> 1. 布尔类型变量的默认值为 `false`。
> 2. Go 语言中不允许将整型强制转换为布尔型。
> 3. 布尔型无法参与数值运算，也无法与其他类型进行转换。

```go
package main

import (
    "fmt"
    "unsafe"
) 

func main() {
    var b = true
    fmt.Println(b, "占用字节：", unsafe.Sizeof(b)) // 占用字节：1
}
```

## 2.4 字符串

Go 语言里的字符串的内部实现使用 UTF-8 编码。字符串的值为双引号中的内容，可以在 Go 语言的源码中直接添加非 ASCII 码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

### 2.4.1 转义字符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

### 2.4.2 多行字符串

Go 语言中要定义一个多行字符串时，就必须使用反引号字符：

```go
s1 := `第一行
第二行
第三行
` 
fmt.Println(s1)
```

### 2.4.3 字符串的常用操作  

**求字符串的长度**  

```go
var str = "this is str"
fmt.Println(len(str))
```

**拼接字符串**  

```go
var str1 = "你好"
var str2 = "golang"
fmt.Println(str1 + str2)

var str3 = fmt.Sprintf("%v %v", str1, str2)
fmt.Println(str3)
```

**分割字符串**  

```go
var str = "123-456-789"
var arr = strings.Split(str, "-")
fmt.Println(arr)
```

**包含字符串**

```go
var str = "this is golang"
var flag = strings.Contains(str, "golang")
fmt.Println(flag)
```

**判断首字符尾字母是否包含指定字符**  

```go
var str = "this is golang"
var flag = strings.HasPrefix(str, "this")
fmt.Println(flag)

var str = "this is golang"
var flag = strings.HasSuffix(str, "go")
fmt.Println(flag)
```

**判断字符串出现的位置**

```go
var str = "this is golang"
var index = strings.Index(str, "is") //从前往后
fmt.Println(index)

var str = "this is golang"
var index = strings.LastIndex(str, "is") //从后往前
fmt.Println(index)
```

**`Join` 拼接字符串**

```go
var str = "123-456-789"
var arr = strings.Split(str, "-")
var str2 = strings.Join(arr, "*")
fmt.Println(str2)
```

## 2.5 字符型

组成每个字符串的元素叫做字符，可以通过遍历字符串元素获得字符。字符用单引号包裹起来，如：

```go
a := 'a'
b := '0'
//当我们直接输出 byte（字符） 的时候输出的是这个字符对应的码值
fmt.Println(a)
fmt.Println(b)
//如果我们要输出这个字符， 需要格式化输出
fmt.Printf("%c--%c", a, b) //%c 相应 Unicode 码点所表示的字符
```

> [!tip]
>
> 字节（byte）：是计算机中数据处理的基本单位，习惯上用大写 B 来表示，1B（byte）= 8bit（位）
>
> **一个汉子占用 3 个字节一个字母占用一个字节**  

```go
a := "m"
fmt.Println(len(a)) // 1

b := "张"
fmt.Println(len(b)) // 3
```

**Go 语言的字符有以下两种：**  

1. `byte` 或者叫 `uint8` 型， 代表了 ASCII 码的一个字符。  
2. `rune` 代表一个 UTF-8 字符，实际是一个 `int32`。  

当需要处理中文、日文或者其他复合字符时， 则需要用到 `rune` 类型。`rune` 类型实际是一个 `int32`。

```go
s := "hello 张三"
for i := 0; i < len(s); i++ { // byte
    fmt.Printf("%v(%c) ", s[i], s[i])
} 
// 104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 229(å) 188(¼) 160( ) 228(ä) 184(¸) 137() 

for _, r := range s { // rune
    fmt.Printf("%v(%c) ", r, r)
} 
// 104(h) 101(e) 108(l) 108(l) 111(o) 32( ) 24352(张) 19977(三) 
```

字符串底层是一个 `byte` 数组， 所以可以和 `[]byte` 类型相互转换。字符串是不能修改的字符串是由 `byte` 字节组成，所以字符串的长度是 `byte` 字节的长度。 rune 类型用来表示 UTF-8 字符， 一个 `rune` 字符由一个或多个 `byte` 组成。

## 2.6 修改字符串

要修改字符串，需要先将其转换成 `[]rune` 或 `[]byte`，完成后再转换为 `string`。无论哪种转换，都会重新分配内存，并复制字节数组。

```go
var s = "hello 张三"
byteStr := []byte(s)
byteStr[0] = 'H'
fmt.Println(string(byteStr)) // Hello 张三

runeStr := []rune(s)
runeStr[len(runeStr)-1] = '五'
fmt.Println(string(runeStr)) // hello 张五
```

## 2.7 数据类型转化

**Go 语言中只有强制类型转换，没有隐式类型转换。**  

### 2.7.1 数值类型之间的相互转换  

```go
var a int8 = 20
var b int16 = 40
var c = int16(a) + b           // 要转换成相同类型才能运行
fmt.Printf("值：%v，类型：%T", c, c) // 值：60，类型：int16
```

```go
var d float32 = 3.2
var e int16 = 6
var f = d + float32(e)
fmt.Printf("值：%v，类型：%T", f, f) // 值：9.2，类型：float32
```

> [!tip]
>
> 转换的时候建议从低位转换成高位，高位转换成低位的时候如果转换不成功就会溢出，和我们想的结果不一样

### 2.7.2 其他类型转换成 String 类型

使用 `fmt.Sprintf` 把其他类型转换成 string 类型  

```go
package main

import "fmt"

func main() {
	var i int = 10
	var f float32 = 3.14
	var c byte = 'a'
	var b bool = true

	var s string

	s = fmt.Sprintf("%d", i)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%f", f)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%t", b)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%c", c)
	fmt.Printf("类型 %T, %v \n", s, s)
}
```

```go
类型 string, 10 
类型 string, 3.140000 
类型 string, true 
类型 string, a 
```

**使用 `strconv` 包里面的几种转换方法进行转换**  

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1、 int 转换成 string
	var num1 int = 20
	s1 := strconv.Itoa(num1)
	fmt.Printf("类型 %T, %v \n", s1, s1)

	// 2、 float 转 string
	var num2 float64 = 20.113123
	s2 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("类型 %T, %v \n", s2, s2)

	// 3、 bool 转 string
	s3 := strconv.FormatBool(true)
	fmt.Printf("类型 %T, %v \n", s3, s3)

	// 4、 int64 转 string
	var num3 int64 = 20
	s4 := strconv.FormatInt(num3, 10)
	fmt.Printf("类型 %T ,strs=%v \n", s4, s4)
}
```

### 2.7.3 String 类型转换成数值类型

**使用 `strconv` 包里面的几种转换方法进行转换**  

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 1、string 类型转 int
	var s1 string = "1234"
	i, _ := strconv.ParseInt(s1, 10, 64)
	fmt.Printf("值：%d，类型：%T\n", i, i)

	// 2、string 转 float
	var s2 string = "3.1415926"
	f, _ := strconv.ParseFloat(s2, 64)
	fmt.Printf("值：%f，类型：%T\n", f, f)

	// 3、string 转 bool
	var s3 string = "true"
	b, _ := strconv.ParseBool(s3)
	fmt.Printf("值：%t，类型：%T\n", b, b)

	// 4、string 转 byte
	var s4 string = "abc"
	for _, r := range s4 {
		fmt.Printf("值：%c，类型：%T\n", r, r)
	}
}
```

> [!caution]
>
> 在 go 语言中数值类型没法直接转换成 `bool` 类型 `bool` 类型也没法直接转换成数值类型

