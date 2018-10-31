package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	str := "this is for SHA1"

	h := sha1.New()

	h.Write([]byte(str))

	bs := h.Sum(nil) //nil?
	fmt.Printf("%x\n",bs)
}
