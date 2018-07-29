package main

import (
	"fmt"
	"time"
)

type test struct {
	name string
	age  int
}

func (t test) change() {
	t.name = "newName"
	t.age = 88
}

func (t *test) change1() {
	t.name = "newName"
	t.age = 100
}

func main() {
	fmt.Println(time.Now())
	t := test{"xiaoming", 18}
	fmt.Println(t)
	fmt.Println("------------")
	t.change1()
	fmt.Println(t)
}
