package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	f, err := os.OpenFile("jsonData.dat", 1, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.Seek(0, io.SeekStart)
	//fmt.Println(*f)
	var s Student
	decoder := json.NewDecoder(f)
	decoder.Decode(&s)
	fmt.Println(s)
}
