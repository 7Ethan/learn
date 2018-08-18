package language

func main() {
	test1(4, 2)

	defer func() {
		recover()
	}()
	panic("------")

	test2()
}

func test2() {
	defer func() {
		recover()
	}()
	panic("panic test")
}

func test1(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		z = x / y
		return
	}()

	println("x / y = ", z)
}
