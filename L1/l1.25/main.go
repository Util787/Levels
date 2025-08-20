package main

import (
	"fmt"
	"time"
)

func main() {
	delay := time.Second * 5

	sem := make(chan struct{})

	fmt.Println("start")

	go func() {
		timer := time.NewTimer(delay)
		<-timer.C
		sem <- struct{}{}
	}()

	<-sem
	fmt.Println("fin")
}
