package main
// Go语言教程: 第15章 反射 - 反射修改结构体属性

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
	fmt.Println(stu.GetInfo())
}
