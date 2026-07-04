package main
// Go语言教程: 第2章 基本数据类型 - strconv 包转换

import (
	"fmt"
	"strconv"
)

func main() {
	//1、 int 转换成 string
	var num1 int = 20
	s1 := strconv.Itoa(num1)
	fmt.Printf("类型 %T, %v \n", s1, s1)

	// 2、 float 转 string
	var num2 float64 = 20.113123
	s2 := strconv.FormatFloat(num2, 'f', 2, 64)
	fmt.Printf("类型 %T, %v \n", s2, s2)

	// 3、 bool 转 string
	s3 := strconv.FormatBool(true)
	fmt.Printf("类型 %T, %v \n", s3, s3)

	// 4、 int64 转 string
	var num3 int64 = 20
	s4 := strconv.FormatInt(num3, 10)
	fmt.Printf("类型 %T ,strs=%v \n", s4, s4)
}
