package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	b := []byte{}
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	hasher.Reset()
	data := []byte("we shall Overcome ! ")
	n, err := hasher.Write(data)
	if err != nil || n != len(data) {
		log.Printf("Hasher write error : %v / %v", n, err)
	}
	checksum := hasher.Sum(b)
	fmt.Printf("result : %x\n", checksum)
}
