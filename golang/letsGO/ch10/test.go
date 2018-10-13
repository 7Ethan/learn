package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int16
	Haha int32
	Age  int8
}

func main() {
	var u User = User{001, 999, 22}
	uType := reflect.TypeOf(u)
	fmt.Println(uType.Align())
}
