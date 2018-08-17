package Language

import (
	"fmt"
	"golang.org/x/tools/go/gcimporter15/testdata"
)

type People interface {
	Speak(string) string
}

type Student struct{}

func (stu *Student) Speak(think string) (talk string) {
	if think == "Excellent" {
		talk = "Nice"
	} else {
		talk = "Hi"
	}
	return
}

func g(x *interface{}) {

}

func main() {
	//var p People  = &Student{}
	var s Student //interface{} 变量不能通过短声明来定义

	//s := Student{}
	a := &s

	g(a)

	//think := "Excellent"
	//fmt.Println(p.Speak(think))
}
