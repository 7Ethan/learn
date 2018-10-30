package language

import (
	"fmt"
	"reflect"
)

var (
	Int    = reflect.TypeOf(0)
	String = reflect.TypeOf("")
)

func main() {
	t := reflect.TypeOf(make(chan int, 3)).Elem()
	fmt.Println(t)
	//	c := reflect.ChanOf(reflect.SendDir,String)
	//	fmt.Println(c)
	//
	//	println("-----"	)
	//
	//	m := reflect.MapOf(String,Int)
	//	fmt.Println(m)
	//
	//	println("-----"	)
	//
	//	s := reflect.SliceOf(Int)
	//	fmt.Println(s)
	//
	//	println("-----"	)
	//
	//	t := struct {
	//		Name string
	//	}{}
	//	p := reflect.PtrTo(reflect.TypeOf(t))
	//	fmt.Println(p)
}
