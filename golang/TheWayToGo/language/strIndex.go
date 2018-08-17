package main

import (
	"fmt"
	//"strings"
)

func main() {
	//str := "hei Sergey"
	//fmt.Println(strings.Index(str,"Ser"))
	var a int = 1
	var p *int
	p = &a

	*p++
	fmt.Println(*p)
}
