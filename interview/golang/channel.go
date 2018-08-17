package main

import "fmt"

func main() {
	done := make(chan struct{})
	ch := make(chan interface{})

	for i := 1; i < 6; i++ {
		go func(idx int) {
			select {
			case ch <- idx * 100:
				fmt.Println("have receive ",  <- ch)
			case <-done:
				// return (do nothing)
			}
		}(i)
	}

	fmt.Println("==>",<-ch)
	close(done)
	fmt.Println("exit~")
	
}
