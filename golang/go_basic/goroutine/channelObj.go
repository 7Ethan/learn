package main

//import "debug/dwarf"

type Request struct {
	data []int
	ret  chan int
}

func NewResquest(data ...int) *Request {
	return &Request{data, make(chan int, 1)}
}

func Process(req *Request) {
	x := 0
	for _, i := range req.data {
		x += i
	}
}
func main() {
}
