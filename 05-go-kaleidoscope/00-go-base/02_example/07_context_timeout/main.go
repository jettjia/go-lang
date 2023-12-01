/*
 * @Author: jettjia jettjia@qq.com
 * @Date: 2023-11-16 15:19:55
 * @LastEditors: jettjia jettjia@qq.com
 * @LastEditTime: 2023-11-16 15:20:00
 * @FilePath: \05-go-kaleidoscope\00-go-base\02_example\07_context_timeout\main.go
 * @Description: context timeout
 */
package main

import (
	"context"
	"fmt"
	"sync"

	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connecting ...")
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	// 设置一个50毫秒的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel() // 通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
