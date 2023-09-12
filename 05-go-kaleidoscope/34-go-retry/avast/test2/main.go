package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/avast/retry-go"
)

func main() {
	var num int
	var err error
	// 定义一个重试策略
	retryStrategy := []retry.Option{
		retry.Delay(100 * time.Millisecond),
		retry.Attempts(5),
		retry.LastErrorOnly(true),
	}

	// 使用重试策略进行重试
	err = retry.Do(
		func() error {
			num, err = CheckNum()
			if err != nil {
				return err
			}
			return nil
		},
		retryStrategy...)

	if err != nil {
		fmt.Println("Error occurred after 5 retries")
	} else {
		fmt.Println(num)
	}
}

// CheckNum 生成一个随机数，如果小于50，则返回错误，如果大于50，则返回这个数
func CheckNum() (num int, err error) {
	fmt.Println("start check number")
	rand.Seed(time.Now().UnixNano())
	num = rand.Intn(100)
	if num < 50 {
		fmt.Println(num)
		return 0, errors.New("test error")
	} else {
		return
	}
}
