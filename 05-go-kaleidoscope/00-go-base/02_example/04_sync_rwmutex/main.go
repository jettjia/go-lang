/*
* @Author: jettjia jettjia@qq.com
* @Date: 2023-11-16 13:53:02
  - @LastEditors: jettjia jettjia@qq.com
  - @LastEditTime: 2023-11-16 14:12:20

* @FilePath: \05-go-kaleidoscope\00-go-base\02_example\sync_rwmutex\main.go
* @Description: 读写互斥锁
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
这种场景下使用读写锁是更好的一种选择。读写锁在Go语言中使用`sync`包中的`RWMutex`类型

读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的`goroutine`如果是获取读锁会继续获得锁，
如果是获取写锁就会等待；当一个`goroutine`获取写锁之后，其他的`goroutine`无论是获取读锁还是写锁都会等待。
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	rwlock sync.RWMutex
)

func main() {
	start := time.Now()

	// 写
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	// 读
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()

	end := time.Now()
	fmt.Println(end.Sub(start))
	fmt.Println("x:", x)
}

func write() {
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(time.Millisecond * 10) // 假设写操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	wg.Done()
}

func read() {
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	wg.Done()
}
