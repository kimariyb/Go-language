package main

import "fmt"

type myInt int

func (i myInt) SayHello() {
	fmt.Printf("Hello, World!\n")
}

type Person struct {
	name string
	age  int
	sex  string
}

// PrintInfo 值类型接受者
func (p Person) PrintInfo() {
	fmt.Printf("Person %v, %d, %v\n", p.name, p.age, p.sex)
}

// SetInfo 指针类型接受者
func (p *Person) SetInfo(name string, age int, sex string) {
	p.name = name
	p.age = age
	p.sex = sex
}

type Address struct {
	Province string
	City     string
}

type User struct {
	Name    string
	Gender  string
	Address Address
}

type Animal struct {
	name string
}

func (a Animal) Eat() {
	fmt.Printf("Animal %s eats!\n", a.name)
}

type Dog struct {
	Age int
	Animal
}

func main() {
	// 1. 第一种方法，很笨
	var p1 Person
	p1.name = "张三"
	p1.age = 18
	p1.sex = "男"
	fmt.Println(p1)

	// 2. 第二种方法
	var p2 = new(Person)
	p2.name = "张三"
	p2.age = 18
	p2.sex = "男"
	fmt.Println(*p2)

	// 3. 第三种方法
	var p3 = Person{
		name: "张三",
		age:  18,
		sex:  "男",
	}
	fmt.Println(p3)

	p1.PrintInfo()
	p1.SetInfo("李四", 19, "女")
	p1.PrintInfo()

	var m1 myInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v, %T\n", m1, m1)

	user1 := User{
		Name:   "张三",
		Gender: "男",
		Address: Address{
			Province: "广东",
			City:     "深圳",
		},
	}
	fmt.Printf("%#v\n", user1)

	dog1 := Dog{
		Age: 1,
		Animal: Animal{
			name: "旺财",
		},
	}
	dog1.Eat()
}
