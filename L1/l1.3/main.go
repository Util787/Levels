package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numWorkers int
	fmt.Println("Enter num of workers")
	fmt.Scan(&numWorkers)

	intChan := make(chan int)

	for range numWorkers {
		go worker(intChan)
	}

	for { // т.к. запись постоянная не стал делать wg и закрывать канал
		intChan <- rand.Intn(10)
	}

}

func worker(ch <-chan int) {
	for val := range ch {
		fmt.Println(val)
	}
}
