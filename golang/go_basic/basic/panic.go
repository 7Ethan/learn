package language

import "fmt"

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err.(string))
		}
	}()
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("test")
	}()

	defer func() {
		panic("will not print")
	}()

	panic("panic Error")
}

func main() {
	//test()
	defer except()
	panic("test panic") //此处不会打印panic信息,因为调用了recover()是的panic已经恢复
}

func except() {
	recover()
}
