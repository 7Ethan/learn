package main

import (
	"fmt"
	"path"
)

func main() {
	dir, file := path.Split("/a/b/c/test.c")
	fmt.Println(dir)
	fmt.Println(file)
}
