package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}
	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn3 := struct {
		name string
		age  int
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	if reflect.DeepEqual(sn1, sn3) {
		fmt.Println("sn1 ==sn3")
	} else {
		fmt.Println("sn1 !=sn3") //可以知道,这是不一样的的
	}

	fmt.Println(unsafe.Alignof(sn1))
	fmt.Println(unsafe.Alignof(sn3))
	// if sn1 == sn3 {
	// 	fmt.Println("sn1 == sn3")
	// }

	// sm1 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}
	// sm2 := struct {
	// 	age int
	// 	m   map[string]string
	// }{age: 11, m: map[string]string{"a": "1"}}

	// if sm1 == sm2 {
	// 	fmt.Println("sm1 == sm2")
	// }
}
