package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1000)

	go func() {
		for {
			value, ok := <-ch
			if !ok {
				fmt.Println("Can't read from channel!")
				return
			}
			fmt.Println(value)
		}
	}()

	go func() {
		fmt.Println("ch <- i")
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 100)
}

/*
ok
ch <- i
Can't read from channel!
panic: send on closed channel

goroutine 5 [running]:
main.main.func1(0xc00006e000)
        ~/go/src/github.com/7Ethan/learn/interview/golang/goroutine/closed_channel.go:12 +0x43
created by main.main
        ~/go/src/github.com/7Ethan/learn/interview/golang/goroutine/closed_channel.go:10 +0x5c
exit status 2
*/
