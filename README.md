# Go Language Tutorial

## 1. 变量和常量

Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。Go 语言中关键字和保留字都不能用作变量名。

Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明。并且 Go 语言的变量声明后必须使用。

### 1.1 变量

#### 1.1.1 var 声明变量

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

#### 1.1.2 一次定义多个变量

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

#### 1.1.3 批量声明变量的时候指定类型

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

#### 1.1.4 变量的初始化

Go 语言在声明变量的时候，会自动对变量对应的内存区域进行初始化操作。每个变量会被初始化成其类型的默认值，例如：整型和浮点型变量的默认值为 `0`。字符串变量的默认值为空字符串。布尔型变量默认为 `false`。切片、函数、指针变量的默认为 `nil`。

当然我们也可在声明变量的时候为其指定初始值。 变量初始化的标准格式如下：

```go
var name string = "zhangsan"
var age int = 18
```

#### 1.1.5 类型推导

有时候我们会将变量的类型省略，这个时候编译器会根据等号右边的值来推导变量的类型完成初始化。

```go
var name = "zhangsan"
var age = 18
```

#### 1.1.6 短变量声明法

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

#### 1.1.7 匿名变量

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

### 1.2 常量

相对于变量，常量是恒定不变的值，多用于定义程序运行期间不会改变的那些值。常量的声明和变量声明非常类似，只是把 `var` 换成了 `const`，**常量在定义的时候必须赋值**。

#### 1.2.1 使用 const 定义常量

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

### 1.3 变量、常量命名规则

1. 变量名称必须由数字、字母、下划线组成。
2. 标识符开头不能是数字。
3. 标识符不能是保留字和关键字。
4. 变量的名字是区分大小写的如: `age` 和 `Age` 是不同的变量。在实际的运用中，也建议，不要用一个单词大小写区分两个变量。
5. 变量名称一定要见名思意：变量名称建议用名词，方法名称建议用动词。
6. 变量命名一般采用驼峰式，当遇到特有名词（缩写或简称，如 DNS）的时候，特有名词根据是否私有全部大写或小写。



## 2. 基本数据类型

Go 语言中数据类型分为：基本数据类型和复合数据类型  

- 基本数据类型有：整型、浮点型、布尔型、字符串
- 复合数据类型有：数组、切片、结构体、函数、map、通道（channel）、接口等。

### 2.1 整型

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

### 2.2 浮点型

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

#### 2.2.1 精度丢失问题

几乎所有的编程语言都有精度丢失这个问题，这是典型的二进制浮点数精度损失问题，在定长条件下，二进制小数和十进制小数互转可能有精度丢失。

```go
m1 := 8.2
m2 := 3.8
fmt.Println(m1 - m2) // 期望是 4.4， 结果打印出了 4.3999999999999995
```

> [!TIP]
>
> 可以使用第三方包来解决精度损失问题 https://github.com/shopspring/decimal

#### 2.2.2 科学计数法表示浮点类型

```go
num1 := 1.234e3
num2 := 1.234e-3

fmt.Println(num1, num2)
```

### 2.3 布尔型

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

### 2.4 字符串

Go 语言里的字符串的内部实现使用 UTF-8 编码。字符串的值为双引号中的内容，可以在 Go 语言的源码中直接添加非 ASCII 码字符，例如：

```go
s1 := "hello"
s2 := "你好"
```

#### 2.4.1 转义字符

Go 语言的字符串常见转义符包含回车、换行、单双引号、制表符等，如下表所示。

| 转义符 |                含义                |
| :----: | :--------------------------------: |
|  `\r`  |         回车符（返回行首）         |
|  `\n`  | 换行符（直接跳到下一行的同列位置） |
|  `\t`  |               制表符               |
|  `\'`  |               单引号               |
|  `\"`  |               双引号               |
|  `\\`  |               反斜杠               |

#### 2.4.2 多行字符串

Go 语言中要定义一个多行字符串时，就必须使用反引号字符：

```go
s1 := `第一行
第二行
第三行
` 
fmt.Println(s1)
```

#### 2.4.3 字符串的常用操作  

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

### 2.5 字符型

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

### 2.6 修改字符串

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

### 2.7 数据类型转化

**Go 语言中只有强制类型转换，没有隐式类型转换。**  

#### 2.7.1 数值类型之间的相互转换  

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

#### 2.7.2 其他类型转换成 String 类型

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

#### 2.7.3 String 类型转换成数值类型

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

## 3. 运算符

### 3.1 算术运算符

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

### 3.2 关系运算符

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

### 3.3 逻辑运算符

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

## 4. 流程控制

### 4.1 if else

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

### 4.2 for

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

### 4.3 for range

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

### 4.4 switch case

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

### 4.5 break

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

### 4.6 continue

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

### 4.7 goto

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

## 5. 数组

数组是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素，这种类型可以是任意的原始类型，比如 int、 string 等，也可以是用户自定义的类型。 

一个数组包含的元素个数被称为数组的长度。 

在 Golang 中数组是一个长度固定的数据类型，数组的长度是类型的一部分，也就是说 `[5]int` 和 `[10]int` 是两个不同的类型。 

### 5.1 数组的声明

```go
var a [3]int
var b [4]int
```

数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，**长度不能变**。数组可以通过下标进行访问，下标是从 `0` 开始，最后一个元素下标是 `len-1`  

### 5.2 数组的初始化

初始化数组时可以使用初始化列表来设置数组元素的值。

```go
var testArray [5]int
var numArray = [3]int{1, 2, 3}
var cityArray = [3]string{"北京", "上海", "广州"}

fmt.Println(testArray) // [0 0 0 0 0]
fmt.Println(numArray) // [1 2 3]
fmt.Println(cityArray) // [北京 上海 广州]
```

除此之外，还可以让编译器根据初始值的个数自行推断数组的长度。

```go
var testArray = [...]int{}
var numArray = [...]int{1, 2, 3}
var cityArray = [...]string{"北京", "上海", "广州"}

fmt.Println(testArray)
fmt.Println(numArray)
fmt.Println(cityArray)
```

我们还可以使用指定索引值的方式来初始化数组。

```go
a := [...]int{1: 1, 3: 5}
fmt.Println(a) // [0 1 0 5]
fmt.Printf("type of a:%T\n", a) // type of a:[4]int
```

### 5.3 数组的遍历

```go
// 1. for 循环遍历
for i := 0; i < len(array); i++ {
    fmt.Println(array[i])
}

// 2. for range 遍历
for _, v := range array {
    fmt.Println(v)
}
```

