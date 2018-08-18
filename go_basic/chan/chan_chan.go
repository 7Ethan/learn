package main

import (
	"fmt"
	"time"
)

type Request struct {
	Name string
	// nsg  chan int
}

var r Request = Request{Name: "baidu"}
var rh chan Request = make(chan Request, 1)

func main() {
	rh <- r
	ch1 := make(chan chan Request, 10)
	go func() {
		for i := 0; i < 10; i++ {
			request := <-rh
			ch1 <- rh
			fmt.Println("ch1 <- rh", " === ", request.Name)
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v := <-ch1
			fmt.Printf(" <- ch1 : %V\n", v)
			rh <- r
		}
	}()

	time.Sleep(time.Second * 1)

}
