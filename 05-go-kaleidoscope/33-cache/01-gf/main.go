package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gctx"
	"time"
)

func main() {
	var (
		ctx = gctx.New()
	)
	// 当键名不存在时写入，设置过期时间1000毫秒
	_, err := gcache.SetIfNotExist(ctx, "k1", "v1", time.Second)
	if err != nil {
		panic(err)
	}

	// 打印当前的键名列表
	keys, err := gcache.Keys(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)

	// 打印当前的键值列表
	values, err := gcache.Values(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(values)

	// 获取指定键值，如果不存在时写入，并返回键值
	value, err := gcache.GetOrSet(ctx, "k2", "v2", 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)

	// 打印当前的键值对
	data1, err := gcache.Data(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(data1)

	// 等待1秒，以便k1:v1自动过期
	time.Sleep(time.Second)

	// 再次打印当前的键值对，发现k1:v1已经过期，只剩下k2:v2
	data2, err := gcache.Data(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(data2)
}