### 5.4 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
    a := [3]int{10, 20, 30}
    modifyArray(a) // 在 modify 中修改的是 a 的副本 x
    fmt.Println(a) // [10 20 30]
    
    b := [3][2]int{
        {1, 1},
        {1, 1},
        {1, 1},
    } 
    modifyArray2(b) // 在 modify 中修改的是 b 的副本 x
    fmt.Println(b) // [[1 1] [1 1] [1 1]]
}
```

> [!note]
>
> 1. 数组支持 `==`、 `!=` 操作符，因为内存总是被初始化过的。
> 2. `[n]*T` 表示指针数组，`*[n]T` 表示数组指针。

## 6. 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

**切片是一个引用类型，它的内部结构包含地址、长度和容量。**  

### 6.1 切片的声明

```go
var a []string
var b []int = []int{1, 2, 3}

fmt.Println(a == nil) //  true
fmt.Println(b == nil) //  false
```

> [!note]
>
> 当你声明了一个切片，却没有给其赋值，此时该切片为默认值 `nil`

> [!important]
>
> 当你声明了一个变量，但却还并没有赋值时，Golang 中会自动给你的变量赋值一个默认零值。 

| 类型类别     | 具体类型                                                     | 零值（默认值） |
| :----------- | :----------------------------------------------------------- | :------------- |
| **布尔型**   | `bool`                                                       | `false`        |
| **整型**     | `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0`            |
| **浮点型**   | `float32`, `float64`                                         | `0.0`          |
| **复数型**   | `complex64`, `complex128`                                    | `(0+0i)`       |
| **字符串**   | `string`                                                     | `""`           |
| **数组**     | `[n]T`                                                       | `[n]T{}`       |
| **切片**     | `[]T`                                                        | `nil`          |
| **映射**     | `map[K]V`                                                    | `nil`          |
| **通道**     | `chan T`                                                     | `nil`          |
| **指针**     | `*T`                                                         | `nil`          |
| **函数**     | `func(...)`                                                  | `nil`          |
| **接口**     | `interface{}`                                                | `nil`          |
| **结构体**   | `struct{}`                                                   | 各字段的零值   |
| **字节类型** | `byte` (`uint8` 别名)                                        | `0`            |
| **符文类型** | `rune` (`int32`别名)                                         | `0`            |
| **错误类型** | `error`                                                      | `nil`          |

### 6.2 切片的遍历

切片的遍历与数组类似，可以使用 `for` 或者 `for range` 遍历。

```go
var strSlice []string = []string{"hello", "world", "golang"}

// 1. 使用 for 遍历切片
for i := 0; i < len(strSlice); i++ {
    fmt.Println(strSlice[i])
}

// 2. 使用 for range 遍历切片
for _, v := range strSlice {
    fmt.Println(v)
}
```

### 6.3 基于数组定义切片

由于切片的底层就是一个数组，所以我们可以基于数组定义切片。  

```go
numArray := [5]int{1, 2, 3, 4, 5}
numSlice := numArray[1:3] // 基于数组创建切片
fmt.Println(numSlice) // [2 3]
```

### 6.4 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的 `len()` 函数求长度，使用内置的 `cap()` 函数求切片的容量。

- 切片的长度就是**它所包含的元素个数**。  
- 切片的容量是从**它的第一个元素开始数，到其底层数组元素末尾的个数**。

```go
s := []int{2, 3, 5, 7, 11, 13}
fmt.Println(s)
fmt.Printf("长度：%v，容量：%v\n", len(s), cap(s))

c := s[:2]
fmt.Println(c)
fmt.Printf("长度：%v，容量：%v\n", len(c), cap(c))

d := s[1:3]
fmt.Println(d)
fmt.Printf("长度：%v，容量：%v\n", len(d), cap(d))

// 输出：
// [2 3 5 7 11 13]
// 长度：6，容量：6
// [2 3]
// 长度：2，容量：6
// [3 5]
// 长度：2，容量：5
```

### 6.5 使用 make() 函数声明

如果需要动态的创建一个切片，我们就需要使用内置的 `make()` 函数。

```go
make([]T, size, cap)
```

其中：

- `[]T`：切片的元素类型
- `size`：切片中的元素数量
- `cap`：切片的容量

```go
var a = make([]int, 2, 10)
fmt.Println(a) // [0 0]
fmt.Println(len(a)) // 2
fmt.Println(cap(a)) // 10
```

> [!caution]
>
> 切片之间是不能比较的，我们不能使用 `==` 操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较操作是和 `nil` 比较。 

### 6.6 切片的基础操作

#### 6.6.1 append() 方法

当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行扩容，此时该切片指向的底层数组就会更换。扩容操作往往发生在`append()` 函数调用时，所以我们通常都需要用原变量接收 `append()` 函数的返回值。

错误写法：

```go
s := []int{1, 2, 3, 5, 6, 7}
s[6] = 8
fmt.Println(s) //index out of range [6] with length 6
```

正确写法：

```go
var numSlice []int
for i := 0; i < 10; i++ {
    numSlice = append(numSlice, i)
}
```

1. `append()` 函数将元素追加到切片的最后并返回该切片。  
2. 切片 `numSlice` 的容量按照 1，2，4，8，16 这样的规则自动进行扩容，每次扩容后都是扩容前的 2 倍。

```go
var citySlice []string
// 添加一个元素
citySlice = append(citySlice, "北京")

// 添加多个元素
citySlice = append(citySlice, "上海", "广州", "深圳")

// 添加切片
citySlice = append(citySlice, []string{"杭州", "苏州"}...)
fmt.Println(citySlice) // [北京 上海 广州 深圳 杭州 苏州]
```

#### 6.6.2 copy() 方法

```go
a := []int{1, 2, 3, 4, 5}
b := a
fmt.Println(a) //[1 2 3 4 5]
fmt.Println(b) //[1 2 3 4 5]
b[0] = 1000
fmt.Println(a) //[1000 2 3 4 5]
fmt.Println(b) //[1000 2 3 4 5]
```

由于切片是引用类型，所以在修改 b 的同时 a 的值也会发生变化。Go 语言内建的 `copy()` 函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()` 函数的使用格式如下：

```go
copy(destSlice, srcSlice []T)
```

其中：

- `destSlice`：目标切片
- `srcSlice`：数据来源切片

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := make([]int, 5)

copy(s2, s1) // 将 s1 拷贝到 s2
fmt.Println(s1) // [1 2 3 4 5]
fmt.Println(s2) // [1 2 3 4 5]

