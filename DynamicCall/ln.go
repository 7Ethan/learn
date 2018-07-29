package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (*Data) Test(x, y int) (int, int) { //add 100
	return x + 100, y + 100
}

func (*Data) Sum(s string, x []int) string {
	c := 0
	for _, i := range x { //index, value
		c += i
	}
	return fmt.Sprintf(s, c)
}

func main() {
	d := new(Data)
	v := reflect.ValueOf(d)

	exec := func(name string, in []reflect.Value) {
		m := v.MethodByName(name)
		//out := m.Call(in)
		out := m.CallSlice(in)
		for _, v := range out {
			fmt.Println(v.Interface())
		}
	}

	exec("Test", []reflect.Value{
		reflect.ValueOf(1),
		reflect.ValueOf(2),
	})

	fmt.Println("---------")

	exec("Sum", []reflect.Value{
		//panic: reflect: Call using int as type string
		reflect.ValueOf("result = %d"), //注意参数!!!是因为少了字符串这个参数
		reflect.ValueOf([]int{100, 200, 400}),
		//reflect.ValueOf(2),
	})
}
