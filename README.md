# Go Language Tutorial

<div align="center">

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?logo=go)](https://go.dev)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)]()

**从零开始的 Go 语言学习教程 — 全面的基础知识 + 丰富的代码示例**

</div>

---

## 📖 简介

本项目是一份完整的 Go 语言入门教程，涵盖从基础语法到高级特性的全部核心知识点。每个章节都配有详细的说明和可直接运行的代码示例，适合 Go 语言初学者系统学习，也适合有经验的开发者快速查阅。

## 🗂️ 目录结构

```
Go-language/
├── README.md              ← 项目索引（本文件）
├── LICENSE                ← MIT 许可证
├── CLAUDE.md              ← 项目文档配置
├── docs/                  ← 17 章独立教程文档
│   ├── 01-basics.md       ← 变量和常量
│   ├── 02-types.md        ← 基本数据类型
│   ├── 03-operators.md    ← 运算符
│   ├── 04-control-flow.md ← 流程控制
│   ├── 05-array-slice.md  ← 数组
│   ├── 06-slice.md        ← 切片
│   ├── 07-map.md          ← 哈希表
│   ├── 08-functions.md    ← 函数
│   ├── 09-pointer.md      ← 指针
│   ├── 10-struct.md       ← 结构体
│   ├── 11-json.md         ← JSON
│   ├── 12-package.md      ← 包
│   ├── 13-interface.md    ← 接口
│   ├── 14-goroutine.md    ← 协程
│   ├── 15-reflection.md   ← 反射
│   ├── 16-io.md           ← IO 流
│   └── 17-generics.md     ← 泛型
└── example/               ← 35 个配套代码示例
    ├── demo01/ ~ demo35/
    └── (每个示例包含 main.go)
```

## 📚 教程目录

| # | 章节 | 文档 | 配套示例 |
|---|------|------|----------|
| 1 | **变量和常量** — var、const、短变量声明、匿名变量 | [📄 查看](docs/01-basics.md) | [demo01](example/demo01) |
| 2 | **基本数据类型** — 整型、浮点、布尔、字符串、字符、类型转换 | [📄 查看](docs/02-types.md) | [demo02~demo07](example/demo02) |
| 3 | **运算符** — 算术、关系、逻辑、短路运算 | [📄 查看](docs/03-operators.md) | [demo08~demo09](example/demo08) |
| 4 | **流程控制** — if/else、for、for range、switch、break/continue/goto | [📄 查看](docs/04-control-flow.md) | [demo10~demo11](example/demo10) |
| 5 | **数组** — 声明、初始化、遍历、值类型特性 | [📄 查看](docs/05-array-slice.md) | [demo12](example/demo12) |
| 6 | **切片** — 声明、扩容、append/copy、排序算法 | [📄 查看](docs/06-slice.md) | [demo13](example/demo13) |
| 7 | **哈希表** — map 声明、增删改查、遍历 | [📄 查看](docs/07-map.md) | [demo14](example/demo14) |
| 8 | **函数** — 参数、返回值、闭包、defer、panic/recover | [📄 查看](docs/08-functions.md) | [demo15~demo17](example/demo15) |
| 9 | **指针** — & 和 \*、new/make 内存分配 | [📄 查看](docs/09-pointer.md) | [demo18](example/demo18) |
| 10 | **结构体** — 定义、方法、接收者、嵌套、继承 | [📄 查看](docs/10-struct.md) | [demo19](example/demo19) |
| 11 | **JSON** — 序列化/反序列化、Tag、嵌套结构体 | [📄 查看](docs/11-json.md) | [demo20](example/demo20) |
| 12 | **包** — go mod、自定义包、init 函数、第三方包 | [📄 查看](docs/12-package.md) | [demo21](example/demo21) |
| 13 | **接口** — 定义、空接口、类型断言、接口嵌套 | [📄 查看](docs/13-interface.md) | [demo22~demo23](example/demo22) |
| 14 | **协程** — goroutine、channel、select、并发安全、锁 | [📄 查看](docs/14-goroutine.md) | [demo24~demo29](example/demo24) |
| 15 | **反射** — TypeOf、ValueOf、结构体反射、修改值 | [📄 查看](docs/15-reflection.md) | [demo30~demo32](example/demo30) |
| 16 | **IO 流** — 文件读写、bufio、ioutil | [📄 查看](docs/16-io.md) | [demo33~demo34](example/demo33) |
| 17 | **泛型** — 类型参数、约束、泛型接口 | [📄 查看](docs/17-generics.md) | [demo35](example/demo35) |

## 🚀 快速开始

### 运行示例

```bash
# 进入任意示例目录
cd example/demo01

# 直接运行
go run main.go
```

### 学习建议

1. 按照 **教程目录** 的顺序逐章学习
2. 每学完一章，运行对应的 demo 示例加深理解
3. 尝试修改示例代码，观察输出变化
4. 结合文档和代码双重学习效果最佳

## 🔗 示例与章节对照表

| 示例 | 对应章节 | 核心知识点 |
|------|----------|-----------|
| demo01 | Ch1 变量和常量 | var、短变量声明、常量 |
| demo02 | Ch2 基本数据类型 | 浮点精度丢失 |
| demo03 | Ch2 基本数据类型 | 字符串修改 (byte/rune) |
| demo04 | Ch2 基本数据类型 | 数值类型转换 |
| demo05 | Ch2 基本数据类型 | fmt.Sprintf 转换 |
| demo06 | Ch2 基本数据类型 | strconv 包使用 |
| demo07 | Ch2 基本数据类型 | string 转数值 |
| demo08 | Ch3 运算符 | 算术运算符 |
| demo09 | Ch3 运算符 | 关系运算符 |
| demo10 | Ch4 流程控制 | for 循环 + break |
| demo11 | Ch4 流程控制 | switch case |
| demo12 | Ch5 数组 | 数组初始化与遍历 |
| demo13 | Ch6 切片 | 切片删除操作 |
| demo14 | Ch7 map | 哈希表操作 |
| demo15 | Ch8 函数 | 可变参数 |
| demo16 | Ch8 函数 | 闭包 |
| demo17 | Ch8 函数 | panic/recover |
| demo18 | Ch9 指针 | 指针取值 |
| demo19 | Ch10 结构体 | 方法、值/指针接收者 |
| demo20 | Ch11 JSON | 序列化与反序列化 |
| demo21 | Ch12 包 | go mod + 自定义包 |
| demo22 | Ch13 接口 | 接口实现 |
| demo23 | Ch13 接口 | 多接口实现 |
| demo24 | Ch14 协程 | GOMAXPROCS 设置 |
| demo25 | Ch14 协程 | 管道 channel |
| demo26 | Ch14 协程 | 并发素数计算 |
| demo27 | Ch14 协程 | 单向管道 |
| demo28 | Ch14 协程 | 互斥锁 |
| demo29 | Ch14 协程 | 读写锁 |
| demo30 | Ch15 反射 | TypeOf / ValueOf |
| demo31 | Ch15 反射 | 结构体字段与方法反射 |
| demo32 | Ch15 反射 | 反射修改结构体 |
| demo33 | Ch16 IO 流 | 文件读取 |
| demo34 | Ch16 IO 流 | 文件写入 |
| demo35 | Ch17 泛型 | 泛型函数与约束 |

## 📄 许可证

本项目基于 [MIT License](LICENSE) 开源。

## 🤝 贡献指南

欢迎提交 Pull Request 或 Issue 来完善本教程！

- 修正文档错误
- 改进代码示例
- 补充新的知识点
- 优化文档排版
