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
