package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch := make(chan int, 64)
	go Producer(2, ch)
	go Producer(700, ch)
	go Comsumer(ch)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	fmt.Println("Quit : %v", <-sig)
}

func Producer(factor int, out chan<- int) {
	for i := 1; i < 10; i++ {
		out <- i * factor
		time.Sleep(time.Duration(time.Millisecond * 100))
	}
}

func Comsumer(in <-chan int) {
	for v := range in {
		fmt.Println("Receive form channel : ", v)
		time.Sleep(time.Duration(time.Second * 1))
	}
}
