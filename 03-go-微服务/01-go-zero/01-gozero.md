# 介绍

文档：https://zeromicro.github.io/go-zero

http://zero.gocn.vip/



# 环境准备

```shell
# install go

# install docker
`curl -sSL https://get.daocloud.io/docker | sh`

# install mysql

# install redis

# install etcd
REGISTRY=quay.io/coreos/etcd
rm -rf /tmp/etcd-data.tmp && mkdir -p /tmp/etcd-data.tmp && \
  docker run \
  -p 2379:2379 \
  -p 2380:2380 \
  --mount type=bind,source=/tmp/etcd-data.tmp,destination=/etcd-data \
  --name etcd \
  quay.io/coreos/etcd:v3.4.14 \
  /usr/local/bin/etcd \
  --name s1 \
  --data-dir /etcd-data \
  --listen-client-urls http://0.0.0.0:2379 \
  --advertise-client-urls http://0.0.0.0:2379 \
  --listen-peer-urls http://0.0.0.0:2380 \
  --initial-advertise-peer-urls http://0.0.0.0:2380 \
  --initial-cluster s1=http://0.0.0.0:2380 \
  --initial-cluster-token tkn \
  --initial-cluster-state new \
  --log-level info \
  --logger zap \
  --log-outputs stderr

docker exec etcd /bin/sh -c "/usr/local/bin/etcdctl version"
docker exec etcd /bin/sh -c "/usr/local/bin/etcdctl endpoint health"
docker exec etcd /bin/sh -c "/usr/local/bin/etcdctl put foo bar"
docker exec etcd /bin/sh -c "/usr/local/bin/etcdctl get foo"

# install protoc
https://github.com/protocolbuffers/protobuf/releases
or

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.14.protoc-3.14.0-linux-x86_64.zip
unzip protoc-3.14.0-linux-x86_64.zip
mv bin/protoc /usr/local/bin/

# install protoc-gen-go
go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2

# install goctl
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl
```



# 快速上手

## Installation

在项目目录下通过如下命令安装：

```shell
# 设置GOPROXY
`go env -w GOPROXY=https://goproxy.cn,direct`
`go env`
```

```shell
`GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero`
```

安装goctl工具

```shell
`GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get -u github.com/tal-tech/go-zero/tools/goctl`
```

快速生成api服务

```shell
`goctl api new greet`
`cd greet`
`go run greet.go -f etc/greet-api.yaml`
```

默认侦听在8888端口（可以在配置文件里修改），可以通过curl请求：

```shell
`curl -i http://localhost:8888/from/you`
   HTTP/1.1 200 OK
   Date: Sun, 30 Aug 2020 15:32:35 GMT
   Content-Length: 0
```

编写业务代码：

- api文件定义了服务对外暴露的路由，可参考[api规范(opens new window)](https://github.com/tal-tech/zero-doc/blob/main/doc/goctl.md)
- 可以在servicecontext.go里面传递依赖给logic，比如mysql, redis等
- 在api定义的get/post/put/delete等请求对应的logic里增加业务处理逻辑



## api目录介绍

```shell
.
├── etc
│   └── greet-api.yaml              // 配置文件
├── go.mod                          // mod文件
├── greet.api                       // api描述文件
├── greet.go                        // main函数入口
└── internal                        
    ├── config  
    │   └── config.go               // 配置声明type
    ├── handler                     // 路由及handler转发
    │   ├── greethandler.go
    │   └── routes.go
    ├── logic                       // 业务逻辑
    │   └── greetlogic.go
    ├── middleware                  // 中间件文件
    │   └── greetmiddleware.go
    ├── svc                         // logic所依赖的资源池
    │   └── servicecontext.go
    └── types                       // request、response的struct，根据api自动生成，不建议编辑
        └── types.go
