package main

import (
	"fmt"
	"time"
)

func main() {
	req := make(chan int, 5)
	for i := 0; i < 5; i++ {
		req <- i
	}
	close(req)

	limiter := time.NewTicker(200 * time.Millisecond)

	for r := range req {
		<-limiter.C
		fmt.Println("Request : ", r, time.Now())
	}
}
