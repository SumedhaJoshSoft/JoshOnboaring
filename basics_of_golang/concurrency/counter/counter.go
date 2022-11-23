package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mu *sync.Mutex) {
	mu.Lock()
	x = x + 1
	mu.Unlock()
	wg.Done()
}
func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg, &mu)
	}
	wg.Wait()
	fmt.Println(x)
}
