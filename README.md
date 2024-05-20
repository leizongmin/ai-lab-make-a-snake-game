# ai-lab-make-a-snake-game

这是一个可在命令行界面上运行的贪吃蛇游戏。本项目现已更新为一个更全面的 Go 项目结构，包括新增的 `pkg`、`internal` 和 `api` 目录，用于组织代码和接口。主 Go 程序仍位于 `cmd/main.go`。`go.mod` 文件已更新以反映新的依赖关系。`.gitignore` 文件现在包括了额外的构建产物和 IDE 文件的模式。

## 项目介绍

本项目是一个在命令行界面上运行的贪吃蛇游戏，旨在提供一个简单而有趣的方式来学习和探索 Go 语言编程。通过参与本项目的开发，您可以了解到 Go 语言的基本语法，模块化编程，以及如何组织和管理一个较为复杂的项目结构。

## 初始化开发环境

要开始开发，您需要先安装 Go 语言环境。请遵循[官方文档](https://golang.org/doc/install)进行安装。安装完成后，您可以通过以下命令来验证安装是否成功：

```bash
go version
```

接下来，克隆本项目到您的本地机器：

```bash
git clone https://github.com/leizongmin/ai-lab-make-a-snake-game.git
```

然后，进入项目目录并运行游戏：

```bash
cd ai-lab-make-a-snake-game
go run cmd/main.go
```

## 游戏控制

游戏支持通过键盘操作进行控制。您可以使用 `ASDW` 键或方向键来控制蛇的移动方向，使用空格键来暂停或开始游戏。

## 游戏结束与重新开始

当游戏结束时，屏幕中央会显示“Game Over”提示。此时，您可以按空格键来重新开始游戏。

## 构建和运行

您可以使用提供的 `Makefile` 来构建项目。运行以下命令来构建可执行文件：

```bash
make dist
```

构建完成后，您可以在 `dist` 目录下找到可执行文件，运行它来启动游戏：

```bash
./dist/snake-game
```

## 贡献代码

我们欢迎任何形式的贡献，无论是新功能，bug 修复，或是文档改进。请首先通过 issue 讨论您想要做的改动，然后您可以 fork 本仓库，进行您的改动，并提交 pull request。我们会尽快进行审查。
