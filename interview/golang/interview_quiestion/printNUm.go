package m面试

import (
	//"fmt"
	//"runtime"
	"fmt"
	"time"
)

func main() {
	chan_int := make(chan bool)
	chan_char := make(chan bool, 1)
	done := make(chan struct{})

	str := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K"}

	//chan_char <- true
	go func() {
		for i := 1; i < 11; i += 2 {
			<-chan_char
			fmt.Print(i)
			fmt.Print(i + 1)
			chan_int <- true
		}
	}()

	go func() {
		for i := 0; i < 10; i += 2 {
			<-chan_int
			fmt.Print(str[i])
			fmt.Print(str[i+1])
			chan_char <- true
		}
		done <- struct{}{}
	}()
	//go func() {
	//	for i := 1;i < 11; i += 2 {
	//		fmt.Print(i)
	//		fmt.Print(i + 1)
	//		runtime.Gosched()
	//	}
	//}()
	//
	//go func() {
	//	for i := 0;i < 10 ; i += 2{
	//		fmt.Print(str[i])
	//		fmt.Print(str[i + 1])
	//		runtime.Gosched()
	//	}
	//}()

	chan_char <- true
	<-done

	time.Sleep(10 * time.Second)
}
