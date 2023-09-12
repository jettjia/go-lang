# Grpc介绍



gRPC 一开始由 Google 开发，是一款语言中立、平台中立、开源的远程过程调用(RPC)系统。

在 gRPC 里客户端应用可以像调用本地对象一样直接调用另一台不同的机器上服务端应用的方法，使得您能够更容易地创建分布式应用和服务。与许多 RPC 系统类似，gRPC 也是基于以下理念：定义一个服务，指定其能够被远程调用的方法（包含参数和返回类型）。在服务端实现这个接口，并运行一个 gRPC 服务器来处理客户端调用。在客户端拥有一个存根能够像服务端一样的方法。



### gRPC的特性

#### 基于HTTP/2

HTTP/2 提供了连接多路复用、双向流、服务器推送、请求优先级、首部压缩等机制。可以节省带宽、降低TCP链接次数、节省CPU，帮助移动设备延长电池寿命等。gRPC 的协议设计上使用了HTTP2 现有的语义，请求和响应的数据使用HTTP Body 发送，其他的控制信息则用Header 表示。

#### IDL使用ProtoBuf

gRPC使用ProtoBuf来定义服务，ProtoBuf是由Google开发的一种数据序列化协议（类似于XML、JSON、hessian）。ProtoBuf能够将数据进行序列化，并广泛应用在数据存储、通信协议等方面。压缩和传输效率高，语法简单，表达力强。

#### 多语言支持

gRPC支持多种语言（C, C++, Python, PHP, Nodejs, C#, Objective-C、Golang、Java），并能够基于语言自动生成客户端和服务端功能库。目前已提供了C版本grpc、Java版本grpc-java 和 Go版本grpc-go，其它语言的版本正在积极开发中，其中，grpc支持C、C++、Node.js、Python、Ruby、Objective-C、PHP和C#等语言，grpc-java已经支持Android开发。

### gRPC优缺点

#### 优点

- protobuf二进制消息，性能好/效率高（空间和时间效率都很不错）
- proto文件生成目标代码，简单易用
- 序列化反序列化直接对应程序中的数据类，不需要解析后在进行映射(XML,JSON都是这种方式)
- 支持向前兼容（新加字段采用默认值）和向后兼容（忽略新加字段），简化升级
- 支持多种语言（可以把proto文件看做IDL文件）
- Netty等一些框架集成

#### 缺点：

- GRPC尚未提供连接池，需要自行实现
- 尚未提供“服务发现”、“负载均衡”机制
- 因为基于HTTP2，绝大部多数HTTP Server、Nginx都尚不支持，即Nginx不能将GRPC请求作为HTTP请求来负载均衡，而是作为普通的TCP请求。（nginx1.9版本已支持）
- Protobuf二进制可读性差（貌似提供了Text_Fromat功能）
  默认不具备动态特性（可以通过动态定义生成消息类型或者动态编译支持）

# 安装环境

## 安装 protobuf

1.下载地址：https://github.com/protocolbuffers/protobuf/releases

根据自身电脑的操作系统，选择最新的releases版本下载

 ![image-20210629221003318](images/image-20210629221003318.png)

2.解压后在bin目录找到protoc.exe，然后把它复制到GOBIN目录下

> 一般操作是把protoc.exe所在的目录配到环境变量里，这里直接把protoc.exe复制到GOBIN目录下，前提是环境变量已经配置了GOBIN环境变量。

3.打开cmd，运行`protoc --version`

成功打印当前版本信息证明安装成功了。

## 安装相关包

安装 golang 的proto工具包

```shell
go get -u github.com/golang/protobuf/proto
```

安装 goalng 的proto编译支持

```shell
go get -u github.com/golang/protobuf/protoc-gen-go

# 这里在win环境，需要将gopath里生成的 
# E:\web\go_work\wingopath\bin\protoc-gen-go.exe
# 拷贝到, go的bin目录下
# E:\dev\go\bin
```

安装 gRPC 包

```
go get -u google.golang.org/grpc
```

## 创建并编译proto文件

1.新建proto文件夹，在里面新建simple.proto文件

E:\web\go_work\wingopath\go-study\src\microtest\protof\simple.proto

```go
 syntax = "proto3";// 协议为proto3

package proto;

// 定义发送请求信息
message SimpleRequest{
    // 定义发送的参数
    // 参数类型 参数名 标识号(不可重复)
    string data = 1;
}

// 定义响应信息
message SimpleResponse{
    // 定义接收的参数
    // 参数类型 参数名 标识号(不可重复)
    int32 code = 1;
    string value = 2;
}

// 定义我们的服务（可定义多个服务,每个服务可定义多个接口）
service Simple{
    rpc Route (SimpleRequest) returns (SimpleResponse){};
}
```

2.编译proto文件

cmd进入simple.proto所在目录，运行以下指令进行编译

```
protoc --go_out=plugins=grpc:./ ./simple.proto
```

注意：

```
go 1.14版本以后，在simple.proto
增加如下
option go_package = "./;errors";
```



## VSCode-proto3插件介绍

使用VSCode的朋友看这里，博主介绍一个VSCode插件，方便对编辑和编译proto文件。

- 扩展程序中搜索 `VSCode-proto3`，然后点击安装。
- 在设置中找到setting.json文件，添加vscode-proto3插件配置

```
   // vscode-proto3插件配置
    "protoc": {
        // protoc.exe所在目录
        "path": "E:\\dev\\go\\bin\\protoc.exe",
        // 保存时自动编译
        "compile_on_save": true,
        "options": [
            // go编译输出指令
            "--go_out=plugins=grpc:."
        ]
    }
  
```

每次编辑完proto文件后，只需要保存，它就会自动帮助完成编译。而且代码有高亮显示，代码自动补全，代码格式化等功能。

 ![](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200409201719878-445176439.gif)



# 简单gRPC

### 前言

gRPC主要有4种请求和响应模式，分别是`简单模式(Simple RPC)`、`服务端流式（Server-side streaming RPC）`、`客户端流式（Client-side streaming RPC）`、和`双向流式（Bidirectional streaming RPC）`。

- `简单模式(Simple RPC)`：客户端发起请求并等待服务端响应。
- `服务端流式（Server-side streaming RPC）`：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。
- `客户端流式（Client-side streaming RPC）`：与服务端数据流模式相反，这次是客户端源源不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。
- `双向流式（Bidirectional streaming RPC）`：双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

本篇文章先介绍简单模式。

### 新建proto文件

主要是定义我们服务的方法以及数据格式，我们使用上一篇的simple.proto文件。

1.定义发送消息的信息

```protobuf
message SimpleRequest{
    // 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
    string data = 1;//发送数据
}
```

2.定义响应信息

```protobuf
message SimpleResponse{
    // 定义接收的参数
    // 参数类型 参数名 标识号(不可重复)
    int32 code = 1;  //状态码
    string value = 2;//接收值
}
```

3.定义服务方法Route

```protobuf
// 定义我们的服务（可定义多个服务,每个服务可定义多个接口）
service Simple{
    rpc Route (SimpleRequest) returns (SimpleResponse){};
}
```

4.编译proto文件

我这里使用上一篇介绍的VSCode-proto3插件，保存后自动编译。

> 指令编译方法，进入simple.proto文件所在目录，运行：
> `protoc --go_out=plugins=grpc:./ ./simple.proto`

### 创建Server端

1.定义我们的服务，并实现Route方法

```go
import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "go-grpc-example/proto"
)
// SimpleService 定义我们的服务
type SimpleService struct{}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}
```

