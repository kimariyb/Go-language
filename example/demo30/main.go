package main
// Go语言教程: 第15章 反射 - TypeOf/ValueOf

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

func ReflectFunc(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("TypeOf: %v, Name: %v, Kind: %v\n", v, v.Name(), v.Kind())
}

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

type Person struct {
	Name string
}

type MyInt int

func main() {
	//var a *float32 // 指针
	//var b MyInt    // 自定义类型
	//var c rune     // 类型别名
	//ReflectFunc(a) // TypeOf: *float32, Name: , Kind: ptr
	//ReflectValue(a)
	//ReflectFunc(b) // TypeOf: main.MyInt, Name: MyInt, Kind: int
	//ReflectFunc(c) // TypeOf: int32, Name: int32, Kind: int32
	//
	//var d Person = Person{"Tom"}
	//ReflectFunc(d) // TypeOf: main.Person, Name: Person, Kind: struct
	//
	//var e = []int{1, 2, 3}
	//ReflectFunc(e) // TypeOf: []int, Name: , Kind: slice

	//var a float32 = 3.14
	//var b int = 100
	//ReflectValue(a) // 3.140000104904175
	//ReflectValue(b) // 100
	//
	//var c string = "hello"
	//ReflectValue(c) // hello
	//
	//// 将 int 类型的原始值转换为 reflect.Value 类型
	//var d = reflect.ValueOf(b)
	//fmt.Printf("Type d: %T\n", d) // Type d: reflect.Value

	var a = 100
	ReflectSetValue(&a)
	fmt.Println(a) // 123

	var b = "你好"
	ReflectSetValue(&b)
	fmt.Println(b) // hello
}
