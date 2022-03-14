# goframe

## gf-cli

```
wget -O gf https://github.com/gogf/gf-cli/releases/download/v1.16.3/gf_linux_amd64 && chmod +x gf && ./gf install
```

```shell
#下载
wget https://goframe.org/cli/linux_amd64/gf 

#安装gf工具到/usr/local/go/bin
./gf install 

#查看版本
gf -v 


# win
https://goframe.org/cli/windows_amd64/gf.exe
```



```shell
# 初始化一个项目
gf init my-app
cd my-app
gf run main.go

# 生成dao
gf gen dao

# 生产model，测试无用
gf gen model -path ./app/model -c config/config.yaml -g sys -t sys_user


# 生产接口文件
gf swagger --pack

# docker
gf docker main.go -p -t jettjia/gf-demos:test  
```



## 简单案例

### 目录结构说明

```shell
├── app
│   ├── amqp
│   ├── api
│   ├── dao
│   ├── define
│   ├── model
│   ├── service
│   ├── shared
│   └── task
├── bin
│   └── linux_amd64
├── boot
│   └── boot.go
├── config
│   └── config.toml
├── docker
├── Dockerfile
├── document
├── go.mod
├── go.sum
├── i18n
├── library
│   ├── middleware
│   └── response
├── main.go
├── packed
│   ├── packed.go
│   └── swagger.go
├── public
│   ├── html
│   ├── plugin
│   └── resource
├── README.MD
├── router
│   └── router.go
├── swagger
│   └── swagger.json
└── template

```



### DB-orm

https://goframe.org/pages/viewpage.action?pageId=1114686



### 定时任务

https://goframe.org/pages/viewpage.action?pageId=1114363#gtimer(任务定时器)-单例任务

```go
package task

import (
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func init()  {
	cronT()
}

func cronT() {
	gcron.AddSingleton("* * * * * *", func() {
		glog.Println("doing")
	})
}
```



### redis

https://goframe.org/display/gf/NoSQL+Redis

```go
package api

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"my-app/library/response"
)

// 注册控制器
var TestRedis = testRedisApi{}

type testRedisApi struct{}

// redis
func (*testRedisApi) TestRedis(r *ghttp.Request) {
	g.Redis().Do("SET", "k", "v")
	v, _ := g.Redis().DoVar("GET", "k")
	response.JsonExit(r, 0, "ok", v.String())
}

```

```
127.0.0.1:8199/test/redis/test-redis
```



### amqp-rabbitmq

https://github.com/streadway/amqp



```go
package api

import (
	"github.com/gogf/gf/net/ghttp"
	"my-app/app/amqp/producer"
	"my-app/library/response"
)

// 注册控制器
var TestMq = testMqApi{}

type testMqApi struct{}


// rabbitmq send
func (*testMqApi) TestMqSend(r *ghttp.Request) {
	poolName := "default"

	msg := r.GetString("msg")

	producer.DemoTest(poolName, msg)

	response.JsonExit(r, 0, "ok", msg)
}

```

```
127.0.0.1:8199/test/mq/test-mq-send
{
    "msg" : "666"
}
```



### ES

github.com/olivere/elastic

https://www.cnblogs.com/Yemilice/p/12887812.html

```go
package api

import (
	"context"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/olivere/elastic/v7"
	"my-app/library/middleware/elasticsearch"
	"my-app/library/response"
	"time"
)

// 注册控制器
var TestEs = testEsApi{}

type testEsApi struct{}

// es create index
func (*testEsApi) CreateIndex(r *ghttp.Request) {
	esUrl := elasticsearch.MidEs.EsConnStr("default")
	esUser := elasticsearch.MidEs.EsUser("default")
	esPassword := elasticsearch.MidEs.EsPassword("default")

	fmt.Println(esUrl)
	fmt.Println(esUser)
	fmt.Println(esPassword)
	fmt.Println("=================")
	client, err := elastic.NewClient(
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(esUser, esPassword),
		elastic.SetSniff(false),
	)
	fmt.Println(err)
	if err != nil {
		response.JsonExit(r, 1, "err", err)
	}

	indexname := "test-index" // indexname 你可以想成表名
	client.CreateIndex(indexname).Do(context.Background())

	response.JsonExit(r, 0, "ok")
}
```

```
127.0.0.1:8199/test/es/create-index
```



### 链路追踪 

https://goframe.org/pages/viewpage.action?pageId=3673727



### 第三方请求封装 

https://goframe.org/display/gf/HTTPClient



### 错误处理

https://goframe.org/pages/viewpage.action?pageId=1114255



### 异常处理

https://goframe.org/pages/viewpage.action?pageId=1114341





### csv导出 

