package main

import (
	"encoding/xml"
	"fmt"
)

type Stu struct {
	XMLName xml.Name `xml:"student"`
	Name    string   //`xml:"name,attr"`
	Age     int      //`xml:"age,attr"`
	Phone   []string //`xml:"phones>phone",`
}

/*
	可以 使用 tag 来指定 Struct 的字段与 xml 标记的对应关系
*/

type ABC string

func main() {
	str := `<?xml version="1.0" encoding="utf-8"?>
				<student name="张三" age="19">
			   <phones>
			   <phone>12345</phone>
			   <phone>67890</phone>
			   </phones>
			   </student>`
	var s Stu
	xml.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}
