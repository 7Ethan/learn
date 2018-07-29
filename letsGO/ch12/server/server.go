package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

const (
	LS  = "LS"
	CD  = "CD"
	PWD = "PWD"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7076")
	CheckError(err)
	listener, err1 := net.ListenTCP("tcp", tcpAddr)
	CheckError(err1)
	for {
		//wait to connect
		conn, err2 := listener.Accept()
		if err != nil {
			fmt.Println(err2)
			continue
		}
		fmt.Println("Receive client's request!")
		go ServerClient(conn)
	}
}

func ServerClient(conn net.Conn) {
	defer conn.Close()
	str := ReadData(conn)
	if str == "" {
		SendData(conn, "接受数据出错")
		return
	}
	fmt.Println("收到命令 : ", str)
	switch str {
	case LS:
		ListDir(conn)
	case PWD:
		Pwd(conn)
	default:
		if str[0:2] == CD {
			Chdir(conn, str[3:])
		} else {
			SendData(conn, "命令错误")
		}
	}
}

func Chdir(conn net.Conn, s string) {
	err := os.Chdir(s)
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, "OK")
	}
}

func ListDir(conn net.Conn) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		SendData(conn, err.Error())
		return
	}
	var str string
	for i, j := 0, len(files); i < j; i++ {
		f := files[i]
		str += f.Name() + "\t"
		if f.IsDir() {
			str += "dir\r\n"
		} else {
			str += "file\r\n"
		}
	}
	SendData(conn, str)
}

func ReadData(conn net.Conn) string {
	var data bytes.Buffer
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			fmt.Println(err)
			return ""
		}
		if buf[n-1] == 0 {
			data.Write(buf[0 : n-1])
			break
		} else {
			data.Write(buf[0:n])
		}
	}
	return string(data.Bytes())
}

func SendData(conn net.Conn, data string) {
	buf := []byte(data)
	buf = append(buf, 0)
	_, err := conn.Write(buf)
	if err != nil {
		fmt.Println(err)
	}
}

func Pwd(conn net.Conn) {
	s, err := os.Getwd()
	if err != nil {
		SendData(conn, err.Error())
	} else {
		SendData(conn, s)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
