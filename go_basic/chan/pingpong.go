package main

import (
	"fmt"
)

func main() {
	ping := make(chan struct{}, 1)
	pong := make(chan struct{}, 1)
	done := make(chan struct{})

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
			fmt.Printf("%c", 'A'+i)
			ping <- struct{}{}
		}
		done <- struct{}{}
	}()

	<-done
	close(ping)
	close(pong)
}
