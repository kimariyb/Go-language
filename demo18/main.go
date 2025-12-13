package main

import "fmt"

func modify1(x int) {
	x = 100
}

func modify2(x *int) {
	*x = 100
}

func main() {
	//var a = 10
	//var b = &a
	////fmt.Printf("%v, %T\n", a, a)
	////fmt.Printf("%v, %T\n", b, b)
	//
	//c := *b // 指针取值（根据指针的值去内存取值）
	//fmt.Printf("type of c:%T\n", c)
	//fmt.Printf("value of c:%v\n", c)

	//a := 10
	//modify1(a)
	//fmt.Printf("%d\n", a) // 10
	//modify2(&a)
	//fmt.Printf("%d\n", a) // 100

	a := new(int)
	b := new(bool)

	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", *a, *a)
	fmt.Printf("%v, %T\n", *b, *b)
}
