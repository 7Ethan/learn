package main

import (
	"sync"
	//"fmt"
	"runtime"
)

func main() {
	//runtime.GOMAXPROCS(1)
	//wg := new(sync.WaitGroup)
	//wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//
	//	for i := 0;i < 5;i++ {
	//		fmt.Println(i)
	//		if i == 5 {
	//			runtime.Gosched()
	//		}
	//	}
	//}()
	//
	//go func() {
	//	defer  wg.Done()
	//	println("hello")
	//}()
	//wg.Wait()

	wg := new(sync.WaitGroup)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 6; i++ {
			println(i)
			if i == 3 {
				runtime.Gosched()
			}
		}
	}()
	go func() {
		defer wg.Done()
		println("Hello, World!")
	}()
	wg.Wait()
}
