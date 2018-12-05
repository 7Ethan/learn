package main

import (
	"fmt"
	"time"
)

func main() {
	ping := make(chan bool)
	msg := make(chan struct{})
	// str := []string{"A","B","C","D","E","F","G","H","I","J","K"}

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Print(i)
			<-ping
			msg <- struct{}{}
		}

	}()

	go func() {
		for j := 0; j < 10; j++ {
			// fmt.Print(str[j])
			fmt.Printf("%c", 'A'+i)
			ping <- true
			<-msg
		}
	}()

	time.Sleep(3 * time.Second)

}
