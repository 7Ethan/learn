package main

import (
	"fmt"
	"log"
	"math/rand"
	"runtime"
	"time"
)

func generator() <-chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Int63n(3000)) * time.Millisecond)
			out <- i
			// fmt.Println()
			func() {
				_, file, line, _ := runtime.Caller(1)
				log.Printf("%s : %d\n", file, line)
			}()
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	tm := time.After(time.Duration(10) * time.Second)
	for {
		select {
		case <-c1:
			fmt.Println("Received data from c1 :", <-c1)
		case <-c2:
			fmt.Println("Received data from c2 :", <-c2)
		case <-tm:
			fmt.Println("Over")
			return
		}
	}
}

// func generator() chan int {
// 	out := make(chan int)
// 	go func() {
// 		i := 0
// 		for {
// 			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
// 			out <- i
// 			i++
// 		}
// 	}()
// 	return out
// }

// func main() {
// 	var c1, c2 = generator(), generator()
// 	tm := time.After(time.Second * 10)
// 	for {
// 		select {
// 		case <-c1:
// 			fmt.Println("Received from c1 value : ", <-c1)
// 		case <-c2:
// 			fmt.Println("Received from c2 value : ", <-c2)
// 		case <-time.After(500 * time.Millisecond):
// 			fmt.Println("timeout!")
// 		case <-tm:
// 			fmt.Println("bye")
// 			return
// 		}
// 	}
// }
