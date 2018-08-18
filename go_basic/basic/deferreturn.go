package language

import "fmt"

func add(a, b int) (x int) {

	x = a + b
	defer func() {
		x += 100
	}()
	return

}

func main() {
	fmt.Print(add(10, 10))
}
