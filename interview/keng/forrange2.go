package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		v := value //没有这行就不能成功赋值
		myMap[index] = &v
	}
	prtMap(myMap)
}

func prtMap(myMap map[int]*int) {
	for key, value := range myMap {
		fmt.Printf("map[%v]=%v\n", key, *value)
	}
}
