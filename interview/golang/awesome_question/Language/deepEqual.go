package Language

import (
	"fmt"
	"reflect"
)

type s struct {
	a, b, c string
}

func main() {
	a := interface{}(&s{"a", "b", "c"})
	b := interface{}(&s{"a", "b", "c"})

	fmt.Println(reflect.DeepEqual(a, b))
}
