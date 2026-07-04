package main
// Go语言教程: 第13章 接口 - 一个结构体实现多个接口

import "fmt"

type A interface {
	Get() string
}

type B interface {
	Set(string, int)
}

type People struct {
	Name string
	Age  int
}

func (p *People) Get() string {
	return fmt.Sprintf("name:%s, age:%d", p.Name, p.Age)
}

func (p *People) Set(name string, age int) {
	p.Name = name
	p.Age = age
}

type Sayer interface {
	Say()
}

type Mover interface {
	Move()
}

type Animal interface {
	Sayer
	Mover
}

type Cat struct {
	Name string
}

func (c Cat) Say() {
	fmt.Println(c.Name, "喵喵喵")
}

func (c Cat) Move() {
	fmt.Println(c.Name, "跑")
}

func main() {
	var people = &People{"张三", 18}

	var p1 A = people // 实现 A 接口
	var p2 B = people // 实现 B 接口

	fmt.Println(p1.Get()) // name:张三, age:18
	p2.Set("王五", 20)
	fmt.Println(p1.Get()) // name:王五, age:20

	var cat = Cat{"小猫"}
	var animal Animal = cat
	animal.Say()
	animal.Move()

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "张三"
	studentInfo["age"] = 18
	studentInfo["married"] = false
	studentInfo["hobby"] = []string{"football", "swimming"}
	// 错误写法：
	// invalid operation: cannot index studentInfo["hobby"] (map index expression of type interface{})
	// fmt.Println(studentInfo["hobby"][1])
	// 正确写法：
	fmt.Println(studentInfo["hobby"].([]string)[1])
}