s2[0] = 100
fmt.Println(s1) // [1 2 3 4 5]
fmt.Println(s2) // [100 2 3 4 5]
```

#### 6.6.3 删除元素

Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。要从切片 a 中删除索引为 index 的元素，操作方法是

```go 
a = append(a[:index], a[index+1:]...)
```

```go
func popString(srcSlice []string, index int) []string {
	// 删除切片中的元素
	return append(srcSlice[:index], srcSlice[index+1:]...)
}
```

### 6.7 切片排序

#### 6.7.1 冒泡排序

核心思想：**相邻比较和交换**，就像水中气泡上浮一样，它通过反复比较相邻元素，将较大的元素逐步向后移动，每一轮都会将当前未排序部分的最大值“冒泡”到正确位置，直到整个序列有序。

```go
func bubbleSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 冒泡排序
	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟冒泡
		isSwap := false
		for j := 0; j < len(numSlice)-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				numSlice[j], numSlice[j+1] = numSlice[j+1], numSlice[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return numSlice
}
```

#### 6.7.2 选择排序

核心思想：**选取最小元素**，它把序列分为已排序和未排序两部分，每次从未排序部分中选出最小的元素，直接将其放到已排序部分的末尾，如同打擂台选冠军一样逐个确定位置，直到所有元素归位。

```go
func selecSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟选择
		minIndex := i
		// 找到最小值的索引
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[j] < numSlice[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
            // 将未排序区间最小值交换到已排序区间的末尾
			numSlice[i], numSlice[minIndex] = numSlice[minIndex], numSlice[i]
		}
	}

	return numSlice
}
```

#### 6.7.3 插入排序

核心思想：**构建有序序列**，就像整理手中的扑克牌，它将序列视为已排序和未排序两部分，每次从未排序部分取出第一个元素，在已排序部分中从后向前扫描，找到合适位置插入，从而逐步扩大有序区间，直到全部有序。

```go
func insertSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 遍历无序区间，从第二个元素开始遍历
	for i := 1; i < len(numSlice); i++ {
		// 保存当前元素
		current := numSlice[i]
		// 找到当前元素应该插入的位置
		j := i - 1

		// 如果 j >= 0 且 numSlice[j] > current，
        // 则说明当前元素在有序区间中，则将当前元素插入有序区间中
		for j >= 0 && numSlice[j] > current {
			numSlice[j+1] = numSlice[j]
			j--
		}
		numSlice[j+1] = current
	}
	return numSlice

}
```

#### 6.7.4 sort 包

对于 `int`、 `float64` 和 `string` 数组或是切片的排序， go 分别提供了 `sort.Ints()`、`sort.Float64s()` 和 `sort.Strings()` 函数，默认都是从小到大排序。

```go
intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

sort.Ints(intList)
sort.Float64s(floatList)
sort.Strings(stringList)
fmt.Println(intList) // [0 1 2 3 4 5 6 7 8 9]
fmt.Println(floatList) // [3.14 4.2 5.9 10.2 12.4 27.81828 31.4 50.7 99.9]
fmt.Println(stringList) // [a b c d f i w x y z]
```

Golang 的 `sort` 包可以使用 `sort.Reverse(slice)` 来调换 `slice.Interface.Less`，所以逆序排序函数可以这么写。

```go
sort.Sort(sort.Reverse(sort.IntSlice(intList)))
fmt.Println(intList) // [9 8 7 6 5 4 3 2 1 0]
```

## 7. 哈希表

### 7.1 哈希表的声明

Go 语言的哈希表可以通过 `map` 实现，`map` 是一种无序的基于 key-value 的数据结构，Go 语言中的 `map` 是引用类型，必须初始化才能使用。

```go
map[KeyType]ValueType
```

其中:

- `KeyType`:表示键的类型。
- `ValueType`:表示键对应的值的类型。

> [!note]
>
> `map` 类型的变量默认初始值为 `nil`，需要使用 `make()` 函数来分配内存。 

```go
make(map[KeyType]ValueType, [cap])
```

> [!caution]
>
> 获取 `map` 的容量不能使用 `cap`，`cap` 返回的是数组切片分配的空间大小，根本不能用于 `map`。要获取 `map`  的容量，可以用 `len()` 函数。

`map` 中的数据都是成对出现的，`map` 的基本使用示例代码如下：

```go
hashMap := make(map[string]int)
hashMap["a"] = 1
hashMap["b"] = 2
hashMap["c"] = 3
fmt.Println(hashMap) // map[a:1 b:2 c:3]
fmt.Println(hashMap["a"]) // 1
```

`map` 也支持在声明的时候填充元素，例如：

```go
userInfo := map[string]string{
    "username": "zhangsan",
    "password": "123456",
}
fmt.Println(userInfo) // map[password:123456 username:zhangsan]

```

### 7.2 判断键是否存在

Go 语言中有个判断 `map` 中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80

// 如果 key 存在 ok 为 true,v 为对应的值；
// 不存在 ok 为 false,v 为值类型的零值
v, ok := scoreMap["张三"]
if ok {
    fmt.Println(v) // 90
} else {
    fmt.Println("没有该用户")
}
```

### 7.3 哈希表的遍历

Go 语言中使用 `for range` 遍历 `map`。  

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80
scoreMap["王五"] = 70

for k, v := range scoreMap {
    fmt.Println(k, v)
}
```

我们只想遍历 key 的时候，可以按下面的写法：

```go
for k := range scoreMap {
    fmt.Println(k)
}
```

> [!caution]
>
> 遍历 `map` 时的元素顺序与添加键值对的顺序无关。  

### 7.4 delete() 方法

使用 `delete()` 内建函数从 `map` 中删除一组键值对，`delete()` 函数的格式如下：

```go
delete(map, key)
```

其中：

- `map`：表示要删除键值对的 `map` 对象
- `key`：表示要删除的键值对的键

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80
scoreMap["王五"] = 70
fmt.Println(scoreMap) // map[张三:90 李四:80 王五:70]

// delete 方法删除
delete(scoreMap, "张三")
fmt.Println(scoreMap) // map[李四:80 王五:70]
```

## 8. 函数

函数是组织好的、可重复使用的、用于执行指定任务的代码块。Go 语言中支持：函数、匿名函数和闭包。

### 8.1 函数的声明

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

### 8.2 函数的调用

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

### 8.3 函数参数

#### 8.3.1 类型简写 

函数的参数中如果相邻变量的类型相同，则可以省略类型，例如：

```go
func intSum(x, y int) int {
	return x + y
}
```

#### 8.3.2 可变参数

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

### 8.4 函数返回值

Go 语言中通过 `return` 关键字向外输出返回值。

#### 8.4.1 函数多返回值  

Go 语言中函数支持多返回值，函数如果有多个返回值时必须用 `()` 将所有返回值包裹起来。

```go
func calc(x, y int) (int, int) {
    sum := x + y
    sub := x - y
    return sum, sub
}
```

#### 8.4.2 返回值命名

函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 `return` 关键字返回。

```go
func calc(x, y int) (sum, sub int) {
    sum = x + y
    sub = x - y
    return
}
```

### 8.5 函数类型与变量

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

### 8.6 高级函数

#### 8.6.1 函数作为参数  

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

#### 8.6.2 函数作为返回值

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

### 8.7 匿名函数和闭包

#### 8.7.1 匿名函数

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

#### 8.7.2 闭包

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

### 8.8 defer

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

#### 8.8.1 defer 执行时机  

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

#### 8.8.2 defer 示例

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

