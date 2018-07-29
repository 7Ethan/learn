package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		err := fmt.Errorf("usage: %s infile.txt outfile.txt", filepath.Base(os.Args[0]))
		//fmt.Println(err.Error())
		return
	}
}
