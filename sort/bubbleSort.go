package main

import "fmt"

func bubbleSort(data []int) {
	fmt.Println("BubbleSort :")
	fmt.Println("Source Data : ", data)

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
		fmt.Println(i, " : middle data -> ", data)
	}
	fmt.Println("Result is : ", data)
}

func main() {
	arr := []int{19, 2, 3, 4, 56, 43, -3, 6}
	bubbleSort(arr)
}
