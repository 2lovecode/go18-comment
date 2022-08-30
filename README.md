# Golang1.18 中文注释版

## 说明
1. 新增的example目录，用来放置go测试代码
2. 新增的notes目录，用来放置阅读源码过程中的笔记

## 怎么配置环境，让源码读起来更方便
1. 建议使用goland
    - jetbrains给开源项目作者开放了免费申请注册码通道[jetbrains开源支持申请](https://www.jetbrains.com.cn/community/opensource/#support)
2. 配置环境变量
    - GOROOT设置为当前项目目录，即path/to/go18-comment
    - goland设置GOROOT时可能会报错，因为当前目录下还没有bin/go，我们需要先编译下当前项目，执行下面命令即可
      - cd src
      - ./make.bash
    - GOPATH也设置为当前项目目录，即path/to/go18-comment

## 怎么使用当前编译的bin/go，运行我们的测试代码
1. 进入src目录
2. 执行 ../bin/go run ../example/hello.go

## 笔记目录
#### [当执行go run xxx.go时发生了什么？](./notes/go_run.md)


--------

# The Go Programming Language

Go is an open source programming language that makes it easy to build simple,
reliable, and efficient software.

![Gopher image](https://golang.org/doc/gopher/fiveyears.jpg)
*Gopher image by [Renee French][rf], licensed under [Creative Commons 3.0 Attributions license][cc3-by].*

Our canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the
BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.

After downloading a binary release, visit https://golang.org/doc/install
for installation instructions.

#### Install From Source

If a binary distribution is not available for your combination of
operating system and architecture, visit
https://golang.org/doc/install/source
for source installation instructions.

### Contributing

Go is the work of thousands of contributors. We appreciate your help!

To contribute, please read the contribution guidelines at https://golang.org/doc/contribute.

Note that the Go project uses the issue tracker for bug reports and
proposals only. See https://golang.org/wiki/Questions for a list of
places to ask questions about the Go language.

[rf]: https://reneefrench.blogspot.com/
[cc3-by]: https://creativecommons.org/licenses/by/3.0/
