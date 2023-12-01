package main

import (
	"context"
	"fmt"
	"time"
)

func PrintTask(ctx context.Context) {

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("A man must walk down many roads.")
		}
	}
}

/*
*
一个 Context 不能拥有 Cancel 方法，同时我们也只能 Done channel 接收数据。
其中的原因是一致的：接收取消信号的函数和发送信号的函数通常不是一个。

	典型的场景是：父操作为子操作操作启动 goroutine，子操作也就不能取消父操作。
*/
func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	defer cancel()

	go PrintTask(ctx)
	go PrintTask(ctx)
	go PrintTask(ctx)

	time.Sleep(3 * time.Second)
	fmt.Println("main exit...")
}
