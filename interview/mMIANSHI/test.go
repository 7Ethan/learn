package main

import "fmt"

func Test(person string) func() string {
	/*
		Do someting
	*/
	return func() string {
		return (person + " is working")
	}
}
func main() {
	p := Test("Sergey")
	fmt.Println(p())
}
