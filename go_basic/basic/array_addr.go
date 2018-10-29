package main

import "fmt"

func main() {
	num := make([]int, 2, 5)
	fmt.Printf("%p\n", num)
	fmt.Printf("%p\n", &num[0])
	fmt.Printf("%p\n", &num)
}

/*
result : 
	0xc00001a180
	0xc00001a180
	0xc00000a080
*/