```



## rpc服务目录

```
.
├── etc             // yaml配置文件
│   └── greet.yaml
├── go.mod
├── greet           // pb.go文件夹①
│   └── greet.pb.go
├── greet.go        // main函数
├── greet.proto     // proto 文件
├── greetclient     // call logic ②
│   └── greet.go
└── internal        
    ├── config      // yaml配置对应的实体
    │   └── config.go
    ├── logic       // 业务代码
    │   └── pinglogic.go
    ├── server      // rpc server
    │   └── greetserver.go
    └── svc         // 依赖资源
        └── servicecontext.go
```



## 技巧

### golang工具

https://zeromicro.github.io/go-zero/intellij.html



## 业务开发-演示工程

文档：https://zeromicro.github.io/go-zero/business-dev.html

### 下载演示工程

https://zeromicro.github.io/go-zero/resource/book.zip

### 目录拆分

```
mall // 工程名称
├── common // 通用库
│   ├── randx
│   └── stringx
├── go.mod
├── go.sum
└── service // 服务存放目录
    ├── afterSale
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── cart
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── order
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── pay
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    ├── product
    │   ├── cmd
    │   │   ├── api
    │   │   └── rpc
    │   └── model
    └── user
        ├── cmd
        │   ├── api
        │   ├── cronjob
        │   ├── rmq
        │   ├── rpc
        │   └── script
        └── model
```

里面都目录说明：https://github.com/tal-tech/zero-doc/blob/main/doc/shorturl.md

```
rpc/transform
├── etc
│   └── transform.yaml              // 配置文件
├── internal
│   ├── config
│   │   └── config.go               // 配置定义
│   ├── logic
│   │   ├── expandlogic.go          // expand 业务逻辑在这里实现
│   │   └── shortenlogic.go         // shorten 业务逻辑在这里实现
│   ├── server
│   │   └── transformerserver.go    // 调用入口, 不需要修改
│   └── svc
│       └── servicecontext.go       // 定义 ServiceContext，传递依赖
├── pb
│   └── transform.pb.go
├── transform.go                    // rpc 服务 main 函数
├── transform.proto
└── transformer
    ├── transformer.go              // 提供了外部调用方法，无需修改
    ├── transformer_mock.go         // mock 方法，测试用
    └── types.go                    // request/response 结构体定义
```



### model生成

https://zeromicro.github.io/go-zero/model-gen.html

```shell
`goctl model mysql ddl -src user.sql -dir . -c`
```

```shell
`goctl model mysql datasource -url="$datasource" -table="user" -c -dir .`
```



### api文件编写

#### 编写user.api文件

```shell
$ `vim service/user/cmd/api/user.api`
```

```go
type (
    LoginReq {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    LoginReply {
        Id           int64 `json:"id"`
        Name         string `json:"name"`
        Gender       string `json:"gender"`
        AccessToken  string `json:"accessToken"`
        AccessExpire int64 `json:"accessExpire"`
        RefreshAfter int64 `json:"refreshAfter"`
    }
)

service user-api {
    @handler login
    post /user/login (LoginReq) returns (LoginReply)
}
```

#### 生成api服务

方式一

```shell
$ `cd book/service/user/cmd/api`
$ `goctl api go -api user.api -dir .`
```



### 业务编码

前面一节，我们已经根据初步需求编写了user.api来描述user服务对外提供哪些服务访问，在本节我们接着前面的步伐， 通过业务编码来讲述go-zero怎么在实际业务中使用。

#### 添加Mysql配置

```
$ vim service/user/cmd/api/internal/config/config.go
```

```go
package config

import (
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql struct{
		DataSource string
	}

	CacheRedis cache.CacheConf
}

```



#### 完善yaml配置

```shell
$ vim service/user/cmd/api/etc/user-api.yaml
```

```shell
Name: user-api
Host: 0.0.0.0
Port: 8888

Mysql:
  DataSource: root:root@tcp(10.4.7.40)/book?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 10.4.7.40:6379
    Pass: ''
    Type: node
```



#### 完善服务依赖

```shell
$ vim service/user/cmd/api/internal/svc/servicecontext.go
```

```go
type ServiceContext struct {
    Config    config.Config
    UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
    conn:=sqlx.NewMysql(c.Mysql.DataSource)
    return &ServiceContext{
        Config: c,
        UserModel: model.NewUserModel(conn,c.CacheRedis),
    }
}
```



#### 填充登录逻辑

```shell
$ vim service/user/cmd/api/internal/logic/loginlogic.go
```

```go
func (l *LoginLogic) Login(req types.LoginReq) (*types.LoginReply, error) {
    if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
        return nil, errors.New("参数错误")
    }

    userInfo, err := l.svcCtx.UserModel.FindOneByNumber(req.Username)
    switch err {
    case nil:
    case model.ErrNotFound:
        return nil, errors.New("用户名不存在")
    default:
        return nil, err
    }

    if userInfo.Password != req.Password {
        return nil, errors.New("用户密码不正确")
    }

    // ---start---
    now := time.Now().Unix()
    accessExpire := l.svcCtx.Config.Auth.AccessExpire
    jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
    if err != nil {
        return nil, err
    }
    // ---end---

    return &types.LoginReply{
        Id:           userInfo.Id,
        Name:         userInfo.Name,
        Gender:       userInfo.Gender,
        AccessToken:  jwtToken,
        AccessExpire: now + accessExpire,
        RefreshAfter: now + accessExpire/2,
    }, nil
}
```



### jwt鉴权

#### 添加配置定义和yaml配置项

```shell
$ vim service/user/cmd/api/internal/config/config.go

