package main

import (
	"log"
	"net"
	"time"
)

func establishConn(i int) net.Conn {
	conn, err := net.Dial("tcp", ":8088")
	if err != nil {
		log.Printf("%d dial error :", i)
		return nil
	}
	log.Println(i, ":connection to server OK")
	return conn
}
func main() {
	var sl []net.Conn
	for i := 0; i < 1000; i++ {
		conn := establishConn(i)
		if conn != nil {
			sl = append(sl, conn)
		}
	}
	time.Sleep(10000 * time.Second)
}
