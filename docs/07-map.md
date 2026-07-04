# 7. 哈希表

## 7.1 哈希表的声明

Go 语言的哈希表可以通过 `map` 实现，`map` 是一种无序的基于 key-value 的数据结构，Go 语言中的 `map` 是引用类型，必须初始化才能使用。

```go
map[KeyType]ValueType
```

其中:

- `KeyType`:表示键的类型。
- `ValueType`:表示键对应的值的类型。

> [!note]
>
> `map` 类型的变量默认初始值为 `nil`，需要使用 `make()` 函数来分配内存。 

```go
make(map[KeyType]ValueType, [cap])
```

> [!caution]
>
> 获取 `map` 的容量不能使用 `cap`，`cap` 返回的是数组切片分配的空间大小，根本不能用于 `map`。要获取 `map`  的容量，可以用 `len()` 函数。

`map` 中的数据都是成对出现的，`map` 的基本使用示例代码如下：

```go
hashMap := make(map[string]int)
hashMap["a"] = 1
hashMap["b"] = 2
hashMap["c"] = 3
fmt.Println(hashMap) // map[a:1 b:2 c:3]
fmt.Println(hashMap["a"]) // 1
```

`map` 也支持在声明的时候填充元素，例如：

```go
userInfo := map[string]string{
    "username": "zhangsan",
    "password": "123456",
}
fmt.Println(userInfo) // map[password:123456 username:zhangsan]
```

## 7.2 判断键是否存在

Go 语言中有个判断 `map` 中键是否存在的特殊写法，格式如下:

```go
value, ok := map[key]
```

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80

// 如果 key 存在 ok 为 true,v 为对应的值；
// 不存在 ok 为 false,v 为值类型的零值
v, ok := scoreMap["张三"]
if ok {
    fmt.Println(v) // 90
} else {
    fmt.Println("没有该用户")
}
```

## 7.3 哈希表的遍历

Go 语言中使用 `for range` 遍历 `map`。  

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80
scoreMap["王五"] = 70

for k, v := range scoreMap {
    fmt.Println(k, v)
}
```

我们只想遍历 key 的时候，可以按下面的写法：

```go
for k := range scoreMap {
    fmt.Println(k)
}
```

> [!caution]
>
> 遍历 `map` 时的元素顺序与添加键值对的顺序无关。  

## 7.4 delete() 方法

使用 `delete()` 内建函数从 `map` 中删除一组键值对，`delete()` 函数的格式如下：

```go
delete(map, key)
```

其中：

- `map`：表示要删除键值对的 `map` 对象
- `key`：表示要删除的键值对的键

```go
scoreMap := make(map[string]int)
scoreMap["张三"] = 90
scoreMap["李四"] = 80
scoreMap["王五"] = 70
fmt.Println(scoreMap) // map[张三:90 李四:80 王五:70]

// delete 方法删除
delete(scoreMap, "张三")
fmt.Println(scoreMap) // map[李四:80 王五:70]
```

