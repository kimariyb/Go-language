package main
// Go语言教程: 第16章 IO流 - 文件读取

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	// 打开文件
	file, err := os.Open(".//.gitignore")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//// 读取文件
	//var buf = make([]byte, 128) // 一次 128 字节
	//read, err := file.Read(buf)
	//if err == io.EOF {
	//	fmt.Println("EOF")
	//	return
	//}
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Printf("read: %d\n", read)
	//fmt.Println(string(buf[:read]))

	//// 循环读取文件
	//var content []byte
	//var buffer = make([]byte, 128)
	//
	//for {
	//	read, err := file.Read(buffer)
	//	if err == io.EOF {
	//		fmt.Println("EOF")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		return
	//	}
	//	content = append(content, buffer[:read]...)
	//}
	//fmt.Println(string(content))

	//reader := bufio.NewReader(file)
	//for {
	//	line, err := reader.ReadString('\n')
	//	if err == io.EOF {
	//		if len(line) != 0 {
	//			fmt.Println(line)
	//		}
	//		fmt.Println("EOF")
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		return
	//	}
	//	fmt.Println(line)
	//}
	contents, err := ioutil.ReadFile(".//.gitignore")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(string(contents))
}
