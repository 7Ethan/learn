package main

import (
	"flag"
	"fmt"
)

func style() {
	methodPtr := flag.String("method", "default", "method of sample")
	value := flag.Int("value", -1, "value of sample")

	flag.Parse()

	fmt.Println(*methodPtr, *value)
}

func style1() {
	var method string
	var value int

	flag.StringVar(&method, "method", "default", "method of sample")
	flag.IntVar(&value, "value", -1, "value of sample")

	flag.Parse()

	fmt.Println(method, value)
}

// func main(){
// 	style1()
// }
