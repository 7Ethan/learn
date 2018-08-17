package main

import (
	"fmt"
	//"runtime"
	"runtime"
	"time"
)

func SayHello() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " : Hello")
		runtime.Gosched()
	}
}

func SayWorld() {
	for i := 0; i < 10; i++ {
		fmt.Println(i, " : World")
		runtime.Gosched()
	}
}
func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	go SayHello()
	go SayWorld()
	//fmt.Print(n)
	time.Sleep(5 * time.Second)

}
