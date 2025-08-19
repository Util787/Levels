package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}

	numChan := make(chan int)
	resChan := make(chan int)

	go func() { // если правильно понял по условию не надо делать отдельные функции для конвейеров
		for _, val := range arr {
			numChan <- val
		}
		close(numChan)
	}()

	go func() {
		for val := range numChan {
			resChan <- val * 2
		}
		close(resChan)
	}()

	for val := range resChan {
		fmt.Println(val)
	}
}
