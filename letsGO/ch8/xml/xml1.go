package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := &Student{"张三", 19}
	result, err := xml.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(result))
}
