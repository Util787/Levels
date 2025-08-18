package main

import (
	"fmt"
	"sync"
)

var arr = [5]int{2, 4, 6, 8, 10}

func main() {
	wg := &sync.WaitGroup{}

	for _, val := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sqr(val) // версия го >1.22 значит можно передавать val напрямую в sqr() внутри анонимной функции, иначе пришлось бы брать val через аргумент в анонимной функции
		}()
	}

	wg.Wait()
}

func sqr(num int) {
	fmt.Println(num * num)
}
