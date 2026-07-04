package main
// Go语言教程: 第5章 数组 - 数组声明与遍历

import "fmt"

func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
	var testArray = [...]int{}
	var numArray = [...]int{1, 2, 3}
	var cityArray = [...]string{"北京", "上海", "广州"}

	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)

	a := [...]int{1: 1, 3: 5}
	fmt.Println(a)                  // [0 1 0 5]
	fmt.Printf("type of a:%T\n", a) // type of a:[4]int

	var array = [...]string{"java", "python", "golang"}

	// 1. for 循环遍历
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	// 2. for range 遍历
	for _, v := range array {
		fmt.Println(v)
	}

	arr1 := [3]int{1, 2, 3}
	modifyArray(arr1)
	fmt.Println(arr1)

	arr2 := [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	modifyArray2(arr2)
	fmt.Println(arr2)
}
