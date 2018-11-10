package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	os.Setenv("sergey", "100")
	fmt.Println(os.Getenv("sergey"))

	for i, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(i, pair)
	}
}
