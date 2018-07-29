package Language

import (
	"fmt"
	"log"
)

func f() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover:%#v", r)
		}
	}()
	panic(2) //遇到panic(2) 则此函数发生panic,函数调用结束,然后延迟调用了recover()
	panic(1)
	fmt.Println("This will not print")
}

func main() {
	f() //结果是 : 2018/03/02 09:40:27 recover:2
}
