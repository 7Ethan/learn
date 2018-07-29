package main

import "fmt"

var (
	firstName string
	lastName  string
)

func main() {
	fmt.Println("please input : ")
	fmt.Scanln(&firstName, &lastName)
	fmt.Println(firstName, lastName)
}
