package main
// Go语言教程: 第2章 基本数据类型 - string转数值类型

import (
	"fmt"
	"strconv"
)

func main() {
	// 1、string 类型转 int
	var s1 string = "1234"
	i, _ := strconv.ParseInt(s1, 10, 64)
	fmt.Printf("值：%d，类型：%T\n", i, i)

	// 2、string 转 float
	var s2 string = "3.1415926"
	f, _ := strconv.ParseFloat(s2, 64)
	fmt.Printf("值：%f，类型：%T\n", f, f)

	// 3、string 转 bool
	var s3 string = "true"
	b, _ := strconv.ParseBool(s3)
	fmt.Printf("值：%t，类型：%T\n", b, b)

	// 4、string 转 byte
	var s4 string = "abc"
	for _, r := range s4 {
		fmt.Printf("值：%c，类型：%T\n", r, r)
	}
}
