# 17. 泛型

## 17.1 泛型的概念

Go 泛型的核心是**类型参数**（Type Parameters）和**类型约束**（Type Constraints）：

- **类型参数**：用方括号`[]`声明在函数名或类型名之后，代表一组可能的类型。
- **类型约束**：限制类型参数的取值范围，确保代码中对类型参数的操作是合法的

```go
// 泛型函数
func 函数名[T 类型约束](参数 T) 返回值 T { ... }

// 泛型类型
type 类型名[T 类型约束] struct {
    字段 T
}
```

Go 语言提供了两个内置的基础约束：

- **`any`**：等价于`interface{}`，表示任意类型。
- **`comparable`**：表示所有可比较的类型（支持`==`和`!=`操作），适用于需要比较操作的场景（如 `map` 的 `key`）

```go
package main

import "fmt"

// Swap 泛型函数：交换两个值
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// 使用示例
func main() {
	x, y := 10, 20
	Swap(&x, &y)      // 类型推断：无需显式指定T=int
	fmt.Println(x, y) // 输出 20 10

	s1, s2 := "hello", "world"
	Swap(&s1, &s2)      // 类型推断：T=string
	fmt.Println(s1, s2) // 输出 world hello
}
```

## 17.2 类型约束

### 17.2.1 联合类型约束

使用`|`操作符指定多个允许的类型，适用于需要限制类型范围的场景：

```go
// 支持int64和float64类型的求和函数
func SumNumbers[K comparable, V int64 | float64](m map[K]V) V {
    var sum V
    for _, v := range m {
        sum += v
    }
    return sum
}
```

### 17.2.2 自定义接口约束

通过接口定义更复杂的类型约束，接口可以包含方法签名和类型集：

```go
// 定义有序类型约束（支持 < 比较操作）
type Ordered interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 | uintptr |
    float32 | float64 |
    string
}

// 泛型函数：查找切片中的最大值
func Max[T Ordered](slice []T) T {
    if len(slice) == 0 {
        panic("slice is empty")
    }
    max := slice[0]
    for _, v := range slice[1:] {
        if v > max {
            max = v
        }
    }
    return max
}
```

### 17.2.3 近似类型约束

使用`~`操作符匹配底层类型为 T 的所有类型（包括自定义类型）：

```go
type MyInt int

// 约束匹配所有底层类型为int的类型
type IntLike interface {
    ~int
}

func Double[T IntLike](x T) T {
    return x * 2
}

func main() {
    var a int = 5
    var b MyInt = 10
    fmt.Println(Double(a))  // 输出 10
    fmt.Println(Double(b))  // 输出 20（MyInt底层是int）
}
```

## 17.3 类型推断

Go 语言泛型支持强大的**类型推断**机制，在大多数情况下可以省略类型参数，让代码更简洁：

### 17.3.1 函数参数类型推断

```go
// 无需显式指定类型参数T
sum := SumNumbers(map[string]int64{"a": 1, "b": 2, "c": 3})
fmt.Println(sum)  // 输出 6
```

### 17.3.2 赋值语句类型推断

```go
// 从右侧推断Stack的类型参数为int
intStack := Stack[int]{}
// 等价于：
var intStack Stack[int]
```

### 17.3.3 方法调用类型推断

```go
intStack.Push(1)  // 从参数1推断T=int
element, ok := intStack.Pop()  // 从返回值推断T=int
```

## 17.4 泛型接口

接口可以包含类型参数，定义更灵活的抽象：

```go
type Parser[T any] interface {
    Parse(data []byte) (T, error)
    String() string
}

// 实现JSON解析器
type JSONParser[T any] struct{}

func (p JSONParser[T]) Parse(data []byte) (T, error) {
    var result T
    err := json.Unmarshal(data, &result)
    return result, err
}

func (p JSONParser[T]) String() string {
    return "JSON parser"
}
```

