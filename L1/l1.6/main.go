package main

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"sync"
)

func main() {
	stopByCondition()
	stopByNotificationChan()
	stopByCtx()
	stopByRuntime()
	stopByOs()
}

func stopByCondition() {
	defer fmt.Println("Stop by condition")
	finish := true

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			if finish {
				return
			}
			fmt.Println("Smth went wrong")
		}
	}()

	wg.Wait()
}

func stopByNotificationChan() {
	defer fmt.Println("Stop by channel")
	finishChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-finishChan:
				return
			default:
				fmt.Println("Smth went wrong")
			}
		}
	}()

	close(finishChan)
	wg.Wait()
}

func stopByCtx() {
	defer fmt.Println("Stop by ctx")
	wg := sync.WaitGroup{}
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	cancel() // для демонстрации

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Println("Smth went wrong")
			}
		}
	}()

	wg.Wait()
}

func stopByRuntime() {
	defer fmt.Println("Stop by runtime")
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		runtime.Goexit()
	}()

	wg.Wait()
}

func stopByOs() {
	defer fmt.Println("Stop by os")

	go func() {
		os.Exit(0)
	}()
}
