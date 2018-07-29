package main

import "fmt"

func selectSort(data []int) {
	fmt.Println("Resource data : ", data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			if data[i] > data[j] {
				data[i], data[j] = data[j], data[i]
			}
		}
		fmt.Printf("Middle Data:%2d=>%d\n", i, data)
	}
	fmt.Printf("Result Data:%d\n", data)
}

func main() {
	arr := []int{1, 2, 3, 4, 56, 43, -3, 6}
	selectSort(arr)
}
