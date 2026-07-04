package main
// Go语言教程: 第7章 哈希表(map) - 增删改查

import "fmt"

func main() {
	// 1. 使用 make 创建 map
	hashMap := make(map[string]int)
	hashMap["a"] = 1
	hashMap["b"] = 2
	hashMap["c"] = 3
	fmt.Println(hashMap)
	fmt.Println(hashMap["a"])

	userInfo := map[string]string{
		"username": "zhangsan",
		"password": "123456",
	}
	fmt.Println(userInfo)

	scoreMap := make(map[string]int)
	scoreMap["张三"] = 90
	scoreMap["李四"] = 80
	scoreMap["王五"] = 70
	fmt.Println(scoreMap)

	// 如果 key 存在 ok 为 true,v 为对应的值；
	// 不存在 ok 为 false,v 为值类型的零值
	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v) // 90
	} else {
		fmt.Println("没有该用户")
	}

	for k, v := range scoreMap {
		fmt.Println(k, v)
	}

	for k := range scoreMap {
		fmt.Println(k)
	}
	delete(scoreMap, "张三")
	fmt.Println(scoreMap)
}
