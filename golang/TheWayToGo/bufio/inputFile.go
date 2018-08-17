package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	inputFile, err := os.Open("input.dat")
	defer inputFile.Close()
	if err != nil {
		fmt.Println("An error occured on opening the inputfile\n" +
			"Does the file exit>")
		return
	}

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readRrr := inputReader.ReadString('\n')
		fmt.Println(inputString)
		//此处要注意执行顺序,否则不能正确输出
		if readRrr == io.EOF {
			return
		}

	}
}
