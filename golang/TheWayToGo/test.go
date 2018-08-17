package main

import (
	"./mySort"
	"fmt"
)

type TTest []int

func main() {
	//data :=
	var data TTest = []int{34, 765, 38, -54, 12, 0, 789, -234, 56, 123, 657}

	fmt.Println("Resource :", data)
	//a := mySort.MySlice(data)
	//mySort.MySort(a)
	//fmt.Println("Sorted : ",a)
	a := mySort.MySorter(data)
	mySort.MySort(a)
	fmt.Println(data)
}

/*
type MySorter interface {
	Lenn()	int
	Less(i,j int)	bool
	Swap(i,j int)
}
*/

func (t TTest) Lenn() int {
	return len(t)
}

func (t TTest) Less(i, j int) bool {
	return t[i] < t[j]
}

func (t TTest) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
