package main

import "fmt"

func main() {
	array := []int{55, 94, 87, 12, 4, 32, 11, 8, 39, 42, 64, 53, 70, 12, 9}
	fmt.Println("before MergeSort", array)
	array = MergeSort(array)
	fmt.Println("after MergeSort:", array)
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	key := len(arr) / 2
	left := MergeSort(arr[:key])
	right := MergeSort(arr[key:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}
