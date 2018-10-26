package language

import (
	"unsafe"
	//"fmt"
	"time"
)

type data struct {
	x [100 * 1024]byte
}

func test() uintptr {
	p := &data{}
	return uintptr(unsafe.Pointer(p))
}
func main() {
	//x := 0x12345678
	//
	//p := unsafe.Pointer(&x)
	//a := (*[4]byte)(p)
	//
	//for i := 0;i < len(a);i++ {
	//	fmt.Printf("%X ",a[i])
	//}

	const N = 1000
	cache := new([N]uintptr)

	for i := 0; i < N; i++ {
		cache[i] = test()
		time.Sleep(time.Millisecond)
	}
}
