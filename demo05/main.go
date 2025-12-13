package main

import "fmt"

func main() {
	var i int = 10
	var f float32 = 3.14
	var c byte = 'a'
	var b bool = true

	var s string

	s = fmt.Sprintf("%d", i)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%f", f)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%t", b)
	fmt.Printf("类型 %T, %v \n", s, s)
	s = fmt.Sprintf("%c", c)
	fmt.Printf("类型 %T, %v \n", s, s)
}
