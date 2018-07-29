package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	select {
	case <-c:
		fmt.Println("receive data")
	case <-time.After(5 * time.Second):
		fmt.Println("timeout")
	}
}