```go
package api

import (
	"encoding/csv"
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"io"
	"log"
	"my-app/library/response"
	"os"
)

// 注册控制器
var TestCsv = testCsvApi{}

type testCsvApi struct{}

type UserTest struct {
	ID   string
	Name string
	Tel  string
	Addr string
}

// csv down
func (*testCsvApi) DownCsv(r *ghttp.Request) {
	Users1 := []UserTest{
		{"1", "贾直接", "adminTel", "武汉硚口1号"},
		{"2", "test", "testTel", "testAddr"},
	}
	StructToCsv("tmp/user.csv", Users1)

	response.JsonExit(r, 0, "ok")
}

// 追加写入
func StructToCsv(filename string, UsersDb []UserTest) {
	newFile, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		newFile.Close()
	}()
	// 写入UTF-8
	newFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	// 写数据到csv文件
	w := csv.NewWriter(newFile)
	header := []string{"ID", "Name", "Tel", "Addr"} //标题
	w.Write(header)
	for _, v := range UsersDb {
		context := []string{
			v.ID,
			v.Name,
			v.Tel,
			v.Addr,
		}
		// data = append(data, context)
		w.Write(context)
	}
	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
	w.Flush()
}

```



```
127.0.0.1:8199/test/csv/down-csv
```



### csv导入

```
127.0.0.1:8199/test/csv/up-csv
```



```go
var Users []*UserTest

// 解析上传csv
func CsvToUp(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	reader := csv.NewReader(f)
	result1 := make([][]string, 0)
	result2 := make([][]string, 0)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		result1 = append(result1, record)
	}
	for k, _ := range result1 {
		if result1[k][1] == "Name" && result1[k][2] == "Tel" && result1[k][3] == "Addr" {
			result2 = append(result1[:k], result1[k+1:]...)
		}
	}
	for _, v := range result2 {
		user := &UserTest{
			ID:   v[0],
			Name: v[1],
			Addr: v[2],
			Tel:  v[3],
		}
		Users = append(Users, user)
	}
	for _, v := range Users {
		fmt.Println(v)
	}
}

```



```
127.0.0.1:8199/test/csv/up-csv
参数：file 文件类型
```





### jwt

https://goframe.org/pages/viewpage.action?pageId=6357048

```
go get github.com/gogf/gf-jwt
```



```shell
# 登录
127.0.0.1:8199/login
{
    "username" : "admin",
    "password" : "admin"
}

# 刷新/续签
curl -s -XPOST 127.0.0.1:8199/user/refresh_token -d 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MTI2ODgsImlhdCI6MTYyODQxMjM4OCwiaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.1KW6vVvKRMLzlt7_2WkxcxbFVDzkDhJXBsrtfFbtwqE'

# 退出/注销
curl -s -XPOST 127.0.0.1:8199/user/logout -d 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTIwMTMyNDgsImlhdCI6MTYxMjAxMjk0OCwiaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.UYKBsnXipeRBR4GENFVnpE09wazagHlCLMnQ7o7EJE4'

# token过期
curl -s -XPOST 127.0.0.1:8199/user/info -d 'token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTIwMTMyNDgsImlhdCI6MTYxMjAxMjk0OCwiaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.UYKBsnXipeRBR4GENFVnpE09wazagHlCLMnQ7o7EJE4'
```



请求接口需要增加header信息

```
Authorization

Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjg0MTY2NjgsImlhdCI6MTYyODQxMzA2OCwiaWQiOjEsInVzZXJuYW1lIjoiYWRtaW4ifQ.TAYRhZ7vU9Q8S-A4z2aniTaaRts6oKp5O40gqnsQJYo
```



### 跨域

router下

```go
// MiddlewareCORS 跨域
func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}
```



### 中间件 

https://goframe.org/pages/viewpage.action?pageId=1114315

```go
package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"my-app/library/response"
)

var MiddlewareParam = new(middlewareTest)

type middlewareTest struct {}

// 这里是模拟一个中间件
// 这里获取 参数里没有 abc = 123
// 就提示不通过
func (m *middlewareTest) Check(r *ghttp.Request) {
	paramStr := r.GetInt("abc")
	if paramStr != 123 {
		response.JsonExit(r, 1, "err", "请求的参数里必须包含有 abc=123")
	}

	r.Middleware.Next()
}
```

route

```go
// 额外的中间件判断
func MiddlewareABCTest(r *ghttp.Request) {
	middleware.MiddlewareParam.Check(r)
	r.Middleware.Next()
}
```



### 全局错误统一处理

```go
package middleware

import (
	"github.com/gogf/gf/net/ghttp"
	"net/http"
)

var MiddlewareErr = new(middlewareErr)

type middlewareErr struct {}


// 这里是对全局返回的错误进行统一处理
// 比如服务器500了，返回给前端的信息就是统一的一句话
func (m *middlewareErr) ErrHandle(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.Status >= http.StatusInternalServerError {
		r.Response.ClearBuffer()
		r.Response.Writeln("哎哟我去，服务器居然开小差了，程序员giegie要罚款了！")
	}
}
```



### oss上传 todo

### 腾讯上传 todo



### websocket

https://goframe.org/pages/viewpage.action?pageId=1114298



### 分布式事务-seata-golang

https://github.com/opentrx/seata-go-samples



### 分布式调度 

https://github.com/ouqiang/gocron



## 案例运行

```
# 注册
127.0.0.1:8199/user/registry
{
    "passport" : "222",
    "password" : "23",
    "nickname" : "jett"
}

# 获取
127.0.0.1:8199/user/profile
{
    "userId" : 1
}

# 修改
127.0.0.1:8199/user/update
{
    "id" : 1,
    "passport" : "222",
    "nickname" : "jett2"
}
```

