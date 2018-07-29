package language

import (
	"fmt"
	_ "github.com/constabulary/gb/testdata/src/d.v1"
)

func main() {
	d := [...]struct {
		str string
		fn  func()
	}{
		{
			str: "test",
			fn: func() {
				println("------")
			},
		},
		{},
	}
	d[0].fn()
	fmt.Println(d)
}
