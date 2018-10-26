package language

import "fmt"

func test(s string, n ...int) string {
	var sum int
	for i := range n {
		sum += i
	}
	return fmt.Sprintf(s, sum)
}

func main() {
	fmt.Print(test("the sum is : %d", 1, 2, 3))
}
