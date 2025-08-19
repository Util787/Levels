package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	fmt.Println(split(arr))
}

func split(arr []float64) map[int][]float64 {
	m := make(map[int][]float64)

	for _, val := range arr {
		k := int(makeKey(val))
		m[k] = append(m[k], val)
	}

	return m
}

func makeKey(num float64) float64 {
	if num < 0 {
		return math.Ceil(num/10) * 10 // чтобы например из -25.4 получить ключ -20 надо округлять в большую сторону с положительными наоборот
	} else {
		return math.Floor(num/10) * 10
	}
}
