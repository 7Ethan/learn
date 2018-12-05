package Language

import (
	"fmt"
	"reflect"
)

const N = 3

func main() {
	m := make(map[int]*int)

	for i := 0; i < N; i++ {
		fmt.Println(reflect.TypeOf(i))
		j := int(i)
		m[i] = &j
	}

	for _, v := range m {
		print(*v)
	}
}