```

```go
type Config struct {
    rest.RestConf
    Mysql struct{
        DataSource string
    }
    CacheRedis cache.CacheConf
    Auth      struct {
        AccessSecret string
        AccessExpire int64
    }
}
```

```shell
$ vim service/user/cmd/api/etc/user-api.yaml

```

```go
Name: user-api
Host: 0.0.0.0
Port: 8888
Mysql:
  DataSource: $user:$password@tcp($url)/$db?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: $host
    Pass: $pass
    Type: node
Auth:
  AccessSecret: $AccessSecret
  AccessExpire: $AccessExpire
```



```shell
$ vim service/user/cmd/api/internal/logic/loginlogic.go

```

```go
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
  claims := make(jwt.MapClaims)
  claims["exp"] = iat + seconds
  claims["iat"] = iat
  claims["userId"] = userId
  token := jwt.New(jwt.SigningMethodHS256)
  token.Claims = claims
  return token.SignedString([]byte(secretKey))
}
```

#### search api使用jwt token鉴权

##### 编写search.api文件

```
$ vim service/search/cmd/api/search.api

```

```go
type (
    SearchReq {
        // 图书名称
        Name string `form:"name"`
    }

    SearchReply {
        Name string `json:"name"`
        Count int `json:"count"`
    }
)

@server(
    jwt: Auth
)
service search-api {
    @handler search
    get /search/do (SearchReq) returns (SearchReply)
}

service search-api {
    @handler ping
    get /search/ping
}
```

##### 生成代码

```shell
goctl api go -api search.api -dir .
```

##### 添加yaml配置项

```
$ vim service/search/cmd/api/etc/search-api.yaml

```

```
Name: search-api
Host: 0.0.0.0
Port: 8889
Auth:
  AccessSecret: $AccessSecret
  AccessExpire: $AccessExpire
```

##### 验证 jwt token

启动user api服务，登录

```shell
  $ cd service/user/cmd/api
  $ go run user.go -f etc/user-api.yaml
  Starting server at 0.0.0.0:8888...
```

```shell
  $ curl -i -X POST \
    http://127.0.0.1:8888/user/login \
    -H 'content-type: application/json' \
    -d '{
      "username":"666",
      "password":"123456"
  }'
  
  HTTP/1.1 200 OK
Content-Type: application/json
Date: Tue, 06 Apr 2021 13:43:56 GMT
Content-Length: 251

{"id":1,"name":"小明","gender":"男","accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTc3MjM4MzYsImlhdCI6MTYxNzcxNjYzNiwidXNlcklkIjoxfQ.mus_vnMuZkTKDaX4Gd0pMLjTSxvuyAIpwI1DFDCXtYY","accessExpire":1617723836,"refreshAfter":1617720236}
```

启动search api服务，调用`/search/do`验证jwt鉴权是否通过

```shell
  $ go run search.go -f etc/search-api.yaml
  Starting server at 0.0.0.0:8889...

