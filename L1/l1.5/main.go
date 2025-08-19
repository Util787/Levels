package main

import (
	"fmt"
	"time"
)

const N = 1

func main() {
	numInts := 100
	ch := make(chan int, numInts) // я сначала сделал запись через горутину но я так понял что она должна быть последовательна именно в main следовательно нужен буфер

	for i := 0; i < numInts; i++ {
		ch <- i
	}
	close(ch)

	go func() {
		for val := range ch {
			fmt.Println(val)
		}
	}()

	<-time.After(time.Second * N) // просто блокирую main для ожидания
}