该方法需要传入RPC的上下文context.Context，它的作用结束`超时`或`取消`的请求。更具体的说请参考[该文章](https://blog.csdn.net/chinawangfei/article/details/86559975)

2.启动gRPC服务器

```go
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

里面每个方法的作用都有注释，这里就不解析了。
运行服务端

```powershell
go run server.go
:8000 net.Listing...
```

### 创建Client端

```go
package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "go-grpc-example/2-simple_rpc/proto"
)

// Address 连接地址
const Address string = ":8000"

var grpcClient pb.SimpleClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
	route()
}

// route 调用服务端Route方法
func route() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "grpc",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	res, err := grpcClient.Route(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res)
}

```

运行客户端

```powershell
go run client.go
code:200 value:"hello grpc"
```

成功调用Server端的Route方法并获取返回的数据。

### 总结

本篇介绍了简单RPC模式，客户端发起请求并等待服务端响应。



# 服务端流式gRPC

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#3541337994)

上一篇介绍了`简单模式RPC`，当数据量大或者需要不断传输数据时候，我们应该使用流式RPC，它允许我们边处理边传输数据。本篇先介绍`服务端流式RPC`。

`服务端流式RPC`：客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流，直到里面没有任何消息。

情景模拟：实时获取股票走势。

1.客户端要获取某原油股的实时走势，客户端发送一个请求

2.服务端实时返回该股票的走势

### 新建proto文件[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#1657918117)

新建server_stream.proto文件

1.定义发送信息

```protobuf
// 定义发送请求信息
message SimpleRequest{
    // 定义发送的参数，采用驼峰命名方式，小写加下划线，如：student_name
    // 请求参数
    string data = 1;
}
```

2.定义接收信息

```protobuf
// 定义流式响应信息
message StreamResponse{
    // 流式响应数据
    string stream_value = 1;
}
```

3.定义服务方法ListValue

服务端流式rpc，只要在响应数据前添加stream即可

```protobuf
// 定义我们的服务（可定义多个服务,每个服务可定义多个接口）
service StreamServer{
    // 服务端流式rpc，在响应数据前添加stream
    rpc ListValue(SimpleRequest)returns(stream StreamResponse){};
}
```

4.编译proto文件

进入server_stream.proto所在目录，运行指令:

```
protoc --go_out=plugins=grpc:./ ./server_stream.proto
```

### 创建Server端[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#814763942)

1.定义我们的服务，并实现ListValue方法

```go
// SimpleService 定义我们的服务
type StreamService struct{}
// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 5; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
	}
	return nil
}
```

初学者可能觉得比较迷惑，ListValue的参数和返回值是怎样确定的。其实这些都是编译proto时生成的.pb.go文件中有定义，我们只需要实现就可以了。

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200413200011972-285642192.png)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200413200011972-285642192.png)

2.启动gRPC服务器

```go
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	// 默认单次接收最大消息长度为`1024*1024*4`bytes(4M)，单次发送消息最大长度为`math.MaxInt32`bytes
	// grpcServer := grpc.NewServer(grpc.MaxRecvMsgSize(1024*1024*4), grpc.MaxSendMsgSize(math.MaxInt32))
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamServerServer(grpcServer, &StreamService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

运行服务端

```powershell
go run server.go
:8000 net.Listing...
```

### 创建Client端[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#1574360768)

1.创建调用服务端ListValue方法

```go
// listValue 调用服务端的ListValue方法
func listValue() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc ",
	}
	// 调用我们的服务(ListValue方法)
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
	}
}
```

2.启动gRPC客户端

```go
// Address 连接地址
const Address string = ":8000"

var grpcClient pb.StreamServerClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewStreamServerClient(conn)
	route()
	listValue()
}
```

运行客户端

```powershell
go run client.go
stream server grpc 0
stream server grpc 1
stream server grpc 2
stream server grpc 3
stream server grpc 4
```

客户端不断从服务端获取数据

### 思考[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#2513901056)

假如服务端不停发送数据，类似获取股票走势实时数据，客户端能自己停止获取数据吗？

答案：可以的

1.我们把服务端的ListValue方法稍微修改

```go
// ListValue 实现ListValue方法
func (s *StreamService) ListValue(req *pb.SimpleRequest, srv pb.StreamServer_ListValueServer) error {
	for n := 0; n < 15; n++ {
		// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
		err := srv.Send(&pb.StreamResponse{
			StreamValue: req.Data + strconv.Itoa(n),
		})
		if err != nil {
			return err
		}
		log.Println(n)
		time.Sleep(1 * time.Second)
	}
	return nil
}
```

2.再把客户端调用ListValue方法的实现稍作修改，就可以得到结果了

```go
// listValue 调用服务端的ListValue方法
func listValue() {
	// 创建发送结构体
	req := pb.SimpleRequest{
		Data: "stream server grpc ",
	}
	// 调用我们的服务(Route方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	stream, err := grpcClient.ListValue(context.Background(), &req)
	if err != nil {
		log.Fatalf("Call ListStr err: %v", err)
	}
	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.StreamValue)
		break
	}
	//可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	//调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	stream.CloseSend()
}
```

只需要调用`CloseSend()`方法，就可以关闭服务端的stream，让它停止发送数据。值得注意的是，调用`CloseSend()`后，若继续调用`Recv()`，会重新激活stream，接着当前的结果继续获取消息。

这能完美解决客户端`暂停`->`继续`获取数据的操作。

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12693749.html#301453772)

本篇介绍了`服务端流式RPC`的简单使用，客户端发起一个请求，服务端不停返回数据，直到服务端停止发送数据或客户端主动停止接收数据为止。



# 客户端流式gRPC

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12696733.html#119357241)

上一篇介绍了`服务端流式RPC`，客户端发送请求到服务器，拿到一个流去读取返回的消息序列。 客户端读取返回的流的数据。本篇将介绍`客户端流式RPC`。

`客户端流式RPC`：与`服务端流式RPC`相反，客户端不断的向服务端发送数据流，而在发送结束后，由服务端返回一个响应。

情景模拟：客户端大量数据上传到服务端。

### 新建proto文件[#](https://www.cnblogs.com/FireworksEasyCool/p/12696733.html#1911893160)

新建client_stream.proto文件

1.定义发送信息

```protobuf
// 定义流式请求信息
message StreamRequest{
    //流式请求参数
    string stream_data = 1;
}
```

2.定义接收信息

```protobuf
// 定义响应信息
message SimpleResponse{
    //响应码
    int32 code = 1;
    //响应值
    string value = 2;
}
```

3.定义服务方法RouteList

客户端流式rpc，只要在请求的参数前添加stream即可

```protobuf
service StreamClient{
    // 客户端流式rpc，在请求的参数前添加stream
    rpc RouteList (stream StreamRequest) returns (SimpleResponse){};
}
```

4.编译proto文件

进入client_stream.proto所在目录，运行指令:

```
protoc --go_out=plugins=grpc:./ ./client_stream.proto
```

### 创建Server端[#](https://www.cnblogs.com/FireworksEasyCool/p/12696733.html#1654923856)

1.定义我们的服务，并实现RouteList方法

```go
// SimpleService 定义我们的服务
type SimpleService struct{}
// RouteList 实现RouteList方法
func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
	}
}
```

2.启动gRPC服务器

```go
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamClientServer(grpcServer, &SimpleService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

运行服务端

```powershell
go run server.go
:8000 net.Listing...
```

### 创建Client端[#](https://www.cnblogs.com/FireworksEasyCool/p/12696733.html#1770128403)

1.创建调用服务端RouteList方法

```go
// routeList 调用服务端RouteList方法
func routeList() {
	//调用服务端RouteList方法，获流
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}
```

2.启动gRPC客户端

```go
// Address 连接地址
const Address string = ":8000"

var streamClient pb.StreamClientClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClientClient(conn)
	routeList()
}
```

运行客户端

```powershell
go run client.go
code:200 value:"hello grpc"
value:"ok"
```

服务端不断从客户端获取到数据

```powershell
stream client rpc 0
stream client rpc 1
stream client rpc 2
stream client rpc 3
stream client rpc 4
```

### 思考[#](https://www.cnblogs.com/FireworksEasyCool/p/12696733.html#416227234)

服务端在没有接受完消息时候能主动停止接收数据吗（很少有这种场景）？

答案：可以的，但是客户端代码需要注意EOF判断

1.我们把服务端的RouteList方法实现稍微修改，当接收到一条数据后马上调用SendAndClose()关闭stream.

```go
// RouteList 实现RouteList方法
func (s *SimpleService) RouteList(srv pb.StreamClient_RouteListServer) error {
	for {
		//从流中获取消息
		res, err := srv.Recv()
		if err == io.EOF {
			//发送结果，并关闭
			return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
		}
		if err != nil {
			return err
		}
		log.Println(res.StreamData)
		return srv.SendAndClose(&pb.SimpleResponse{Value: "ok"})
	}
}
```

2.再把客户端调用RouteList方法的实现稍作修改

```go
// routeList 调用服务端RouteList方法
func routeList() {
	//调用服务端RouteList方法，获流
	stream, err := streamClient.RouteList(context.Background())
	if err != nil {
		log.Fatalf("Upload list err: %v", err)
	}
	for n := 0; n < 5; n++ {
		//向流中发送消息
		err := stream.Send(&pb.StreamRequest{StreamData: "stream client rpc " + strconv.Itoa(n)})
		//发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream，此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("RouteList get response err: %v", err)
	}
	log.Println(res)
}
```

客户端Send()需要检测err是否为EOF，因为当服务端在消息没接收完前主动调用SendAndClose()关闭stream，若此时客户端继续执行Send()，则会返回EOF错误。



# 双向流式gRPC

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12698194.html#175745178)

上一篇介绍了`客户端流式RPC`，客户端不断的向服务端发送数据流，在发送结束或流关闭后，由服务端返回一个响应。本篇将介绍`双向流式RPC`。

`双向流式RPC`：客户端和服务端双方使用读写流去发送一个消息序列，两个流独立操作，双方可以同时发送和同时接收。

###### 情景模拟：双方对话（可以一问一答、一问多答、多问一答，形式灵活）。

### 新建proto文件[#](https://www.cnblogs.com/FireworksEasyCool/p/12698194.html#2756698280)

新建both_stream.proto文件

1.定义发送信息

```protobuf
// 定义流式请求信息
message StreamRequest{
    //流请求参数
    string question = 1;
}
```

2.定义接收信息

```protobuf
// 定义流式响应信息
message StreamResponse{
    //流响应数据
    string answer = 1;
}
```

3.定义服务方法Conversations

双向流式rpc，只要在请求的参数前和响应参数前都添加stream即可

```protobuf
service Stream{
    // 双向流式rpc，同时在请求参数前和响应参数前加上stream
    rpc Conversations(stream StreamRequest) returns(stream StreamResponse){};
}
```

4.编译proto文件

进入both_stream.proto所在目录，运行指令:

```
protoc --go_out=plugins=grpc:./ ./both_stream.proto
```

### 创建Server端[#](https://www.cnblogs.com/FireworksEasyCool/p/12698194.html#2539949657)

1.定义我们的服务，并实现RouteList方法

这里简单实现对话中一问一答的形式

```go
// StreamService 定义我们的服务
type StreamService struct{}
// Conversations 实现Conversations方法
func (s *StreamService) Conversations(srv pb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&pb.StreamResponse{
			Answer: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Question,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Question)
	}
}
```

2.启动gRPC服务器

```go
const (
	// Address 监听地址
	Address string = ":8000"
	// Network 网络通信协议
	Network string = "tcp"
)

