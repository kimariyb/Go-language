package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// 创建一个文件
	file, err := os.OpenFile(
		"./demo34/test.txt",   // 文件名
		os.O_CREATE|os.O_RDWR, // 创建和读写模式
		0666)

	if err != nil {
		return
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			return
		}
	}(file)

	//contents := "你好 Golang\n"
	//_, err = file.Write([]byte(contents))
	//if err != nil {
	//	return
	//} // 写入切片
	//_, err = file.WriteString("直接写入") // 写入字符串
	//if err != nil {
	//	return
	//}

	//writer := bufio.NewWriter(file)
	//
	//_, err = writer.WriteString("\nHello World\n")
	//if err != nil {
	//	return
	//}
	//err = writer.Flush()
	//if err != nil {
	//	return
	//} // 将缓存中的内容写入文件

	str := "hello world"
	err = ioutil.WriteFile("./demo34/test.txt", []byte(str), 0666)
	if err != nil {
		return
	}
}
