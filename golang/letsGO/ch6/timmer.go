package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	var test time.Time
	fmt.Println(time.Now())
	c := time.After(5 * time.Second) //返回只读的channel,不能写入数据
	fmt.Println(c)
	test = <-c
	fmt.Println(reflect.TypeOf(test))

	tm := time.NewTimer(10 * time.Second)
	t := <-tm.C
	fmt.Println(t)
}
