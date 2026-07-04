package main
// Go语言教程: 第6章 切片 - 元素删除操作

import (
	"fmt"
	"sort"
)

func popString(srcSlice []string, index int) []string {
	// 删除切片中的元素
	return append(srcSlice[:index], srcSlice[index+1:]...)
}

func bubbleSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 冒泡排序
	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟冒泡
		isSwap := false
		for j := 0; j < len(numSlice)-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				numSlice[j], numSlice[j+1] = numSlice[j+1], numSlice[j]
				isSwap = true
			}
		}
		// 将 isSwap 检查移到内层循环外部
		if !isSwap {
			break
		}
	}
	return numSlice
}

func selecSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟选择
		minIndex := i
		// 找到最小值的索引
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[j] < numSlice[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			numSlice[i], numSlice[minIndex] = numSlice[minIndex], numSlice[i]
		}
	}

	return numSlice
}

func insertSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 遍历无序区间，从第二个元素开始遍历
	for i := 1; i < len(numSlice); i++ {
		// 保存当前元素
		current := numSlice[i]
		// 找到当前元素应该插入的位置
		j := i - 1

		// 如果 j >= 0 且 numSlice[j] > current，则说明当前元素在有序区间中，则将当前元素插入有序区间中
		for j >= 0 && numSlice[j] > current {
			numSlice[j+1] = numSlice[j]
			j--
		}
		numSlice[j+1] = current
	}
	return numSlice

}

func main() {
	// 声明切片
	//var a []string

	var a = make([]int, 2, 10)
	fmt.Println(a)      // [0 0]
	fmt.Println(len(a)) // 2
	fmt.Println(cap(a)) // 10

	var b []int = []int{1, 2, 3}

	fmt.Println(a == nil) //  true
	fmt.Println(b == nil) //  false

	var strSlice []string = []string{"hello", "world", "golang"}

	// 1. 使用 for 遍历切片
	for i := 0; i < len(strSlice); i++ {
		fmt.Println(strSlice[i])
	}

	// 2. 使用 for range 遍历切片
	for _, value := range strSlice {
		fmt.Println(value)
	}

	//numArray := [5]int{1, 2, 3, 4, 5}
	//numSlice := numArray[1:3] // 基于数组创建切片
	//fmt.Println(numSlice)     // [2 3]

	s := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(s)
	fmt.Printf("长度：%v，容量：%v\n", len(s), cap(s))

	c := s[:2]
	fmt.Println(c)
	fmt.Printf("长度：%v，容量：%v\n", len(c), cap(c))

	d := s[1:3]
	fmt.Println(d)
	fmt.Printf("长度：%v，容量：%v\n", len(d), cap(d))

	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, 5)
	copy(s2, s1)    // 将 s1 拷贝到 s2
	fmt.Println(s1) // [1 2 3 4 5]
	fmt.Println(s2) // [1 2 3 4 5]
	s2[0] = 100
	fmt.Println(s1) // [1 2 3 4 5]
	fmt.Println(s2) // [100 2 3 4 5]

	//var numSlice []int
	//for i := 0; i < 10; i++ {
	//	numSlice = append(numSlice, i)
	//	fmt.Println(numSlice)
	//}

	var citySlice []string
	// 添加一个元素
	citySlice = append(citySlice, "北京")

	// 添加多个元素
	citySlice = append(citySlice, "上海", "广州", "深圳")

	// 添加切片
	citySlice = append(citySlice, []string{"杭州", "苏州"}...)
	fmt.Println(citySlice)

	fmt.Println(popString(citySlice, 1))

	fmt.Println(bubbleSort([]int{5, 3, 2, 4, 1}))
	fmt.Println(selecSort([]int{5, 3, 2, 4, 1}))
	fmt.Println(insertSort([]int{5, 3, 2, 4, 1}))

	intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
	floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
	stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

	sort.Ints(intList)
	sort.Float64s(floatList)
	sort.Strings(stringList)
	fmt.Println(intList)    // [0 1 2 3 4 5 6 7 8 9]
	fmt.Println(floatList)  // [3.14 4.2 5.9 10.2 12.4 27.81828 31.4 50.7 99.9]
	fmt.Println(stringList) // [a b c d f i w x y z]

	// 降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(intList)))
	fmt.Println(intList)
}
