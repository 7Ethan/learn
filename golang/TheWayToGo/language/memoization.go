package main

import (
	. "fmt"
	"reflect"
	"runtime"
	_ "time"
)

type name struct {
	key   string `key--------`
	value int    `value`
}

var names = make([]name, 10)

const LIM = 41

var fib [LIM]uint64

func main() {
	names = []name{{"Sergey", 1}, {"jay", 2}}
	t := reflect.TypeOf(names)
	Println(t.Elem().String())
	println("-----------")

	n := name{
		"hahah",
		3,
	}

	println(reflect.TypeOf(n).Field(0).Tag)
	//var result uint64
	//start := time.Now()
	//for i := 0; i < LIM ; i++  {
	//	result = fibonacci(i)
	//	fmt.Printf("fibonacci[%d] is : %d \n",i,result)
	//}
	//end := time.Now()
	//delta := end.Sub(start)
	//fmt.Println("time is : ",delta)

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	Printf(" %d kb \n", m.Alloc/1024)
}

func fibonacci(n int) (res uint64) {
	if fib[n] != 0 {
		res = fib[n]
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fib[n] = res
	return
}
