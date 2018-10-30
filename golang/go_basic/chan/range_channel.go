package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 10)

	ch <- "one"
	ch <- "two"
	close(ch)

	for data := range ch {
		fmt.Println(data)
	}
	// close(ch)//it must NOT be here!!!! // it will cause all goroutine deadlock!!
}
