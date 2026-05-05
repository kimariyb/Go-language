package main

import (
	"encoding/json"
	"fmt"
)

//type Student struct {
//	ID     int
//	name   string // 私有属性不能被 json 包访问
//	Age    int
//	Gender string
//}

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
	var s Student = Student{
		ID:     1,
		Name:   "张三",
		Age:    18,
		Gender: "男",
	}
	fmt.Printf("student:%#v\n", s)
	sByte, err := json.Marshal(s)
	if err != nil {
		// 处理错误
		fmt.Println(err)
	}
	sStr := string(sByte)
	fmt.Println(sStr) // {"ID":1,"Age":18,"Gender":"男"}

	var jsonStr = `{"ID":1,"Name":"李四","Age":18,"Gender":"男"}`
	var student Student
	err = json.Unmarshal([]byte(jsonStr), &student)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("student:%#v\n", student) // student:main.Student{ID:1, Name:"李四", Age:18, Gender:"男"}

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
