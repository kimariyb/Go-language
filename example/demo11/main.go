package main
// Go语言教程: 第4章 流程控制 - switch/case

import "fmt"

func main() {
	suffix := ".a"
	switch suffix {
	case ".html":
		fmt.Println("text/html")

	case ".css":
		fmt.Println("text/css")

	case ".js":
		fmt.Println("text/javascript")

	default:
		fmt.Println("格式错误")

	}

	n := 2
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}

lable:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				break lable
			}
			fmt.Println("i, j", i, j)
		}
	}

here:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue here
			}
			fmt.Println("i, j", i, j)
		}
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				goto breakTag
			}
			fmt.Println("i, j", i, j)
		}
	}
	return
breakTag:
	fmt.Println("breakTag")
}