### 8.9 panic()/recover()

Go 语言使用 `panic()/recover()` 模式来处理错误。`panic()` 可以在任何地方引发，但 `recover()` 只有在 `defer` 调用的函数中有效。 

#### 8.9.1 panic()/recover() 的基本使用

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

#### 8.9.2 defer/recover() 实现异常处理

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

#### 8.9.3 defer/panic()/recover() 抛出异常

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

## 9. 指针

指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另一个变量的内存地址。

Go 语言中的指针操作非常简单，我们只需要记住两个符号：`&`（取地址）和 `*`（根据地址取值）

### 9.1 指针地址和指针类型  

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

### 9.2 指针取值

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

### 9.3 new()/make()

在 Go 语言中对于引用类型的变量，在使用时不仅要声明它，还要为它分配内存空间。而对于值类型的声明不需要分配内存空间，是因为它们在声明的时候已经默认分配好了内存空间。 Go 语言中 `new()` 和 `make()` 是内建的两个函数，主要用来分配内存。

#### 9.3.1 new() 函数分配内存

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

#### 9.3.2 make() 函数分配内存  

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

#### 9.3.3 new() 与 make() 的区别 

1. 二者都是用来做内存分配的。
2. `make()` 只用于 `slice`、`map` 以及 `channel` 的初始化，返回的还是这三个引用类型本身。
3. 而 `new()` 用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。

## 10. 结构体

Golang 中没有类的概念，Golang 中的结构体和其他语言中的类有点相似。和其他面向对象语言中的类相比，Golang 中的结构体具有更高的扩展性和灵活性。

### 10.1 自定义类型和类型别名

#### 10.1.1 自定义类型

在 Go 语言中有一些基本的数据类型，如 `string`、`int`、`float`、`bool`等数据类型，Go 语言中可以使用 `type` 关键字来定义自定义类型。

```go
type myInt int
```

#### 10.1.2 类型别名  

`TypeAlias` 只是 `Type` 的别名，本质上 `TypeAlias` 与 `Type` 是同一个类型。 

```go
type TypeAlias = Type
```

`rune` 和 `byte` 就是类型别名，他们的底层定义如下  

```go
type byte = uint8
type rune = int32
```

#### 10.1.3 自定义类型和类型别名的区别  

```go
package main

import "fmt"

// 类型定义
type newInt int

// 类型别名
type myInt = int

func main() {
	var a newInt
	var b myInt
	fmt.Printf("type of a:%T\n", a) // type of a:main.newInt
	fmt.Printf("type of b:%T\n", b) // type of b:int
}

```

### 10.2 结构体的初始化

#### 10.2.1 结构体的声明

使用 `type` 和 `struct` 关键字来定义结构体，具体代码格式如下：

```go
type structName struct {
    fieldName fieldType
    ...
}
```

其中：

- `structName`：类型名，表示自定义结构体的名称，在同一个包内不能重复。  
- `fieldName`：字段名，表示结构体字段名。结构体中的字段名必须唯一。  
- `fieldType`：字段类型，表示结构体字段的具体类型。  

```go
type person struct {
    name string
    city string
    age int8
}
```

> [!caution]
>
> 结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，在其他的包里面可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用。

#### 10.2.2 结构体实例化

只有当结构体实例化时，才会真正地分配内存。也就是必须实例化后才能使用结构体的字段。可以使用 `var` 关键字声明结构体类型。  

```go
package main

import "fmt"

type person struct {
    name string
    city string
    age int
} 

func main() {
    var p1 person
    p1.name = "张三"
    p1.city = "北京"
    p1.age = 18
    fmt.Printf("p1=%v\n", p1) // p1={张三 北京 18}
    fmt.Printf("p1=%#v\n", p1) // p1=main.person{name:"张三", city:"北京", age:18}
}
```

我们还可以通过使用 `new()` 对结构体进行实例化，得到的是结构体的地址。格式如下：

```go
package main

import "fmt"

type person struct {
    name string
    city string
    age int
} 

func main() {
    var p2 = new(person)
    p2.name = "张三"
    p2.age = 20
    p2.city = "北京"
    fmt.Printf("%T\n", p2) // *main.person
    fmt.Printf("p2=%#v\n", p2) // p2=&main.person{name:"张三", city:"北京", age:20}
}
```

> [!caution]
>
> 在 Golang 中支持对结构体指针直接使用 `.` 来访问结构体的成员。`p2.name = "张三"` 其实在底层是 `(*p2).name = "张三"`

也可以通过键值对初始化的方法。格式如下：

```go
package main

import "fmt"

type person struct {
    name string
    city string
    age int
} 

func main() {
    var p3 = &person{
        name: "zhangsan",
        city: "北京",
        age: 18,
	} 
    fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"zhangsan", city:"北京", age:18}
}
```

### 10.3 结构体方法和接收者  

在 Go 语言中，没有类的概念但是可以给类型（结构体，自定义类型）定义方法。所谓方法就是定义了接收者的函数。 

接收者的概念就类似于其他语言中的 `this` 或者 `self`。

```go
func (ReceiveVar ReceiveType) funcName (para Type) (return Type) {
	funcBody
}
```

其中：

- `ReceiveVar`：接收者变量，接收者中的参数变量名在命名时，官方建议使用接收者类型名的第一个小写字母，而不是`self`、`this` 之类的命名。例如，`Person` 类型的接收者变量应该命名为 `p`，`Connector` 类型的接收者变量应该命名为 `c` 等。
- `ReceiveType`：接收者类型，接收者类型和参数类似，可以是指针类型和非指针类型。

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func (p Person) PrintInfo() {
	fmt.Println("name:", p.name)
	fmt.Println("age:", p.age)
	fmt.Println("sex:", p.sex)
}

func main() {
	var p1 Person
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"

	p1.PrintInfo()
}

// 输出：
// name: 张三
// age: 18
// sex: 男
```

**值类型作为接收者时**，Go 语言会在代码运行时将接收者的值复制一份。在值类型接收者的方法中可以获取接收者的成员值， 但修改操作只是针对副本，无法修改接收者变量本身。

**指针类型作为接收者时**，由于指针的特性，调用方法时修改接收者指针的任意成员变量，在方法结束后，修改都是有效的。这种方式就十分接近于其他语言中面向对象中的 `this` 或者 `self`。

```go
package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

// PrintInfo 值类型接受者
func (p Person) PrintInfo() {
	fmt.Printf("Person %v, %d, %v\n", p.name, p.age, p.sex)
}

// SetInfo 指针类型接受者
func (p *Person) SetInfo(name string, age int, sex string) {
	p.name = name
	p.age = age
	p.sex = sex
}

