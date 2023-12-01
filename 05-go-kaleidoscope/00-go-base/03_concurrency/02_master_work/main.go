package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	routineCountTotal = 5 //限制线程数
)

func main() {
	var numberTasks = [5]string{"13456755448", " 13419385751", "13419317885", " 13434343439", "13438522395"}
	client = &http.Client{}
	beg := time.Now()

	wg := &sync.WaitGroup{}
	tasks := make(chan string)
	results := make(chan string)
	//receiver接受响应并处理的函数块, 也可以单独写在一个函数
	go func() {
		for result := range results {
			if result == "" {
				close(results)
			} else {
				fmt.Println("result:", result)
			}
		}
	}()
	for i := 0; i < routineCountTotal; i++ {
		wg.Add(1)
		go worker(wg, tasks, results)
	}
	//分发任务
	for _, task := range numberTasks {
		tasks <- task
	}

	tasks <- ""   //worker结束标志
	wg.Wait()     //同步结束
	results <- "" // result结束标志

	fmt.Printf("time consumed: %fs", time.Now().Sub(beg).Seconds())
}

func worker(group *sync.WaitGroup, tasks chan string, result chan string) {
	for task := range tasks {
		if task == "" {
			close(tasks)
		} else {
			respBody, err := NumberQueryRequest(task)
			if err != nil {
				fmt.Printf("error occurred in NumberQueryRequest: %s\n", task)
				result <- err.Error()
			}
			result <- string(respBody)
		}
	}
	group.Done()
}

var client *http.Client

func NumberQueryRequest(keyword string) (body []byte, err error) {
	url := fmt.Sprintf("https://api.binstd.com/shouji/query?appkey=df2720f76a0991fa&shouji=%s", keyword)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
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
