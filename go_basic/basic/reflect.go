package main

import (
	"fmt"
	"reflect"
)

type S struct{}

func (S) Sval() {
	fmt.Println("sval")
}

func (*S) Sptr() {
	fmt.Println("sptr")
}

func method(a interface{}) {
	t := reflect.TypeOf(a)
	// _ := reflect.ValueOf(a)
	for i, n := 0, t.NumMethod(); i < n; i++ {
		m := t.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func main() {
	s := &S{}
	arr := [9]int{}
	// method(s)
	t := reflect.TypeOf(s)
	fmt.Println("number of t's method:", t.NumMethod())
	fmt.Println(reflect.ValueOf(arr).Type())
}
