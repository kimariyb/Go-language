# 13. 接口

## 13.1 接口的声明

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

## 13.2 空接口

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

### 13.2.1 空接口作为函数的参数

使用空接口实现可以接收任意类型的函数参数。

```go
func show(a interface{}) {
	fmt.Printf("type:%T value:%v\n", a, a)
}
```

### 13.2.2 map 的值实现空接口  

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

### 13.2.3 切片实现空接口  

```go
var slice = []interface{}{"张三", 20, true, 32.2}
fmt.Println(slice)
```

## 13.3 类型断言  

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

## 13.4 结构体值接收者和指针接收者实现接口的区别

### 13.4.1 值接收者  

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

### 13.4.2 指针接收者

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

## 13.5 一个结构体实现多个接口  

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

## 13.6 接口嵌套

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

