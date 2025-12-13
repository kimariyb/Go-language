package main

import "fmt"

func adder() func(int) int {
	var x int

	return func(y int) int {
		x += y
		return x
	}
}

func Calc(base int) (func(int) int, func(int) int) {
	add := func(y int) int {
		base += y
		return base
	}
	sub := func(y int) int {
		base -= y
		return base
	}
	return add, sub
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	//var f = adder()
	//
	//fmt.Println(f(5)) // 5
	//fmt.Println(f(5)) // 10
	//fmt.Println(f(5)) // 15

	//f1, f2 := Calc(10)
	//fmt.Println(f1(5), f2(5)) // 15 10
	//fmt.Println(f1(5), f2(5)) // 15 10
	//fmt.Println(f1(5), f2(5)) // 15 10

	//fmt.Println("Start")
	//defer fmt.Println(1)
	//defer fmt.Println(2)
	//defer fmt.Println(3)
	//fmt.Println("End")

	//fmt.Println("====")
	//fmt.Println(f1())
	//fmt.Println(f2())
	//fmt.Println(f3())
	//fmt.Println(f4())

	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}
