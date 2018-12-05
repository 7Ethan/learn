package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan bool)
	message := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			<-ping
			fmt.Print(i)
			ping <- true
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-ping
			fmt.Printf("%c", 'A'+i)
			ping <- true
		}
		message <- struct{}{}
	}()

	ping <- true
	<-message
	time.Sleep(3 * time.Second)
}
