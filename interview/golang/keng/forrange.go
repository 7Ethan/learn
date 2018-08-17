package main

import (
	"fmt"
	"time"
)

func main() {
	str := []string{"I", "am", "Sergey"}
	for _, v := range str {
		go func(v string) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(3 * time.Second)
}
