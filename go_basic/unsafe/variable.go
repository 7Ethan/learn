package main

func main() {
	var (
		a struct{}
		b int = 1
	)
	_ = a
	_ = b
	println(&a, &b)
}
