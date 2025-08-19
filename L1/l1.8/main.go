package main

import (
	"fmt"
)

func setBit(n int64, i uint, bit uint) int64 {
	if bit == 1 {
		return n | (1 << i)
	} else {
		return n &^ (1 << i)
	}
}

func main() {
	var n int64 = 5  // 0101
	var i uint = 0   // индекс первого бита
	var bit uint = 0 // 0 или 1

	res := setBit(n, i, bit)
	fmt.Println("Src: ", n)
	fmt.Println("Res: ", res)
}
