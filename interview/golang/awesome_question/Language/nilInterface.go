package Language

import (
	"fmt"
	"reflect"
)

type S struct{}

func (s S) F() {}

type IF interface {
	F()
}

func InitType() S {
	var s S
	return s
}

func InitPointer() *S {
	var s *S
	return s
}
func InitEfaceType() interface{} {
	var s S
	return s
}

func InitEfacePointer() interface{} {
	var s *S
	return s
}

func InitIfaceType() IF {
	var s S
	return s
}

func InitIfacePointer() IF {
	var s *S
	return s
}

func NilInterface() {
	// println(InitType() == nil)	//can not convert nil to type S
	//fmt.Println(S{} == nil)
	fmt.Println(InitType())
	println(InitPointer() == nil) //true
	fmt.Println(InitPointer())
	fmt.Println(reflect.TypeOf(InitPointer()).String())
	fmt.Println()

	println(InitEfaceType() == nil) //false
	fmt.Println(InitEfacePointer())
	fmt.Println(reflect.TypeOf(InitEfacePointer()).String())
	fmt.Println()

	println(InitEfacePointer() == nil) //false
	fmt.Println(InitEfacePointer())
	fmt.Println()

	println(InitIfaceType() == nil)    //false
	println(InitIfacePointer() == nil) //false
}
