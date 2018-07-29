package main

import (
	"fmt"
	"runtime"
)

func main() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf(" %d kb \n", m.Alloc/1024)
	fmt.Println("hello")
}
