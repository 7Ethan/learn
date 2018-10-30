package main

import (
	//"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	buffer, err := os.OpenFile("1.txt", os.O_RDWR|os.O_APPEND, 0666)
	defer buffer.Close()
	if err != nil {
		fmt.Println("xxxx")
	}
	//fmt.Println(string(buffer))
	fmt.Fprintf(buffer, "456------") //必须使用os.O_APPEND才能追加,否则会只能覆盖
	io.Copy(os.Stdout, buffer)

	//var b  bytes.Buffer
	//b.Write([]byte("123"))
	//
	//fmt.Fprintf(&b,"456")
	//io.Copy(os.Stdout,&b)
}
