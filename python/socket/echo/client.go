package main

import (
	"fmt"
	"net"
	"time"
)

//事实证明，使用 go cli 往Python server 发送数据完全没问题，只要是字节流都可以。

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error :", err)
	}
	defer conn.Close()

	for i := 0;i < 10;i++{
		go func(i int) {
			_,err = conn.Write([]byte(fmt.Sprintf("Hello ^_^ Message from {goroutine:%d} ! \n",i)))
			if err != nil{
				fmt.Println("Client write data Error : ",err)
			}
		}(i)
	}

	var buffer [1024]byte
	readBuffer := buffer[:10]
	_,err = conn.Read(readBuffer)
	if err != nil{
		fmt.Println("Client receive date Error : ",err)
	}
	fmt.Println(string(readBuffer[0:]))
	time.Sleep(time.Second * 2)
}
