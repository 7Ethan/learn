package main

import (
	"log"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatal("error listen :", err)
	}
	defer l.Close()
	log.Println("listen OK! ")

	var i int
	for {
		time.Sleep(10 * time.Second)
		if _, err := l.Accept(); err != nil {
			log.Fatal("accept error : ", err)
		}
		i++
		log.Printf("%d accept a new connection .", i)
	}
}
