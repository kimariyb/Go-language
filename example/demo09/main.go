package main
// Go语言教程: 第3章 运算符 - 关系运算符

import "fmt"

func main() {
	var n1 int = 9
	var n2 int = 8

	// 1、== 运算符
	fmt.Println(n1 == n2) // false
	// 2、!= 运算符
	fmt.Println(n1 != n2) // true
	// 3、> 运算符
	fmt.Println(n1 > n2) // true
	// 4、>= 运算符
	fmt.Println(n1 >= n2) // true
	// 5、< 运算符
	fmt.Println(n1 < n2) // flase
	// 6、<= 运算符
	fmt.Println(n1 <= n2) // flase

	// 1、&& 逻辑运算符
	var age int = 40
	if age > 30 && age < 50 {
		fmt.Println("ok1")
	}
	if age > 30 && age < 40 {
		fmt.Println("ok2")
	}

	// 2、|| 逻辑运算符
	if age > 30 || age < 50 {
		fmt.Println("ok3")
	}
	if age > 30 || age < 40 {
		fmt.Println("ok4")
	}
	// 3、! 逻辑运算符
	if age > 30 {
		fmt.Println("ok5")
	}
	if !(age > 30) {
		fmt.Println("ok6")
	}

	// 示例1：&& 短路
	fmt.Println("=== && 短路演示 ===")

	a, b := 10, 20

	// 第一个条件为 false，第二个条件不会执行
	if a > 100 && b > 0 {
		fmt.Println("条件成立")
	} else {
		fmt.Println("条件不成立，b>0不会被执行")
	}

	// 示例2：|| 短路
	fmt.Println("\n=== || 短路演示 ===")

	// 第一个条件为true，第二个条件不会执行
	if a < 100 || b < 0 {
		fmt.Println("条件成立，b<0不会被执行")
	}
}
