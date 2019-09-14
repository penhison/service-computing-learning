服务计算第二次作业

## 安装 Golang 开发环境

教程为[如何使用Go编程](https://go-zh.org/doc/code.html)

基本按照教程复现了一遍过程，在$GOPATH/github.com/penhison/目录下建立了hello和stringutil两个包，并写了hello.go，reverse.go和reverse_test.go三个文件，github仓库保存在https://github.com/penhison/golang-learning

要点：
1. 建立PATH
```
$ mkdir $HOME/work
$ export GOPATH=$HOME/gowork
$ export PATH=$PATH:$GOPATH/bin
```
这里使用了$HOME目录下的gowork作为GOPATH
gopath的帮助可以在命令行下输入`go help gopath`获取

2. 文件结构
GOPATH目录下的文件夹一般有src,bin,pkg三个，src用于存放源代码，我们主要关注于src目录，本次测试中我使用`github.com/penhison`作为基本路径。

3. 编译并运行文件
首先在基本路径下创建包目录，建立代码文件，如本次实验创建了hello/hello.go,然后使用命令
```
go install github.com/penhison/hello
```
构建并安装hello的包，如果已经在hello文件夹下可以使用`go install`直接构建并安装。

4. 引用其他包
引用包的方法为`import "package path"`如import "github.com/penhison/stringutil",搜索的路径为标准库路径和$GOPATH/src

使用远程库可以先使用`go get path`获取远程包到本地的$GOPATH/src，然后和本地包使用相同的方法导入。

4. 测试
测试文件名必须以`_test`结尾，测试函数以`Test`开头，可运行`go help test` 或从 testing 包文档中查看具体说明。