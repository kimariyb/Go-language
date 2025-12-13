package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	//var a int8 = 120
	//fmt.Printf("%T\n", a)
	//fmt.Println(unsafe.Sizeof(a))

	fmt.Printf("%f\n", math.Pi)   //默认保留 6 位小数
	fmt.Printf("%.2f\n", math.Pi) //保留 2 位小数

	m1 := 8.2
	m2 := 3.8
	fmt.Println(m1 - m2) // 期望是 4.4， 结果打印出了 4.3999999999999995

	// 科学计数法表示浮点数
	num1 := 1.234e3
	num2 := 1.234e-3

	fmt.Println(num1, num2)

	//var b = true
	//fmt.Println(b, "占用字节：", unsafe.Sizeof(b))

	var str = "123-456-789"
	var arr = strings.Split(str, "-")
	fmt.Println(arr)
	var str2 = strings.Join(arr, "*")
	fmt.Println(str2)

	a := 'a'
	b := '0'
	//当我们直接输出 byte（字符） 的时候输出的是这个字符对应的码值
	fmt.Println(a)
	fmt.Println(b)
	//如果我们要输出这个字符， 需要格式化输出
	fmt.Printf("%c--%c", a, b) //%c 相应 Unicode 码点所表示的字符
}
