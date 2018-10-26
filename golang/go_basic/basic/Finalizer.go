package language

import (
	"fmt"
	"runtime"
	"time"
)

type data struct {
	x [100 * 1024]byte
	o *data
}

func test() {
	var a, b data
	a.o = &b
	b.o = &a
	runtime.SetFinalizer(&a, func(d *data) { fmt.Printf("a %p final.\n", d) })
	runtime.SetFinalizer(&b, func(d *data) { fmt.Printf("b %p final.\n", d) })
}

func main() {
	for {
		test()
		time.Sleep(time.Millisecond)
	}
}
