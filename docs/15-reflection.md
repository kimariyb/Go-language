# 15. 反射

## 15.1 什么是发射

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

## 15.2 reflect.TypeOf()

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

### 15.2.1 Name 和 Kind

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

## 15.3 reflect.ValueOf()  

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

## 15.4 结构体反射

### 15.4.1 与结构体相关的方法

任意值通过 `reflect.TypeOf()` 获得反射对象信息后，如果它的类型是结构体，可以通过反射值对象 `reflect.Type` 的 `NumField()` 和 `Field()` 方法获得结构体成员的详细信息。

- `Field(i int) StructField`：根据索引，返回索引对应的结构体字段的信息
- `NumField() int`：返回结构体成员字段数量。 
- `FieldByName(name string) (StructField, bool)`：根据给定字符串返回字符串对应的结构体字段的信息。
- `FieldByIndex(index []int) StructField`：多层成员访问时，根据 `[]int` 提供的每个结构体的字段索引，返回字段的信息。 
- `FieldByNameFunc(match func(string) bool) (StructField,bool)`：根据传入的匹配函数匹配需要的字段。
- `NumMethod() int`：返回该类型的方法集中方法的数目
- `Method(int) Method`：返回该类型方法集中的第 i 个方法
- `MethodByName(string) (Method, bool)`：根据方法名返回该类型方法集中的方法

### 15.4.2 StructField 类型

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

### 15.4.3 结构体反射实例

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

