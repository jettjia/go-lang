# linux环境Golang配置

```shell
cd ~/Downloads
wget https://studygolang.com/dl/golang/go1.16.linux-amd64.tar.gz
sudo tar zxvf go1.16.linux-amd64.tar.gz -C /usr/local

vim .bashrc
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
source .bashrc


```

# 设置 goproxy

```
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct
```

# 使用 go module

```
go env -w GO111MODULE=on
```

# 交叉编译

交叉编译依赖下面几个环境变量：
`$GOARCH` 目标平台（编译后的目标平台）的处理器架构（386、amd64、arm）
`$GOOS` 目标平台（编译后的目标平台）的操作系统（darwin、freebsd、linux、windows）

## Mac

```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

## Linux

```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

## Windows

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build



# 检测是否配置成功

```
go version
go env
```

