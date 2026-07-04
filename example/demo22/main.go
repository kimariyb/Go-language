package main
// Go语言教程: 第13章 接口 - 接口实现

import "fmt"

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string， value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type", v)
	}
}

type Usber interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println(p.Name, "手机开始工作")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止工作")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机开始工作")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "相机停止工作")
}

func main() {
	// 创建手机对象
	phone := Phone{"小米"}
	var usbPhone Usber = phone
	usbPhone.Start() // 小米 手机开始工作
	usbPhone.Stop()  // 小米 手机停止工作

	camera := Camera{"尼康"}
	var usbCamera Usber = camera
	usbCamera.Start() // 尼康 相机开始工作
	usbCamera.Stop()  // 尼康 相机停止工作

	// 定义一个空接口 x, x 变量可以接收任意的数据类型
	var x interface{}
	s := "你好 golang"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	b := true
	x = b
	fmt.Printf("type:%T value:%v\n", x, x)

	//var x interface{}
	//x = "Hello world"
	//v, ok := x.(string)
	//if ok {
	//	fmt.Println(v)
	//} else {
	//	fmt.Println("类型断言失败")
	//}

	phone1 := Phone{"小米"}
	var p1 Usber = phone1
	p1.Start() // 小米 手机开始工作
	p1.Stop()  // 小米 手机停止工作

	phone2 := &Phone{"苹果"}
	var p2 Usber = phone2
	p2.Start() // 小米 手机开始工作
	p2.Stop()  // 小米 手机停止工作
}
