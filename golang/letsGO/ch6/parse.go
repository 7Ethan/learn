package main

import (
	"fmt"
	"time"
)

func main() {
	d, err := time.Parse("01-02-2006", "23-08-2017")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(d)
}
