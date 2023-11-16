package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	go func() {
		fmt.Println("go work...")
		time.Sleep(time.Second * 5)
		ch <- struct{}{}
	}()

	<-ch

	fmt.Println("finished")
}
