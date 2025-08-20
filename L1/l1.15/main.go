package main

import (
	"fmt"
	"strings"
	"sync"
)

var (
	justString string
	mu         sync.Mutex
)

func createHugeString(n int) string {
	b := strings.Builder{}
	for range n {
		b.WriteString("W") //просто для примера
	}
	return b.String()
}

func createJustString() {
	v := createHugeString(1 << 10)

	runes := []rune(v)

	mu.Lock()
	defer mu.Unlock()

	if len(runes) >= 100 { // можно было бы взять функцию RuneCountInString из utf8 но тогда пришлось бы дважды переводить в руны
		justString = string(runes[:100])
	}
}

func main() {
	createJustString()

	mu.Lock() // на случай если надо будет запускать createJustString в горутинах
	fmt.Println(justString)
	mu.Unlock()
}
