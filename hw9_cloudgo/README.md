# 开发 web 服务程序
作业地址： [https://pmlpml.github.io/ServiceComputingOnCloud/ex-cloudgo-start](https://pmlpml.github.io/ServiceComputingOnCloud/ex-cloudgo-start)

扩展要求博客链接：[https://blog.csdn.net/panghch/article/details/103056515](https://blog.csdn.net/panghch/article/details/103056515)

## 作业要求
- 基本要求
	1. 编程 web 服务程序 类似 cloudgo 应用。
		- 要求有详细的注释
		- 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
	2. 使用 curl 测试，将测试结果写入 README.md
	3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

- 扩展要求
选择以下一个或多个任务，以博客的形式提交。

	1. 选择 net/http 源码，通过源码分析、解释一些关键功能实现
	2. 选择简单的库，如 mux 等，通过源码分析、解释它是如何实现扩展的原理，包括一些 golang 程序设计技巧。
	3. 在 docker hub 申请账号，从 github 构建 cloudgo 的 docker 镜像，最后在 Amazon 云容器服务中部署。
	4. 实现 Github - Travis CI - Docker hub - Amazon “不落地”云软件开发流水线
	5. 其他 web 开发话题

## 实验过程
本次实验基于cloudgo的基础上添加了使用GET方法模拟注册和登录的功能。
PS1：真正的注册和登录绝对不能使用url来传递密码等需要保密的信息，一般会使用POST方法，数据加密放在请求体上，这里只是简单地实验一下简单的url解析。
PS2: 为了简单起见，这次实验并没有进行数据持久化，注册用户数据保存在内存中。

程序设计：
1. url路由
negroni和mux两个库能够让我们轻松地对url进行路由，路由的方法类似于：
```
mx.HandleFunc("/login", loginHandler(formatter)).Methods("GET")
```
其中mx是mux的变量，HandleFunc是路由的函数，第一个参数是路由的路径，第二个参数是路由的接口实现。

基本实现过程为：
```
func NewServer() *negroni.Negroni {

    formatter := render.New(render.Options{
        IndentJSON: true,
    })

    n := negroni.Classic()
    mx := mux.NewRouter()

    initRoutes(mx, formatter)

    n.UseHandler(mx)
    return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
    mx.HandleFunc("/hello/{id}", helloHandler(formatter)).Methods("GET")
    mx.HandleFunc("/register", registerHandler(formatter)).Methods("GET")
    mx.HandleFunc("/login", loginHandler(formatter)).Methods("GET")
}
```

先定义变量，然后将路径与对应的路由实现绑定。

2. 路由接口实现
路由接口实现部分是服务器后端的逻辑部分，所有的算法逻辑基本都在这里实现。

这里使用了render库的Render，官方文档如下：
```
HTML: Uses the html/template package to render HTML templates.
JSON: Uses the encoding/json package to marshal data into a JSON-encoded response.
XML: Uses the encoding/xml package to marshal data into an XML-encoded response.
Binary data: Passes the incoming data straight through to the http.ResponseWriter.
Text: Passes the incoming string straight through to the http.ResponseWriter.
```
官方例子：
```
// main.go
package main

import (
    "encoding/xml"
    "net/http"

    "github.com/unrolled/render"  // or "gopkg.in/unrolled/render.v1"
)

type ExampleXml struct {
    XMLName xml.Name `xml:"example"`
    One     string   `xml:"one,attr"`
    Two     string   `xml:"two,attr"`
}

func main() {
    r := render.New()
    mux := http.NewServeMux()

    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Welcome, visit sub pages now."))
    })

    mux.HandleFunc("/data", func(w http.ResponseWriter, req *http.Request) {
        r.Data(w, http.StatusOK, []byte("Some binary data here."))
    })

    mux.HandleFunc("/text", func(w http.ResponseWriter, req *http.Request) {
        r.Text(w, http.StatusOK, "Plain text here")
    })

    mux.HandleFunc("/json", func(w http.ResponseWriter, req *http.Request) {
        r.JSON(w, http.StatusOK, map[string]string{"hello": "json"})
    })

    mux.HandleFunc("/jsonp", func(w http.ResponseWriter, req *http.Request) {
        r.JSONP(w, http.StatusOK, "callbackName", map[string]string{"hello": "jsonp"})
    })

    mux.HandleFunc("/xml", func(w http.ResponseWriter, req *http.Request) {
        r.XML(w, http.StatusOK, ExampleXml{One: "hello", Two: "xml"})
    })

    mux.HandleFunc("/html", func(w http.ResponseWriter, req *http.Request) {
        // Assumes you have a template in ./templates called "example.tmpl"
        // $ mkdir -p templates && echo "<h1>Hello {{.}}.</h1>" > templates/example.tmpl
        r.HTML(w, http.StatusOK, "example", "World")
    })

    http.ListenAndServe("127.0.0.1:3000", mux)
}
```

render库非常简单易用，本次直接使用纯文本作为服务器返回的类型。

具体逻辑为用户包含用户名和密码，用户名是主键，

注册时判断用户名是否存在，若存在则返回用户名已被使用的信息，否则将用户信息写入数据库，返回注册成功的信息。

登录时判断用户名与密码是否对应，若用户名不存在或密码不正确则返回用户名或密码错误的信息，否则登录成功。

## 实验结果
注册使用GET方法，url为/register, 参数为name和password,类似于
```
http://host/register?name=username&password=123456
```

登录使用GET方法，url为/login, 参数为name和password,类似于
```
http://host/login?name=username&password=123456
```

服务器返回值如上所述。


## 使用 curl 工具访问 web 程序
首先启动服务器,默认端口8080，可使用-p来指定
```
go run main.go
```

### 使用用户名qwe,密码123登录
```
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ curl -v "127.0.0.1:8080/login?name=qwe&password=123"
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /login?name=qwe&password=123 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Wed, 13 Nov 2019 06:50:49 GMT
< Content-Length: 23
<
* Connection #0 to host 127.0.0.1 left intact
login fail! username or password incorrect
```
登录失败，用户名或密码错误

### 使用用户名qwe,密码123注册
```
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ curl -v 127.0.0.1:8080/register?name=qwe&password=123
[1] 525
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ *   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /register?name=qwe HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Wed, 13 Nov 2019 06:51:24 GMT
< Content-Length: 23
<
* Connection #0 to host 127.0.0.1 left intact
user  register success!
```
注册成功。


### 再次使用用户名qwe,密码123注册
```
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ curl -v "127.0.0.1:8080/register?name=qwe&password=123"
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /register?name=qwe&password=123 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Wed, 13 Nov 2019 06:58:09 GMT
< Content-Length: 30
<
* Connection #0 to host 127.0.0.1 left intact
qwe has already been register!
```
注册失败，用户已被注册。

### 再次使用用户名qwe,密码123登录
```
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ curl -v "127.0.0.1:8080/login?name=qwe&password=123"
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
> GET /login?name=qwe&password=123 HTTP/1.1
> Host: 127.0.0.1:8080
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Wed, 13 Nov 2019 07:00:11 GMT
< Content-Length: 23
<
* Connection #0 to host 127.0.0.1 left intact
user qwe login success!
```
登录成功。

## 压力测试
Ubuntu下安装 Apache web 压力测试程序
```
sudo apt install apache2-utils
```

### ab简介 
ab的全称是Apache Bench，是Apache自带的网络压力测试工具

ab命令对发出负载的计算机要求很低，不会占用很高CPU和内存，但也能给目标服务器产生巨大的负载，能实现基础的压力测试。

在进行压力测试时，最好与服务器使用交换机直连，以获取最大的网络吞吐量。

ab命令最基本的参数是-n和-c：

	-n 执行的请求数量
	-c 并发请求个数

其他参数：

	-t 测试所进行的最大秒数
	-p 包含了需要POST的数据的文件
	-T POST数据所使用的Content-type头信息
	-k 启用HTTP KeepAlive功能，即在一个HTTP会话中执行多个请求，默认时，不启用KeepAlive功能

### 测试
对登录功能进行100个并发，1000次请求的测试：
```
penhison@DESKTOP-FQOBQPU:/mnt/c/Users/penhison/Desktop$ ab -n 1000 -c 100 "127.0.0.1:8080/login?name=qwe&password=123"
This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking 127.0.0.1 (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:
Server Hostname:        127.0.0.1
Server Port:            8080

Document Path:          /login?name=qwe&password=123
Document Length:        23 bytes

Concurrency Level:      100
Time taken for tests:   0.650 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      140000 bytes
HTML transferred:       23000 bytes
Requests per second:    1537.65 [#/sec] (mean)
Time per request:       65.034 [ms] (mean)
Time per request:       0.650 [ms] (mean, across all concurrent requests)
Transfer rate:          210.23 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0   12   8.0     13      36
Processing:     8   50  16.0     51      89
Waiting:        8   46  17.1     45      87
Total:         23   62  11.8     63     108

Percentage of the requests served within a certain time (ms)
  50%     63
  66%     66
  75%     69
  80%     70
  90%     78
  95%     82
  98%     87
  99%     90
 100%    108 (longest request)
```

结果解释：
重要参数为：

Time per request:       65.034 [ms] (mean) 用户平均等待时间65.034ms

Time per request:       0.650 [ms] (mean, across all concurrent requests) 服务器每个请求的响应时间0.650ms

Requests per second:    1537.65 [#/sec] (mean) 吞吐率1537.65个请求每秒