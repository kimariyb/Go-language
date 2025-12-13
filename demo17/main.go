package main

import (
	"errors"
	"fmt"
)

func f1() {
	fmt.Println("hello")
}

func f2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()
	panic("panic")
}

func f3() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	num1 := 10
	num2 := 0
	res := num1 / num2
	fmt.Println("res=", res)
}

func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	}
	return errors.New("读取文件错误")
}

func f4() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	err := readFile("xxx.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	f1()
	f2()
	fmt.Println("world")

	f3()
	f4()
	fmt.Println("continue...")
}
