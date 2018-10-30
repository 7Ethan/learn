package language

func main() {
	x := []int{1, 3, 4, 5}
	switch {
	case x[0] > 2:
		println(">2")
	case x[1] > 2:
		println(" < 2")
	default:
		println("--------")
	}
}
