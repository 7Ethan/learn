package main

import (
	"fmt"

	A "./A"
	// 多次引用同一个包,一个程序只会执行一次init()函数
	B "./B"
	C "./C"
)

func main() {
	fmt.Print("A -->hello(): ")
	A.Hello_A()
	B.Hello_B()
	fmt.Println(C.Hello_C())
}

/**

在同一个package中，可以多个文件中定义init方法
在同一个go文件中，可以重复定义init方法
在同一个package中，不同文件中的init方法的执行按照文件名先后执行各个文件中的init方法
在同一个文件中的多个init方法，按照在代码中编写的顺序依次执行不同的init方法

*/
