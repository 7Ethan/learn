package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	fmt.Println(now)
	sec := now.Unix()
	nanos := now.UnixNano()
	fmt.Println("Second : ",sec, " ,nanos :",nanos)

	millis := nanos / 100000
	fmt.Println(millis)

	fmt.Println("-------------")
	fmt.Println(time.Unix(sec,0))
	fmt.Println(time.Unix(0,nanos))
}
