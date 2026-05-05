package main

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
