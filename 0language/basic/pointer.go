package language

type x struct {
}

func (*x) test() {
	println("x.test()")
}

func main() {
	p := &x{}
	p.test()
}
