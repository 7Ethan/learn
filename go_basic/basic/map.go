package language

import "fmt"

func main() {
	n_map := make(map[int]bool)
	for i := 1; i <= 10; i++ {
		n_map[i] = true
	}

	//for num, _ := range n_map {
	//	fmt.Print(num)
	//}
	fmt.Print("\n")
	for num, _ := range n_map {
		fmt.Print(num)
	}
}