func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	log.Println(Address + " net.Listing...")
	// 新建gRPC服务器实例
	grpcServer := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterStreamServer(grpcServer, &StreamService{})

	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

运行服务端

```powershell
go run server.go
:8000 net.Listing...
```

### 创建Client端[#](https://www.cnblogs.com/FireworksEasyCool/p/12698194.html#1676515263)

1.创建调用服务端Conversations方法

```go
// conversations 调用服务端的Conversations方法
func conversations() {
	//调用服务端的Conversations方法，获取流
	stream, err := streamClient.Conversations(context.Background())
	if err != nil {
		log.Fatalf("get conversations stream err: %v", err)
	}
	for n := 0; n < 5; n++ {
		err := stream.Send(&pb.StreamRequest{Question: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Answer)
	}
	//最后关闭流
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream err: %v", err)
	}
}
```

2.启动gRPC客户端

```go
// Address 连接地址
const Address string = ":8000"

var streamClient pb.StreamClient

func main() {
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	streamClient = pb.NewStreamClient(conn)
	conversations()
}
```

运行客户端，获取到服务端的应答

```powershell
go run client.go
from stream server answer: the 1 question is stream client rpc 0
from stream server answer: the 2 question is stream client rpc 1
from stream server answer: the 3 question is stream client rpc 2
from stream server answer: the 4 question is stream client rpc 3
from stream server answer: the 5 question is stream client rpc 4
```

服务端获取到来自客户端的提问

```powershell
from stream client question: stream client rpc 0
from stream client question: stream client rpc 1
from stream client question: stream client rpc 2
from stream client question: stream client rpc 3
from stream client question: stream client rpc 4
```



# gRPC超时设置

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12702959.html#3733307752)

gRPC默认的请求的超时时间是很长的，当你没有设置请求超时时间时，所有在运行的请求都占用大量资源且可能运行很长的时间，导致服务资源损耗过高，使得后来的请求响应过慢，甚至会引起整个进程崩溃。

