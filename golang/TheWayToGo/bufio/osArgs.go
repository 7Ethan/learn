package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str := "hello"
	if len(os.Args) > 0 {
		str += strings.Join(os.Args, " ")
	}
	fmt.Println(str)
}
