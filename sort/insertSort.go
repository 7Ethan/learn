package main

import "fmt"

func insertSort(data []int) {
	fmt.Println("Insert Sort ")
	fmt.Println("resource data : ", data)

	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] { //insert
			for j := i; j > 0; j-- {
				if data[j] < data[j-1] {
					data[j], data[j-1] = data[j-1], data[j]
				}
			}
		}
		fmt.Println("middle resource : ", data)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 56, 43, -3, 6}
	insertSort(arr)
}