```

我们先不传jwt token，看看结果

```shell
  $ curl -i -X GET \
    'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0'
  
  HTTP/1.1 401 Unauthorized
  Date: Mon, 08 Feb 2021 10:41:57 GMT
  Content-Length: 0
```

很明显，jwt鉴权失败了，返回401的statusCode，接下来我们带一下jwt token（即用户登录返回的`accessToken`）

```shell
  $ curl -i -X GET \
    'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
    -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI4NjcwNzQsImlhdCI6MTYxMjc4MDY3NCwidXNlcklkIjoxfQ.JKa83g9BlEW84IiCXFGwP2aSd0xF3tMnxrOzVebbt80'
```



##### 获取jwt token中携带的信息

go-zero从jwt token解析后会将用户生成token时传入的kv原封不动的放在http.Request的Context中，因此我们可以通过Context就可以拿到你想要的值

```
$ vim service/search/cmd/api/internal/logic/searchlogic.go

```

添加一个log来输出从jwt解析出来的userId。

```go
func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
    logx.Infof("userId: %v",l.ctx.Value("userId"))// 这里的key和生成jwt token时传入的key一致
    return &types.SearchReply{}, nil
}
```

运行结果

```
{"@timestamp":"2021-02-09T10:29:09.399+08","level":"info","content":"userId: 1"}

```



### 中间件使用

这里以`search`服务为例来演示中间件的使用

#### 路由中间件

* 重新编写`search.api`文件，添加`middleware`声明

```
  $ cd service/search/cmd/api
  $ vim search.api
```

```go
  @server(
      jwt: Auth
      middleware: Example // 路由中间件声明
  )
  service search-api {
      @handler search
      get /search/do (SearchReq) returns (SearchReply)
  }
```

* 重新生成api代码

```
  $ goctl api go -api search.api -dir .

```

生成完后会在`internal`目录下多一个`middleware`的目录，这里即中间件文件，后续中间件的实现逻辑也在这里编写。



* 完善资源依赖`ServiceContext`

```
 $ vim service/search/cmd/api/internal/svc/servicecontext.go

```

```go
 type ServiceContext struct {
     Config config.Config
     Example rest.Middleware
 }

 func NewServiceContext(c config.Config) *ServiceContext {
     return &ServiceContext{
         Config: c,
         Example: middleware.NewExampleMiddleware().Handle,
     }
 }
```

* 编写中间件逻辑 这里仅添加一行日志，内容example middle，如果服务运行输出example middle则代表中间件使用起来了。

```
  $ vim service/search/cmd/api/internal/middleware/examplemiddleware.go

```

```go
  func (m *ExampleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          logx.Info("example middle")
          next(w, r)
      }
  }
```

* 启动服务验证

```
  {"@timestamp":"2021-02-09T11:32:57.931+08","level":"info","content":"example middle"}
```



### rpc编写与调用

#### rpc服务编写

##### 编译proto文件

```
  $ vim service/user/cmd/rpc/user.proto

```

```go
  syntax = "proto3";

  package user;

  message IdReq{
    int64 id = 1;
  }

  message UserInfoReply{
    int64 id = 1;
    string name = 2;
    string number = 3;
    string gender = 4;
  }

  service user {
    rpc getUser(IdReq) returns(UserInfoReply);
  }
```

生成rpc服务代码

```shell
$ cd service/user/cmd/rpc
$ goctl rpc proto -src user.proto -dir .
```

##### 添加配置及完善yaml配置项

```
  $ vim service/user/cmd/rpc/internal/config/config.go

```

```go
  type Config struct {
      zrpc.RpcServerConf
      Mysql struct {
          DataSource string
      }
      CacheRedis cache.CacheConf
  }
```

```go
  $ vim /service/user/cmd/rpc/etc/user.yaml

```

```yaml
Name: user.rpc
ListenOn: 10.4.7.40:8080

