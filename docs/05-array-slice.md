# 5. 数组

数组是指一系列同一类型数据的集合。数组中包含的每个数据被称为数组元素，这种类型可以是任意的原始类型，比如 int、 string 等，也可以是用户自定义的类型。 

一个数组包含的元素个数被称为数组的长度。 

在 Golang 中数组是一个长度固定的数据类型，数组的长度是类型的一部分，也就是说 `[5]int` 和 `[10]int` 是两个不同的类型。 

## 5.1 数组的声明

```go
var a [3]int
var b [4]int
```

数组的长度必须是常量，并且长度是数组类型的一部分。一旦定义，**长度不能变**。数组可以通过下标进行访问，下标是从 `0` 开始，最后一个元素下标是 `len-1`  

## 5.2 数组的初始化

初始化数组时可以使用初始化列表来设置数组元素的值。

```go
var testArray [5]int
var numArray = [3]int{1, 2, 3}
var cityArray = [3]string{"北京", "上海", "广州"}

fmt.Println(testArray) // [0 0 0 0 0]
fmt.Println(numArray) // [1 2 3]
fmt.Println(cityArray) // [北京 上海 广州]
```

除此之外，还可以让编译器根据初始值的个数自行推断数组的长度。

```go
var testArray = [...]int{}
var numArray = [...]int{1, 2, 3}
var cityArray = [...]string{"北京", "上海", "广州"}

fmt.Println(testArray)
fmt.Println(numArray)
fmt.Println(cityArray)
```

我们还可以使用指定索引值的方式来初始化数组。

```go
a := [...]int{1: 1, 3: 5}
fmt.Println(a) // [0 1 0 5]
fmt.Printf("type of a:%T\n", a) // type of a:[4]int
```

## 5.3 数组的遍历

```go
// 1. for 循环遍历
for i := 0; i < len(array); i++ {
    fmt.Println(array[i])
}

// 2. for range 遍历
for _, v := range array {
    fmt.Println(v)
}
```

## 5.4 数组是值类型

数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

```go
func modifyArray(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func main() {
    a := [3]int{10, 20, 30}
    modifyArray(a) // 在 modify 中修改的是 a 的副本 x
    fmt.Println(a) // [10 20 30]
    
    b := [3][2]int{
        {1, 1},
        {1, 1},
        {1, 1},
    } 
    modifyArray2(b) // 在 modify 中修改的是 b 的副本 x
    fmt.Println(b) // [[1 1] [1 1] [1 1]]
}
```

> [!note]
>
> 1. 数组支持 `==`、 `!=` 操作符，因为内存总是被初始化过的。
> 2. `[n]*T` 表示指针数组，`*[n]T` 表示数组指针。

