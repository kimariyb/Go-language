package main
// Go语言教程: 第4章 流程控制 - for循环/break

import "fmt"

func main() {
	var score int = 80

	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	k := 1
	for {
		if k <= 10 {
			fmt.Println(k)
		} else {
			break // break 就是跳出这个 for 循环
		}
		k++
	}

	s := "hello world"
	for index, value := range s {
		fmt.Printf("%d:%c\n", index, value)
	}

	slice := []int{1, 2, 3, 4, 5}
	for index, value := range slice {
		fmt.Printf("%d:%d\n", index, value)
	}

}