Mysql:
  DataSource: root:root@tcp(10.4.7.40)/book?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai
CacheRedis:
  - Host: 10.4.7.40:6379
    Pass: ""
    Type: node

Etcd:
  Hosts:
    - 10.4.7.40:2379
  Key: user.rpc
```

##### 添加资源依赖

```
  $ vim service/user/cmd/rpc/internal/svc/servicecontext.go

```

```go
  type ServiceContext struct {
      Config    config.Config
      UserModel model.UserModel
  }

  func NewServiceContext(c config.Config) *ServiceContext {
      conn := sqlx.NewMysql(c.Mysql.DataSource)
      return &ServiceContext{
          Config: c,
          UserModel: model.NewUserModel(conn, c.CacheRedis),
      }
  }
```

##### 添加rpc逻辑

```go
  $ service/user/cmd/rpc/internal/logic/getuserlogic.go

```

```go
  func (l *GetUserLogic) GetUser(in *user.IdReq) (*user.UserInfoReply, error) {
      one, err := l.svcCtx.UserModel.FindOne(in.Id)
      if err != nil {
          return nil, err
      }

      return &user.UserInfoReply{
          Id:     one.Id,
          Name:   one.Name,
          Number: one.Number,
          Gender: one.Gender,
      }, nil
  }
```



#### 使用rpc

接下来我们在search服务中调用user rpc

##### 添加UserRpc配置及yaml配置项

```
  $ vim service/search/cmd/api/internal/config/config.go

```

```go
  type Config struct {
      rest.RestConf
      Auth struct {
          AccessSecret string
          AccessExpire int64
      }
      UserRpc zrpc.RpcClientConf
  }
```

```
  $ vim service/search/cmd/api/etc/search-api.yaml

```

```yaml
Name: search-api
Host: 0.0.0.0
Port: 8889
Auth:
  AccessSecret: "b60d7387-a0ef-eadc-7361-2b24a8b56851
                 0d5ef0d6-7a5d-4c8e-53bd-9689a3b9a37c
                 47f7e6d4-2dfd-8417-ae7a-4b59a0a924f8
                 9d61a9bb-d156-e9e8-59f5-619a6cc32c53
                 d4076129-ebe0-0ead-8e66-e5b35514f0eb
                 7f2ae9fb-ff3b-007e-5b59-df1578766122
                 11401e5e-fde7-e2e6-e92b-bf5605ae42d7
                 fd9ae522-6006-2aa9-7678-b23b17c884d4"
  AccessExpire: 7200

UserRpc:
  Etcd:
    Hosts:
      - 10.4.7.40:2379
    Key: user.rpc
```

##### 添加依赖

```
  $ vim service/search/cmd/api/internal/svc/servicecontext.go

```

```go
  type ServiceContext struct {
      Config  config.Config
      Example rest.Middleware
      UserRpc userclient.User
  }

  func NewServiceContext(c config.Config) *ServiceContext {
      return &ServiceContext{
          Config:  c,
          Example: middleware.NewExampleMiddleware().Handle,
          UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
      }
  }
```

##### 补充逻辑

```
  $ vim /service/search/cmd/api/internal/logic/searchlogic.go

```

```go
  func (l *SearchLogic) Search(req types.SearchReq) (*types.SearchReply, error) {
      userIdNumber := json.Number(fmt.Sprintf("%v", l.ctx.Value("userId")))
      logx.Infof("userId: %s", userIdNumber)
      userId, err := userIdNumber.Int64()
      if err != nil {
          return nil, err
      }

      // 使用user rpc
      _, err = l.svcCtx.UserRpc.GetUser(l.ctx, &userclient.IdReq{
          Id: userId,
      })
      if err != nil {
          return nil, err
      }

      return &types.SearchReply{
          Name:  req.Name,
          Count: 100,
      }, nil
  }
```



#### 启动并验证服务

- 启动etcd、redis、mysql
- 启动user rpc

```shell
  $ cd /service/user/cmd/rpc
  $ go run user.go -f etc/user.yaml
