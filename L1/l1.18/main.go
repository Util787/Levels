package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	count atomic.Uint64
}

func (c *Counter) Increment() {
	c.count.Add(1)
}

func main() {
	counter := Counter{}
	numGor := 5

	wg := sync.WaitGroup{}

	for i := 0; i < numGor; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()
	fmt.Println(counter.count.Load())
}
