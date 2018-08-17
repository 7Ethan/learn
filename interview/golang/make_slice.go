package main

import "fmt"

func main() {
	list := new([]int)
	// new list 返回的是一个切片的指针....
	// list = append(list, 1)
	fmt.Printf("%#T,%#v", list, list)
}

/**
 * 此外 []int形式一般是切片形式,[n]int才是数组形式,
 * 数组的类型包括数据类型以及数组的大小,而且数组大小已经定义,不能在改变.
 */
