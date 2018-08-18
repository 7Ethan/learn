package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := new(sync.WaitGroup)
	wg.Add(3)

	sem := make(chan int, 1)
	for i := 0; i < 3; i++ {
		go func(id int) {
			defer wg.Done()
			sem <- 1
			for x := 0; x < 3; x++ {
				println(id, x)
			}
			<-sem
		}(i)
	}
	wg.Wait()
}
