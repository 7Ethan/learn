package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	result, err := reader.ReadString('`')
	if err != nil {
		fmt.Println("read err :", err)
	}
	// fmt.Println("result: ", string(result))
	fmt.Println("result: ", result)
}
