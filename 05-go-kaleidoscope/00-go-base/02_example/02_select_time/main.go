/*
* @Author: jettjia jettjia@qq.com
* @Date: 2023-11-16 13:31:39
* @LastEditors: jettjia jettjia@qq.com
* @LastEditTime: 2023-11-16 13:33:46
* @FilePath: \05-go-kaleidoscope\00-go-base\example\select_time\main.go
* @Description:
有default语句且其他case语句不满足，则执行default语句输出"case default"；
如果没有default语句，那么1秒之后，
第二个case语句满足，则会输出"case timeout"，
这也是select常见的用法之一：判断超时。
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch4 := make(chan int) //无缓存通道

	select {
	case <-ch4:
		fmt.Println("case 4")
	case <-time.After(time.Second * 1):
		fmt.Println("time out")
	}
}
