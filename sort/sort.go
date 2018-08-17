package main

import (
	"fmt"
)

func main() {
	arr1 := []int{55, 94, 87, 12, 4, 32, 11, 8, 39, 42, 64, 53, 70, 12, 9}
	// BubbleSort(arr1)
	// InsertSort(arr1)
	// Selectsort(arr1)
	Shellsort(arr1)
	fmt.Println(arr1)
}

func BubbleSort(data []int) {
	fmt.Println("Resource data : ", data)
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	fmt.Printf("Result Data:%d\n", data)
}

func Selectsort(data []int) {
	fmt.Println("SelectSort")
	fmt.Println("Resource data : ", data)
	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data)-i-1; j++ {
			if data[j] < data[i] {
				data[j], data[i] = data[i], data[j]
			}
		}
	}
	fmt.Printf("Result Data:%d\n", data)
}

func InsertSort(data []int) {
	fmt.Println("InsertSort.")
	fmt.Println("Resource data : ", data)
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			for j := i; j > 0; j-- {
				if data[j] < data[j-1] {
					data[j], data[j-1] = data[j-1], data[j]
				}
			}
		}
	}
	fmt.Printf("Result Data:%d\n", data)
}

func Shellsort(arr []int) {
	h := 0
	for h <= len(arr)/3 {
		h = h*3 + 1
	}

	for h >= 1 {
		for i := h; i < len(arr); i++ {
			for j := i; j >= h && arr[j] > arr[j-1]; j -= h {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		h /= 3
	}
}
