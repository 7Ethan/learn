package main

import (
	"fmt"
	_ "time"
)

func main() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch) //main执行完会杀死所有协程,此时这个线程的死循环也终结了
	//time.Sleep(9 * time.Second)

}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}
func getData(ch chan string) {
	var input string
	// time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
