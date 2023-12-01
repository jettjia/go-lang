/*
* @Author: jettjia jettjia@qq.com
* @Date: 2023-11-16 14:18:09
* @LastEditors: jettjia jettjia@qq.com
* @LastEditTime: 2023-11-16 15:16:09
* @FilePath: \05-go-kaleidoscope\00-go-base\02_example\05_context_cancel\main.go
* @Description: context cancel

gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。
gen的调用者在使用生成的整数之后需要取消上下文，以免`gen`启动的内部goroutine发生泄漏
*/
package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当我们取完需要的整数后调用cancel

	for n := range gen(ctx) {
		fmt.Println("n:", n)
		if n == 5 {
			break
		}
	}

}

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)

	n := 1
	go func() {

		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()

	return dst
}
