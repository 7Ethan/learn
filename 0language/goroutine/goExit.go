package main

import (
	"runtime"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer println("a.defer")

		func() {
			defer println("b.defer")
			runtime.Goexit()
			println("defer")
		}()
		println("done")
	}()
	wg.Wait()
}
