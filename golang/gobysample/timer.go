package main

import (
	"fmt"
	"time"
)

func main() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("timer 1 timeout!")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 goroutine!")
	}()

	stop := timer2.Stop()
	if stop {
		fmt.Println("timer has stop.")
	}
}
