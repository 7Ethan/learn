package main

import (
	"bufio"
	"fmt"
	"os"
)

var inputReader bufio.Reader
var input string

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input : ")
	input, err := inputReader.ReadString('z') //z后面的不读
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(input)
}
