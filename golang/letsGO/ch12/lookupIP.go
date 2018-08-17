package main

import (
	"fmt"
	"net"
)

func main() {
	ips, err := net.LookupIP("www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ips)
}
