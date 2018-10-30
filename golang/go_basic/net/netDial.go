package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("begin to connect .")
	conn, err := net.DialTimeout("tcp", ":8088", 10)
	if err != nil {
		log.Fatal("dial Error : ", err)
		return
	}
	defer conn.Close()
	log.Println("dial ok .")
}
