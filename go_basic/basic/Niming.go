package language

import (
	"fmt"
)

func main() {
	// --- function variable ---
	//fn := func() { println("hello")}
	//fn()

	// --- function collection ---
	//fns := [](func(x int)int){
	//	func(x int) int {return x + 1},
	//	func(x int) int {return x +2},
	//	func(x int) int {return x + 3},
	//}
	//fns2 := [](func(a,b int)int){
	//	func(a, b int) int {
	//		return a + b
	//	},
	//}
	////fns = [](0)
	//println(fns[3](1))
	//println(fns2[0](1,2))

	//---------function as field
	//d := struct {
	//	str string
	//	fn func()
	//}{
	//	fn : func() {
	//		println("-------")
	//	},
	//}
	////d.fn()
	////fmt.Println(d.fn())
	//d.str = "string"
	////d.fn()
	//fmt.Println(d)
	//d.fn()

	//-------channel of function----------
	//fc := make(chan func()string,2)
	//fc <- func() string {
	//	return "-----------"
	//}
	//fc <- func() string {
	//	return "==========="
	//}
	//println((<-fc)())
	//println((<-fc)())
	f := test()
	f()
}

func test() func() {
	x := 100
	fmt.Printf("x1 (%p) = %d\n", &x, x)
	return func() {
		fmt.Printf("x2 (%p) = %d\n", &x, x)
	}
}
