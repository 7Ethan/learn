package main

import "fmt"

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

func qsort(values []int, left, right int) {
	//key := arr[start]
	//i,j := start,end
	//
	//for i <= j  {
	//	for arr[j] > key {
	//		j--
	//	}
	//	for arr[i] < key {
	//		i++
	//	}
	//
	//	if i <= j {
	//		arr[i],arr[j] = arr[j],arr[i]
	//		i++
	//		j--
	//	}
	//
	//	if start < j {
	//		qsort(arr,start,j)
	//	}
	//	if end > i {
	//		qsort(arr,i,end)
	//	}
	//	fmt.Println(arr)
	//}

	temp := values[left]
	p := left
	i, j := left, right

	for i <= j {
		for j >= p && values[j] >= temp {
			j--
		}
		if j >= p {
			values[p] = values[j]
			p = j
		}

		for i <= p && values[i] <= temp {
			i++
		}
		if i <= p {
			values[p] = values[i]
			p = i
		}

	}

	values[p] = temp

	if p-left > 1 {
		qsort(values, left, p-1)
	}
	if right-p > 1 {
		qsort(values, p+1, right)
	}

	fmt.Println(values)

}

func main() {
	//arr := []int{70,75,69,32,88,18,16,58}
	//QuickSort(arr,0,len(arr) - 1)
	//fmt.Println(arr)

	arr := []int{70, 75, 69, 32, 88, 18, 16, 58}
	qsort(arr, 0, len(arr)-1)
	//	//array := []int{3, 7, 9, 8, 38, 93, 12, 222, 45, 93, 23, 84, 65, 2}
	//
	//	//qsort(array,0,len(array) - 1 )
	//	//fmt.Println(arr)
	//	//fmt.Println(array)
}
