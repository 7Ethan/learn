package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	outputFile, outputErr := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputErr != nil {
		fmt.Printf("An error occur with file opening")
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputString := "hello World "

	for i := 0; i < 10; i++ {
		outputWriter.WriteString(outputString)
	}
	outputWriter.Flush()
}
