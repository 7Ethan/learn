package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	done := make(chan struct{})
	ch := make(chan interface{})
	workcount := 5
	for i := 1; i < workcount; i++ {
		wg.Add(1)
		go doIt(i, ch, done, &wg)
	}

	for i := 1; i < workcount; i++ {
		ch <- i
	}

	close(done)
	wg.Wait()
	fmt.Println("Main routine shutdown.")
}

func doIt(workerID int, ch <-chan interface{}, done <-chan struct{}, wg *sync.WaitGroup) {
	fmt.Println(workerID, " start !")
	defer wg.Done()
	for {
		select {
		case <-done:
			fmt.Printf("%d have been shutdown.\n", workerID)
			return
		case <-ch:
			fmt.Printf("%d is working .\n", workerID)
		}
	}
}
