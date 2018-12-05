package Language

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func main() {
	fmt.Println(f())
	//function variable
	fn := func() {
		fmt.Println("hello")
	}
	fn()
	fmt.Printf("%T\n", fn) //打印func()
}
