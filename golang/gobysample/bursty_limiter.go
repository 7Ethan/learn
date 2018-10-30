package main

import (
	"fmt"
	"time"
)

func main() {
	burstyLimiter := make(chan time.Time, 3)

	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}
	go func() {
		for t := range time.Tick(500 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	burstyRequest := make(chan int, 10)
	for i := 0; i < 10; i++ {
		burstyRequest <- i
	}
	close(burstyRequest)
	for job := range burstyRequest {
		<-burstyLimiter
		fmt.Println("Request : ", job)
		// 首先有三个 request 瞬间打印出来，接下来 500 * time.Millisecond 的时间间隔处理一个request
	}
}
