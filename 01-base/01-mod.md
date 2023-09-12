## **初始化**

```text
 go mod init
```

## **添加依赖包**

```text
 go mod tidy
```

下载位置: **GOPATH/pkg/mod**

拉取缺少模块,移除不用的模块

**将依赖包放到当前vendor目录**

```text
 go mod vendor
```

将依赖复制到vendor下

**下载依赖包**

```text
 go mod download
```

**检验依赖**

```text
 go mod verify
```

**打印模块[依赖图](https://www.zhihu.com/search?q=依赖图&search_source=Entity&hybrid_search_source=Entity&hybrid_search_extra={"sourceType"%3A"answer"%2C"sourceId"%3A2311361609})**

```text
 go mod graph
```

其实也不用太担心,一般只要go get以后,本地都有备份,fakerjs事件发生概率也不是很高,而且github和gitee也会有其备份