package main

import (
	"fmt"
	"log"
	"runtime"
)

type Where func()

var where Where = func() {
	_, file, line, _ := runtime.Caller(3)
	log.Printf("%s : %d ", file, line)
}

func f() (ret int) {
	defer func() {
		ret++
	}()
	where()
	return 1
}

func main() {
	fmt.Println(f())
	// start := time.Now()
}
