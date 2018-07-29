package main

import (
	"fmt"
	"os"
)

func main() {
	path := "/Users/shangan/Desktop/GO/src/Golang"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf(" %s is a directory", path)
	}
}
