/*
 * @Author: jettjia jettjia@qq.com
 * @Date: 2023-11-16 15:17:38
 * @LastEditors: jettjia jettjia@qq.com
 * @LastEditTime: 2023-11-16 15:18:56
 * @FilePath: \05-go-kaleidoscope\00-go-base\02_example\06_context_deadline\main.go
 * @Description: context deadline
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	}
}
