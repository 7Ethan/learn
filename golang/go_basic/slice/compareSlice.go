package main

import "fmt"

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		//i < len(a) && i < len(b)遍历两个切片长度中较小的次数
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func main() {
	ba := []byte("afadwe")
	bb := []byte("afadwe")
	fmt.Println(Compare(ba, bb))
}
