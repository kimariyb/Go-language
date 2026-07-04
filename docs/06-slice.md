# 6. 切片

切片（Slice）是一个拥有相同类型元素的可变长度的序列。它是基于数组类型做的一层封装。它非常灵活，支持自动扩容。

**切片是一个引用类型，它的内部结构包含地址、长度和容量。**  

## 6.1 切片的声明

```go
var a []string
var b []int = []int{1, 2, 3}

fmt.Println(a == nil) //  true
fmt.Println(b == nil) //  false
```

> [!note]
>
> 当你声明了一个切片，却没有给其赋值，此时该切片为默认值 `nil`

> [!important]
>
> 当你声明了一个变量，但却还并没有赋值时，Golang 中会自动给你的变量赋值一个默认零值。 

| 类型类别     | 具体类型                                                     | 零值（默认值） |
| :----------- | :----------------------------------------------------------- | :------------- |
| **布尔型**   | `bool`                                                       | `false`        |
| **整型**     | `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8`, `uint16`, `uint32`, `uint64` | `0`            |
| **浮点型**   | `float32`, `float64`                                         | `0.0`          |
| **复数型**   | `complex64`, `complex128`                                    | `(0+0i)`       |
| **字符串**   | `string`                                                     | `""`           |
| **数组**     | `[n]T`                                                       | `[n]T{}`       |
| **切片**     | `[]T`                                                        | `nil`          |
| **映射**     | `map[K]V`                                                    | `nil`          |
| **通道**     | `chan T`                                                     | `nil`          |
| **指针**     | `*T`                                                         | `nil`          |
| **函数**     | `func(...)`                                                  | `nil`          |
| **接口**     | `interface{}`                                                | `nil`          |
| **结构体**   | `struct{}`                                                   | 各字段的零值   |
| **字节类型** | `byte` (`uint8` 别名)                                        | `0`            |
| **符文类型** | `rune` (`int32`别名)                                         | `0`            |
| **错误类型** | `error`                                                      | `nil`          |

## 6.2 切片的遍历

切片的遍历与数组类似，可以使用 `for` 或者 `for range` 遍历。

```go
var strSlice []string = []string{"hello", "world", "golang"}

// 1. 使用 for 遍历切片
for i := 0; i < len(strSlice); i++ {
    fmt.Println(strSlice[i])
}

// 2. 使用 for range 遍历切片
for _, v := range strSlice {
    fmt.Println(v)
}
```

## 6.3 基于数组定义切片

由于切片的底层就是一个数组，所以我们可以基于数组定义切片。  

```go
numArray := [5]int{1, 2, 3, 4, 5}
numSlice := numArray[1:3] // 基于数组创建切片
fmt.Println(numSlice) // [2 3]
```

## 6.4 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的 `len()` 函数求长度，使用内置的 `cap()` 函数求切片的容量。

- 切片的长度就是**它所包含的元素个数**。  
- 切片的容量是从**它的第一个元素开始数，到其底层数组元素末尾的个数**。

```go
s := []int{2, 3, 5, 7, 11, 13}
fmt.Println(s)
fmt.Printf("长度：%v，容量：%v\n", len(s), cap(s))

c := s[:2]
fmt.Println(c)
fmt.Printf("长度：%v，容量：%v\n", len(c), cap(c))

d := s[1:3]
fmt.Println(d)
fmt.Printf("长度：%v，容量：%v\n", len(d), cap(d))

// 输出：
// [2 3 5 7 11 13]
// 长度：6，容量：6
// [2 3]
// 长度：2，容量：6
// [3 5]
// 长度：2，容量：5
```

## 6.5 使用 make() 函数声明

如果需要动态的创建一个切片，我们就需要使用内置的 `make()` 函数。

```go
make([]T, size, cap)
```

其中：

- `[]T`：切片的元素类型
- `size`：切片中的元素数量
- `cap`：切片的容量

```go
var a = make([]int, 2, 10)
fmt.Println(a) // [0 0]
fmt.Println(len(a)) // 2
fmt.Println(cap(a)) // 10
```

> [!caution]
>
> 切片之间是不能比较的，我们不能使用 `==` 操作符来判断两个切片是否含有全部相等元素。切片唯一合法的比较操作是和 `nil` 比较。 

## 6.6 切片的基础操作

### 6.6.1 append() 方法

当底层数组不能容纳新增的元素时，切片就会自动按照一定的策略进行扩容，此时该切片指向的底层数组就会更换。扩容操作往往发生在`append()` 函数调用时，所以我们通常都需要用原变量接收 `append()` 函数的返回值。

错误写法：

```go
s := []int{1, 2, 3, 5, 6, 7}
s[6] = 8
fmt.Println(s) //index out of range [6] with length 6
```

正确写法：

```go
var numSlice []int
for i := 0; i < 10; i++ {
    numSlice = append(numSlice, i)
}
```

1. `append()` 函数将元素追加到切片的最后并返回该切片。  
2. 切片 `numSlice` 的容量按照 1，2，4，8，16 这样的规则自动进行扩容，每次扩容后都是扩容前的 2 倍。

