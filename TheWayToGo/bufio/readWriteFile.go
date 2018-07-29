package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	inputFile := "input.dat"
	outputFile := "input_copy.txt"

	inputSlice, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "file error : %s", err)
	}
	fmt.Printf("%s", string(inputSlice))
	werr := ioutil.WriteFile(outputFile, inputSlice, 0644)
	if werr != nil {
		panic(werr.Error())
	}
}
