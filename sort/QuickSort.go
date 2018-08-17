package main

import (
	"fmt"
	"math/rand"
)

func QuickSort(arr []int, start, end int) {
	if start < end {
		key := arr[(start+end)/2]
		i, j := start, end

		for i <= j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}

			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
				fmt.Println(arr)
			}
		}

		if start < j {
			QuickSort(arr, start, j)
		}

		if end > i {
			QuickSort(arr, i, end)
		}
	}
}

func qsort(arr []int, start, end int) {
	if start < end {
		key := arr[rand.Intn(len(arr)-1)]
		i, j := start, end

		for arr[i] < key {
			i++
		}
		for arr[j] > key {
			j--
		}
		if i <= j {
			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}
		if start < j {
			qsort(arr, start, j)
		}
		if end > i {
			qsort(arr, i, end)
		}
	}
}

func main() {
	//arr := []int{70,75,69,32,88,18,16,58}
	//QuickSort(arr,0,len(arr) - 1)
	//fmt.Println(arr)

	arr := []int{70, 75, 69, 32, 88, 89, -456, 23, 0, 45, 654, 23, 7, 52, 76, 18, 16, 58}
	array := []int{19, 2, 3, 4, 56, 43, -3, 6}

	qsort(arr, 0, len(arr)-1)
	qsort(array, 0, len(array)-1)
	//	//array := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//
	//	//qsort(array,0,len(array) - 1 )
	fmt.Println(arr)
	//	//fmt.Println(array)
	fmt.Println(array)

}
