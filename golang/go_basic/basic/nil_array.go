package main

import (
	"fmt"
)

func main() {
	var a []int
	b := a[:]
	fmt.Println(b == nil)
	if b == nil {
		println("b is nil")
	} else {
		println("b is not nil")
	}
	if a == nil {
		fmt.Println("a is nil")
	} else {
		println("a not nil")
	}

	/**
	true
	b is nil
	a is nil
	*/
}
