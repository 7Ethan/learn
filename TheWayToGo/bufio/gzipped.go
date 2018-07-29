package main

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"os"
)

func main() {
	fName := "MyFile.gz"
	var r *bufio.Reader
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}

	fz, err := gzip.NewReader(file)
	if err != nil {
		r = bufio.NewReader(file)
	} else {
		r = bufio.NewReader(fz)
	}

	for {
		line, err := r.ReadString('\n')
		if err != nil {
			fmt.Println("Done reading file ")
			os.Exit(0)
		}
		fmt.Println(line)
	}

}
