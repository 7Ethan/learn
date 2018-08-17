package main

import (
	"fmt"
	"time"
	_ "time"
)

func f(in chan int) {
	fmt.Println(<-in)
}

func f2(c chan int, num int) {
	c <- num
}

func main() {
	out := make(chan int)
	go f2(out, 3)
	go f(out) //新开一个协程用来接收通道里的数据,此时另外的线程能够感知数据已经被读取,从而解决了阻塞
	time.Sleep(3 * time.Second)
}
