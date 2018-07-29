package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string
	Age  int
}

func main() {
	s := Student{"xiaoming", 18}
	rv := reflect.ValueOf(s)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	reFiled := rv.FieldByName("Age")
	fmt.Println(reFiled.String())
}
