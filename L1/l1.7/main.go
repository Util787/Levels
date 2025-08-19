package main

import (
	"fmt"
	"sync"
)

type MyMap struct {
	mu sync.RWMutex
	m  map[string]int
}

func NewMyMap() *MyMap {
	return &MyMap{
		m: make(map[string]int),
	}
}

func (m *MyMap) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.m[key] = value
}

func (m *MyMap) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.m[key]
	return val, ok
}

func main() {
	MyMap := NewMyMap()
	wg := &sync.WaitGroup{}

	// тестирую на одном и том же ключе
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			MyMap.Set("key1", i)
		}()
	}

	for i := 0; i < 10; i++ {
		if val, ok := MyMap.Get("key1"); ok {
			fmt.Printf("key1 = %d\n", val)
		}
	}

	wg.Wait()
}