为了避免这种情况，我们的服务应该设置超时时间。前面的[入门教程](https://github.com/Bingjian-Zhu/go-grpc-example)提到，当客户端发起请求时候，需要传入上下文`context.Context`，用于结束`超时`或`取消`的请求。

本篇以[简单RPC](https://bingjian-zhu.github.io/2020/04/10/Go-gRPC教程-简单RPC（二）/)为例，介绍如何设置gRPC请求的超时时间。

### 客户端请求设置超时时间[#](https://www.cnblogs.com/FireworksEasyCool/p/12702959.html#1165223018)

修改调用服务端方法

1.把超时时间设置为当前时间+3秒

```go
	clientDeadline := time.Now().Add(time.Duration(3 * time.Second))
	ctx, cancel := context.WithDeadline(ctx, clientDeadline)
	defer cancel()
```

2.响应错误检测中添加超时检测

```go
       // 传入超时时间为3秒的ctx
	res, err := grpcClient.Route(ctx, &req)
	if err != nil {
		//获取错误状态
		statu, ok := status.FromError(err)
		if ok {
			//判断是否为调用超时
			if statu.Code() == codes.DeadlineExceeded {
				log.Fatalln("Route timeout!")
			}
		}
		log.Fatalf("Call Route err: %v", err)
	}
	// 打印返回值
	log.Println(res.Value)
```

完整的[client.go](https://github.com/Bingjian-Zhu/go-grpc-example/blob/master/6-grpc_deadlines/client/client.go)代码

### 服务端判断请求是否超时[#](https://www.cnblogs.com/FireworksEasyCool/p/12702959.html#4224553538)

当请求超时后，服务端应该停止正在进行的操作，避免资源浪费。

```go
// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
	data := make(chan *pb.SimpleResponse, 1)
	go handle(ctx, req, data)
	select {
	case res := <-data:
		return res, nil
	case <-ctx.Done():
		return nil, status.Errorf(codes.Canceled, "Client cancelled, abandoning.")
	}
}

func handle(ctx context.Context, req *pb.SimpleRequest, data chan<- *pb.SimpleResponse) {
	select {
	case <-ctx.Done():
		log.Println(ctx.Err())
		runtime.Goexit() //超时后退出该Go协程
	case <-time.After(4 * time.Second): // 模拟耗时操作
		res := pb.SimpleResponse{
			Code:  200,
			Value: "hello " + req.Data,
		}
		// //修改数据库前进行超时判断
		// if ctx.Err() == context.Canceled{
		// 	...
		// 	//如果已经超时，则退出
		// }
		data <- &res
	}
}
```

一般地，在写库前进行超时检测，发现超时就停止工作。

完整[server.go](https://github.com/Bingjian-Zhu/go-grpc-example/tree/master/6-grpc_deadlines/server/server.go)代码

### 运行结果[#](https://www.cnblogs.com/FireworksEasyCool/p/12702959.html#3765392802)

服务端：

```powershell
:8000 net.Listing...
goroutine still running
```

客户端：

```powershell
Route timeout!
```

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12702959.html#1241113338)

超时时间的长短需要根据自身服务而定，例如返回一个`hello grpc`，可能只需要几十毫秒，然而处理大量数据的同步操作则可能要很长时间。需要考虑多方面因素来决定这个超时时间，例如系统间端到端的延时，哪些RPC是串行的，哪些是可以并行的等等。





# gRPC-TLS认证+自定义方法认证

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12710325.html#2305183513)

前面篇章的gRPC都是明文传输的，容易被篡改数据。本章将介绍如何为gRPC添加安全机制，包括TLS证书认证和Token认证。

### TLS证书认证[#](https://www.cnblogs.com/FireworksEasyCool/p/12710325.html#4226401740)

###### 什么是TLS

TLS（Transport Layer Security，安全传输层)，TLS是建立在`传输层`TCP协议之上的协议，服务于应用层，它的前身是SSL（Secure Socket Layer，安全套接字层），它实现了将应用层的报文进行加密后再交由TCP进行传输的功能。

###### TLS的作用

TLS协议主要解决如下三个网络安全问题。

- 保密(message privacy)，保密通过加密encryption实现，所有信息都加密传输，第三方无法嗅探；
- 完整性(message integrity)，通过MAC校验机制，一旦被篡改，通信双方会立刻发现；
- 认证(mutual authentication)，双方认证,双方都可以配备证书，防止身份被冒充；

#### 生成私钥

生成RSA私钥：`openssl genrsa -out server.key 2048`

> 生成RSA私钥，命令的最后一个参数，将指定生成密钥的位数，如果没有指定，默认512

生成ECC私钥：`openssl ecparam -genkey -name secp384r1 -out server.key`

> 生成ECC私钥，命令为椭圆曲线密钥参数生成及操作，本文中ECC曲线选择的是secp384r1

#### 生成公钥

```
openssl req -new -x509 -sha256 -key server.key -out server.pem -days 3650
```

> openssl req：生成自签名证书，-new指生成证书请求、-sha256指使用sha256加密、-key指定私钥文件、-x509指输出证书、-days 3650为有效期

此后则输入证书拥有者信息

```
Country Name (2 letter code) [AU]:CN
State or Province Name (full name) [Some-State]:XxXx
Locality Name (eg, city) []:XxXx
Organization Name (eg, company) [Internet Widgits Pty Ltd]:XX Co. Ltd
Organizational Unit Name (eg, section) []:Dev
Common Name (e.g. server FQDN or YOUR name) []:go-grpc-example
Email Address []:xxx@xxx.com
```

#### 服务端构建TLS证书并认证

```go
func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server.pem", "../pkg/tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	// 新建gRPC服务器实例,并开启TLS认证
	grpcServer := grpc.NewServer(grpc.Creds(creds))
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

- `credentials.NewServerTLSFromFile`：从输入证书文件和密钥文件为服务端构造TLS凭证
- `grpc.Creds`：返回一个ServerOption，用于设置服务器连接的凭证。

完整[server.go](https://github.com/Bingjian-Zhu/go-grpc-example/tree/master/7-grpc_security/server/server.go)代码

#### 客户端配置TLS连接

```go
var grpcClient pb.SimpleClient

func main() {
	//从输入的证书文件中为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("../pkg/tls/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
	}
	defer conn.Close()

	// 建立gRPC连接
	grpcClient = pb.NewSimpleClient(conn)
}
```

- `credentials.NewClientTLSFromFile`：从输入的证书文件中为客户端构造TLS凭证。
- `grpc.WithTransportCredentials`：配置连接级别的安全凭证（例如，TLS/SSL），返回一个DialOption，用于连接服务器。

完整[client.go](https://github.com/Bingjian-Zhu/go-grpc-example/tree/master/7-grpc_security/client/client.go)代码

到这里，已经完成TLS证书认证了，gRPC传输不再是明文传输。此外，添加自定义的验证方法能使gRPC相对更安全。下面以Token认证为例，介绍gRPC如何添加自定义验证方法。

### Token认证[#](https://www.cnblogs.com/FireworksEasyCool/p/12710325.html#2565296531)

客户端发请求时，添加Token到上下文`context.Context`中，服务器接收到请求，先从上下文中获取Token验证，验证通过才进行下一步处理。

#### 客户端请求添加Token到上下文中

```go
type PerRPCCredentials interface {
    GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
    RequireTransportSecurity() bool
}
```

gRPC 中默认定义了 `PerRPCCredentials`，是提供用于自定义认证的接口，它的作用是将所需的安全认证信息添加到每个RPC方法的上下文中。其包含 2 个方法：

- `GetRequestMetadata`：获取当前请求认证所需的元数据
- `RequireTransportSecurity`：是否需要基于 TLS 认证进行安全传输

接下来我们实现这两个方法

```go
// Token token认证
type Token struct {
	AppID     string
	AppSecret string
}

// GetRequestMetadata 获取当前请求认证所需的元数据（metadata）
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_id": t.AppID, "app_secret": t.AppSecret}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return true
}
```

然后再客户端中调用Dial时添加自定义验证方法进去

```go
//构建Token
	token := auth.Token{
		AppID:     "grpc_token",
		AppSecret: "123456",
	}
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
```

完整[client.go](https://github.com/Bingjian-Zhu/go-grpc-example/tree/master/7-grpc_security/client/client.go)代码

#### 服务端验证Token

首先需要从上下文中获取元数据，然后从元数据中解析Token进行验证

```go
// Check 验证token
func Check(ctx context.Context) error {
	//从上下文中获取元数据
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "获取Token失败")
	}
	var (
		appID     string
		appSecret string
	)
	if value, ok := md["app_id"]; ok {
		appID = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}
	if appID != "grpc_token" || appSecret != "123456" {
		return status.Errorf(codes.Unauthenticated, "Token无效: app_id=%s, app_secret=%s", appID, appSecret)
	}
	return nil
}

