package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	Name string
}

var single *singleton
var once sync.Once

func GetInstance() *singleton {
	once.Do(func() {
		single = &singleton{
			Name: "sergey",
		}
	})
	return single
}

func main() {
	// result := GetInstance()
	fmt.Printf("---> %p\n", GetInstance())
	a := test()
	fmt.Println(a)
}

func test() *singleton {
	fmt.Printf("%p\n", GetInstance())
	return GetInstance()
}
