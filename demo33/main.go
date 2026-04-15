package main

import (
	"fmt"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open("D:\\project\\Go-language\\.gitignore")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	if err != nil {
		return
	}

	// 读取文件
	var buf = make([]byte, 1024)
	read, err := file.Read(buf)
	if err != nil {
		return
	}
	fmt.Println(string(buf[:read]))

	var array = [1024]byte{}
	fmt.Printf("%T, %T\n", array, array[:])

}
