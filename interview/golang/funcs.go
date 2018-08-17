package main

func test(x int) (func(), func()) {
	return func() {
			println(x)
			x += 10
		}, func() {
			println(x)
		}
}

func main() {
	a, b := test(100)
	// b()
	// a()
	//执行结果都是一百
	//------------------

	a()
	b()
	//执行结果是 100 110
	//原因是a先执行,而x的作用域再函数内是全局的,所以a执行后被修改了
}
