package Language

import (
	"math/rand"
	"sync"
)

const N = 10

func main() {
	m := make(map[int]int)

	mu := sync.Mutex{}
	wg := &sync.WaitGroup{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			m[rand.Int()] = rand.Int()
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
}
