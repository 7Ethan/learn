package language

import (
	"fmt"
	//"unsafe"
)

func main() {
	c := make(chan string, 10)
	<-c
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := data[1:3]
	s2 := append(s, 100, 200)

	fmt.Println(data, "  ", (data))
	fmt.Println(s, "  ", &s)
	fmt.Println(s2, "  ", &s2)

	fmt.Printf(" %p \n ", &data)
	fmt.Printf(" %p \n ", &s)
	fmt.Printf(" %p \n ", &s2)

	//println(unsafe.Sizeof(&data))
	//println(unsafe.Sizeof(&s))
	//println(unsafe.Sizeof(&s2))
}