```

* 启动search api

```shell
$ cd service/search/cmd/api
$ go run search.go -f etc/search-api.yaml
```

* 验证服务

```shell
  $ curl -i -X GET \
    'http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0' \
    -H 'authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTI4NjcwNzQsImlhdCI6MTYxMjc4MDY3NCwidXNlcklkIjoxfQ.JKa83g9BlEW84IiCXFGwP2aSd0xF3tMnxrOzVebbt80'
   
  HTTP/1.1 200 OK
  Content
  -Type: application/json
  Date: Tue, 09 Feb 2021 06:05:52 GMT
  Content-Length: 32

  {"name":"西游记","count":100}
```



### 错误处理

#### user api之login

在之前，我们在登录逻辑中处理用户名不存在时，直接返回来一个error。我们来登录并传递一个不存在的用户名看看效果。

```shell
curl -X POST \
  http://127.0.0.1:8888/user/login \
  -H 'content-type: application/json' \
  -d '{
    "username":"1",
    "password":"123456"
}'

HTTP/1.1 400 Bad Request
Content-Type: text/plain; charset=utf-8
X-Content-Type-Options: nosniff
Date: Tue, 09 Feb 2021 06:38:42 GMT
Content-Length: 19

用户名不存在
```

接下来我们将其以json格式进行返回



#### 自定义错误

首先在common中添加一个`baseerror.go`文件，并填入代码

```go
  package errorx

  const defaultCode = 1001

  type CodeError struct {
      Code int    `json:"code"`
      Msg  string `json:"msg"`
  }

  type CodeErrorResponse struct {
      Code int    `json:"code"`
      Msg  string `json:"msg"`
  }

  func NewCodeError(code int, msg string) error {
      return &CodeError{Code: code, Msg: msg}
  }

  func NewDefaultError(msg string) error {
      return NewCodeError(defaultCode, msg)
  }

  func (e *CodeError) Error() string {
      return e.Msg
  }

  func (e *CodeError) Data() *CodeErrorResponse {
      return &CodeErrorResponse{
          Code: e.Code,
          Msg:  e.Msg,
      }
  }
```

将登录逻辑中错误用CodeError自定义错误替换

```shell
vim user/cmd/api/internal/logic/loginlogic.go
```

```go
  if len(strings.TrimSpace(req.Username)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
          return nil, errorx.NewDefaultError("参数错误")
      }

      userInfo, err := l.svcCtx.UserModel.FindOneByNumber(req.Username)
      switch err {
      case nil:
      case model.ErrNotFound:
          return nil, errorx.NewDefaultError("用户名不存在")
      default:
          return nil, err
      }

      if userInfo.Password != req.Password {
          return nil, errorx.NewDefaultError("用户密码不正确")
      }

      now := time.Now().Unix()
      accessExpire := l.svcCtx.Config.Auth.AccessExpire
      jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
      if err != nil {
          return nil, err
      }

      return &types.LoginReply{
          Id:           userInfo.Id,
          Name:         userInfo.Name,
          Gender:       userInfo.Gender,
          AccessToken:  jwtToken,
          AccessExpire: now + accessExpire,
          RefreshAfter: now + accessExpire/2,
      }, nil
```

#### 开启自定义错误

```
  $ vim service/user/cmd/api/user.go

```

```go
  func main() {
      flag.Parse()

      var c config.Config
      conf.MustLoad(*configFile, &c)

      ctx := svc.NewServiceContext(c)
      server := rest.MustNewServer(c.RestConf)
      defer server.Stop()

      handler.RegisterHandlers(server, ctx)

      // 自定义错误
      httpx.SetErrorHandler(func(err error) (int, interface{}) {
          switch e := err.(type) {
          case *errorx.CodeError:
              return http.StatusOK, e.Data()
          default:
              return http.StatusInternalServerError, nil
          }
      })

      fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
      server.Start()
  }
```





## CI/CD



## 服务部署



## 日志收集

## 链路追踪

https://zeromicro.github.io/go-zero/trace.html



## 服务监控

