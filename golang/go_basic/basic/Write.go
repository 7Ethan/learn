package main

import (
	"fmt"
	"log"
)

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	*p = slice
	return len(data), nil
}

func main() {
	var str ByteSlice = ByteSlice{'a', 's', 'd'}
	num, err := str.Write([]byte("???"))
	if err != nil {
		log.Fatal("ERROR")
	}
	fmt.Println(num)
	fmt.Printf("resoure : %s", str)
}
