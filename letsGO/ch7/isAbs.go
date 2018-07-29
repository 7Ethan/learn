package main

import (
	"fmt"
	"path"
	"strings"
)

func main() {
	fmt.Println(path.IsAbs("/a/b/c"))
	fmt.Println(path.IsAbs(strings.Replace("c:\\Windows\\system", "\\", "/", -1)))
	/*
		Go只支持linux的绝对路径
	*/
}
