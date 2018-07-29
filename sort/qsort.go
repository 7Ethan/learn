package main

import (
	"fmt"
	"math/rand"
)

func qsort(s []int) []int {
	if len(s) < 1 {
		return s
	}

	mark := s[rand.Intn(len(s))]
	var mid []int
	var left []int
	var right []int

	for _, n := range s {
		switch {
		case n > mark:
			right = append(right, n)
		case n < mark:
			left = append(left, n)
		default:
			mid = append(mid, n)
		}
	}

	left = qsort(left)
	right = qsort(right)

	var result []int
	result = append(left, mid...)
	result = append(result, right...)

	return result
}

func main() {
	arr := []int{345, 89, -456, 23, 0, 45, 654, 23, 7, 52, 769}
	fmt.Println(qsort(arr))
}
