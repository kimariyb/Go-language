package main
// Go语言教程: 第2章 基本数据类型 - 数值类型转换

import "fmt"

func main() {
	var a int8 = 20
	var b int16 = 40
	var c = int16(a) + b           // 要转换成相同类型才能运行
	fmt.Printf("值：%v，类型：%T", c, c) //值： 60--类型 int16

	fmt.Println()
	var d float32 = 3.2
	var e int16 = 6
	var f = d + float32(e)

	fmt.Printf("值：%v，类型：%T", f, f)
}
