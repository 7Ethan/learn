package main

import "fmt"

func main() {
	data := make(chan int)
	exit := make(chan bool)

	go func() {
		for d := range data {
			fmt.Println(d)
		}
		fmt.Println("processed")
		exit <- true
	}()

	for i := 0; i < 10; i++ {
		data <- i
	}
	close(data)

	fmt.Println(" send over")
	<-exit
}