// Route 实现Route方法
func (s *SimpleService) Route(ctx context.Context, req *pb.SimpleRequest) (*pb.SimpleResponse, error) {
    //检测Token是否有效
	if err := Check(ctx); err != nil {
		return nil, err
	}
	res := pb.SimpleResponse{
		Code:  200,
		Value: "hello " + req.Data,
	}
	return &res, nil
}
```

- `metadata.FromIncomingContext`：从上下文中获取元数据

完整[server.go](https://github.com/Bingjian-Zhu/go-grpc-example/tree/master/7-grpc_security/server/server.go)代码

服务端代码中，每个服务的方法都需要添加Check(ctx)来验证Token，这样十分麻烦。gRPC拦截器，能很好地解决这个问题。gRPC拦截器功能类似中间件，拦截器收到请求后，先进行一些操作，然后才进入服务的代码处理。

### 服务端添加拦截器[#](https://www.cnblogs.com/FireworksEasyCool/p/12710325.html#640111819)

```go
func main() {
	// 监听本地端口
	listener, err := net.Listen(Network, Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}
	// 从输入证书文件和密钥文件为服务端构造TLS凭证
	creds, err := credentials.NewServerTLSFromFile("../pkg/tls/server.pem", "../pkg/tls/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}
	//普通方法：一元拦截器（grpc.UnaryInterceptor）
	var interceptor grpc.UnaryServerInterceptor
	interceptor = func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		//拦截普通方法请求，验证Token
		err = Check(ctx)
		if err != nil {
			return
		}
		// 继续处理请求
		return handler(ctx, req)
	}
	// 新建gRPC服务器实例,并开启TLS认证和Token认证
	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptor))
	// 在gRPC服务器注册我们的服务
	pb.RegisterSimpleServer(grpcServer, &SimpleService{})
	log.Println(Address + " net.Listing whth TLS and token...")
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatalf("grpcServer.Serve err: %v", err)
	}
}
```

- `grpc.UnaryServerInterceptor`：为一元拦截器，只会拦截简单RPC方法。流式RPC方法需要使用流式拦截器`grpc.StreamInterceptor`进行拦截。

客户端发起请求，当Token不正确时候，会返回

```powershell
Call Route err: rpc error: code = Unauthenticated desc = Token无效: app_id=grpc_token, app_secret=12345
```

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12710325.html#1907787152)

本篇介绍如何为gRPC添加TLS证书认证和自定义认证，从而让gRPC更安全。添加gRPC拦截器，从而省略在每个方法前添加Token检测代码，使代码更简洁。



# gRPC-middleware使用

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#1159261083)

上篇介绍了gRPC中TLS认证和自定义方法认证，最后还简单介绍了gRPC拦截器的使用。gRPC自身只能设置一个拦截器，所有逻辑都写一起会比较乱。本篇简单介绍[go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)的使用，包括`grpc_zap`、`grpc_auth`和`grpc_recovery`。

### go-grpc-middleware简介[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#619292056)

go-grpc-middleware封装了认证（auth）, 日志（ logging）, 消息（message）, 验证（validation）, 重试（retries） 和监控（retries）等拦截器。

- 安装 `go get github.com/grpc-ecosystem/go-grpc-middleware`
- 使用

```go
import "github.com/grpc-ecosystem/go-grpc-middleware"
myServer := grpc.NewServer(
    grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
        grpc_ctxtags.StreamServerInterceptor(),
        grpc_opentracing.StreamServerInterceptor(),
        grpc_prometheus.StreamServerInterceptor,
        grpc_zap.StreamServerInterceptor(zapLogger),
        grpc_auth.StreamServerInterceptor(myAuthFunction),
        grpc_recovery.StreamServerInterceptor(),
    )),
    grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
        grpc_ctxtags.UnaryServerInterceptor(),
        grpc_opentracing.UnaryServerInterceptor(),
        grpc_prometheus.UnaryServerInterceptor,
        grpc_zap.UnaryServerInterceptor(zapLogger),
        grpc_auth.UnaryServerInterceptor(myAuthFunction),
        grpc_recovery.UnaryServerInterceptor(),
    )),
)
```

`grpc.StreamInterceptor`中添加流式RPC的拦截器。
`grpc.UnaryInterceptor`中添加简单RPC的拦截器。

### grpc_zap日志记录[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#3004283364)

1.创建zap.Logger实例

```go
func ZapInterceptor() *zap.Logger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("failed to initialize zap logger: %v", err)
	}
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}
```

2.把zap拦截器添加到服务端

```go
grpcServer := grpc.NewServer(
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
		)),
	)
```

3.日志分析

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200421150239484-1399156155.png)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200421150239484-1399156155.png)
各个字段代表的意思如下：

```json
{
	  "level": "info",						// string  zap log levels
	  "msg": "finished unary call",					// string  log message

	  "grpc.code": "OK",						// string  grpc status code
	  "grpc.method": "Ping",					/ string  method name
	  "grpc.service": "mwitkow.testproto.TestService",              // string  full name of the called service
	  "grpc.start_time": "2006-01-02T15:04:05Z07:00",               // string  RFC3339 representation of the start time
	  "grpc.request.deadline": "2006-01-02T15:04:05Z07:00",         // string  RFC3339 deadline of the current request if supplied
	  "grpc.request.value": "something",				// string  value on the request
	  "grpc.time_ms": 1.345,					// float32 run time of the call in ms

	  "peer.address": {
	    "IP": "127.0.0.1",						// string  IP address of calling party
	    "Port": 60216,						// int     port call is coming in on
	    "Zone": ""							// string  peer zone for caller
	  },
	  "span.kind": "server",					// string  client | server
	  "system": "grpc",						// string

	  "custom_field": "custom_value",				// string  user defined field
	  "custom_tags.int": 1337,					// int     user defined tag on the ctx
	  "custom_tags.string": "something"				// string  user defined tag on the ctx
}
```

4.把日志写到文件中

上面日志是在控制台输出的，现在我们把日志写到文件中，修改`ZapInterceptor`方法。

```go
import (
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// ZapInterceptor 返回zap.logger实例(把日志写到文件中)
func ZapInterceptor() *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  "log/debug.log",
		MaxSize:   1024, //MB
		LocalTime: true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		zap.NewAtomicLevel(),
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	grpc_zap.ReplaceGrpcLogger(logger)
	return logger
}
```

### grpc_auth认证[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#3334683034)

go-grpc-middleware中的grpc_auth默认使用`authorization`认证方式，以authorization为头部，包括`basic`, `bearer`形式等。下面介绍`bearer token`认证。`bearer`允许使用`access key`（如JSON Web Token (JWT)）进行访问。

1.新建grpc_auth服务端拦截器

```go
// TokenInfo 用户信息
type TokenInfo struct {
	ID    string
	Roles []string
}

// AuthInterceptor 认证拦截器，对以authorization为头部，形式为`bearer token`的Token进行验证
func AuthInterceptor(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}
	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, " %v", err)
	}
	//使用context.WithValue添加了值后，可以用Value(key)方法获取值
	newCtx := context.WithValue(ctx, tokenInfo.ID, tokenInfo)
	//log.Println(newCtx.Value(tokenInfo.ID))
	return newCtx, nil
}

//解析token，并进行验证
func parseToken(token string) (TokenInfo, error) {
	var tokenInfo TokenInfo
	if token == "grpc.auth.token" {
		tokenInfo.ID = "1"
		tokenInfo.Roles = []string{"admin"}
		return tokenInfo, nil
	}
	return tokenInfo, errors.New("Token无效: bearer " + token)
}

