# 10. 结构体

Golang 中没有类的概念，Golang 中的结构体和其他语言中的类有点相似。和其他面向对象语言中的类相比，Golang 中的结构体具有更高的扩展性和灵活性。

## 10.1 自定义类型和类型别名

### 10.1.1 自定义类型

在 Go 语言中有一些基本的数据类型，如 `string`、`int`、`float`、`bool`等数据类型，Go 语言中可以使用 `type` 关键字来定义自定义类型。

```go
type myInt int
```

### 10.1.2 类型别名  

`TypeAlias` 只是 `Type` 的别名，本质上 `TypeAlias` 与 `Type` 是同一个类型。 

```go
type TypeAlias = Type
```

`rune` 和 `byte` 就是类型别名，他们的底层定义如下  

```go
type byte = uint8
type rune = int32
```

### 10.1.3 自定义类型和类型别名的区别  

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

## 10.2 结构体的初始化

### 10.2.1 结构体的声明

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

### 10.2.2 结构体实例化

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

## 10.3 结构体方法和接收者  

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

## 10.4 给任意类型添加方法  

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

## 10.5 结构体的匿名字段  

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

## 10.6 嵌套结构体  

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

## 10.7 结构体的继承

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

