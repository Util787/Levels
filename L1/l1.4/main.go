package main

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	var numWorkers int
	fmt.Println("Enter num of workers")
	fmt.Scan(&numWorkers)

	intChan := make(chan int)

	// В задании просят обосновать выбор, я так понял подразумевается выбор между передачей контекста и самого канала quit в горутины,
	// как по мне контекст более абстрактный и понятный(сразу ясно что надо обрабатывать завершение или прокидывать дальше) чем какой-то финиш канал
	baseCtx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdownWg := &sync.WaitGroup{} // обычно в грейсфул я просто дожидаюсь err ответа от функции завершения чего-либо но тут учитывая условие только такой способ удостовериться в завершении придумал

	for range numWorkers {
		shutdownWg.Add(1)
		go func() {
			defer shutdownWg.Done()
			worker(baseCtx, intChan)
		}()
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT)

	for {
		select {
		case <-quit:
			slog.Info("Shutting down gracefully...")
			cancel()
			shutdownWg.Wait()
			slog.Info("Shutdown complete")
			return

		default:
			intChan <- rand.Intn(10)
			time.Sleep(time.Second) // для наглядности
		}
	}

}

func worker(ctx context.Context, ch <-chan int) {
	for {
		select {
		case <-ctx.Done():
			slog.Info("Stopping worker")
			return
		case val := <-ch:
			fmt.Println(val)
		}
	}
}