func main() {
	var p1 Person
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"

	p1.PrintInfo() // Person 张三, 18, 男
	p1.SetInfo("李四", 19, "女")
	p1.PrintInfo() // Person 李四, 19, 女
}
```

> [!caution]
>
> Go 文档不推荐使用结构体在值接收器和指针接收器上都有方法。 

### 10.4 给任意类型添加方法  

在 Go 语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。

```go
package main

import "fmt"

type myInt int

func (i myInt) SayHello() {
	fmt.Printf("Hello, World!\n")
}

func main() {
	var m1 myInt
	m1.SayHello() // Hello, World!
	m1 = 100
	fmt.Printf("%#v, %T\n", m1, m1) // 100, main.myInt
}
```

> [!caution]
>
> 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法。  

### 10.5 结构体的匿名字段  

结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段。

```go
//Person 结构体 Person 类型
type Person struct {
    string
    int
} 

func main() {
    p1 := Person{
        "小王子",
        18,
    } 
    fmt.Printf("%#v\n", p1) // main.Person{string:"北京", int:18}
	fmt.Println(p1.string, p1.int) // 北京 18
}
```

匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个。

### 10.6 嵌套结构体  

一个结构体中可以嵌套包含另一个结构体或结构体指针。

```go
type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

func main() {
	user1 := User{
		Name:   "张三",
		Gender: "男",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("%#v\n", user1)
}

// 输出：
// main.User{Name:"张三", Gender:"男",
// Address:main.Address{Province:"广东", City:"深圳"}}
```

### 10.7 结构体的继承

Go 语言中使用结构体也可以实现其他编程语言中的继承。

```go
type Animal struct {
	name string
}

func (a Animal) Eat() {
	fmt.Printf("Animal %s eats!\n", a.name)
}

type Dog struct {
	Age int
	Animal
}

func main() {
	dog1 := Dog{
		Age: 1,
		Animal: Animal{
			name: "旺财",
		},
	}
	dog1.Eat()
}

// 输出：
// Animal 旺财 eats!
```

## 11. JSON

### 11.1 JSON 序列化和反序列化

Golang 中的序列化和反序列化主要通过 `"encoding/json"` 包中的 `json.Marshal()` 和 `json.Unmarshal()` 方法实现。

#### 11.1.1 序列化

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	name   string // 私有属性不能被 json 包访问
	Age    int
	Gender string
}

func main() {
	var s Student = Student{
		ID:     1,
		name:   "张三", // 私有属性不能被 json 包访问
		Age:    18,
		Gender: "男",
	}

	sByte, err := json.Marshal(s)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	sStr := string(sByte)
	fmt.Println(sStr) // {"ID":1,"Age":18,"Gender":"男"}
}
```

#### 11.1.2 反序列化

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int
	Name   string
	Age    int
	Gender string
}

func main() {
	var jsonStr = `{"ID":1,"Name":"李四","Age":18,"Gender":"男"}`
	var student Student
    err := json.Unmarshal([]byte(jsonStr), &student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("student:%#v\n", student) // student:main.Student{ID:1, Name:"李四", Age:18, Gender:"男"}
}
```

### 11.2 结构体标签 Tag

Tag 是结构体的元信息，可以在运行的时候通过反射的机制读取出来。Tag 在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

```go
`key1 : "value1"`
```

> [!caution]
>
> 为结构体编写 Tag 时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在 key 和 value 之间添加空格。

```go
type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}
```

### 11.3 嵌套结构体的 JSON

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

func main() {
	c := &Class{
		Name:     "1班",
		Students: make([]Student, 0, 200),
	}

	for i := 0; i < 10; i++ {
		stu := Student{
			ID:     i + 1,
			Name:   fmt.Sprintf("学号%d", i+1),
			Age:    18,
			Gender: "男",
		}
		c.Students = append(c.Students, stu)
	}

	// 序列化
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("class:%#v\n", c)
	fmt.Printf("data:%#v\n", string(data))
}
```

## 12. 包

包（package）是多个 Go 源码的集合，是一种高级的代码复用方案。Golang 中的包可以分为三种：1、系统内置包；2、自定义包；3、第三方包。

### 12.1 包管理工具 go mod

Go1.11 版本之后无需手动配置环境变量，使用 `go mod` 管理项目，也不需要非得把项目放到 `$GOPATH` 指定目录下，你可以在你磁盘的任何位置新建一个项目, Go1.13 以后可以彻底不要 `$GOPATH` 了。

#### 12.1.1 go mod init 初始化项目  

实际项目开发中我们首先要在我们项目目录中用 `go mod` 命令生成一个 go.mod 文件管理我们项目的依赖。

```sh
go mod init demo21
```

#### 12.1.2 其他命令

| 命令              | 用途             | 常用参数     |
| ----------------- | ---------------- | ------------ |
| `go mod init`     | 初始化模块       |              |
| `go mod tidy`     | 整理依赖         |              |
| `go mod download` | 下载依赖         |              |
| `go mod graph`    | 查看依赖图       |              |
| `go mod why`      | 查看依赖原因     | `-m`模块级别 |
| `go mod vendor`   | 创建 vendor 目录 |              |

### 12.2 自定义包

包（package）是多个 Go 源码的集合，一个包可以简单理解为一个存放多个 `.go` 文件的文件夹。该文件夹下面的所有 go 文件都要在代码的第一行添加如下代码，声明该文件归属的包。

```go
package Name
```

> [!caution]
>
> - 一个文件夹下面直接包含的文件只能归属一个 `package`， 同样一个 `package` 的文件不能在多个文件夹下。
> - 包名可以不和文件夹的名字一样， 包名不能包含 `-` 符号。
> - 包名为 `main` 的包为应用程序的入口包，这种包编译后会得到一个可执行文件，而编译不包含 `main` 包的源代码则不会得到可执行文件。

#### 12.2.1 定义包

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

#### 12.2.2 导入包

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

### 12.3 init() 初始化函数  

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

### 12.4 第三方包

我们可以在 https://pkg.go.dev/ 查找看常见的 Golang 第三方包。

#### 12.4.1 初始化项目

```sh
go mod init project
```

#### 12.4.2 下载安装这个包（非必须）  

例如要下载解决 `float` 精度损失的包 `decimal`。

```sh
go get github.com/shopspring/decimal
```

#### 12.4.3 go mod tidy 下载丢失的包  

`go mod tidy` 命令可以增加丢失的 `module`，同时去掉未用的 `module`（推荐）

```sh
go mod tidy
```

## 13. 接口

### 13.1 接口的声明

在 Golang 中接口（interface）是一种类型，一种抽象的类型。接口（interface）是一组函数 method 的集合，**Golang 中的接口不能包含任何变量**。

在 Golang 中接口中的所有方法都没有方法体，接口定义了一个对象的行为规范，只定义规范不实现。接口体现了程序设计的多态和高内聚低耦合的思想。

接口的定义格式如下：

```go
type interfacName interface{
    func1 (para) return 1
    func2 (para) return 2
    ...
}
```

其中：

