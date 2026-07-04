package main
// Go语言教程: 第12章 包 - go mod + 自定义包

import (
	"demo21/test"
	"fmt"
)

var x int = 10

const pi = 3.14

func init() {
	fmt.Printf("x is %d\n", x)
}

func main() {
	add := test.Add(10, 20)
	fmt.Println(add)

	sub := test.Sub(10, 20)
	fmt.Println(sub)

	fmt.Printf("pi is %.2f\n", pi)
}