```go
var citySlice []string
// 添加一个元素
citySlice = append(citySlice, "北京")

// 添加多个元素
citySlice = append(citySlice, "上海", "广州", "深圳")

// 添加切片
citySlice = append(citySlice, []string{"杭州", "苏州"}...)
fmt.Println(citySlice) // [北京 上海 广州 深圳 杭州 苏州]
```

### 6.6.2 copy() 方法

```go
a := []int{1, 2, 3, 4, 5}
b := a
fmt.Println(a) //[1 2 3 4 5]
fmt.Println(b) //[1 2 3 4 5]
b[0] = 1000
fmt.Println(a) //[1000 2 3 4 5]
fmt.Println(b) //[1000 2 3 4 5]
```

由于切片是引用类型，所以在修改 b 的同时 a 的值也会发生变化。Go 语言内建的 `copy()` 函数可以迅速地将一个切片的数据复制到另外一个切片空间中，`copy()` 函数的使用格式如下：

```go
copy(destSlice, srcSlice []T)
```

其中：

- `destSlice`：目标切片
- `srcSlice`：数据来源切片

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := make([]int, 5)

copy(s2, s1) // 将 s1 拷贝到 s2
fmt.Println(s1) // [1 2 3 4 5]
fmt.Println(s2) // [1 2 3 4 5]

s2[0] = 100
fmt.Println(s1) // [1 2 3 4 5]
fmt.Println(s2) // [100 2 3 4 5]
```

### 6.6.3 删除元素

Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素。要从切片 a 中删除索引为 index 的元素，操作方法是

```go 
a = append(a[:index], a[index+1:]...)
```

```go
func popString(srcSlice []string, index int) []string {
	// 删除切片中的元素
	return append(srcSlice[:index], srcSlice[index+1:]...)
}
```

## 6.7 切片排序

### 6.7.1 冒泡排序

核心思想：**相邻比较和交换**，就像水中气泡上浮一样，它通过反复比较相邻元素，将较大的元素逐步向后移动，每一轮都会将当前未排序部分的最大值“冒泡”到正确位置，直到整个序列有序。

```go
func bubbleSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 冒泡排序
	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟冒泡
		isSwap := false
		for j := 0; j < len(numSlice)-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				numSlice[j], numSlice[j+1] = numSlice[j+1], numSlice[j]
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
	return numSlice
}
```

### 6.7.2 选择排序

核心思想：**选取最小元素**，它把序列分为已排序和未排序两部分，每次从未排序部分中选出最小的元素，直接将其放到已排序部分的末尾，如同打擂台选冠军一样逐个确定位置，直到所有元素归位。

```go
func selecSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	for i := 0; i < len(numSlice)-1; i++ {
		// 第 i 趟选择
		minIndex := i
		// 找到最小值的索引
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[j] < numSlice[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
            // 将未排序区间最小值交换到已排序区间的末尾
			numSlice[i], numSlice[minIndex] = numSlice[minIndex], numSlice[i]
		}
	}

	return numSlice
}
```

### 6.7.3 插入排序

核心思想：**构建有序序列**，就像整理手中的扑克牌，它将序列视为已排序和未排序两部分，每次从未排序部分取出第一个元素，在已排序部分中从后向前扫描，找到合适位置插入，从而逐步扩大有序区间，直到全部有序。

```go
func insertSort(numSlice []int) []int {
	// 边界条件
	if len(numSlice) <= 1 {
		return numSlice
	}

	// 遍历无序区间，从第二个元素开始遍历
	for i := 1; i < len(numSlice); i++ {
		// 保存当前元素
		current := numSlice[i]
		// 找到当前元素应该插入的位置
		j := i - 1

		// 如果 j >= 0 且 numSlice[j] > current，
        // 则说明当前元素在有序区间中，则将当前元素插入有序区间中
		for j >= 0 && numSlice[j] > current {
			numSlice[j+1] = numSlice[j]
			j--
		}
		numSlice[j+1] = current
	}
	return numSlice

}
```

### 6.7.4 sort 包

对于 `int`、 `float64` 和 `string` 数组或是切片的排序， go 分别提供了 `sort.Ints()`、`sort.Float64s()` 和 `sort.Strings()` 函数，默认都是从小到大排序。

```go
intList := []int{2, 4, 3, 5, 7, 6, 9, 8, 1, 0}
floatList := []float64{4.2, 5.9, 12.4, 10.2, 50.7, 99.9, 31.4, 27.81828, 3.14}
stringList := []string{"a", "c", "b", "z", "x", "w", "y", "d", "f", "i"}

sort.Ints(intList)
sort.Float64s(floatList)
sort.Strings(stringList)
fmt.Println(intList) // [0 1 2 3 4 5 6 7 8 9]
fmt.Println(floatList) // [3.14 4.2 5.9 10.2 12.4 27.81828 31.4 50.7 99.9]
fmt.Println(stringList) // [a b c d f i w x y z]
```

Golang 的 `sort` 包可以使用 `sort.Reverse(slice)` 来调换 `slice.Interface.Less`，所以逆序排序函数可以这么写。

```go
sort.Sort(sort.Reverse(sort.IntSlice(intList)))
fmt.Println(intList) // [9 8 7 6 5 4 3 2 1 0]
```

