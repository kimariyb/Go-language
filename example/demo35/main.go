package main
// Go语言教程: 第17章 泛型 - 泛型函数与约束

import "fmt"

// Swap 泛型函数：交换两个值
func Swap[T any](a, b *T) {
	*a, *b = *b, *a
}

// 使用示例
func main() {
	x, y := 10, 20
	Swap(&x, &y)      // 类型推断：无需显式指定T=int
	fmt.Println(x, y) // 输出 20 10

	s1, s2 := "hello", "world"
	Swap(&s1, &s2)      // 类型推断：T=string
	fmt.Println(s1, s2) // 输出 world hello
}
