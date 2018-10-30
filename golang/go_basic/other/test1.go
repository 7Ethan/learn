package main

import "fmt"

func main() {
	jay := person{10, "Jay"}
	fmt.Println(jay)
	modify(jay)
	fmt.Println(jay)
}
func modify(p person) {
	p.age = p.age + 10
}

type person struct {
	age  int
	name string
}
