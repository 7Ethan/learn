package main

import (
	"encoding/xml"
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
			<Student>
			<Name>张三</Name>
			<Age>19</Age>
			</Student>`
	var s student
	xml.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
