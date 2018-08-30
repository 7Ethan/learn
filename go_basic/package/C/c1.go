package C

import (
	"fmt"

	"../A"
)

func init() {
	fmt.Println("Package C->init()")
}

func Hello_C() string {
	A.Hello_A()
	return fmt.Sprintf("package C using package A ->")
}
