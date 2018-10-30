package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	jobs := make(chan int, 5)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job : ", j)
			} else {
				fmt.Println("All jobs have been done.")
				done <- struct{}{}
				return
			}
		}
	}()

	for i := 0; i < 3; i++ {
		jobs <- i
	}
	close(jobs) // must close channel
	<-done
}
