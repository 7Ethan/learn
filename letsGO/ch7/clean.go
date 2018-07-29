package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Clean("/a/b/../c"))
	fmt.Println(path.Clean("/a/b/../././c"))
}
