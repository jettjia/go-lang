package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/panjf2000/ants/v2"
)

type taskFunc func()

func main() {
	// 创建一个WorkPool对象，用于并发执行任务
	pool, _ := ants.NewPool(10) // 设置最大并发数：10
	defer pool.Release()        // 确保在程序退出时停止工作池

	// 创建一个等待组，用于等待所有任务完成
	var wg sync.WaitGroup
	wg.Add(10) // 任务数量为10

	// 提交任务到工作池，并使用等待组等待任务完成
	for i := 1; i <= 10; i++ {
		go pool.Submit(taskFuncWrapper(i, &wg)) // 提交任务到工作池，并传入任务ID作为参数
	}

	wg.Wait() // 等待所有任务完成
	fmt.Println("所有请求完成")
	fmt.Printf("running goroutines: %d\n", ants.Running())
}

// 定义请求处理函数
func taskFuncWrapper(id int, wg *sync.WaitGroup) taskFunc {
	return func() {
		url := fmt.Sprintf("http://www.baidu.com?id=%d", id) // 构造请求URL
		resp, err := http.Get(url)                           // 发送GET请求
		if err != nil {
			fmt.Printf("请求失败：%s\n", err)
			return
		}
		defer resp.Body.Close()

		// 处理响应内容（这里仅打印状态码）
		fmt.Printf("请求%d响应状态码：%d\n", id, resp.StatusCode)

		wg.Done()
	}
}
