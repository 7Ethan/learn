/**
* @Author:  Sergey Jay
* @Email :  zshangan@iCloud.com
* @Create:  09/10/2017 21:36
* Copyright (c) 2017 Sergey Jay All rights reserved.
* Description:
 */

package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	//test := 123.00
	//fmt.Println(-test)
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)

	a = f // a MyFloat implements Abser
	fmt.Println(a.Abs())
}
