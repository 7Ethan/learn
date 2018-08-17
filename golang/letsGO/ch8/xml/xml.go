package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Name string
	age  int
}

func main() {
	f, err := os.Create("data1.dat")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()
	s := &student{"zhang san", 19}
	encoder := xml.NewEncoder(f)
	encoder.Encode(s)
	f.Seek(0, io.SeekStart)

	decoder := xml.NewDecoder(f)
	var s1 student
	decoder.Decode(&s1)
	fmt.Println(s1)
}