- `interfacName`：接口名，使用 `type` 将接口定义为自定义的类型名。Go 语言的接口在命名时，一般会在单词后面添加 er，如有写操作的接口叫 Writer，有字符串功能的接口叫 Stringer 等。接口名最好要能突出该接口的类型含义。
- `func`：当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。
- `para、return`：参数列表、返回值列表，参数列表和返回值列表中的参数变量名可以省略。  

```go
package main

import "fmt"

type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "相机停止工作")
}

func main() {
	// 创建手机对象
	phone := Phone{"小米"}
	var usbPhone Usber = phone // 表示 phone 实现了 Usber 接口
	usbPhone.Start() // 小米 手机开始工作
	usbPhone.Stop()  // 小米 手机停止工作

	camera := Camera{"尼康"} 
	var usbCamera Usber = camera // 表示 camera 实现了 Usber 接口
	usbCamera.Start() // 尼康 相机开始工作
	usbCamera.Stop()  // 尼康 相机停止工作
}
```

> [!caution]
>
> Golang 中的接口也是一种数据类型，不需要显示实现。**只需要一个变量含有接口类型中的所有方法，那么这个变量就实现了这个接口。**

### 13.2 空接口

Golang 中的接口可以不定义任何方法，没有定义任何方法的接口就是空接口。**空接口表示没有任何约束，因此任何类型变量都可以实现空接口。用空接口可以表示任意数据类型。**

```go
func main() {
    // 定义一个空接口 x, x 变量可以接收任意的数据类型
    var x interface{}
    s := "你好 golang"
    x = s
    fmt.Printf("type:%T value:%v\n", x, x)
    i := 100
    x = i
    fmt.Printf("type:%T value:%v\n", x, x)
    b := true
    x = b
    fmt.Printf("type:%T value:%v\n", x, x)
}
```

> [!tip]
>
> `any()` 是空接口的别名，在所有方面都等效于空接口。

#### 13.2.1 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

#### 13.2.2 map 的值实现空接口  

使用空接口实现可以保存任意值的字典。

```go
var studentInfo = make(map[string]interface{})
studentInfo["name"] = "张三"
studentInfo["age"] = 18
studentInfo["married"] = false
studentInfo["hobby"] = []string{"football", "swimming"}
// 错误写法：
// invalid operation: cannot index studentInfo["hobby"] (map index expression of type interface{})
// fmt.Println(studentInfo["hobby"][1]) 
// 正确写法：
fmt.Println(studentInfo["hobby"].([]string)[1])
```

#### 13.2.3 切片实现空接口  

```go
var slice = []interface{}{"张三", 20, true, 32.2}
fmt.Println(slice)
```

### 13.3 类型断言  

一个接口的值（简称接口值）是由一个具体类型和具体类型的值两部分组成的。这两部分分别称为接口的动态类型和动态值。

```go
x.(T)
```

其中：

- `x`：表示类型为 `interface{}` 的变量
- `T`：表示断言 `x` 可能是的类型

该语法返回两个参数，第一个参数是 `x` 转化为 `T` 类型后的变量，第二个值是一个布尔值，若为 `true` 则表示断言成功，为 `false` 则表示断言失败。  

```go
var x interface{}
x = "Hello world"
v, ok := x.(string)
if ok {
    fmt.Println(v)
} else {
    fmt.Println("类型断言失败")
}
```

> [!caution]
>
> `x.(type)` 语句也可以判断一个变量的类型，但是只能结合 `switch` 语句使用  

```go
func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string， value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type", v)
	}
}
```

> [!note]
>
> 只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。不要为了接口而写接口，那样只会增加不必要的抽象，导致不必要的运行时损耗。

### 13.4 结构体值接收者和指针接收者实现接口的区别

#### 13.4.1 值接收者  

如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量。

```go
type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

func main() {
    // 正确写法
	phone1 := Phone{"小米"}
	var p1 Usber = phone1
	p1.Start() // 小米 手机开始工作
	p1.Stop()  // 小米 手机停止工作
	
    // 正确写法
	phone2 := &Phone{"苹果"}
	var p2 Usber = phone2
	p2.Start() // 苹果 手机开始工作
	p2.Stop()  // 苹果 手机停止工作
}
```

#### 13.4.2 指针接收者

如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量，**结构体值类型没法赋值给接口变量**。

```go
type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p *Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p *Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

func main() {
    // 错误写法
	phone1 := Phone{"小米"}
	var p1 Usber = phone1
	p1.Start() // 小米 手机开始工作
	p1.Stop()  // 小米 手机停止工作
	
    // 正确写法
	phone2 := &Phone{"苹果"}
	var p2 Usber = phone2
	p2.Start() // 苹果 手机开始工作
	p2.Stop()  // 苹果 手机停止工作
}
```

### 13.5 一个结构体实现多个接口  

```go
package main

import "fmt"

type A interface {
	Get() string
}

type B interface {
	Set(string, int)
}

type People struct {
	Name string
	Age  int
}

func (p *People) Get() string {
	return fmt.Sprintf("name:%s, age:%d", p.Name, p.Age)
}

func (p *People) Set(name string, age int) {
	p.Name = name
	p.Age = age
}

func main() {
	var people = &People{"张三", 18}

	var p1 A = people // 实现 A 接口
	var p2 B = people // 实现 B 接口

	fmt.Println(p1.Get()) // name:张三, age:18
	p2.Set("王五", 20)
	fmt.Println(p1.Get()) // name:王五, age:20
}
```

### 13.6 接口嵌套

```go
package main

import "fmt"

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

type Animal interface {
	Sayer
	Mover
}

type Cat struct {
	Name string
}

func (c Cat) Say() {
	fmt.Println(c.Name, "喵喵喵")
}

func (c Cat) Move() {
	fmt.Println(c.Name, "跑")
}

func main() {
	var cat = Cat{"小猫"}
	var animal Animal = cat
	animal.Say() // 小猫 喵喵喵
	animal.Move() // 小猫 跑
}
```

## 14. 协程

### 14.1 高并发基础

#### 14.1.1 线程与进程

- **进程（Process）**就是程序在操作系统中的一次执行过程，**是系统进行资源分配和调度的基本单位**，进程是一个动态概念，是程序在执行过程中分配和管理资源的基本单位，每一个进程都有一个自己的地址空间。 **一个进程至少有 5 种基本状态，它们是：初始态、执行态、等待状态、就绪状态以及终止状态。**
- **线程**是进程的一个执行实例，**是程序执行的最小单元**，它是比进程更小的能独立运行的基本单位。一个进程可以创建多个线程，同一个进程中的多个线程可以并发执行，一个程序要运行的话至少有一个进程。

#### 14.1.2 并行与并发

- **并发：**多个线程同时竞争一个位置，竞争到的才可以执行，每一个时间段只有一个线程在执行。
- **并行：**多个线程可以同时执行，每一个时间段，可以有多个线程同时执行。  

