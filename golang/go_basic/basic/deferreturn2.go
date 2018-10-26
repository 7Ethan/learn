package language

func add(a, b int) (x int) {
	defer func() {
		println(x)
	}()
	x = a + b
	return x + 100
}

func main() {
	println(add(1, 2))
}
