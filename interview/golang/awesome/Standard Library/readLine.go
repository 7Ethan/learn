package main

import (
	"os"
	//"io/ioutil"
	//"fmt"
	"bufio"
	"io"
)

func processLine(line []byte) {
	os.Stdout.Write(line)
}

func ReadLine(filPath string, handleFunc func([]byte)) error {
	f, err := os.Open(filPath)
	if err != nil {
		return err
	}
	defer f.Close()

	bufRd := bufio.NewReader(f)

	for {
		line, err := bufRd.ReadBytes('\n')
		handleFunc(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}

	return nil
}
func main() {
	//file,err := os.OpenFile("test.txt",os.O_RDWR,0666)
	//if err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	////bufio.Scanner.S
	//
	//fmt.Println(ioutil.ReadAll(file))
	ReadLine("test.txt", processLine)
}
