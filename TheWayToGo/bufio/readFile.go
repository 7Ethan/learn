package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input_copy.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var col1, col2, col3 []string
	for {
		var s1, s2, s3 string

		_, err := fmt.Fscanln(file, &s1, &s2, &s3)
		if err != nil {
			break
		}
		col1 = append(col1, s1)
		col2 = append(col2, s2)
		col3 = append(col3, s3)
	}

	fmt.Println(col1)
	fmt.Println(col2)
	fmt.Println(col3)
}
