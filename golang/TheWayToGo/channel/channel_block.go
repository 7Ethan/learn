package main

import (
	"fmt"
	"time"
)

/*
不将通道作为参数传递给协程，而用函数来生成一个通道
并返回(工厂角色);函数内有个匿名函数被协程调用
*/
func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(1 * time.Second)
}

//通道工厂模式
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}
