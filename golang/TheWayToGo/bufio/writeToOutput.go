package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "input.dat"
	outputFile := "output.txt"

	inputSlice, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Read file error %s :", err)
	}
	fmt.Printf("read from input file : %s", inputSlice)

	writeErr := ioutil.WriteFile(outputFile, inputSlice, 0644)
	if writeErr != nil {
		panic(writeErr.Error())
	}

}
