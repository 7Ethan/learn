package Language

import (
	"fmt"
)

func test() {
	s := []int{1, 2, 3, 4, 6, 9}
	ss := s[1:4]
	fmt.Println(ss)
	ss = append(ss, 4)

	for _, v := range ss {
		v += 10
	}

	for i := range ss {
		ss[i] += 10
	}

	fmt.Println(s)
	fmt.Println(ss)
}
