package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/*  打开 D:\\新建文本文档.txt 文件，
	如果文件不存在将会新建，
	如果已 存在，新写入的内容将追加到文件尾
	*/
	f, err := os.OpenFile("/Users/shangan/Desktop/GO/src/Golang/letsGO/ch7/test.txt",
		os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	f.WriteString("\r\nCome on \r\n")
	buf := make([]byte, 1024)
	var str string
	/*重置文件指针，否则读不到内容的。*/
	f.Seek(0, io.SeekStart)
	for {
		n, ferr := f.Read(buf)
		if ferr != nil && ferr != io.EOF {
			fmt.Println(ferr.Error())
			break
		}
		if n == 0 {
			break
		}
		fmt.Println(n)
		str += string(buf[0:n])
	}
	fmt.Println(str)
	f.Close()
}
