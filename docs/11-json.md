# 11. JSON

## 11.1 JSON 序列化和反序列化

Golang 中的序列化和反序列化主要通过 `"encoding/json"` 包中的 `json.Marshal()` 和 `json.Unmarshal()` 方法实现。

### 11.1.1 序列化

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

### 11.1.2 反序列化

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

## 11.2 结构体标签 Tag

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

## 11.3 嵌套结构体的 JSON

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

