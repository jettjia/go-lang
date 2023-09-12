package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/avast/retry-go"
)

func main() {
	err := retry.Do(
		func() error {
			_, err := http.Get("https://www.baidu99999.com")
			return err
		},
		retry.Delay(time.Second),          //设置重试之间的延迟时间
		retry.Attempts(3),                 //设置重试次数的最大值
		retry.DelayType(retry.FixedDelay), //设置重试延迟的类型
	)
	if err != nil {
		fmt.Println("请求失败：", err)
	} else {
		fmt.Println("请求成功")
	}
}
