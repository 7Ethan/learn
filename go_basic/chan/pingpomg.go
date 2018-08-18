package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan struct{}, 1)
	pong := make(chan struct{}, 1)

	// convey := make(chan int, 10)
	ping <- struct{}{}

	go func() {
		for i := 0; i < 10; i++ {
			<-ping
			fmt.Print(i)
			pong <- struct{}{}
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-pong
			fmt.Printf("%c", 65+i)
			ping <- struct{}{}
		}
	}()

	time.Sleep(time.Second * 3)
	close(ping)
	close(pong)
}
