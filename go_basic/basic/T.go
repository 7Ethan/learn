package language

import "fmt"

func main() {
	x := 1234
	p := &x
	*p++
	a := *p
	fmt.Println(a)
}
