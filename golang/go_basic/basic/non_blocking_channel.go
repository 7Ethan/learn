package main

import (
	"fmt"
)

func main() {
	message := make(chan string)
	// signal := make(chan bool)

	// go func() {
		
	// 	tmp := <-message
	// 	fmt.Println(tmp)
	// }()

	msg := "Hello"
	select {
	case message <- msg:
		fmt.Println("received message")
	default:
		fmt.Println("nothing.")
	}
}