通俗的讲**多线程程序在单核 CPU 上面运行就是并发**，**多线程程序在多核 CPU 上运行就是并行**，如果线程数大于 CPU 核数，则多线程程序在多个 CPU 上面运行既有并行又有并发。

### 14.2 Goroutine 的定义

**Golang 中的主线程：**（可以理解为线程/也可以理解为进程），在一个 Golang 程序的主线程上可以起多个协程。Golang 中多协程可以实现并行或者并发。

**协程：**可以理解为用户级线程，这是对内核透明的，也就是系统并不知道有协程的存在，是完全由用户自己的程序进行调度的。

Golang 的一大特色就是从语言层面原生支持协程，在函数或者方法前面加 `go` 关键字就可创建一个协程。可以说 Golang 中的协程就是goroutine 。

> [!tip]
>
> Golang 中的多协程有点类似其他语言中的多线程。

**多协程和多线程：**Golang 中每个 goroutine (协程) 默认占用内存远比 Java、 C 的线程少。OS 线程（操作系统线程）一般都有固定的栈内存（通常为 2MB 左右），一个 goroutine (协程) 占用内存非常小，只有 2KB 左右，多协程 goroutine 切换调度开销方面远比线程要少。

### 14.3 Goroutine 的使用

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

### 14.4 设置 CPU 数量

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

### 14.5 管道

管道（Channel）是 Golang 在语言级别上提供的 goroutine 间的通讯方式，我们可以使用 `channel` 在多个 goroutine 之间传递消息。如果说 goroutine 是 Go 程序并发的执行体，`channel` 就是它们之间的连接。`channel` 是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。

> [!tip]
>
> Golang 的并发模型是 CSP（Communicating Sequential Processes），提倡**通过通信共享内存**而不是**通过共享内存而实现通信**。

Go 语言中的 `channel` 是一种特殊的类型，总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序。 

#### 14.5.1 管道的声明

`channel` 是一种类型，一种引用类型。声明管道类型的格式如下：

```go
var ch1 chan Type
```

声明的管道后需要使用 `make()` 函数初始化之后才能使用。  

```go
make(chan Type, cap)
```

#### 14.5.2 管道的操作

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

#### 14.5.3 管道阻塞

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

#### 14.5.4 遍历管道

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

### 14.6 单向管道

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

### 14.7 select

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

### 14.8 并发安全

#### 14.8.1 互斥锁

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

#### 14.8.2 读写互斥锁  

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

## 15. 反射

### 15.1 什么是发射

**反射是指在程序运行期间对程序本身进行访问和修改的能力。** 正常情况程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。 

支持反射的语言可以在程序编译期将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期获取类型的反射信息，并且有能力修改它们。

> [!important]
>
> Golang 中反射可以实现以下功能：
>
> 1. 反射可以在程序运行期间动态的获取变量的各种信息，比如变量的类型类别。
> 2. 如果是结构体，通过反射还可以获取结构体本身的信息，比如结构体的字段、结构体的方法、结构体的 tag。
> 3. 通过反射，可以修改变量的值，可以调用关联的方法。

Go 语言中的变量是分为两部分的：

1. **类型信息：**预先定义好的元信息。
2. **值信息：**程序运行过程中可动态变化的。

在 Golang 的反射机制中，任何接口值都由是一个具体类型和具体类型的值两部分组成的。

在 Golang 中，反射的相关功能由内置的 `reflect` 包提供，任意接口值在反射中都可以理解为由 `reflect.Type` 和 `reflect.Value` 两部分组成，并且 `reflect` 包提供了 `reflect.TypeOf()` 和 `reflect.ValueOf()` 两个重要函数来获取任意对象的 Value 和 Type。

### 15.2 reflect.TypeOf()

在 Go 语言中，使用 `reflect.TypeOf()` 函数可以接受任意 `interface{}` 参数，可以获得任意值的类型对象 `reflect.Type`，程序通过类型对象可以访问任意值的类型信息。

```go
package main

import (
	"fmt"
	"reflect"
)

func ReflectFunc(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}

func main() {
	// 通过 reflect 包获取变量的详细信息
	var a int = 10
	ReflectFunc(a) // int
	var b string = "hello world"
	ReflectFunc(b) // string
}
```

#### 15.2.1 Name 和 Kind

在反射中的类型还分为两种：类型（Type）和种类（Kind）。 因为在 Go 语言中我们可以使用 `type` 关键字构造很多自定义类型，而 `Kind` 就是指底层的类型，但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到 `Kind` 。举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。

Go 语言的反射中像数组、切片、Map、指针等类型的变量，它们的 `.Name()` 都是返回空。

```go
package main

import (
	"fmt"
	"reflect"
)

func ReflectFunc(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("TypeOf: %v, Name: %v, Kind: %v\n", v, v.Name(), v.Kind())
}

type Person struct {
	Name string
}

type MyInt int

func main() {
	var a *float32 // 指针
	var b MyInt    // 自定义类型
	var c rune     // 类型别名
	ReflectFunc(a) // TypeOf: *float32, Name: , Kind: ptr
	ReflectFunc(b) // TypeOf: main.MyInt, Name: MyInt, Kind: int
	ReflectFunc(c) // TypeOf: int32, Name: int32, Kind: int32

	var d Person = Person{"Tom"}
	ReflectFunc(d) // TypeOf: main.Person, Name: Person, Kind: struct

	var e = []int{1, 2, 3}
	ReflectFunc(e) // TypeOf: []int, Name: , Kind: slice
}
```

### 15.3 reflect.ValueOf()  

`reflect.ValueOf()` 返回的是 `reflect.Value` 类型，其中包含了原始值的值信息。`reflect.Value` 与原始值之间可以互相转换。

```go
package main

import (
	"fmt"
	"reflect"
)

func ReflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int:
		fmt.Println(v.Int())
	case reflect.Float32:
		fmt.Println(v.Float())
	case reflect.String:
		fmt.Println(v.String())
	case reflect.Bool:
		fmt.Println(v.Bool())
	default:
		fmt.Println("unsupported type")
	}
}

func main() {
	var a float32 = 3.14
	var b int = 100
	ReflectValue(a) // 3.140000104904175
	ReflectValue(b) // 100

	var c string = "hello"
	ReflectValue(c) // hello
    
    // 将 int 类型的原始值转换为 reflect.Value 类型
	var d = reflect.ValueOf(b)
	fmt.Printf("Type d: %T\n", d) // Type d: reflect.Value
}
```

想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的 

`Elem()` 方法来获取指针对应的值。

```go
package main

import (
	"fmt"
	"reflect"
)

func ReflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int {
		v.Elem().SetInt(123)
	} else if v.Elem().Kind() == reflect.String {
		v.Elem().SetString("hello")
	} else {
		fmt.Println("unsupported type")
	}
}

func main() {
	var a = 100
	ReflectSetValue(&a)
	fmt.Println(a) // 123

	var b = "你好"
	ReflectSetValue(&b)
	fmt.Println(b) // hello
}
```

