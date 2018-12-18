package main

import (
	"encoding/json"
	"fmt"
)

// func CallBack() string {
// 	return "Call back function!"
// }

type CallBack func() string

type Info struct {
	Email string
}

type Student struct {
	Id    int64
	Name  string
	Other *Info
	Func  CallBack
}

func main() {
	info := &Info{"123@qq.com"}
	stu := Student{Id: 10001, Name: "ehrn", Other: info, Func: func() string {
		return "Call back function"
	}}
	fmt.Println(stu)

	result, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("Err : ", err)	//Err :  json: unsupported type: main.CallBack
									// go 的json 标准库不支持编解码 function
	}
	fmt.Println(string(result))
}
