package main

import "fmt"

func main() {
	prime := sieve()
	for {
		fmt.Println(<-prime)
	}

}

func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; i < 1000; i++ {
			ch <- i
		}
	}()
	return ch
}

func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch
			ch = filter(ch, prime)
			out <- prime
		}
	}()
	return out
}