> [!caution]
>
> 在使用 `Elem()` 之前，**一定要判断 `Kind()`**，否则非常容易触发 `Panic`。

### 15.4 结构体反射

#### 15.4.1 与结构体相关的方法

任意值通过 `reflect.TypeOf()` 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 `reflect.Type` 的 `NumField()` 和 `Field()` 方法获得结构体成员的详细信息。

- `Field(i int) StructField`：根据索引，返回索引对应的结构体字段的信息
- `NumField() int`：返回结构体成员字段数量。 
- `FieldByName(name string) (StructField, bool)`：根据给定字符串返回字符串对应的结构体字段的信息。
- `FieldByIndex(index []int) StructField`：多层成员访问时，根据 `[]int` 提供的每个结构体的字段索引，返回字段的信息。 
- `FieldByNameFunc(match func(string) bool) (StructField,bool)`：根据传入的匹配函数匹配需要的字段。
- `NumMethod() int`：返回该类型的方法集中方法的数目
- `Method(int) Method`：返回该类型方法集中的第 i 个方法
- `MethodByName(string) (Method, bool)`：根据方法名返回该类型方法集中的方法

#### 15.4.2 StructField 类型

`StructField` 类型用来描述结构体中的一个字段的信息。`StructField` 的定义如下：

```go
type StructField struct {
    // 参见 http://golang.org/ref/spec#Uniqueness_of_identifiers
    Name string // Name 是字段的名字
    PkgPath string //PkgPath 是非导出字段的包路径， 对导出字段该字段为""
    Type Type // 字段的类型
    Tag StructTag // 字段的标签
    Offset uintptr // 字段在结构体中的字节偏移量
    Index []int // 用于 Type.FieldByIndex 时的索引切片
    Anonymous bool // 是否匿名字段
}
```

#### 15.4.3 结构体反射实例

当我们使用反射得到一个结构体数据之后可以通过索引依次获取其字段信息，也可以通过字段名去获取指定的字段信息。

**获取结构体属性：**  

```go
package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s *Student) GetInfo() string {
	info := "Name:" + s.Name + " Age:" + strconv.Itoa(s.Age) + " Score:" + strconv.Itoa(s.Score)
	fmt.Println("GetInfo被调用:", info)
	return info
}

func (s *Student) SetInfo(name string, age int, score int) {
	s.Name = name
	s.Age = age
	s.Score = score
	fmt.Println("SetInfo被调用")
}

func (s *Student) Print() {
	fmt.Println("Print被调用: 打印。。。")
}

func PrintStructField(s interface{}) {
	t := reflect.TypeOf(s)

	// 处理指针情况：如果是指针，取其指向的类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		fmt.Println("It is not a struct")
		return
	}

	// 1. 通过索引获取字段
	field0 := t.Field(0)
	fmt.Printf("字段0: 名称=%s, 类型=%s, JSON tag=%s\n", field0.Name, field0.Type, field0.Tag.Get("json"))

	// 2. 通过名称获取字段
	if field1, ok := t.FieldByName("Name"); ok {
		fmt.Printf("字段Name: 名称=%s, 类型=%s, JSON tag=%s\n", field1.Name, field1.Type, field1.Tag.Get("json"))
	}

	// 3. 获取字段总数
	fmt.Println("该结构体的字段数：", t.NumField())
}

func PrintStructMethod(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	// 所以如果要调用指针接收者的方法，传入的必须是指针
	if t.Kind() != reflect.Ptr || t.Elem().Kind() != reflect.Struct {
		fmt.Println("Expected a pointer to a struct")
		return
	}

	// 1. 通过类型变量里面的 Method 可以获取结构体的方法
	// 1. 获取方法信息 (按 ASCII 排序)
	// 注意：确保至少有一个方法，否则 Method(0) 会 Panic
	if t.NumMethod() > 0 {
		var tMethod = t.Method(0)
		fmt.Printf("第0个方法: 名称=%s, 类型=%s\n", tMethod.Name, tMethod.Type)
	}

	// 2. 获取方法总数
	fmt.Println("该结构体的方法数：", t.NumMethod())

	fmt.Println(">>> 调用 Print 方法:")
	mPrint := v.MethodByName("Print")
	if mPrint.IsValid() {
		mPrint.Call(nil)
	} else {
		fmt.Println("Print 方法不存在")
	}

	// 4. 执行方法传入参数
	fmt.Println(">>> 调用 SetInfo 方法:")
	mSetInfo := v.MethodByName("SetInfo")
	if mSetInfo.IsValid() {
		var params []reflect.Value
		params = append(params, reflect.ValueOf("李四"))
		params = append(params, reflect.ValueOf(22))
		params = append(params, reflect.ValueOf(100))
		mSetInfo.Call(params)
	}

	// 5. 执行方法获取返回值
	fmt.Println(">>> 调用 GetInfo 方法:")
	mGetInfo := v.MethodByName("GetInfo")
	if mGetInfo.IsValid() {
		results := mGetInfo.Call(nil)
		// 打印返回值 (Results 是 []reflect.Value)
		if len(results) > 0 {
			fmt.Println("返回值:", results[0].String())
		}
	}
}

func main() {
	stu1 := &Student{Name: "小王子", Age: 18, Score: 90}

	// PrintStructField 可以传值也可以传指针（只要函数内处理好了）
	// 这里演示传值
	PrintStructField(*stu1)

	// PrintStructMethod 必须传指针，因为 Student 的方法接收者是 *Student
	PrintStructMethod(stu1)
}
```

**修改结构体方法：**

```go
package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	return fmt.Sprintf("姓名：%v，年龄：%v，分数：%v", s.Name, s.Age, s.Score)
}

func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Elem().Kind() != reflect.Struct {
		fmt.Println("It is not a struct")
		return
	}

	name := v.Elem().FieldByName("Name")
	name.SetString("李四")

	age := v.Elem().FieldByName("Age")
	age.SetInt(30)
}

func main() {
	stu := &Student{
		Name:  "张三",
		Age:   20,
		Score: 90,
	}

	reflectChangeStruct(stu)
	fmt.Println(stu.GetInfo()) // 姓名：李四，年龄：30，分数：90
}
```

> [!important]
>
> 反射是一个强大并富有表现力的工具，能让我们写出更灵活的代码。但是反射不应该被滥用，原因有以下两个。
>
> 1. 基于反射的代码是极其脆弱的，反射中的类型错误会在真正运行的时候才会引发 `panic`，那很可能是在代码写完的很长时间之后。
> 2. 大量使用反射的代码通常难以理解。

## 16. IO 流

### 16.1 打开和关闭文件

#### 16.1.1 
