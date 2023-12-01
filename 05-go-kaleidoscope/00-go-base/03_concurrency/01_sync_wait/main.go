package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

func main() {
	var numberTasks = [5]string{"13456755448", " 13419385751", "13419317885", " 13434343439", "13438522395"}
	client = &http.Client{}
	beg := time.Now()

	wg := &sync.WaitGroup{}
	for _, keyword := range numberTasks {
		wg.Add(1)
		go func(keyword string, group *sync.WaitGroup) {
			body, err := NumberQueryRequest(keyword)
			if err != nil {
				fmt.Printf("error occurred in query keyword: %s, error: %s\n", keyword, err.Error())
			} else {
				fmt.Printf("search %s success, data size is %d\n, body is %s\n", keyword, len(body), string(body))
			}
			group.Done()
		}(keyword, wg)
	}

	wg.Wait() //waitGroup在主线程等待所有任务完成
	fmt.Printf("time consumed: %fs", time.Now().Sub(beg).Seconds())
}

var client *http.Client

// GET请求的通用写法, 注意复用client即可.
func NumberQueryRequest(keyword string) (body []byte, err error) {
	url := fmt.Sprintf("https://api.binstd.com/shouji/query?appkey=df2720f76a0991fa&shouji=%s", keyword)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
	if err != nil {
		return nil, err
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		data, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("response status code is not OK, response code is %d, body:%s", resp.StatusCode, string(data))
	}
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
