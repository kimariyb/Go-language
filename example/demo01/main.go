package main
// Go语言教程: 第1章 变量和常量 - var、短变量声明、常量

import "fmt"

func main() {

	// 1. 使用 var 关键字定义变量
	var name string = "James"

	fmt.Println("Hello,", name)

	// 2. 使用 := 语法定义变量
	age := 18
	fmt.Println("I am", age, "years old")

	// 3. fmt 中的 Print、Println、Printf
	fmt.Print("Hello, ")
	fmt.Print("world")

	fmt.Println("Hello, ")
	fmt.Println("world")

	fmt.Printf("I am %d years old\n", age)

	var num int64
	num = 123
	fmt.Printf("值: %d, 类型: %T", num, num)

}
