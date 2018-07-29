package main

import "fmt"

func producer(c chan<- int) {
	defer close(c)
	for i := 0; i < 10; i++ {
		c <- i
	}
}

/*
	chan<- int 只能发送,从这个通道读取操作是语法错误
	<-chan bool 只能读取,往这个通道发送内容是语法错误
*/

func comsumer(c <-chan int, f chan<- bool) {
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			break
		}
	}
	f <- true
}

func main() {
	buf := make(chan int)
	flag := make(chan bool)
	go producer(buf)
	go comsumer(buf, flag)
	fmt.Println(<-flag)
}
