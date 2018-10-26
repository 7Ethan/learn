package main

import "fmt"

type request struct {
	args       []int
	f          func([]int) int
	resultchan chan int
}

func sum(a []int) (s int) {
	for _, v := range a {
		s += v
	}
	return s
}

func serverHandle(queue chan *request) {
	for req := range queue {
		req.resultchan <- req.f(req.args)
	}
}

func main() {
	request := &request{[]int{3, 4, 5}, sum, make(chan int)} // Send request
	clientRequests <- request
	// Wait for response.
	fmt.Printf("answer: %d\n", <-request.resultchan)
}