//从token中获取用户唯一标识
func userClaimFromToken(tokenInfo TokenInfo) string {
	return tokenInfo.ID
}
```

代码中的对token进行简单验证并返回模拟数据。

2.客户端请求添加`bearer token`

实现和上篇的自定义认证方法大同小异。gRPC 中默认定义了 `PerRPCCredentials`，是提供用于自定义认证的接口，它的作用是将所需的安全认证信息添加到每个RPC方法的上下文中。其包含 2 个方法：

- `GetRequestMetadata`：获取当前请求认证所需的元数据
- `RequireTransportSecurity`：是否需要基于 TLS 认证进行安全传输

接下来我们实现这两个方法

```go
// Token token认证
type Token struct {
	Value string
}

const headerAuthorize string = "authorization"

// GetRequestMetadata 获取当前请求认证所需的元数据
func (t *Token) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{headerAuthorize: t.Value}, nil
}

// RequireTransportSecurity 是否需要基于 TLS 认证进行安全传输
func (t *Token) RequireTransportSecurity() bool {
	return true
}
```

> 注意：这里要以`authorization`为头部，和服务端对应。

发送请求时添加token

```go
//从输入的证书文件中为客户端构造TLS凭证
	creds, err := credentials.NewClientTLSFromFile("../tls/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	//构建Token
	token := auth.Token{
		Value: "bearer grpc.auth.token",
	}
	// 连接服务器
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&token))
```

> 注意：Token中的Value的形式要以`bearer token值`形式。因为我们服务端使用了`bearer token`验证方式。

3.把grpc_auth拦截器添加到服务端

```go
grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
		)),
	)
```

写到这里，服务端都会拦截请求并进行`bearer token`验证，使用`bearer token`是规范了与`HTTP`请求的对接，毕竟gRPC也可以同时支持`HTTP`请求。

### grpc_recovery恢复[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#1413184962)

把gRPC中的`panic`转成`error`，从而恢复程序。

1.直接把grpc_recovery拦截器添加到服务端

最简单使用方式

```go
grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_recovery.StreamServerInterceptor,
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
            grpc_recovery.UnaryServerInterceptor(),
		)),
	)
```

2.自定义错误返回

当`panic`时候，自定义错误码并返回。

```go
// RecoveryInterceptor panic时返回Unknown错误吗
func RecoveryInterceptor() grpc_recovery.Option {
	return grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
		return grpc.Errorf(codes.Unknown, "panic triggered: %v", p)
	})
}
```

添加grpc_recovery拦截器到服务端

```go
grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
            grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)
```

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12750339.html#2627961103)

本篇介绍了`go-grpc-middleware`中的`grpc_zap`、`grpc_auth`和`grpc_recovery`拦截器的使用。`go-grpc-middleware`中其他拦截器可参考[GitHub](https://github.com/grpc-ecosystem/go-grpc-middleware)学习使用。



# gRPC-proto数据验证

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12761033.html#2258195911)

上篇介绍了[go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware)的`grpc_zap`、`grpc_auth`和`grpc_recovery`使用，本篇将介绍`grpc_validator`，它可以对gRPC数据的输入和输出进行验证。

### 创建proto文件，添加验证规则[#](https://www.cnblogs.com/FireworksEasyCool/p/12761033.html#377645860)

这里使用第三方插件[go-proto-validators](https://github.com/mwitkow/go-proto-validators)自动生成验证规则。

```
go get github.com/mwitkow/go-proto-validators
```

1.新建simple.proto文件

```protobuf
syntax = "proto3";

package proto;

import "github.com/mwitkow/go-proto-validators/validator.proto";

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
}

message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z]{2,5}$"}];
  // proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}

service Simple{
  rpc Route (InnerMessage) returns (OuterMessage){};
}
```

代码`import "github.com/mwitkow/go-proto-validators/validator.proto"`，文件`validator.proto`需要`import "google/protobuf/descriptor.proto";`包，不然会报错。

`google/protobuf`地址：https://github.com/protocolbuffers/protobuf/tree/master/src/google/protobuf/descriptor.proto

把`src`文件夹中的`protobuf`目录下载到GOPATH目录下。

2.编译simple.proto文件

```
go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators
```

指令编译：`protoc --govalidators_out=. --go_out=plugins=grpc:./ ./simple.proto`

> 或者使用`VSCode-proto3`插件，[第一篇](https://www.cnblogs.com/FireworksEasyCool/p/12669371.html)有介绍。只需要添加`"--govalidators_out=."`即可。

```
    // vscode-proto3插件配置
    "protoc": {
        // protoc.exe所在目录
        "path": "C:\\Go\\bin\\protoc.exe",
        // 保存时自动编译
        "compile_on_save": true,
        "options": [
            // go编译输出指令
            "--go_out=plugins=grpc:.",
            "--govalidators_out=."
        ]
    },
```

编译完成后，自动生成`simple.pb.go`和`simple.validator.pb.go`文件，`simple.pb.go`文件不再介绍，我们看下`simple.validator.pb.go`文件。

```go
// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: go-grpc-example/9-grpc_proto_validators/proto/simple.proto

package proto

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	regexp "regexp"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if !(this.SomeInteger > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be greater than '0'`, this.SomeInteger))
	}
	if !(this.SomeInteger < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be less than '100'`, this.SomeInteger))
	}
	if !(this.SomeFloat >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeFloat", fmt.Errorf(`value '%v' must be greater than or equal to '0'`, this.SomeFloat))
	}
	if !(this.SomeFloat <= 1) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeFloat", fmt.Errorf(`value '%v' must be lower than or equal to '1'`, this.SomeFloat))
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile(`^[a-z]{2,5}$`)

func (this *OuterMessage) Validate() error {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		return github_com_mwitkow_go_proto_validators.FieldError("ImportantString", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,5}$"`, this.ImportantString))
	}
	if nil == this.Inner {
		return github_com_mwitkow_go_proto_validators.FieldError("Inner", fmt.Errorf("message must exist"))
	}
	if this.Inner != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Inner); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Inner", err)
		}
	}
	return nil
}
```

里面自动生成了`message`中属性的验证规则。

### 把`grpc_validator`验证拦截器添加到服务端[#](https://www.cnblogs.com/FireworksEasyCool/p/12761033.html#2179867576)

```go
grpcServer := grpc.NewServer(cred.TLSInterceptor(),
	grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_validator.StreamServerInterceptor(),
	        grpc_auth.StreamServerInterceptor(auth.AuthInterceptor),
			grpc_zap.StreamServerInterceptor(zap.ZapInterceptor()),
			grpc_recovery.StreamServerInterceptor(recovery.RecoveryInterceptor()),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		    grpc_validator.UnaryServerInterceptor(),
		    grpc_auth.UnaryServerInterceptor(auth.AuthInterceptor),
			grpc_zap.UnaryServerInterceptor(zap.ZapInterceptor()),
            grpc_recovery.UnaryServerInterceptor(recovery.RecoveryInterceptor()),
		)),
	)
```

运行后，当输入数据验证失败后，会有以下错误返回

```powershell
Call Route err: rpc error: code = InvalidArgument desc = invalid field SomeInteger: value '101' must be less than '100'
```

### 其他类型验证规则设置[#](https://www.cnblogs.com/FireworksEasyCool/p/12761033.html#3828274910)

`enum`验证

```protobuf
syntax = "proto3";
package proto;
import "github.com/mwitkow/go-proto-validators/validator.proto";

message SomeMsg {
  Action do = 1 [(validator.field) = {is_in_enum : true}];
}

enum Action {
  ALLOW = 0;
  DENY = 1;
  CHILL = 2;
}
```

`UUID`验证

```protobuf
syntax = "proto3";
package proto;
import "github.com/mwitkow/go-proto-validators/validator.proto";

