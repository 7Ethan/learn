package main

import (
	"reflect"
)

func init() {
}

func main() {
	println(typeof("Hello"))
}

func typeof(v interface{}) string {
	return reflect.TypeOf(v).String()
}
