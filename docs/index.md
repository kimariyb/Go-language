# Go Language Tutorial

<div align="center">

从零开始的 Go 语言入门教程，涵盖从基础语法到高级特性的全部核心知识点。

</div>

---

## 📖 关于本教程

本教程是一份完整的 Go 语言学习资料，适合以下读者：

- **Go 语言初学者** — 从零开始，系统学习
- **有其他语言经验的开发者** — 快速掌握 Go 语法
- **有经验的 Go 开发者** — 查阅核心知识点

## 📚 章节导航

| # | 章节 | 内容概要 |
|---|------|----------|
| 1 | **[变量和常量](01-basics.md)** | var、const、短变量声明、匿名变量 |
| 2 | **[基本数据类型](02-types.md)** | 整型、浮点、布尔、字符串、字符、类型转换 |
| 3 | **[运算符](03-operators.md)** | 算术、关系、逻辑、短路运算 |
| 4 | **[流程控制](04-control-flow.md)** | if/else、for、for range、switch、break/continue/goto |
| 5 | **[数组](05-array-slice.md)** | 声明、初始化、遍历、值类型特性 |
| 6 | **[切片](06-slice.md)** | 声明、扩容、append/copy、排序算法 |
| 7 | **[哈希表](07-map.md)** | map 声明、增删改查、遍历 |
| 8 | **[函数](08-functions.md)** | 参数、返回值、闭包、defer、panic/recover |
| 9 | **[指针](09-pointer.md)** | & 和 \*、new/make 内存分配 |
| 10 | **[结构体](10-struct.md)** | 定义、方法、接收者、嵌套、继承 |
| 11 | **[JSON](11-json.md)** | 序列化/反序列化、Tag、嵌套结构体 |
| 12 | **[包](12-package.md)** | go mod、自定义包、init 函数、第三方包 |
| 13 | **[接口](13-interface.md)** | 定义、空接口、类型断言、接口嵌套 |
| 14 | **[协程](14-goroutine.md)** | goroutine、channel、select、并发安全、锁 |
| 15 | **[反射](15-reflection.md)** | TypeOf、ValueOf、结构体反射、修改值 |
| 16 | **[IO 流](16-io.md)** | 文件读写、bufio、ioutil |
| 17 | **[泛型](17-generics.md)** | 类型参数、约束、泛型接口 |

## 🚀 快速开始

所有代码示例都可以在 `example/` 目录下找到，每个示例可以直接运行：

```bash
cd example/demo01
go run main.go
```

## 🛠️ 示例索引

| 示例 | 对应章节 | 知识点 |
|------|----------|--------|
| [demo01](https://github.com/kimariyb/Go-language/tree/main/example/demo01) | 第1章 | var、短变量声明、常量 |
| [demo02](https://github.com/kimariyb/Go-language/tree/main/example/demo02) | 第2章 | 浮点精度丢失 |
| [demo03](https://github.com/kimariyb/Go-language/tree/main/example/demo03) | 第2章 | 字符串修改 |
| [demo04](https://github.com/kimariyb/Go-language/tree/main/example/demo04) | 第2章 | 数值类型转换 |
| [demo05](https://github.com/kimariyb/Go-language/tree/main/example/demo05) | 第2章 | fmt.Sprintf |
| [demo06](https://github.com/kimariyb/Go-language/tree/main/example/demo06) | 第2章 | strconv 包 |
| [demo07](https://github.com/kimariyb/Go-language/tree/main/example/demo07) | 第2章 | string 转数值 |
| [demo08](https://github.com/kimariyb/Go-language/tree/main/example/demo08) | 第3章 | 算术运算符 |
| [demo09](https://github.com/kimariyb/Go-language/tree/main/example/demo09) | 第3章 | 关系运算符 |
| [demo10](https://github.com/kimariyb/Go-language/tree/main/example/demo10) | 第4章 | for 循环 + break |
| [demo11](https://github.com/kimariyb/Go-language/tree/main/example/demo11) | 第4章 | switch case |
| [demo12](https://github.com/kimariyb/Go-language/tree/main/example/demo12) | 第5章 | 数组初始化与遍历 |
| [demo13](https://github.com/kimariyb/Go-language/tree/main/example/demo13) | 第6章 | 切片删除操作 |
| [demo14](https://github.com/kimariyb/Go-language/tree/main/example/demo14) | 第7章 | 哈希表操作 |
| [demo15](https://github.com/kimariyb/Go-language/tree/main/example/demo15) | 第8章 | 可变参数 |
| [demo16](https://github.com/kimariyb/Go-language/tree/main/example/demo16) | 第8章 | 闭包 |
| [demo17](https://github.com/kimariyb/Go-language/tree/main/example/demo17) | 第8章 | panic/recover |
| [demo18](https://github.com/kimariyb/Go-language/tree/main/example/demo18) | 第9章 | 指针取值 |
| [demo19](https://github.com/kimariyb/Go-language/tree/main/example/demo19) | 第10章 | 方法、值/指针接收者 |
| [demo20](https://github.com/kimariyb/Go-language/tree/main/example/demo20) | 第11章 | JSON 序列化与反序列化 |
| [demo21](https://github.com/kimariyb/Go-language/tree/main/example/demo21) | 第12章 | go mod + 自定义包 |
| [demo22](https://github.com/kimariyb/Go-language/tree/main/example/demo22) | 第13章 | 接口实现 |
| [demo23](https://github.com/kimariyb/Go-language/tree/main/example/demo23) | 第13章 | 多接口实现 |
| [demo24](https://github.com/kimariyb/Go-language/tree/main/example/demo24) | 第14章 | GOMAXPROCS |
| [demo25](https://github.com/kimariyb/Go-language/tree/main/example/demo25) | 第14章 | 管道 channel |
| [demo26](https://github.com/kimariyb/Go-language/tree/main/example/demo26) | 第14章 | 并发素数计算 |
| [demo27](https://github.com/kimariyb/Go-language/tree/main/example/demo27) | 第14章 | 单向管道 |
| [demo28](https://github.com/kimariyb/Go-language/tree/main/example/demo28) | 第14章 | 互斥锁 |
| [demo29](https://github.com/kimariyb/Go-language/tree/main/example/demo29) | 第14章 | 读写锁 |
| [demo30](https://github.com/kimariyb/Go-language/tree/main/example/demo30) | 第15章 | TypeOf / ValueOf |
| [demo31](https://github.com/kimariyb/Go-language/tree/main/example/demo31) | 第15章 | 结构体字段与方法反射 |
| [demo32](https://github.com/kimariyb/Go-language/tree/main/example/demo32) | 第15章 | 反射修改结构体 |
| [demo33](https://github.com/kimariyb/Go-language/tree/main/example/demo33) | 第16章 | 文件读取 |
| [demo34](https://github.com/kimariyb/Go-language/tree/main/example/demo34) | 第16章 | 文件写入 |
| [demo35](https://github.com/kimariyb/Go-language/tree/main/example/demo35) | 第17章 | 泛型函数与约束 |

---

> 💡 **学习建议**: 按章节顺序阅读，每章配合对应的 demo 示例运行效果更佳。