message UUIDMsg {
  // user_id must be a valid version 4 UUID.
  string user_id = 1 [(validator.field) = {uuid_ver: 4, string_not_empty: true}];
}
```

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12761033.html#3698822124)

`go-grpc-middleware`中`grpc_validator`集成`go-proto-validators`，我们只需要在编写proto时设好验证规则，并把`grpc_validator`添加到gRPC服务端，就能完成gRPC的数据验证，很简单也很方便。



# gRPC-gRPC转换HTTP

### 前言[#](https://www.cnblogs.com/FireworksEasyCool/p/12782137.html#2271984965)

我们通常把`RPC`用作内部通信，而使用`Restful Api`进行外部通信。为了避免写两套应用，我们使用[grpc-gateway](https://github.com/grpc-ecosystem/grpc-gateway)把`gRPC`转成`HTTP`。服务接收到`HTTP`请求后，`grpc-gateway`把它转成`gRPC`进行处理，然后以`JSON`形式返回数据。本篇代码以上篇为基础，最终转成的`Restful Api`支持`bearer token`验证、数据验证，并添加`swagger`文档。

### gRPC转成HTTP[#](https://www.cnblogs.com/FireworksEasyCool/p/12782137.html#1971576044)

#### 编写和编译proto

1.编写simple.proto

```protobuf
syntax = "proto3";

package proto;

import "github.com/mwitkow/go-proto-validators/validator.proto";
import "go-grpc-example/10-grpc-gateway/proto/google/api/annotations.proto";

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
}

message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z]{2,5}$"}];
  // proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}

service Simple{
  rpc Route (InnerMessage) returns (OuterMessage){
      option (google.api.http) ={
          post:"/v1/example/route"
          body:"*"
      };
  }
}
```

可以看到，`proto`变化不大，只是添加了API的路由路径

```protobuf
      option (google.api.http) ={
          post:"/v1/example/route"
          body:"*"
      };
```

2.编译`simple.proto`

`simple.proto`文件引用了`google/api/annotations.proto`（[来源](https://github.com/grpc-ecosystem/grpc-gateway/tree/master/third_party/googleapis/google/api)），先要把它编译了。我这里是把`google/`文件夹直接复制到项目中的`proto/`目录中进行编译。发现`annotations.proto`引用了`google/api/http.proto`，那把它也编译了。

进入`annotations.proto`所在目录，编译：

```powershell
protoc --go_out=plugins=grpc:./ ./http.proto
protoc --go_out=plugins=grpc:./ ./annotations.proto
```

进入`simple.proto`所在目录，编译：

```powershell
#生成simple.validator.pb.go和simple.pb.go
protoc --govalidators_out=. --go_out=plugins=grpc:./ ./simple.proto
#生成simple.pb.gw.go
protoc --grpc-gateway_out=logtostderr=true:./ ./simple.proto
```

以上完成`proto`编译，接着修改服务端代码。

#### 服务端代码修改

1.`server/`文件夹下新建`gateway/`目录，然后在里面新建`gateway.go`文件

```go
package gateway

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	pb "go-grpc-example/10-grpc-gateway/proto"
	"go-grpc-example/10-grpc-gateway/server/swagger"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

// ProvideHTTP 把gRPC服务转成HTTP服务，让gRPC同时支持HTTP
func ProvideHTTP(endpoint string, grpcServer *grpc.Server) *http.Server {
	ctx := context.Background()
	//获取证书
	creds, err := credentials.NewClientTLSFromFile("../tls/server.pem", "go-grpc-example")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	//添加证书
	dopts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	//新建gwmux，它是grpc-gateway的请求复用器。它将http请求与模式匹配，并调用相应的处理程序。
	gwmux := runtime.NewServeMux()
	//将服务的http处理程序注册到gwmux。处理程序通过endpoint转发请求到grpc端点
	err = pb.RegisterSimpleHandlerFromEndpoint(ctx, gwmux, endpoint, dopts)
	if err != nil {
		log.Fatalf("Register Endpoint err: %v", err)
	}
	//新建mux，它是http的请求复用器
	mux := http.NewServeMux()
	//注册gwmux
	mux.Handle("/", gwmux)
	log.Println(endpoint + " HTTP.Listing whth TLS and token...")
	return &http.Server{
		Addr:      endpoint,
		Handler:   grpcHandlerFunc(grpcServer, mux),
		TLSConfig: getTLSConfig(),
	}
}

// grpcHandlerFunc 根据不同的请求重定向到指定的Handler处理
func grpcHandlerFunc(grpcServer *grpc.Server, otherHandler http.Handler) http.Handler {
	return h2c.NewHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.ProtoMajor == 2 && strings.Contains(r.Header.Get("Content-Type"), "application/grpc") {
			grpcServer.ServeHTTP(w, r)
		} else {
			otherHandler.ServeHTTP(w, r)
		}
	}), &http2.Server{})
}

// getTLSConfig获取TLS配置
func getTLSConfig() *tls.Config {
	cert, _ := ioutil.ReadFile("../tls/server.pem")
	key, _ := ioutil.ReadFile("../tls/server.key")
	var demoKeyPair *tls.Certificate
	pair, err := tls.X509KeyPair(cert, key)
	if err != nil {
		grpclog.Fatalf("TLS KeyPair err: %v\n", err)
	}
	demoKeyPair = &pair
	return &tls.Config{
		Certificates: []tls.Certificate{*demoKeyPair},
		NextProtos:   []string{http2.NextProtoTLS}, // HTTP2 TLS支持
	}
}
```

它主要作用是把不用的请求重定向到指定的服务处理，从而实现把`HTTP`请求转到`gRPC`服务。

2.gRPC支持HTTP

```go
    //使用gateway把grpcServer转成httpServer
	httpServer := gateway.ProvideHTTP(Address, grpcServer)
	if err = httpServer.Serve(tls.NewListener(listener, httpServer.TLSConfig)); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
```

#### 使用postman测试

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426190917132-305705093.gif)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426190917132-305705093.gif)

在动图中可以看到，我们的`gRPC`服务已经同时支持`RPC`和`HTTP`请求了，而且API接口支持`bearer token`验证和数据验证。为了方便对接，我们把API接口生成`swagger`文档。

### 生成swagger文档[#](https://www.cnblogs.com/FireworksEasyCool/p/12782137.html#3458078179)

#### 生成swagger文档-simple.swagger.json

1.安装`protoc-gen-swagger`

```
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
```

2.编译生成simple.swagger.json

到simple.proto文件目录下，编译：
`protoc --swagger_out=logtostderr=true:./ ./simple.proto`

> 再次提一下，本人在VSCode中使用`VSCode-proto3`插件，[第一篇](https://www.cnblogs.com/FireworksEasyCool/p/12669371.html)有介绍，只要保存，就会自动编译，很方便，无需记忆指令。完整配置如下：

```
    // vscode-proto3插件配置
    "protoc": {
        // protoc.exe所在目录
        "path": "C:\\Go\\bin\\protoc.exe",
        // 保存时自动编译
        "compile_on_save": true,
        "options": [
            "--go_out=plugins=grpc:.",//在当前目录编译输出.pb.go文件
            "--govalidators_out=.",//在当前目录编译输出.validator.pb文件
            "--grpc-gateway_out=logtostderr=true:.",//在当前目录编译输出.pb.gw.go文件
            "--swagger_out=logtostderr=true:."//在当前目录编译输出.swagger.json文件
        ]
    }
