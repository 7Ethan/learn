package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "https://www.7ethan.top"

	u,err := url.Parse(str)
	if err != nil{
		panic(err)
	}
	fmt.Println(u.Host)
}
