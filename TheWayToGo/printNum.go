package main

import (
	"fmt"
	"time"
)

func main() {
	chan_int := make(chan bool)
	chan_char := make(chan bool, 1)
	done := make(chan struct{})
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	go func() {
		for i := 0; i < 10; i += 2 {
			<-chan_char
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_int <- true
		}
	}()

	go func() {
		for i := 0; i < 10; i += 2 {
			<-chan_int
			fmt.Print(str[i])
			fmt.Print(str[i+1])
			chan_char <- true
		}
		done <- struct{}{}
	}()

	chan_char <- true
	<-done
	time.Sleep(3 * time.Second)
}