```

编译生成后把需要的文件留下，不需要的删掉。

#### 把swagger-ui转成Go代码，备用

1.下载`swagger-ui`

[下载地址](https://github.com/swagger-api/swagger-ui)，把`dist`目录下的所有文件拷贝我们项目的`server/swagger/swagger-ui/`目录下。

2.把`Swagger UI`转换为Go代码

安装`go-bindata`：
`go get -u github.com/jteeuwen/go-bindata/...`

回到`server/`所在目录，运行指令把`Swagger UI`转成Go代码。
`go-bindata --nocompress -pkg swagger -o swagger/datafile.go swagger/swagger-ui/...`

- 这步有坑，必须要回到`main`函数所在的目录运行指令，因为生成的Go代码中的`_bindata` 映射了`swagger-ui`的路径，程序是根据这些路径来找页面的。如果没有在`main`函数所在的目录运行指令，则生成的路径不对，会报404，无法找到页面。本项目`server/`端的`main`函数在`server.go`中，所以在`server/`所在目录下运行指令。

```go
var _bindata = map[string]func() (*asset, error){
	"swagger/swagger-ui/favicon-16x16.png": swaggerSwaggerUiFavicon16x16Png,
	"swagger/swagger-ui/favicon-32x32.png": swaggerSwaggerUiFavicon32x32Png,
	"swagger/swagger-ui/index.html": swaggerSwaggerUiIndexHtml,
	"swagger/swagger-ui/oauth2-redirect.html": swaggerSwaggerUiOauth2RedirectHtml,
	"swagger/swagger-ui/swagger-ui-bundle.js": swaggerSwaggerUiSwaggerUiBundleJs,
	"swagger/swagger-ui/swagger-ui-bundle.js.map": swaggerSwaggerUiSwaggerUiBundleJsMap,
	"swagger/swagger-ui/swagger-ui-standalone-preset.js": swaggerSwaggerUiSwaggerUiStandalonePresetJs,
	"swagger/swagger-ui/swagger-ui-standalone-preset.js.map": swaggerSwaggerUiSwaggerUiStandalonePresetJsMap,
	"swagger/swagger-ui/swagger-ui.css": swaggerSwaggerUiSwaggerUiCss,
	"swagger/swagger-ui/swagger-ui.css.map": swaggerSwaggerUiSwaggerUiCssMap,
	"swagger/swagger-ui/swagger-ui.js": swaggerSwaggerUiSwaggerUiJs,
	"swagger/swagger-ui/swagger-ui.js.map": swaggerSwaggerUiSwaggerUiJsMap,
}
```

#### 对外提供swagger-ui

1.在`swagger/`目录下新建`swagger.go`文件

```go
package swagger

import (
	"log"
	"net/http"
	"path"
	"strings"

	assetfs "github.com/elazarl/go-bindata-assetfs"
)

//ServeSwaggerFile 把proto文件夹中的swagger.json文件暴露出去
func ServeSwaggerFile(w http.ResponseWriter, r *http.Request) {
	if !strings.HasSuffix(r.URL.Path, "swagger.json") {
		log.Printf("Not Found: %s", r.URL.Path)
		http.NotFound(w, r)
		return
	}

	p := strings.TrimPrefix(r.URL.Path, "/swagger/")
	// "../proto/"为.swagger.json所在目录
	p = path.Join("../proto/", p)

	log.Printf("Serving swagger-file: %s", p)

	http.ServeFile(w, r, p)
}

//ServeSwaggerUI 对外提供swagger-ui
func ServeSwaggerUI(mux *http.ServeMux) {
	fileServer := http.FileServer(&assetfs.AssetFS{
		Asset:    Asset,
		AssetDir: AssetDir,
		Prefix:   "swagger/swagger-ui", //swagger-ui文件夹所在目录
	})
	prefix := "/swagger-ui/"
	mux.Handle(prefix, http.StripPrefix(prefix, fileServer))
}
```

2.注册`swagger`

在`gateway.go`中添加如下代码

```go
	//注册swagger
	mux.HandleFunc("/swagger/", swagger.ServeSwaggerFile)
	swagger.ServeSwaggerUI(mux)
```

到这里我们已经完成了`swagger`文档的添加工作了，由于谷歌浏览器不能使用自己制作的TLS证书，所以我们用火狐浏览器进行测试。

用火狐浏览器打开：https://127.0.0.1:8000/swagger-ui/

在最上面地址栏输入：https://127.0.0.1:8000/swagger/simple.swagger.json

然后就可以看到swagger生成的API文档了。

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426200056107-1097342377.png)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426200056107-1097342377.png)

还有个问题，我们使用了bearer token进行接口验证的，怎么把`bearer token`也添加到swagger中呢？
最后我在`grpc-gateway`GitHub上的这个[Issues](https://github.com/grpc-ecosystem/grpc-gateway/issues/1089)找到解决办法。

#### 在swagger中配置`bearer token`

1.修改`simple.proto`文件

```protobuf
syntax = "proto3";

package proto;

import "github.com/mwitkow/go-proto-validators/validator.proto";
import "go-grpc-example/10-grpc-gateway/proto/google/api/annotations.proto";
import "go-grpc-example/10-grpc-gateway/proto/google/options/annotations.proto";

message InnerMessage {
  // some_integer can only be in range (1, 100).
  int32 some_integer = 1 [(validator.field) = {int_gt: 0, int_lt: 100}];
  // some_float can only be in range (0;1).
  double some_float = 2 [(validator.field) = {float_gte: 0, float_lte: 1}];
}

message OuterMessage {
  // important_string must be a lowercase alpha-numeric of 5 to 30 characters (RE2 syntax).
  string important_string = 1 [(validator.field) = {regex: "^[a-z]{2,5}$"}];
  // proto3 doesn't have `required`, the `msg_exist` enforces presence of InnerMessage.
  InnerMessage inner = 2 [(validator.field) = {msg_exists : true}];
}

option (grpc.gateway.protoc_gen_swagger.options.openapiv2_swagger) = {
  security_definitions: {
    security: {
      key: "bearer"
      value: {
        type: TYPE_API_KEY
        in: IN_HEADER
        name: "Authorization"
        description: "Authentication token, prefixed by Bearer: Bearer <token>"
      }
    }
  }

  security: {
    security_requirement: {
      key: "bearer"
    }
  }

  info: {
		title: "grpc gateway sample";
		version: "1.0";	
		license: {
			name: "MIT";			
		};
  }

  schemes: HTTPS
};

service Simple{
  rpc Route (InnerMessage) returns (OuterMessage){
      option (google.api.http) ={
          post:"/v1/example/route"
          body:"*"
      };
      // //禁用bearer token
      // option (grpc.gateway.protoc_gen_swagger.options.openapiv2_operation) = {
      //   security: { } // Disable security key
      // };
  }
}
```

2.重新编译生成simple.swagger.json

大功告成！

#### 验证测试

1.添加`bearer token`

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201427202-713287948.gif)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201427202-713287948.gif)

2.调用接口，正确返回数据

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201751447-2114446576.png)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201751447-2114446576.png)

3.传递不合规则的数据，返回违反数据验证逻辑错误

[![img](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201907769-1849889200.png)](https://img2020.cnblogs.com/blog/1508611/202004/1508611-20200426201907769-1849889200.png)

### 总结[#](https://www.cnblogs.com/FireworksEasyCool/p/12782137.html#3277538882)

本篇介绍了如何使用`grpc-gateway`让`gRPC`同时支持HTTP，最终转成的`Restful Api`支持`bearer token`验证、数据验证。同时生成`swagger`文档，方便API接口对接。