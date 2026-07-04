package main
// Go语言教程: 第2章 基本数据类型 - 字符串修改(byte/rune)

import (
	"fmt"
)

func main() {

	var s = "hello 张三"

	byteStr := []byte(s)
	fmt.Println(len(byteStr))
	byteStr[0] = 'H'
	fmt.Println(string(byteStr))

	runeStr := []rune(s)
	fmt.Println(len(runeStr))

	runeStr[len(runeStr)-1] = '五'
	fmt.Println(string(runeStr))
}
