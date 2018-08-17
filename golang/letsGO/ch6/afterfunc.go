package main

import (
	"fmt"
	"time"
)

func test() {
	fmt.Println("Test time")
}

func main() {
	fmt.Println(time.Now())
	time.AfterFunc(3*time.Second, test)
	//var str string
	//fmt.Scan(&str)
	time.Sleep(10 * time.Second)
}
