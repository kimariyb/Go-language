# Go Language Tutorial — 项目手册

## 项目概述

Go 语言入门教程，涵盖从基础语法到高级特性（泛型、并发、反射等）的全部核心知识点。

## 目录结构

```
Go-language/
├── README.md          ← 项目索引（含章节→示例映射表）
├── LICENSE            ← MIT 许可证
├── CLAUDE.md          ← 本文件
├── .gitignore         ← Git 忽略规则
├── docs/              ← 17 章教程文档（Markdown）
│   ├── 01-basics.md
│   ├── 02-types.md
│   ├── ...
│   └── 17-generics.md
└── example/           ← 35 个配套示例（Go module）
    ├── demo01/
    ├── ...
    └── demo35/
```

## 关键约定

- **教程文档**: 位于 `docs/`，按编号 `01-17` 排序，每章一个 `.md` 文件
- **代码示例**: 位于 `example/demoXX/`，每个示例包含 `main.go`，可直接 `go run`
- **文档与示例映射**: 见 README.md 中的对照表
- **文档格式**: GitHub Flavored Markdown，包含 `> [!note]`, `> [!caution]`, `> [!tip]` 等提示块
- **代码块**: 使用 `go` 语言标注

## 开发命令

```bash
# 运行示例
cd example/demoXX && go run main.go

# 运行所有示例（验证不报错）
for d in example/demo*/; do (cd "$d" && go build .); done

# 查看提交状态
git status

# 提交变更
git add -A && git commit -m "描述"
```
