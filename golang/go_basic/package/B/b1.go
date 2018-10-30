package B

import (
	"fmt"

	A "../A"
)

func init() {
	fmt.Println("Package B -> init()")
}

func Hello_B() {
	A.Hello_A()
	fmt.Println("B say hello")
}
