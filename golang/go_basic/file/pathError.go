package main

import (
	"fmt"
	"os"
	"syscall"
)

func main() {
	for try := 0; try < 2; try++ {
		file, err := os.Create("hello.txt")
		if err == nil {
			return
		}
		if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
			fmt.Println("delete temfile")
			continue
		}
		return
	}
}
