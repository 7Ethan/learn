// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Person struct {
// 	Name string `xml:",attr"`
// 	Age int  `xml:"age"`
// 	}

// func main()  {
// 	p := Person{"sergey",10}
// 	var data []byte
// 	var err error
// 	if data,err = xml.MarshalIndent(p,"","  "); err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(string(data))

// 	p2 := new(Person)

// 	if err = xml.Unmarshal(data,p2);err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println(p2)
// }

