package main

import "fmt"

func intSum(a int, b int) (res int) {
	res = a + b
	return res
}

func intSum2(x ...int) int {
	fmt.Printf("%T\n", x) // []int
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

type Calculation func(int, int) int

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func doCalc(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func main() {
	res := intSum(1, 2)
	fmt.Println(res)
	res2 := intSum2(1, 2, 3, 4, 5)
	fmt.Println(res2)

	var c Calculation = add
	fmt.Printf("%T\n", c) // main.Calculation
	fmt.Println(c(1, 2))  // 3

	var a = doCalc("+")
	fmt.Println(a(5, 2)) // 7

	fmt.Println(
		func(x, y int) int {
			return x + y
		}(10, 20))
}
