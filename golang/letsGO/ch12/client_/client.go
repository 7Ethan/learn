package client_

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
	"strings"
	_ "unicode"
)

const (
	LS   = "LS"
	CD   = "CD"
	PWD  = "PWD"
	QUIT = "QUIT"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(" Please input cmd : ")
		line, err := reader.ReadString('\n')
		checkError(err)
		line = strings.TrimSpace(line)
		line = strings.ToUpper(line)
		arr := strings.SplitN(line, " ", 2)
		fmt.Println(arr)
		switch arr[0] {
		case LS:
			SendRequest(LS)
		case CD:
			SendRequest(CD + " " + strings.TrimSpace(arr[1]))
		case PWD:
			SendRequest(PWD)
		case QUIT:
			fmt.Println("Program exit ")
			return
		default:
			fmt.Println("Wrong cmd!")
		}
	}
}

//send request
func SendRequest(cmd string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:7076")
	checkError(err)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)
	SendData(conn, cmd)
	fmt.Println(ReadData(conn))
	conn.Close()

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

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
