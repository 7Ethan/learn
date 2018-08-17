package main

import "fmt"

const (
	x = iota
	y
	z = "zz"
	k
	p = iota
	q
)

func main() {
	fmt.Println(x, y, z, k, p, q)
	// 0 1 zz zz 4 5
}
