package main

import (
	"fmt"
)

/*
*
* @description:
1、三个case都是通信
2、三个case都是channal表达式(2次写入和1次读取)，都会被求值；
3、三个case都满足，则随机挑选一个执行，故多次执行后输出可能不同。
* @return {*}
*/
func main() {
	ch1 := make(chan int, 4)
	ch2 := make(chan int, 4)
	ch3 := make(chan int, 4)
	ch3 <- 3

	select {
	case ch1 <- 1:
		fmt.Println("case 1")
		fmt.Println("ch1 is:", <-ch1)
	case ch2 <- 2:
		fmt.Println("case 2")
		fmt.Println("ch2 is:", <-ch2)
	case <-ch3:
		fmt.Println("case 3")
	default:
		fmt.Println("default")
	}
}
